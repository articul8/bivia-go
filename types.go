// Package bivia provides a Go client for the Bivia calibrated decision API.
package bivia

// ── Input references ──────────────────────────────────────────────

// ImageRef is a reference to an image, either by URL or base64-encoded bytes.
type ImageRef struct {
	URL  string `json:"url,omitempty"`
	B64  string `json:"b64,omitempty"`
	MIME string `json:"mime,omitempty"`
}

// AudioRef is a reference to an audio clip, either by URL or base64-encoded bytes.
type AudioRef struct {
	URL  string `json:"url,omitempty"`
	B64  string `json:"b64,omitempty"`
	MIME string `json:"mime,omitempty"`
}

// DocumentRef is a reference to a document for text extraction.
type DocumentRef struct {
	URL      string `json:"url,omitempty"`
	B64      string `json:"b64,omitempty"`
	Filename string `json:"filename,omitempty"`
	MIME     string `json:"mime,omitempty"`
}

// VideoRef is a reference to a video clip.
type VideoRef struct {
	URL     string `json:"url,omitempty"`
	B64     string `json:"b64,omitempty"`
	MIME    string `json:"mime,omitempty"`
	MaxEdge *int   `json:"max_edge,omitempty"`
}

// SignalRef is a raw numeric signal for data-science decisions.
type SignalRef struct {
	Name     string        `json:"name,omitempty"`
	Data     [][]float64   `json:"data"`
	Channels *int          `json:"channels,omitempty"`
	Fs       *float64      `json:"fs,omitempty"`
	Dtype    string        `json:"dtype,omitempty"`
}

// ── Decision request params ───────────────────────────────────────

// BooleanParams are the parameters for a yes/no decision.
type BooleanParams struct {
	Question     string        `json:"question"`
	State        string        `json:"state,omitempty"`
	Images       []ImageRef    `json:"images,omitempty"`
	Audio        []AudioRef    `json:"audio,omitempty"`
	Documents    []DocumentRef `json:"documents,omitempty"`
	Domain       *string       `json:"domain,omitempty"`
	Mode         *string       `json:"mode,omitempty"`
	Think        *bool         `json:"think,omitempty"`
	AllowAbstain *bool         `json:"allow_abstain,omitempty"`
}

// ChoiceParams are the parameters for a pick-one decision.
type ChoiceParams struct {
	Question     string        `json:"question"`
	Choices      []any         `json:"choices"`
	State        string        `json:"state,omitempty"`
	Images       []ImageRef    `json:"images,omitempty"`
	Audio        []AudioRef    `json:"audio,omitempty"`
	Documents    []DocumentRef `json:"documents,omitempty"`
	Domain       *string       `json:"domain,omitempty"`
	Mode         *string       `json:"mode,omitempty"`
	Think        *bool         `json:"think,omitempty"`
	AllowAbstain *bool         `json:"allow_abstain,omitempty"`
}

// DecideParams are the parameters for an auto-detect decision.
type DecideParams struct {
	Question     string        `json:"question"`
	State        string        `json:"state,omitempty"`
	Outcomes     []any         `json:"outcomes,omitempty"`
	Images       []ImageRef    `json:"images,omitempty"`
	Audio        []AudioRef    `json:"audio,omitempty"`
	Documents    []DocumentRef `json:"documents,omitempty"`
	Domain       *string       `json:"domain,omitempty"`
	Mode         *string       `json:"mode,omitempty"`
	Think        *bool         `json:"think,omitempty"`
	AllowAbstain *bool         `json:"allow_abstain,omitempty"`
}

