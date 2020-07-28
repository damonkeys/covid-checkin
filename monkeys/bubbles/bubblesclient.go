package bubblesclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// Categories describe a raw business based grouping for a message
var Categories = [...]string{"CASHFLOW", "STANDARD", "MERCHANT"}

// Audiences describe who should see the message supposedly (who sees the message in reality is decided by the reader of the message)
var Audiences = [...]string{"USERS", "NONE"}

// SenderIDs is an identifier describing the originator ot the message in a non technical way
var SenderIDs = [...]string{"PAYMENTS", "PROFILE"}

// SaveMessageRequest is used to be send to bubbles and stored as a message
type SaveMessageRequest struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	UserID   string `json:"userID"`
	Category string `json:"category"`
	SenderID string `json:"senderID"`
	Audience string `json:"audience"`
}

// SaveMessage saves a message to bubbles via http(s). If the message wasn't saved it returns an error
func SaveMessage(saveMessageRequest *SaveMessageRequest) error {

	err := paramIsNotEmpty(saveMessageRequest.Title)
	if err != nil {
		return err
	}
	err = paramIsNotEmpty(saveMessageRequest.Text)
	if err != nil {
		return err
	}
	err = paramIsNotEmpty(saveMessageRequest.UserID)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(saveMessageRequest)
	if err != nil {
		return err
	}
	_, err = sendToBubbles(bytes)

	return err
}

func sendToBubbles(jsonString []byte) (*http.Response, error) {
	//TODO resolve url from config
	const url = "https://checkin.chckr.de/b/msg"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil

}

func paramIsNotEmpty(param string) error {
	const explanation = "A param is not there"
	if len(param) == 0 {
		return errors.New(explanation)
	}
	if param == "" {
		return errors.New(explanation)
	}
	return nil
}
