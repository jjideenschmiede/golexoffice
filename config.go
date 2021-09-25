//**********************************************************
//
// This file is part of lexoffice.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package golexoffice

import (
	"io"
	"net/http"
)

const (
	baseURL = "https://api.lexoffice.io"
)

// Config is to define the request data
type Config struct {
	Path, Method, Token string
	ContentType         string
	Body                io.Reader
}

// Send is to send a new request
func (c Config) Send() (*http.Response, error) {

	// Set url
	url := baseURL + c.Path

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(c.Method, url, c.Body)
	if err != nil {
		return nil, err
	}

	// Define header
	request.Header.Set("Authorization", "Bearer "+c.Token)
	request.Header.Set("Content-Type", c.ContentType)
	request.Header.Set("Accept", "application/json")

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
