package garbageservice

import "context"

type Cleaner interface {
	Clean(ctx context.Context) error
}

func New(cleaners ...Cleaner) *Service {
	return &Service{
		Cleaners: cleaners,
	}
}

type Service struct {
	Cleaners []Cleaner
}
