package qr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type (

	// Call represents an API call to the qr code api and contains a handler, that deals with the result and the api url to use
	Call struct {
		apiURL  string
		Handler ResultHandler
	}

	//FileHandler has data where to store the qr code file (hint env QR_CODE_FILE_PATH) and
	// under which filename to store the qr code in the path.
	FileHandler struct {
		Filename string
		envPath  string
	}

	// ResultHandler handles content :-) In the concrete case we are using this interface to
	// deal with different qr code api call results: A SVG []byte is a text but a PNG []byte needs proper encoding.
	// In Addtions Implementations can decide wether or not to log the data and wether or not to store it and more.
	ResultHandler interface {
		Handle(content []byte) error
	}
)

// Handle deals with data that has a file as adestination. This means it stores the data
// onto the given envPath with the given Filename.
func (result FileHandler) Handle(content []byte) error {
	path := result.envPath + "/" + result.Filename
	err := ioutil.WriteFile(path, content, 0755)
	if err != nil {
		return err
	}
	return nil
}

//This is not the real api url. To use the real api we need to sign up on rapidapi.com
//Should not be a problem but this url taken from sniffing the web qr code generator
//is easier to use. Change if necessary (ip blocking or sth like that)
const apiURL = "https://qr-generator.qrcode.studio/qr/custom"
const apiJSON = `{
	"data":"%s",
	"config":{
	"body":"circle",
	"logo":"%s",
	"logoMode":	"default"
	},
	"size":1000,
	"download":false,
	"file":"%s"
	} `

// NewPixiCall is a convinience method to create a call (Handler) that stores to the (static) qr directory on the pixi server.
// The FileHandler will have its default file name set to default.png so a user of this call should overwritte Filename in the FileHandler
// Use this method to create qr codes that shall be served as business related images via call.Retrieve(..)
func NewPixiCall() *Call {
	path := os.Getenv("QR_CODE_FILE_PATH")
	return &Call{
		apiURL: apiURL,
		Handler: &FileHandler{
			envPath:  path,
			Filename: "default.png",
		},
	}
}

func configureQr(url string, logoURL string, pngOrSVG string) string {
	return fmt.Sprintf(apiJSON, url, logoURL, strings.ToLower(pngOrSVG))
}

// Retrieve gets the qr from the https://www.qrcode-monkey.com/qr-code-api-with-logo api
// Example: Retrieve("http://google.com", "", "svg")
func (c *Call) Retrieve(url string, logoURL string, pngOrSVG string) error {
	json := configureQr(url, logoURL, pngOrSVG)
	err := c.doAPIRequest(json)
	return err
}

// doAPIRequest makes the request, reads the response and sends it to the resultHandler.handle() method
func (c *Call) doAPIRequest(apiJSON string) error {
	req, err := http.NewRequest(http.MethodPost, c.apiURL, strings.NewReader(apiJSON))
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = c.Handler.Handle(body)
	return err

}
