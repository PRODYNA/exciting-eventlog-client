package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type logEntry struct {
	Source  string `json:"source"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type EventLogger interface {
	Error(message string)
	Info(message string)
	Warn(message string)
}

type eventLogger struct {
	Url    string
	Source string
	client http.Client
}

type nopLogger struct {

}

func (n nopLogger) Error(msg string) {}
func (n nopLogger) Info(msg string) {}
func (n nopLogger) Warn(msg string) {}

func NewNopLogger() EventLogger {
	return &nopLogger{}
}

func NewEventLogger(url, src string) EventLogger {
	return &eventLogger{
		Url:    url,
		Source: src,
		client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (e eventLogger) Error(msg string) {
	err := e.writeEventLog(&logEntry{
		Source:  e.Source,
		Type:    "Error",
		Message: msg,
	})
	if err != nil {
		fmt.Println("Error while logging Event", err.Error())
	}
}

func (e eventLogger) Info(msg string) {
	err := e.writeEventLog(&logEntry{
		Source:  e.Source,
		Type:    "Info",
		Message: msg,
	})
	if err != nil {
		fmt.Println("Error while logging Event", err.Error())
	}
}

func (e eventLogger) Warn(msg string) {
	err := e.writeEventLog(&logEntry{
		Source:  e.Source,
		Type:    "Warn",
		Message: msg,
	})
	if err != nil {
		fmt.Println("Error while logging Event", err.Error())
	}
}

func (e eventLogger) writeEventLog(log *logEntry) error {

	b, err := json.Marshal(log)

	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", "http://" + e.Url, bytes.NewReader(b))

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
