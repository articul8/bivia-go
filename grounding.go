package bivia

import "context"

// GroundingService provides access to image and video grounding endpoints.
type GroundingService struct {
	client *Client
}

// Image performs open-vocabulary visual grounding, locating objects
// matching the question in the provided images. Returns bounding boxes
// with coordinates normalized to [0, 1000].
func (s *GroundingService) Image(ctx context.Context, params *GroundImageParams) (*GroundDecision, error) {
	return doRequest[*GroundDecision](s.client, ctx, "POST", "/ground", params)
}

// Video scans a video clip for objects or events matching the query.
// Use Kind "object" for temporal intervals with bounding boxes, or
// "event" for onset timestamps.
func (s *GroundingService) Video(ctx context.Context, params *GroundVideoParams) (*VideoGroundDecision, error) {
	return doRequest[*VideoGroundDecision](s.client, ctx, "POST", "/ground/video", params)
}
