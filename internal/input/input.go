package input

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// Represent the structure of the JSON input
type websiteInput struct {
	URLs []string `json:"urls"`
}

// Reads the JSON file from the disk
func ReadJSONFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}
	return data, nil
}

// Converts JSON bytes into a websiteInput struct
func ParseJSON(data []byte) (websiteInput, error) {
	var input websiteInput

	if err := json.Unmarshal(data, &input); err != nil {
		return input, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return input, nil
}

// validate an individual URL
func ValidateURL(rawURL string) error {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return fmt.Errorf("invalid URL %q: %w", rawURL, err)
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("unsupported scheme in URL: %s", rawURL)
	}

	if parsedURL.Hostname() == "" {
		return fmt.Errorf("missing hostname in URL: %s", rawURL)
	}

	return nil
}

// Validates all URLs in the input.
func ValidateURLs(urls []string) ([]string, error) {
	if len(urls) == 0 {
		return nil, fmt.Errorf("no URLs provided in input file")
	}

	var validURLs []string
	var validationErrors []string

	for _, rawURL := range urls {
		rawURL = strings.TrimSpace(rawURL)

		if err := ValidateURL(rawURL); err != nil {
			validationErrors = append(validationErrors, err.Error())
			continue
		}

		validURLs = append(validURLs, rawURL)
	}

	if len(validationErrors) > 0 {
		return validURLs, fmt.Errorf(
			"found %d invalid URLs: %s",
			len(validationErrors),
			strings.Join(validationErrors, "; "),
		)
	}

	return validURLs, nil
}

// Main entry point for loading website URLs.
func LoadWebsiteURLs(filePath string) ([]string, error) {
	data, err := ReadJSONFile(filePath)
	if err != nil {
		return nil, err
	}

	input, err := ParseJSON(data)
	if err != nil {
		return nil, err
	}

	urls, err := ValidateURLs(input.URLs)
	if err != nil {
		return urls, err
	}

	return urls, nil
}