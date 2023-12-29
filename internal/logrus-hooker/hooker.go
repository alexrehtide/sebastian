package logrushooker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
	"github.com/sirupsen/logrus"
)

type SessionProvider interface {
	Inject(ctx context.Context) *model.Session
}

type LogStorage interface {
	Create(ctx context.Context, ops model.CreateLogOptions) (uint, error)
}

func New(logStorage LogStorage, sessionProvider SessionProvider) *Hooker {
	return &Hooker{LogStorage: logStorage, SessionProvider: sessionProvider}
}

type Hooker struct {
	LogStorage      LogStorage
	SessionProvider SessionProvider
}

func (h *Hooker) Fire(entry *logrus.Entry) error {
	var createLogOptions model.CreateLogOptions

	createLogOptions.Level = entry.Level.String()
	createLogOptions.CreatedAt = entry.Time
	createLogOptions.Message = entry.Message

	if len(entry.Data) != 0 {
		data := make(map[string]interface{})
		for k, v := range entry.Data {
			if err, isError := v.(error); logrus.ErrorKey == k && v != nil && isError {
				data[k] = err.Error()
			} else {
				data[k] = v
			}
		}
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("failed marshal json data: %w", err)
		}
		createLogOptions.Data = string(dataBytes)
	}

	if entry.Context != nil {
		session := h.SessionProvider.Inject(entry.Context)
		if session != nil {
			createLogOptions.SessionID = session.ID
			createLogOptions.AccountID = session.AccountID
			entry.Data["currentSessionId"] = session.ID
			entry.Data["currentAccountId"] = session.AccountID
		}
	}

	_, err := h.LogStorage.Create(context.Background(), createLogOptions)
	if err != nil {
		return fmt.Errorf("failed to send log entry to db: %w", err)
	}

	return nil
}

func (h *Hooker) Levels() []logrus.Level {
	return logrus.AllLevels
}
