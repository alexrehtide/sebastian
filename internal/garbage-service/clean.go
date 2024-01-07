package garbageservice

import (
	"context"
	"fmt"
)

func (s *Service) Clean(ctx context.Context) error {
	for _, c := range s.Cleaners {
		if err := c.Clean(ctx); err != nil {
			fmt.Print(err) // TODO: implement error handler
		}
	}
	return nil
}
