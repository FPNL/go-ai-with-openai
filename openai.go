package main

import (
	"fmt"
	"io"
	"net/http"
)

type API interface {
	Method() string
	Path() string
	JSON() io.Reader
}

var ApiKey = ""

const Domain = "https://api.openai.com"

// const openAiURL = "https://api.openai.com/v1/chat/completions"
// const openAiURL = "https://api.openai.com/v1/images/generations"

func openai(api API) (responseBody []byte, err error) {
	request, err := http.NewRequest(api.Method(), Domain+api.Path(), api.JSON())
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest() error: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+ApiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll() error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status code: %d", response.StatusCode)
	}

	return
}
