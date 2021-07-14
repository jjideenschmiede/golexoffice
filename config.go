//**********************************************************
//
// This file is part of lexoffice.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package lexoffice

import (
	"bytes"
	"net/http"
)

const (
	resourceUrl = "https://api.lexoffice.io"
)

// Request is to define the request data
type Request struct {
	Path, Method, Token string
	Body                []byte
}

// Send is to send a new request
func (r *Request) Send() (*http.Response, error) {

	// Set url
	url := resourceUrl + r.Path

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(r.Method, url, bytes.NewBuffer(r.Body))
	if err != nil {
		return nil, err
	}

	// Define header
	request.Header.Set("Authorization", "Bearer "+r.Token)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
