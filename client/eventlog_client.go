package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
}

type EventLogger struct {
	Url    string
	client http.Client
}

func NewEventLogger(url string) *EventLogger {
	return &EventLogger{
		Url: url,
		client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (e EventLogger) WriteEventLog(log LogEntry) error {

	b, err := json.Marshal(log)

	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", e.Url, bytes.NewReader(b))

	if err != nil {
		return err
	}

	res, err := e.client.Do(r)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	return nil
}
