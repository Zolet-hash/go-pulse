package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type WebsiteInput struct {
	URLs []string `json:"urls"`
}

func main() {
	urls := GenerateURLs(50)

	if err := SaveURLs(urls, "websites.json"); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Successfully generated %d URLs\n", len(urls))
	fmt.Println("File: websites.json")
}

func GenerateURLs(count int) []string {
	urls := make([]string, 0, count)

	for i := 1; i <= count; i++ {
		url := fmt.Sprintf("https://site-%04d.example.com", i)
		urls = append(urls, url)
	}
	return urls
}

func SaveURLs(urls []string, filename string) error {
	input := WebsiteInput{
		URLs: urls,
	}

	data, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
