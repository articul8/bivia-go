package bivia

import "context"

// ExtractService provides access to the extract endpoint.
type ExtractService struct {
	client *Client
}

// Extract extracts the exact value of a single field from a document.
// Set Question to the field name (e.g., "termination notice period").
func (s *ExtractService) Extract(ctx context.Context, params *ExtractParams) (*ExtractDecision, error) {
	return doRequest[*ExtractDecision](s.client, ctx, "POST", "/extract", params)
}
