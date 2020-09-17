package qr

import (
	"testing"
)

//TestHandler stores the result in an []byte for further checks during test
type TestHandler struct {
	result *[]byte
}

func (h *TestHandler) Handle(content []byte) error {
	h.result = &content
	return nil
}

func TestRetrieve(t *testing.T) {
	c := NewPixiCall()
	testHandler := &TestHandler{}
	c.Handler = testHandler

	c.Retrieve("http://nzz.ch", "", "png")

	if len(*testHandler.result) == 0 {
		t.Fatal("Result should not have a length of 0")
	}

}
