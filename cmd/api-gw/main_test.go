package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

// Тест для функции runMain
func TestRunMain(t *testing.T) {
	// Создаем фэйковый контекст
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаем фэйковый сервер
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer.Close()

	// Мокаем env.Setup и возвращаем фэйковые значения
	envSetup = func(ctx context.Context) (*env.Env, error) {
		return &env.Env{
			ApiGWHTTPServer: &http.Server{Addr: testServer.URL},
			Config:          &env.Config{ApiGWService: &env.ServiceConfig{Addr: "testAddr"}},
		}, nil
	}
	defer func() { envSetup = env.Setup }()

	err := runMain(ctx)
	if err != nil {
		t.Errorf("runMain() returned an error: %v", err)
	}
}

// Тест для обработчика сигналов
func TestSignalHandler(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Добавляем задержку, чтобы имитировать работу сервера
	}))
	defer testServer.Close()

	envSetup = func(ctx context.Context) (*env.Env, error) {
		return &env.Env{
			ApiGWHTTPServer: &http.Server{Addr: testServer.URL},
			Config:          &env.Config{ApiGWService: &env.ServiceConfig{Addr: "testAddr"}},
		}, nil
	}
	defer func() { envSetup = env.Setup }()

	go func() {
		time.Sleep(1 * time.Second) // Ждем некоторое время перед отправкой сигнала завершения
		cancel()                    // Отправляем сигнал завершения
	}()

	go func() {
		defer wg.Done()
		err := signalHandler(ctx, &wg)
		if err != nil {
			t.Errorf("signalHandler() returned an error: %v", err)
		}
	}()

	wg.Wait()
}