// ScoreParams are the parameters for a numeric score decision.
type ScoreParams struct {
	Question     string        `json:"question"`
	State        string        `json:"state,omitempty"`
	Range        []float64     `json:"range,omitempty"`
	Buckets      *int          `json:"buckets,omitempty"`
	Images       []ImageRef    `json:"images,omitempty"`
	Audio        []AudioRef    `json:"audio,omitempty"`
	Documents    []DocumentRef `json:"documents,omitempty"`
	Domain       *string       `json:"domain,omitempty"`
	Mode         *string       `json:"mode,omitempty"`
	Think        *bool         `json:"think,omitempty"`
	AllowAbstain *bool         `json:"allow_abstain,omitempty"`
}

// SetParams are the parameters for a multi-select decision.
type SetParams struct {
	Question     string        `json:"question"`
	Options      []any         `json:"options"`
	State        string        `json:"state,omitempty"`
	Images       []ImageRef    `json:"images,omitempty"`
	Audio        []AudioRef    `json:"audio,omitempty"`
	Documents    []DocumentRef `json:"documents,omitempty"`
	Domain       *string       `json:"domain,omitempty"`
	Mode         *string       `json:"mode,omitempty"`
	Think        *bool         `json:"think,omitempty"`
	AllowAbstain *bool         `json:"allow_abstain,omitempty"`
}

// SpanParams are the parameters for a text span extraction.
// Note: /span always runs in fast mode. Sending mode="deep" or think=true is rejected (422).
type SpanParams struct {
	Question  string        `json:"question"`
	State     string        `json:"state,omitempty"`
	Images    []ImageRef    `json:"images,omitempty"`
	Documents []DocumentRef `json:"documents,omitempty"`
	Domain    *string       `json:"domain,omitempty"`
}

// GroundImageParams are the parameters for image grounding.
type GroundImageParams struct {
	Question  string        `json:"question,omitempty"`
	Images    []ImageRef    `json:"images,omitempty"`
	Documents []DocumentRef `json:"documents,omitempty"`
	State     string        `json:"state,omitempty"`
	Domain    *string       `json:"domain,omitempty"`
}

// GroundVideoParams are the parameters for video grounding.
type GroundVideoParams struct {
	Video          VideoRef `json:"video"`
	Query          string   `json:"query"`
	Kind           *string  `json:"kind,omitempty"`
	CoarseStep     *float64 `json:"coarse_step,omitempty"`
	FineResolution *float64 `json:"fine_resolution,omitempty"`
	BoxesOnHits    *bool    `json:"boxes_on_hits,omitempty"`
	MaxProbes      *int     `json:"max_probes,omitempty"`
	AllowAbstain   *bool    `json:"allow_abstain,omitempty"`
}

// BatchQuestion defines a single question within a batch request.
type BatchQuestion struct {
	Type         string `json:"type,omitempty"`
	Instructions string `json:"instructions,omitempty"`
	Question     string `json:"question,omitempty"`
	Criteria     any    `json:"criteria,omitempty"`
	Options      []any  `json:"options,omitempty"`
	Levels       []any  `json:"levels,omitempty"`
	Buckets      *int   `json:"buckets,omitempty"`
}

// BatchParams are the parameters for a batch (systemone) request.
type BatchParams struct {
	Model     string                   `json:"model,omitempty"`
	State     string                   `json:"state,omitempty"`
	Images    []ImageRef               `json:"images,omitempty"`
	Audio     []AudioRef               `json:"audio,omitempty"`
	Documents []DocumentRef            `json:"documents,omitempty"`
	Signals   []SignalRef              `json:"signals,omitempty"`
	Questions map[string]BatchQuestion `json:"questions,omitempty"`
	Mode      *string                  `json:"mode,omitempty"`
}

// DataDecisionParams are the parameters for a data-science decision.
type DataDecisionParams struct {
	Question     string        `json:"question"`
	Signals      []SignalRef   `json:"signals,omitempty"`
	State        string        `json:"state,omitempty"`
	Images       []ImageRef    `json:"images,omitempty"`
	Documents    []DocumentRef `json:"documents,omitempty"`
	Options      []any         `json:"options,omitempty"`
	Kind         *string       `json:"kind,omitempty"`
	Domain       *string       `json:"domain,omitempty"`
	Mode         *string       `json:"mode,omitempty"`
	Think        *bool         `json:"think,omitempty"`
	AllowAbstain *bool         `json:"allow_abstain,omitempty"`
}

