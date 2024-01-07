package eventservice

import "github.com/sirupsen/logrus"

func New(logger *logrus.Logger) *Service {
	return &Service{
		Logger: logger,
	}
}

type Service struct {
	Logger *logrus.Logger
}
