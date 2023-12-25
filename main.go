package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	const baseURL string = "http://localhost:8000"
	fmt.Println("Sample Go API")

	// Perform a simple GET request
	SimpleGetRequest(baseURL)

	// Perform a simple POST request with a JSON payload
	SimplePostRequest(baseURL)

	// Perform a simple POST form request
	SimplePostFormRequest(baseURL)
}

// SimpleGetRequest performs a simple HTTP GET request
func SimpleGetRequest(baseURL string) {
	// Construct the complete URL for the GET request
	getURL := baseURL + "/get"

	// Perform the GET request
	response, err := http.Get(getURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Print the status and content length of the response
	fmt.Println("Status:", response.Status)
	fmt.Println("Content Length:", response.ContentLength)

	// Read the response body and print it
	content, _ := io.ReadAll(response.Body)
	var stringContent strings.Builder
	stringContent.Write(content)
	fmt.Println(stringContent.String())
}

// SimplePostRequest performs a simple HTTP POST request with a JSON payload
func SimplePostRequest(baseURL string) {
	// Construct the complete URL for the POST request
	postURL := baseURL + "/post"

	// Fake JSON payload
	requestBody := strings.NewReader(`
	{
		"name":"aswin",
		"Designation":"Engineer",
		"Tech":"Python"
	}	
	`)

	// Perform the POST request with the JSON payload
	response, err := http.Post(postURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Read the response body and print it
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

// SimplePostFormRequest performs a simple HTTP POST form request
func SimplePostFormRequest(baseURL string) {
	// Construct the complete URL for the POST form request
	postFormURL := baseURL + "/postform"

	// Define form data
	data := url.Values{}
	data.Add("Name", "Aswin")
	data.Add("Tech", "Java")

	// Perform the POST form request with the form data
	response, err := http.PostForm(postFormURL, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Read the response body and print it
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
