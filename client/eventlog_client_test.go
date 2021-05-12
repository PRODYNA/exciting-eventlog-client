package client

import (
	"net/http"
	"testing"
)

type Mock struct {
}

func (m Mock) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{
		Status:     "200",
		StatusCode: 200,
	}, nil
}

func TestLog(t *testing.T) {

	l := eventLogger{
		Url: "http://localhost:8080/eventlog",
		Source: "test",
		client: http.Client{
			Transport: Mock{},
		},
	}

	l.Error("Error")



}
