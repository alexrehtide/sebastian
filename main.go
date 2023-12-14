package main

import "github.com/alexrehtide/sebastian/internal/server"

func main() {
	s := server.NewServer()

	if err := s.Serve(3333); err != nil {
		panic(err)
	}
}
