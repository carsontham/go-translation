package httpapi

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

const (
	// HeaderAcceptEncoding ...
	HeaderAcceptEncoding = "Accept-Encoding"
	// HeaderRapidAPIKey ...
	HeaderRapidAPIKey = "X-RapidAPI-Key"
	// HeaderRapidAPIHost ...
	HeaderRapidAPIHost = "X-RapidAPI-Host"
	//HeaderContentType ...
	HeaderContentType = "content-type"
)

// DecorateRequestFunc ...
type DecorateRequestFunc func(req *http.Request) (*http.Request, error)

// ContentType ...
func ContentType(req *http.Request) (*http.Request, error) {
	req.Header.Add(HeaderContentType, "application/json")
	return req, nil
}

// AcceptEncoding ...
func AcceptEncoding(req *http.Request) (*http.Request, error) {
	req.Header.Add(HeaderAcceptEncoding, "application/gzip")
	return req, nil
}

// RapidAPIKey ...
func RapidAPIKey(req *http.Request) (*http.Request, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	key := os.Getenv("RAPID_API_KEY")
	req.Header.Add(HeaderRapidAPIKey, key)
	return req, nil
}

// RapidAPIHost ...
func RapidAPIHost(req *http.Request) (*http.Request, error) {
	req.Header.Add(HeaderRapidAPIHost, "deep-translate1.p.rapidapi.com")
	return req, nil
}

// DecorateRequest ...
func DecorateRequest(fns ...DecorateRequestFunc) DecorateRequestFunc {
	return func(req *http.Request) (*http.Request, error) {
		for _, fn := range fns {
			var err error
			req, err = fn(req)
			if err != nil {
				return nil, err
			}
		}
		return req, nil
	}
}
