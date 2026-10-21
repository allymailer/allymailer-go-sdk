package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Platform is the SDK and its version in the "platform" query parameter of every request.
const Platform = "go-sdk.1.1.2"

// APIClient is the low level transport: every API method is a POST request with a JSON body that returns JSON.
type APIClient struct {
	config *Configuration
}

func (c *APIClient) platform() string {
	if c.config.Platform != "" {
		return c.config.Platform + "," + Platform
	}
	return Platform
}

// NewAPIClient creates the transport.
func NewAPIClient(config *Configuration) *APIClient {
	return &APIClient{config: config}
}

// Call sends request (nil for an empty body) to the method path, e.g. "/email/send.json",
// and decodes the response into response (may be nil).
func (c *APIClient) Call(ctx context.Context, path string, request interface{}, response interface{}) error {
	body := []byte("{}")
	if request != nil {
		encoded, err := json.Marshal(request)
		if err != nil {
			return &APIError{Message: "can not encode the request: " + err.Error(), Err: err}
		}
		body = encoded
	}

	endpoint := strings.TrimRight(c.config.BaseURL, "/") + "/" + strings.TrimLeft(path, "/") +
		"?" + url.Values{"platform": {c.platform()}}.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return &APIError{Message: "can not create the request: " + err.Error(), Err: err}
	}
	for name, value := range c.config.Headers {
		req.Header.Set(name, value)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.config.UserAgent)
	req.Header.Set("X-API-KEY", c.config.APIKey)

	httpClient := c.config.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return &APIError{Message: "HTTP request failed: " + err.Error(), Err: err}
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIError{Message: "can not read the response: " + err.Error(), HTTPStatus: resp.StatusCode, Err: err}
	}

	var envelope struct {
		Status  string          `json:"status"`
		Message string          `json:"message"`
		Code    json.RawMessage `json:"code"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return &APIError{
			Message:    fmt.Sprintf("invalid JSON in response: %s: %s", err, quote(raw)),
			HTTPStatus: resp.StatusCode,
			Body:       raw,
			Err:        err,
		}
	}
	if resp.StatusCode >= 400 || envelope.Status == "error" {
		apiErr := &APIError{Message: envelope.Message, HTTPStatus: resp.StatusCode, Body: raw}
		if apiErr.Message == "" {
			apiErr.Message = "API error"
		}
		var code int
		if json.Unmarshal(envelope.Code, &code) == nil {
			apiErr.Code = &code
		}
		return apiErr
	}

	if response != nil {
		if err := json.Unmarshal(raw, response); err != nil {
			return &APIError{
				Message:    "can not decode the response: " + err.Error(),
				HTTPStatus: resp.StatusCode,
				Body:       raw,
				Err:        err,
			}
		}
	}
	return nil
}

// quote cuts a response body for an error message.
func quote(raw []byte) string {
	const maxQuote = 500
	if len(raw) <= maxQuote {
		return string(raw)
	}
	return fmt.Sprintf("%s... (%d bytes)", raw[:maxQuote], len(raw))
}
