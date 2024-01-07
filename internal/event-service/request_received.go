package eventservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/sirupsen/logrus"
)

func (s *Service) RequestReceived(ctx context.Context, evt model.RequestReceived) {
	s.Logger.
		WithContext(ctx).
		WithFields(logrus.Fields{
			"ip":   evt.IP,
			"body": evt.Body,
		}).
		Debugf("Request received: %s", evt.Path)
}
