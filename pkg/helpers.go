package pkg

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// RemoveSubDomain receives a host and remove its subdomain, if exists.
func RemoveSubDomain(fullURL string) (string, error) {

	// add http if fullURL doesn't have it, this is for validation only, won't be used on http requests
	if !strings.HasPrefix(fullURL, "http") {
		fullURL = "https://" + fullURL
	}

	// check if received fullURL is valid before parsing it
	err := IsValidURL(fullURL)
	if err != nil {
		return "", err
	}

	// build URL
	fullPathURL, err := url.ParseRequestURI(fullURL)
	if err != nil {
		return "", err
	}

	splitHost := strings.Split(fullPathURL.Host, ".")

	if len(splitHost) < 2 {
		return "", fmt.Errorf("the fullURL (%s) is invalid", fullURL)
	}

	lastURLPart := splitHost[len(splitHost)-2:]
	hostWithSpace := strings.Join(lastURLPart, " ")
	// set fullURL only without subdomain
	fullPathURL.Host = strings.ReplaceAll(hostWithSpace, " ", ".")

	// build URL without subdomain
	result := fullPathURL.Scheme + "://" + fullPathURL.Host

	// check if new URL is still valid
	err = IsValidURL(result)
	if err != nil {
		return "", err
	}

	return fullPathURL.Host, nil
}

// IsValidURL checks if a URL is valid
func IsValidURL(rawURL string) error {

	if len(rawURL) == 0 {
		return errors.New("rawURL cannot be empty string")
	}

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil || parsedURL == nil {
		return fmt.Errorf("the URL (%s) is invalid, error = %v", rawURL, err)
	}
	return nil
}
