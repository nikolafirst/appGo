package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"work/internal/database/users"
	"work/internal/env"

	"github.com/google/uuid"
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
	_ = e
	create, err := e.UsersRepository.Create(
		ctx, users.CreateUserReq{
			ID:       uuid.New(),
			Username: "random",
			Password: "password",
		},
	)
	if err != nil {
		return err
	}

	found, err := e.UsersRepository.FindByID(ctx, create.ID)
	if err != nil {
		return err
	}

	foundBy, err := e.UsersRepository.FindByUsername(ctx, "random")
	if err != nil {
		return err
	}
	fmt.Println(create, found, foundBy)
	return nil
}
