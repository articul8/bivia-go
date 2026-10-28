package bivia

import "context"

// DecisionsService provides access to the six decision endpoints:
// Boolean, Choice, Create (decide), Score, MultiSelect (set), and Span.
type DecisionsService struct {
	client *Client
}

// Boolean makes a yes/no decision with calibrated probability.
func (s *DecisionsService) Boolean(ctx context.Context, params *BooleanParams) (*CategoricalDecision, error) {
	return doRequest[*CategoricalDecision](s.client, ctx, "POST", "/boolean", params)
}

// Choice picks one option from a list of choices (2-26 items).
func (s *DecisionsService) Choice(ctx context.Context, params *ChoiceParams) (*CategoricalDecision, error) {
	return doRequest[*CategoricalDecision](s.client, ctx, "POST", "/choice", params)
}

// Create auto-detects the best decision type from the question.
// Use Outcomes to constrain the options; otherwise the model infers them.
func (s *DecisionsService) Create(ctx context.Context, params *DecideParams) (*CategoricalDecision, error) {
	return doRequest[*CategoricalDecision](s.client, ctx, "POST", "/decide", params)
}

// Score returns a numeric value on a scale with per-level probability mass.
func (s *DecisionsService) Score(ctx context.Context, params *ScoreParams) (*ScoreDecision, error) {
	return doRequest[*ScoreDecision](s.client, ctx, "POST", "/score", params)
}

// MultiSelect selects zero or more options from a list (independent yes/no per option).
func (s *DecisionsService) MultiSelect(ctx context.Context, params *SetParams) (*SetDecision, error) {
	return doRequest[*SetDecision](s.client, ctx, "POST", "/set", params)
}

// Span locates text spans in the input that answer the question.
// Note: /span always runs in fast mode.
func (s *DecisionsService) Span(ctx context.Context, params *SpanParams) (*SpanDecision, error) {
	return doRequest[*SpanDecision](s.client, ctx, "POST", "/span", params)
}