// ExtractParams are the parameters for exact field extraction from documents.
type ExtractParams struct {
	Question  string        `json:"question,omitempty"`
	Fields    []string      `json:"fields,omitempty"`
	State     string        `json:"state,omitempty"`
	Documents []DocumentRef `json:"documents,omitempty"`
}

// ── Response types ────────────────────────────────────────────────

// EvidenceSpan is a character range in the input text that supports the decision.
type EvidenceSpan struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Text  string `json:"text"`
}

// Evidence contains evidence supporting the decision.
type Evidence struct {
	Spans []EvidenceSpan `json:"spans"`
}

// Verification is the result of the deep-mode verification step.
type Verification struct {
	Performed bool   `json:"performed"`
	Result    string `json:"result"`
}

// DocumentNote is per-document extraction status.
type DocumentNote struct {
	Document  string `json:"document"`
	Status    string `json:"status"`
	Chars     *int   `json:"chars,omitempty"`
	Tokens    *int   `json:"tokens,omitempty"`
	Truncated *bool  `json:"truncated,omitempty"`
	Reason    *string `json:"reason,omitempty"`
}

// CategoricalDecision is the response for boolean, choice, and decide endpoints.
type CategoricalDecision struct {
	Type           string             `json:"type"`
	Value          any                `json:"value"`
	Probability    *float64           `json:"probability,omitempty"`
	Distribution   map[string]float64 `json:"distribution,omitempty"`
	Status         string             `json:"status"`
	AbstainReason  *string            `json:"abstain_reason,omitempty"`
	Mode           string             `json:"mode"`
	ThinkAllowed   *bool              `json:"think_allowed,omitempty"`
	Thought        bool               `json:"thought"`
	Verification   Verification       `json:"verification"`
	Evidence       Evidence           `json:"evidence"`
	LatencyMs      float64            `json:"latency_ms"`
	Model          string             `json:"model"`
	InputTokens    *int               `json:"input_tokens,omitempty"`
	ImageTokens    *int               `json:"image_tokens,omitempty"`
	DocumentTokens *int               `json:"document_tokens,omitempty"`
	BillableTokens *int               `json:"billable_tokens,omitempty"`
	DocumentNotes  []DocumentNote     `json:"document_notes,omitempty"`
}

// BooleanDecision is the response for the boolean endpoint (type aliases for clarity).
type BooleanDecision = CategoricalDecision

// ChoiceDecision is the response for the choice endpoint.
type ChoiceDecision = CategoricalDecision

// DecideDecision is the response for the decide (auto-detect) endpoint.
type DecideDecision = CategoricalDecision

// ScoreDecision is the response for the score endpoint.
type ScoreDecision struct {
	Type           string             `json:"type"`
	Value          *float64           `json:"value,omitempty"`
	Levels         map[string]float64 `json:"levels,omitempty"`
	Status         string             `json:"status"`
	AbstainReason  *string            `json:"abstain_reason,omitempty"`
	Mode           string             `json:"mode"`
	ThinkAllowed   *bool              `json:"think_allowed,omitempty"`
	Thought        bool               `json:"thought"`
	Verification   Verification       `json:"verification"`
	Evidence       Evidence           `json:"evidence"`
	LatencyMs      float64            `json:"latency_ms"`
	Model          string             `json:"model"`
	InputTokens    *int               `json:"input_tokens,omitempty"`
	ImageTokens    *int               `json:"image_tokens,omitempty"`
	DocumentTokens *int               `json:"document_tokens,omitempty"`
	BillableTokens *int               `json:"billable_tokens,omitempty"`
	DocumentNotes  []DocumentNote     `json:"document_notes,omitempty"`
}

