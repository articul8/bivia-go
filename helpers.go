package bivia

// String returns a pointer to the given string value.
// Useful for setting optional *string fields on request params.
func String(s string) *string { return &s }

// Int returns a pointer to the given int value.
// Useful for setting optional *int fields on request params.
func Int(i int) *int { return &i }

// Float64 returns a pointer to the given float64 value.
// Useful for setting optional *float64 fields on request params.
func Float64(f float64) *float64 { return &f }

// Bool returns a pointer to the given bool value.
// Useful for setting optional *bool fields on request params.
func Bool(b bool) *bool { return &b }
