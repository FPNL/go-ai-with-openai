package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Image struct {
	RevisedPrompt string `json:"revised_prompt"`
	Url           string `json:"url"`
}

type CreateImageResponse struct {
	Created int     `json:"created"`
	Data    []Image `json:"data"`
}

var _ API = CreateImage{}

type CreateImage struct {
	// A text description of the desired image(s). The maximum length is 1000 characters for dall-e-2 and 4000 characters for dall-e-3.
	Prompt string `json:"prompt,omitempty"` // Required

	// Defaults to dall-e-2
	// The model to use for image generation.
	Model Model `json:"model,omitempty"` // Optional

	// Defaults to 1
	// The number of images to generate. Must be between 1 and 10. For dall-e-3, only n=1 is supported.
	N *int `json:"n,omitempty"` // Optional

	// Defaults to standard
	// The quality of the image that will be generated. hd creates images with finer details and greater consistency across the image. This param is only supported for dall-e-3.
	Quality string `json:"quality,omitempty"` // Optional

	// Defaults to url
	// The format in which the generated images are returned. Must be one of url or b64_json. URLs are only valid for 60 minutes after the image has been generated.
	ResponseFormat *string `json:"response_format,omitempty"` // Optional

	// Defaults to 1024x1024
	// The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024 for dall-e-2. Must be one of 1024x1024, 1792x1024, or 1024x1792 for dall-e-3 models.
	Size *string `json:"size,omitempty"` // Optional

	// Defaults to vivid
	// The style of the generated images. Must be one of vivid or natural. Vivid causes the model to lean towards generating hyper-real and dramatic images. Natural causes the model to produce more natural, less hyper-real looking images. This param is only supported for dall-e-3.
	Style *string `json:"style,omitempty"` // Optional

	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more.
	User string `json:"user,omitempty"` // Optional
}

func (receiver CreateImage) JSON() io.Reader {
	jsonBody, _ := json.Marshal(receiver)
	return bytes.NewReader(jsonBody)
}

func (receiver CreateImage) Method() string {
	return http.MethodPost
}

func (receiver CreateImage) Path() string {
	return "/v1/images/generations"
}

func (receiver CreateImage) Do() (output CreateImageResponse, err error) {
	responseBody, err := openai(receiver)
	if err != nil {
		return output, fmt.Errorf("openai(%v): %s, error: %w", receiver, responseBody, err)
	}

	if err = json.Unmarshal(responseBody, &output); err != nil {
		return output, fmt.Errorf("json.Unmarshal() error: %w", err)
	}

	return output, err
}
