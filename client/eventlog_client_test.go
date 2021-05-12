package client

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)


type Mock struct {

}

func (m Mock) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{
		Status: "200",
		StatusCode: 200,
	}, nil
}

func TestLog(t *testing.T) {

	l := NewEventLogger("http://anyhost:8080")
	l.client = http.Client{
		Transport: Mock{},
	}
	entry := LogEntry{
		Timestamp: time.Now(),
		Source:    "source",
		Type:      "type",
		Message:   "message",
		Status:    "status",
	}

	err := l.WriteEventLog(entry)

	assert.Nil(t, err)

}
