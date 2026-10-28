# bivia-go

[![CI](https://github.com/articul8/bivia-go/actions/workflows/ci.yml/badge.svg)](https://github.com/articul8/bivia-go/actions/workflows/ci.yml)

Go client for [Bivia](https://articul8.ai/bivia) by [Articul8 AI](https://articul8.ai). See [bivia](https://github.com/articul8/bivia) for full documentation and examples.

Zero external dependencies — stdlib only.

## Install

```bash
go get github.com/articul8/bivia-go
```

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	bivia "github.com/articul8/bivia-go"
)

func main() {
	// Reads $A8_API_KEY from the environment by default.
	client := bivia.NewClient()

	ctx := context.Background()

	// Boolean decision
	decision, err := client.Decisions.Boolean(ctx, &bivia.BooleanParams{
		Question: "Is this urgent?",
		State:    "I was double-charged and support never replied.",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Answer: %v (p=%.2f)\n", decision.Value, *decision.Probability)
}
```

## Usage

### Decisions

```go
// Choice — pick one from a list
choice, _ := client.Decisions.Choice(ctx, &bivia.ChoiceParams{
	Question: "Which team handles this?",
	Choices:  []any{"billing", "refunds", "technical"},
	State:    "I need a copy of my invoice.",
})

// Score — numeric value on a scale
score, _ := client.Decisions.Score(ctx, &bivia.ScoreParams{
	Question: "How urgent is this?",
	State:    "Customer threatened to cancel today.",
	Range:    []float64{0, 1},
	Buckets:  bivia.Int(11),
})

// Auto-detect — let the model choose the decision type
decide, _ := client.Decisions.Create(ctx, &bivia.DecideParams{
	Question: "Most likely diagnosis?",
	State:    "Patient has fever, cough, and right lower lobe infiltrate.",
	Outcomes: []any{"common cold", "pneumonia", "healthy"},
})

// Multi-select — zero or more options
set, _ := client.Decisions.MultiSelect(ctx, &bivia.SetParams{
	Question: "Which actions apply?",
	Options:  []any{"cancel", "refund", "status"},
	State:    "Order arrived damaged and the customer wants money back.",
})

// Span — locate text in the input
span, _ := client.Decisions.Span(ctx, &bivia.SpanParams{
	Question: "Where is the auto-renewal clause?",
	State:    "This agreement renews automatically for one year unless either party cancels.",
})
```

### Grounding

```go
// Image grounding — locate objects in images
ground, _ := client.Grounding.Image(ctx, &bivia.GroundImageParams{
	Question: "the invoice total",
	Images:   []bivia.ImageRef{{URL: "https://example.com/invoice.png"}},
})

// Video grounding — locate objects or events in video
video, _ := client.Grounding.Video(ctx, &bivia.GroundVideoParams{
	Video: bivia.VideoRef{URL: "https://example.com/clip.mp4"},
	Query: "the red car",
	Kind:  bivia.String("object"),
})
```

### Batch

```go
// Multiple typed questions in one request
batch, _ := client.Batch.Create(ctx, &bivia.BatchParams{
	State: "I was double-charged and support never replied.",
	Questions: map[string]bivia.BatchQuestion{
		"team": {
			Type:         "choice",
			Instructions: "Which team?",
			Criteria: map[string]string{
				"billing": "Billing disputes and invoices",
				"tech":    "Product defects",
			},
		},
		"urgent": {
			Type:         "boolean",
			Instructions: "Is this urgent?",
		},
	},
})
```

### Data science

```go
// Decision from raw numeric signals
ds, _ := client.Data.Decide(ctx, &bivia.DataDecisionParams{
	Question: "Which analysis method fits this signal?",
	Kind:     bivia.String("choice"),
	Options:  []any{"spectral analysis (FFT)", "linear regression"},
	Signals: []bivia.SignalRef{{
		Name: "temperature",
		Data: [][]float64{{72.1, 73.5, 89.2, 91.0, 88.5, 72.0}},
		Fs:   bivia.Float64(1),
	}},
})
```

### Model discovery

```go
models, _ := client.Models.List(ctx)
caps, _ := client.Models.Capabilities(ctx)
```

## Configuration

```go
client := bivia.NewClient(
	bivia.WithAPIKey("your-api-key"),
	bivia.WithBaseURL("https://custom-endpoint.example.com/bivia"),
	bivia.WithTimeout(30 * time.Second),
	bivia.WithMaxRetries(3),
)
```

### Helper functions

Use pointer helpers for optional fields:

```go
bivia.String("deep")    // *string
bivia.Int(5)            // *int
bivia.Float64(0.5)      // *float64
bivia.Bool(true)        // *bool
```

## Error handling

API errors are returned as `*bivia.APIError`:

```go
decision, err := client.Decisions.Boolean(ctx, params)
if err != nil {
	var apiErr *bivia.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("Status: %d, Message: %s\n", apiErr.Status, apiErr.Message)
	}
}
```

## License

See [LICENSE](LICENSE).
