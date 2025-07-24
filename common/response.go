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
	if err != nil || (resp.RawResponse != nil && !containsStatusCode(expected_status_code, resp.RawResponse.StatusCode())) {
		return mcp.NewToolResultText(fmt.Sprintf("An error occurred, Server responded with status code %v and response %v", resp.RawResponse.StatusCode(), resp.RawResponse.String())), err
	}
	return mcp.NewToolResultText(fmt.Sprintf("%v", resp.FilteredReponse)), nil
}

func containsStatusCode(codes []int, statusCode int) bool {
	for _, code := range codes {
		if code == statusCode {
			return true
		}
	}
	return false
}
