package main

import (
	"strings"
	"testing"
)

func TestHandlerCreation(t *testing.T) {
	flags := &Flags{
		pngOrsvg: "svg",
	}
	c := flags.createPngOrSVGCall()

	resultHandler := c.Handler
	if resultHandler == nil {
		t.Fatal("Handler does not exist")
	}

}

func TestSVGSwitch(t *testing.T) {
	flags := &Flags{
		pngOrsvg: "svg",
	}
	c := flags.createPngOrSVGCall()
	cfh := c.Handler.(*CliqrFileHandler)
	if !strings.Contains(cfh.pathAndFileName, "qr.svg") {
		t.Fatal("Did not find .svg filename in call created with svg flags.")
	}
}
