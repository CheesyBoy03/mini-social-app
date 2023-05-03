package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	minisocialapp "github.com/CheesyBoy03/mini-social-app"
	"github.com/CheesyBoy03/mini-social-app/config"
	"github.com/CheesyBoy03/mini-social-app/pkg/handler"
	"github.com/CheesyBoy03/mini-social-app/pkg/repository"
	"github.com/CheesyBoy03/mini-social-app/pkg/service"
)

func main() {
	config.Init()
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     config.Cfg.String("database.host"),
		Port:     config.Cfg.Int("database.port"),
		Username: config.Cfg.String("database.username"),
		Password: config.Cfg.String("database.password"),
		DBName:   config.Cfg.String("database.name"),
	})
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewServices(service.Dependencies{
		Repos:      repos,
		HashSalt:   config.Cfg.String("server.hash_salt"),
		SigningKey: []byte(config.Cfg.String("server.signing_key")),
	})
	handlers := handler.NewHandler(services)

	server := minisocialapp.NewServer()
	go func() {
		if err := server.Run(config.Cfg.String("server.port"), handlers.Init()); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		log.Fatal(err)
	}

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
