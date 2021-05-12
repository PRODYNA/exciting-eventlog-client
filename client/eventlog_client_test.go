package client

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
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

	l := NewEventLogger("http://localhost:8080/eventlog")
	l.client = http.Client{
		Transport: Mock{},
	}
	entry := NewLogEntry("unittest","log","really passed","201")

	err := l.WriteEventLog(entry)

	assert.Nil(t, err)

}
