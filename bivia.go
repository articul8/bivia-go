// Package bivia provides a Go client for the Bivia calibrated decision API
// at https://api.articul8.ai/bivia.
//
// Zero external dependencies -- stdlib only (net/http, encoding/json, etc.).
//
// Usage:
//
//	client := bivia.NewClient() // reads A8_API_KEY from env
//	decision, err := client.Decisions.Boolean(ctx, &bivia.BooleanParams{
//	    Question: "Is this urgent?",
//	    State:    "I was double-charged and support never replied.",
//	})
package bivia

import (
	"net/http"
	"os"
	"time"
)

const defaultBaseURL = "https://api.articul8.ai/bivia"

// Client is the Bivia API client. Use NewClient to create one.
type Client struct {
	// Decisions provides access to the six decision endpoints:
	// Boolean, Choice, Create (decide), Score, MultiSelect (set), Span.
	Decisions *DecisionsService

	// Grounding provides access to image and video grounding endpoints.
	Grounding *GroundingService

	// Batch provides access to the batch (systemone) endpoint.
	Batch *BatchService

	// Data provides access to the data-science decision endpoint.
	Data *DataService

	// Extract provides access to the extract (field extraction) endpoint.
	Extract *ExtractService

	// Models provides access to model discovery and capability endpoints.
	Models *ModelsService

	apiKey     string
	baseURL    string
	timeout    time.Duration
	maxRetries int
	httpClient *http.Client
}

// NewClient creates a new Bivia API client.
//
// By default it reads the API key from $A8_API_KEY and connects to
// https://api.articul8.ai/bivia. Use WithAPIKey, WithBaseURL, etc.
// to override.
func NewClient(opts ...Option) *Client {
	c := &Client{
		apiKey:     os.Getenv("A8_API_KEY"),
		baseURL:    defaultBaseURL,
		timeout:    60 * time.Second,
		maxRetries: 2,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.httpClient = &http.Client{
		Timeout: c.timeout,
	}

	c.Decisions = &DecisionsService{client: c}
	c.Grounding = &GroundingService{client: c}
	c.Extract = &ExtractService{client: c}
	c.Batch = &BatchService{client: c}
	c.Data = &DataService{client: c}
	c.Models = &ModelsService{client: c}

	return c
}