// SetDecision is the response for the set (multi-select) endpoint.
type SetDecision struct {
	Type           string             `json:"type"`
	Value          []any              `json:"value,omitempty"`
	Probabilities  map[string]float64 `json:"probabilities,omitempty"`
	Status         string             `json:"status"`
	AbstainReason  *string            `json:"abstain_reason,omitempty"`
	Mode           string             `json:"mode"`
	ThinkAllowed   *bool              `json:"think_allowed,omitempty"`
	Thought        bool               `json:"thought"`
	Verification   Verification       `json:"verification"`
	Evidence       Evidence           `json:"evidence"`
	LatencyMs      float64            `json:"latency_ms"`
	Model          string             `json:"model"`
	InputTokens    *int               `json:"input_tokens,omitempty"`
	ImageTokens    *int               `json:"image_tokens,omitempty"`
	DocumentTokens *int               `json:"document_tokens,omitempty"`
	BillableTokens *int               `json:"billable_tokens,omitempty"`
	DocumentNotes  []DocumentNote     `json:"document_notes,omitempty"`
}

// SpanDecision is the response for the span endpoint.
type SpanDecision struct {
	Type           string         `json:"type"`
	Value          []EvidenceSpan `json:"value,omitempty"`
	Status         string         `json:"status"`
	AbstainReason  *string        `json:"abstain_reason,omitempty"`
	Mode           string         `json:"mode"`
	ThinkAllowed   *bool          `json:"think_allowed,omitempty"`
	Thought        bool           `json:"thought"`
	Verification   Verification   `json:"verification"`
	Evidence       Evidence       `json:"evidence"`
	LatencyMs      float64        `json:"latency_ms"`
	Model          string         `json:"model"`
	InputTokens    *int           `json:"input_tokens,omitempty"`
	ImageTokens    *int           `json:"image_tokens,omitempty"`
	DocumentTokens *int           `json:"document_tokens,omitempty"`
	BillableTokens *int           `json:"billable_tokens,omitempty"`
	DocumentNotes  []DocumentNote `json:"document_notes,omitempty"`
}

// GroundRegion is a detected region in an image.
type GroundRegion struct {
	Label string    `json:"label"`
	Box   []float64 `json:"box"`
	Score *float64  `json:"score,omitempty"`
}

// GroundDecision is the response for the ground (image grounding) endpoint.
// Unlike other decision types, this does not include the full decision envelope.
type GroundDecision struct {
	Type           string         `json:"type"`
	Value          []GroundRegion `json:"value"`
	Count          int            `json:"count"`
	Found          bool           `json:"found"`
	Grounded       bool           `json:"grounded"`
	InputTokens    *int           `json:"input_tokens,omitempty"`
	ImageTokens    *int           `json:"image_tokens,omitempty"`
	BillableTokens *int           `json:"billable_tokens,omitempty"`
}

// ExtractValue is a single extracted field value with its source offset.
type ExtractValue struct {
	Value      *string  `json:"value"`
	Start      *int     `json:"start,omitempty"`
	End        *int     `json:"end,omitempty"`
	Confidence *float64 `json:"confidence,omitempty"`
	Found      bool     `json:"found"`
}

// ExtractDecision is the response for the extract endpoint.
type ExtractDecision struct {
	Type     string                  `json:"type"`
	Values   map[string]ExtractValue `json:"values,omitempty"`
	Grounded bool                   `json:"grounded"`
}

// VideoHit is a temporal hit in a video scan.
type VideoHit struct {
	Start      *float64       `json:"start,omitempty"`
	End        *float64       `json:"end,omitempty"`
	PeakT      *float64       `json:"peak_t,omitempty"`
	OnsetT     *float64       `json:"onset_t,omitempty"`
	Confidence *float64       `json:"confidence,omitempty"`
	Boxes      []GroundRegion `json:"boxes,omitempty"`
}

