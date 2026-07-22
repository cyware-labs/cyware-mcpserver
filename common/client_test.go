package common

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"resty.dev/v3"
)

func TestMakeRequestWithContextHonorsTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &APIClient{
		Client:   resty.New(),
		BASE_URL: server.URL,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	_, err := client.MakeRequestWithContext(ctx, http.MethodGet, "/", nil, nil, nil, nil)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context deadline exceeded, got %v", err)
	}
}

func TestMCPToolResponseHandlesNilRawResponseError(t *testing.T) {
	expectedErr := context.DeadlineExceeded

	result, err := MCPToolResponse(&APIResponse{}, []int{http.StatusOK}, expectedErr)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected deadline error, got %v", err)
	}

	content, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected text content, got %T", result.Content[0])
	}
	if !strings.Contains(content.Text, expectedErr.Error()) {
		t.Fatalf("expected response text to include error, got %q", content.Text)
	}
}
