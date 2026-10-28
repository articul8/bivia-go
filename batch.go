package bivia

import "context"

// BatchService provides access to the batch (systemone) endpoint.
type BatchService struct {
	client *Client
}

// Create sends multiple typed questions against the same context in a single request.
// Each question is independently typed (boolean, choice, score, set, or box) and
// independently answered.
func (s *BatchService) Create(ctx context.Context, params *BatchParams) (*BatchResponse, error) {
	return doRequest[*BatchResponse](s.client, ctx, "POST", "/v1/systemone", params)
}
