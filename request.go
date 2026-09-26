package terminal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) newRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	u := c.baseURL + path
	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, buf)
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if c.secretKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.secretKey)
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
			Raw:        respBody,
		}
		if err := json.Unmarshal(respBody, apiErr); err != nil || apiErr.Message == "" {
			apiErr.Message = string(respBody)
		}
		return apiErr
	}

	if v != nil {
		if err := json.Unmarshal(respBody, v); err != nil {
			return fmt.Errorf("failed to unmarshal response payload: %w", err)
		}
	}

	return nil
}

// buildURLWithQuery formats path with query parameters if present.
func buildURLWithQuery(path string, queryParams map[string]string) string {
	if len(queryParams) == 0 {
		return path
	}
	values := url.Values{}
	for k, v := range queryParams {
		if v != "" {
			values.Add(k, v)
		}
	}
	encoded := values.Encode()
	if encoded == "" {
		return path
	}
	return path + "?" + encoded
}
