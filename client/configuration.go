package client

import (
	"net/http"
	"strings"
	"time"
)

// DefaultBaseURL is the base URL of the API.
const DefaultBaseURL = "https://cp.allymailer.com/ua/transactional/api/v1"

// API hosts, e.g. NewClient("YOUR_API_KEY", HostEU1) of the root package.
const (
	HostCP = "cp.allymailer.com"
)

// BasePath is the path of the API on every host.
const BasePath = "/ua/transactional/api/v1"

// BaseURLFor returns the base URL for an API host: "eu1.unione.io" -> "https://eu1.unione.io" + BasePath;
// a full URL is kept as is.
func BaseURLFor(host string) string {
	if strings.Contains(host, "://") {
		return strings.TrimRight(host, "/")
	}
	return "https://" + strings.Trim(host, "/") + BasePath
}

// Configuration holds the API connection settings.
type Configuration struct {
	// APIKey is the user or project API key, sent in the X-API-KEY header.
	APIKey string

	// BaseURL of the API.
	BaseURL string

	// UserAgent sent with every request.
	UserAgent string

	// HTTPClient used to send the requests. Its Transport (an http.RoundTripper) is the hook for retries,
	// logging, metrics, proxies: wrap http.DefaultTransport.
	HTTPClient *http.Client

	// Headers are extra HTTP headers of every request; the SDK headers (Content-Type, Accept, User-Agent, the
	// API key) take precedence over them.
	Headers map[string]string

	// Platform is the package that uses the SDK and its version, e.g. "nethttp.1.0.0": it goes to the "platform"
	// query parameter before the SDK one (Platform).
	Platform string
}

// NewConfiguration returns the default configuration for the API key.
func NewConfiguration(apiKey string) *Configuration {
	return &Configuration{
		APIKey:     apiKey,
		BaseURL:    DefaultBaseURL,
		UserAgent:  "allymailer-go-sdk/1.1.1",
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}
}
