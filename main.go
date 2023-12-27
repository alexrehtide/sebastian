package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	httpserver "github.com/alexrehtide/sebastian/internal/http-server"
	"github.com/jmoiron/sqlx"
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

	db, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Connection with db failed: %v", err)
	}

	server := httpserver.New(db)

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		return server.Listen(":3000")
	})
	g.Go(func() error {
		<-gCtx.Done()
		return server.Shutdown()
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
