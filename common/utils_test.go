package common

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

// JSON numbers decode into interface{} as float64, so numeric params such as
// {"page": 1} previously panicked ExtractParams via an unchecked .(string).
func TestExtractParamsCoercesNonStringValues(t *testing.T) {
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{
		"params": map[string]interface{}{
			"page":      float64(1),
			"page_size": float64(50),
			"direction": "all",
			"active":    true,
		},
	}

	got := ExtractParams(req, []string{"page", "page_size", "direction", "active"})
	want := map[string]string{
		"page":      "1",
		"page_size": "50",
		"direction": "all",
		"active":    "true",
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("param %q = %q, want %q", k, got[k], v)
		}
	}
}

func TestExtractParamsMissingParamsKey(t *testing.T) {
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{}
	if got := ExtractParams(req, []string{"page"}); len(got) != 0 {
		t.Errorf("expected empty map, got %v", got)
	}
}

func TestParseTimeoutDurationAcceptsNumericRepresentations(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want time.Duration
	}{
		{name: "float64", arg: float64(1.5), want: 1500 * time.Millisecond},
		{name: "float32", arg: float32(2), want: 2 * time.Second},
		{name: "int", arg: int(3), want: 3 * time.Second},
		{name: "int64", arg: int64(4), want: 4 * time.Second},
		{name: "uint64", arg: uint64(5), want: 5 * time.Second},
		{name: "json number", arg: json.Number("6.25"), want: 6250 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTimeoutDuration(tt.arg)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestParseTimeoutDurationRejectsInvalidValues(t *testing.T) {
	tests := []struct {
		name string
		arg  any
	}{
		{name: "string", arg: "1"},
		{name: "bad json number", arg: json.Number("bad")},
		{name: "zero", arg: 0},
		{name: "negative", arg: -1},
		{name: "nan", arg: math.NaN()},
		{name: "infinity", arg: math.Inf(1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ParseTimeoutDuration(tt.arg); err == nil {
				t.Fatalf("expected error")
			}
		})
	}
}
