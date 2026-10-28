package bivia

import "context"

// DataService provides access to the data-science decision endpoint.
type DataService struct {
	client *Client
}

// Decide makes a decision based on raw numeric signals. The server profiles the
// signal data and folds the analysis into context before deciding.
func (s *DataService) Decide(ctx context.Context, params *DataDecisionParams) (*CategoricalDecision, error) {
	return doRequest[*CategoricalDecision](s.client, ctx, "POST", "/ds", params)
}
