package common

import (
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"resty.dev/v3"
)

type APIResponse struct {
	RawResponse     *resty.Response
	FilteredReponse any
}

func JsonifyResponse(obj any) any {
	resp, _ := json.Marshal(obj)
	return string(resp)
}

func MCPToolResponse(resp *APIResponse, expected_status_code []int, err error) (*mcp.CallToolResult, error) {
	if err != nil && (resp == nil || resp.RawResponse == nil) {
		return mcp.NewToolResultText(fmt.Sprintf("An error occurred: %v", err)), err
	}
	if err != nil || (resp.RawResponse != nil && !ContainsStatusCode(expected_status_code, resp.RawResponse.StatusCode())) {
		// resp.RawResponse is nil when the request failed before any response
		// was received (e.g. timeout / connection reset). Dereferencing it here
		// would panic the tool handler, so report the transport error instead.
		if resp.RawResponse == nil {
			return mcp.NewToolResultText(fmt.Sprintf("An error occurred before a response was received: %v", err)), err
		}
		return mcp.NewToolResultText(fmt.Sprintf("An error occurred, Server responded with status code %v and response %v", resp.RawResponse.StatusCode(), resp.RawResponse.String())), err
	}
	return mcp.NewToolResultText(fmt.Sprintf("%v", resp.FilteredReponse)), nil
}

func ContainsStatusCode(codes []int, statusCode int) bool {
	for _, code := range codes {
		if code == statusCode {
			return true
		}
	}
	return false
}
