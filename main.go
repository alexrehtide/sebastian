package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	httpserver "github.com/alexrehtide/sebastian/internal/http-server"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "3769"
	dbname   = "postgres"
)

func main() {
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log := logrus.New()

	db, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname))
	if err != nil {
		log.Fatalf("Connection with db failed: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Connection with db failed: %v", err)
	}

	client, err := mongo.Connect(mainCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	server := httpserver.New(client, db, log)

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
