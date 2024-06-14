package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) < 2 {
		log.Fatal("Please provide an API key")
	}

	ApiKey = os.Args[1]

	n := 1
	size := "1024x1024"
	// responseFormat := "url"
	// style := "vivid"

	api := CreateImage{
		Model:  ModelDallE3,
		Prompt: "bird in the sky, and sky is green like grass, sun beyond the sky is sunset",
		N:      &n,
		Size:   &size,
		// ResponseFormat: &responseFormat,
		// Style:          &style,
	}

	output, err := api.Do()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(output.Data[0].Url)
}
