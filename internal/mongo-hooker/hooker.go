package mongohooker

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionProvider interface {
	Inject(ctx context.Context) *model.Session
}

func New(collection *mongo.Collection, sessionProvider SessionProvider) *Hooker {
	return &Hooker{c: collection, SessionProvider: sessionProvider}
}

type Hooker struct {
	c               *mongo.Collection
	SessionProvider SessionProvider
}

func (h *Hooker) Fire(entry *logrus.Entry) error {
	data := make(logrus.Fields)
	data["level"] = entry.Level.String()
	data["time"] = entry.Time
	data["msg"] = entry.Message

	if entry.Context != nil {
		session := h.SessionProvider.Inject(entry.Context)
		if session != nil {
			entry.Data["currentSessionId"] = session.ID
			entry.Data["currentAccountId"] = session.AccountID
		}
	}

	for k, v := range entry.Data {
		if err, isError := v.(error); logrus.ErrorKey == k && v != nil && isError {
			data[k] = err.Error()
		} else {
			data[k] = v
		}
	}

	_, err := h.c.InsertOne(context.Background(), bson.M(data))

	if err != nil {
		return fmt.Errorf("failed to send log entry to mongodb: %w", err)
	}

	return nil
}

func (h *Hooker) Levels() []logrus.Level {
	return logrus.AllLevels
}
