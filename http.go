package bivia

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// doRequest sends an HTTP request and decodes the JSON response into T.
func doRequest[T any](c *Client, ctx context.Context, method, path string, body any) (T, error) {
	var zero T
	var lastErr error
	attempts := 1 + c.maxRetries

	for attempt := 0; attempt < attempts; attempt++ {
		result, retry, err := doOnce[T](c, ctx, method, path, body)
		if err == nil {
			return result, nil
		}
		lastErr = err
		if !retry || attempt >= attempts-1 {
			return zero, err
		}

		// Check for Retry-After header in APIError
		delay := backoff(attempt)
		if apiErr, ok := err.(*APIError); ok && (apiErr.Status == 429 || apiErr.Status == 503) {
			if ra, ok2 := apiErr.Body.(string); ok2 {
				if secs, parseErr := strconv.ParseFloat(ra, 64); parseErr == nil {
					delay = time.Duration(secs * float64(time.Second))
				}
			}
		}

		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return zero, ctx.Err()
		case <-timer.C:
		}
	}

	return zero, lastErr
}

func doOnce[T any](c *Client, ctx context.Context, method, path string, body any) (T, bool, error) {
	var zero T

	// Build request
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return zero, false, fmt.Errorf("bivia: failed to marshal request: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return zero, false, fmt.Errorf("bivia: failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Timeout and connection errors are retryable
		return zero, true, fmt.Errorf("bivia: request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, false, fmt.Errorf("bivia: failed to read response: %w", err)
	}

	// Success
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		var result T
		if err := json.Unmarshal(respBody, &result); err != nil {
			return zero, false, fmt.Errorf("bivia: failed to decode response: %w", err)
		}
		return result, false, nil
	}

	// Build API error
	apiErr := &APIError{
		Status:  resp.StatusCode,
		Message: fmt.Sprintf("HTTP %d", resp.StatusCode),
	}

	var parsed any
	if json.Unmarshal(respBody, &parsed) == nil {
		apiErr.Body = parsed
		// Try to extract a message from the body
		if m, ok := parsed.(map[string]any); ok {
			if detail, ok := m["detail"]; ok {
				switch v := detail.(type) {
				case string:
					apiErr.Message = v
				case []any:
					// FastAPI-style validation errors
					msgs := make([]string, 0, len(v))
					for _, item := range v {
						if entry, ok := item.(map[string]any); ok {
							msg, _ := entry["msg"].(string)
							msgs = append(msgs, msg)
						}
					}
					if len(msgs) > 0 {
						apiErr.Message = fmt.Sprintf("%v", msgs)
					}
				}
			} else if errMsg, ok := m["error"]; ok {
				if s, ok := errMsg.(string); ok {
					apiErr.Message = s
				}
			}
		}
	}

	// Retry on 429, 503, and 5xx
	retryable := resp.StatusCode == 429 || resp.StatusCode == 503 || resp.StatusCode >= 500
	if retryable {
		if ra := resp.Header.Get("Retry-After"); ra != "" {
			apiErr.Body = ra // Carry Retry-After for the retry loop
		}
	}

	return zero, retryable, apiErr
}

// backoff returns a jittered exponential backoff duration.
func backoff(attempt int) time.Duration {
	base := 0.5 * math.Pow(2, float64(attempt))
	if base > 8.0 {
		base = 8.0
	}
	// Add jitter: 75% to 100% of base
	jitter := base * (0.75 + 0.25*rand.Float64())
	return time.Duration(jitter * float64(time.Second))
}