// VideoGroundDecision is the response for the video grounding endpoint.
type VideoGroundDecision struct {
	Type           string     `json:"type"`
	Kind           string     `json:"kind"`
	DurationS      float64    `json:"duration_s"`
	Scanned        *int       `json:"scanned,omitempty"`
	Truncated      *bool      `json:"truncated,omitempty"`
	Hits           []VideoHit `json:"hits"`
	InputTokens    *int       `json:"input_tokens,omitempty"`
	ImageTokens    *int       `json:"image_tokens,omitempty"`
	BillableTokens *int       `json:"billable_tokens,omitempty"`
}

// ── Batch response types ──────────────────────────────────────────

// BatchChoiceAnswer is a choice answer within a batch response.
type BatchChoiceAnswer struct {
	Type          string             `json:"type"`
	Choice        any                `json:"choice,omitempty"`
	Confidence    *float64           `json:"confidence,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
	Abstained     bool               `json:"abstained"`
}

// BatchBooleanAnswer is a boolean answer within a batch response.
type BatchBooleanAnswer struct {
	Type        string   `json:"type"`
	Noul        *float64 `json:"noul,omitempty"`
	Probability *float64 `json:"probability,omitempty"`
	Confidence  *float64 `json:"confidence,omitempty"`
	Abstained   bool     `json:"abstained"`
}

// BatchScoreAnswer is a score answer within a batch response.
type BatchScoreAnswer struct {
	Type          string             `json:"type"`
	Score         *float64           `json:"score,omitempty"`
	Confidence    *float64           `json:"confidence,omitempty"`
	Legend        map[string]string  `json:"legend,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
	Distribution  map[string]float64 `json:"distribution,omitempty"`
	Abstained     bool               `json:"abstained"`
}

// BatchSetAnswer is a set (multi-select) answer within a batch response.
type BatchSetAnswer struct {
	Type          string             `json:"type"`
	Selected      []any              `json:"selected,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
	Abstained     bool               `json:"abstained"`
}

// BatchBoxAnswer is a bounding-box answer within a batch response.
type BatchBoxAnswer struct {
	Type       string             `json:"type"`
	Box        map[string]float64 `json:"box,omitempty"`
	Label      *string            `json:"label,omitempty"`
	Confidence *float64           `json:"confidence,omitempty"`
	Abstained  bool               `json:"abstained"`
}

// Usage contains token usage for a batch request.
type Usage struct {
	Questions     int  `json:"questions"`
	InputTokens   int  `json:"input_tokens"`
	OutputTokens  *int `json:"output_tokens,omitempty"`
	BillableTokens int `json:"billable_tokens"`
}

// BatchResponse is the response for the batch (systemone) endpoint.
// Answers is a map of question name to raw JSON; use json.Unmarshal to
// decode individual answers into the appropriate BatchXxxAnswer type.
type BatchResponse struct {
	Model   string         `json:"model"`
	Answers map[string]any `json:"answers"`
	Usage   Usage          `json:"usage"`
}

// ── Model discovery ───────────────────────────────────────────────

// ModelInfo describes a single available model.
type ModelInfo struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	ReleaseDate *string `json:"release_date,omitempty"`
}

// ModelList is the response for the list models endpoint.
type ModelList struct {
	Models []ModelInfo `json:"models"`
}

// ModalityInfo contains feature flags for a modality.
type ModalityInfo struct {
	Available          *bool `json:"available,omitempty"`
	GroundingAvailable *bool `json:"grounding_available,omitempty"`
}

// Capabilities is the response for the capabilities endpoint.
type Capabilities struct {
	Model         *string        `json:"model,omitempty"`
	DecisionTypes []string       `json:"decision_types,omitempty"`
	Modalities    map[string]any `json:"modalities,omitempty"`
	Evidence      map[string]any `json:"evidence,omitempty"`
	Abstention    *bool          `json:"abstention,omitempty"`
}
