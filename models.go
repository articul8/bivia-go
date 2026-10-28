package bivia

import "context"

// ModelsService provides access to model discovery and capability endpoints.
type ModelsService struct {
	client *Client
}

// List returns the models available on this deployment.
func (s *ModelsService) List(ctx context.Context) (*ModelList, error) {
	return doRequest[*ModelList](s.client, ctx, "GET", "/v1/models", nil)
}

// Capabilities returns the model's supported decision types, input modalities,
// evidence modes, and whether abstention is enabled.
func (s *ModelsService) Capabilities(ctx context.Context) (*Capabilities, error) {
	return doRequest[*Capabilities](s.client, ctx, "GET", "/capabilities", nil)
}
