package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alexrehtide/sebastian/pkg/client"
)

func main() {
	c := client.NewClient()

	if err := c.Connect("localhost:3333"); err != nil {
		panic(err)
	}
	defer c.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Authenticate(ctx, client.AuthenticateRequest{Email: "admin@admin.ru", Password: "test1234"})
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
