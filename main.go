package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	httpserver "github.com/alexrehtide/sebastian/internal/http-server"
	"github.com/alexrehtide/sebastian/migrations"
	"github.com/alexrehtide/sebastian/platform/config"
	mongoconnection "github.com/alexrehtide/sebastian/platform/database/mongo-connection"
	sqlconnection "github.com/alexrehtide/sebastian/platform/database/sql-connection"
	"github.com/alexrehtide/sebastian/platform/migrator"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log := logrus.New()

	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := sqlconnection.New(sqlconnection.PostgresOptions{
		User:     config.PostgresUser,
		Password: config.PostgresPassword,
		Host:     config.PostgresHost,
		Port:     config.PostgresPort,
		DBName:   config.PostgresName,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Connection with db failed: %v", err)
	}

	migrator, err := migrator.New(
		migrator.MigratorOptions{
			DBConn:          sqlDB,
			MigrationsTable: "migrations_table",
		},
		migrations.FS,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := migrator.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	mongoDB, err := mongoconnection.New(mainCtx, mongoconnection.MongoOptions{
		Host: config.MongoHost,
		Port: config.MongoPort,
		Name: config.MongoName,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := httpserver.New(mongoDB, sqlDB, log)

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		return server.Listen(":3000")
	})
	g.Go(func() error {
		<-gCtx.Done()
		return server.Shutdown()
	})

	log.WithField("port", 3000).Infof("Application started")

	if err := g.Wait(); err != nil {
		log.Errorf("exit reason: %s \n", err)
	}
}
