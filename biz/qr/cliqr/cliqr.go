package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/damonkeys/covid-checkin/biz/qr"
)

//Flags helps holding the values of the flags
type Flags struct {
	url      string
	logoURL  string
	pngOrsvg string
}

//CliqrFileHandler allows us to reconfigure file management differently from pixi qr directory and this gives us flexibility for this cli tool
type CliqrFileHandler struct {
	// Example: "~/foo.png" or "./qr.svg" or "/tmp/qr.png"
	pathAndFileName string
}

//Handle writes qr code data to pathAndFileName
func (h CliqrFileHandler) Handle(content []byte) error {
	err := ioutil.WriteFile(h.pathAndFileName, content, 0755)
	if err != nil {
		return err
	}
	return nil
}

const exampleURL = "https://example.com"

func main() {
	flags := &Flags{}
	flags.initFlags()

	if flags.url == exampleURL {
		flag.Usage()
		flag.PrintDefaults()
		fmt.Print("\n\nNo valid url given. I create an example qr code, that you can find as file qr.png\n")
	}
	c := flags.createPngOrSVGCall()
	c.Retrieve(flags.url, flags.logoURL, flags.pngOrsvg)
}

func (flags *Flags) initFlags() {
	flag.PrintDefaults()
	flag.StringVar(&flags.url, "u", exampleURL, "Specify url to encode into qr code")
	flag.StringVar(&flags.logoURL, "l", "", "Specify log to apply into the middle of the qr code (can be left empty)")
	flag.StringVar(&flags.pngOrsvg, "t", "svg", "Specify qr image file type: png oder svg. Default to svg")
	flag.Usage = func() {
		fmt.Printf("\n\nUsage of the Program: \n")
		fmt.Printf("./cliqr -u https://google.com -l http://http://placekitten.com/200/200 -t png \n\n")
	}
	flag.Parse()
}

func (flags *Flags) createPngOrSVGCall() *qr.Call {
	c := qr.NewPixiCall()
	if strings.ToLower(flags.pngOrsvg) == "svg" {
		//  overwrite default storage behaviour for pixi
		c.Handler = &CliqrFileHandler{
			pathAndFileName: "./qr.svg",
		}
		return c
	}
	c.Handler = &CliqrFileHandler{
		pathAndFileName: "./qr.png",
	}
	return c
}
