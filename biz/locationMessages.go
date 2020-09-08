package main

type (
	// LocationResponse defines the returning data after a locations-data-request
	LocationResponse struct {
		Successful bool   `json:"successful"`
		Name       string `json:"name"`
	}
)
