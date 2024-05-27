package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os/signal"
	"sync"
	"syscall"

	"appGo/internal/env"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	if err := runMain(ctx); err != nil {
		log.Fatal(err)
	}
}

func runMain(ctx context.Context) error {
	e, err := env.Setup(ctx)
	if err != nil {
		return fmt.Errorf("setup.Setup: %w", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(4)

	logger := e.Logger
	grpcServer := e.LinksGRPCServer

	go func() {
		defer wg.Done()
		<-ctx.Done()
		// если посылаем сигнал завершения то завершаем работу нашего сервера
		grpcServer.Stop()
	}()

	go func() {
		defer wg.Done()
		<-ctx.Done()

		if err := e.AMQPCloser(); err != nil {
			logger.Error("amqp closer", slog.Any("err", err))
			return
		}
	}()

	// Создаем воркер для прослушки очереди и обновления
	go func() {
		defer wg.Done()
		if err := e.LinkUpdater.Run(ctx); err != nil {
			slog.Error("link updater Run: %w", err)
		}
	}()

	go func() {
		defer wg.Done()

		slog.Info(fmt.Sprintf("links grpc was started %s", e.Config.LinksService.GRPCServer.Addr))

		lis, err := net.Listen("tcp", e.Config.LinksService.GRPCServer.Addr)
		if err != nil {
			slog.Error("net Listen", slog.Any("err", err))
			return
		}

		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("net Listen", slog.Any("err", err))
			return
		}
	}()

	wg.Wait()

	return nil
}
