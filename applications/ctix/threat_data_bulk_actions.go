package ctix

import (
	"context"

	"github.com/cyware-labs/cyware-mcpserver/applications/ctix/helpers"
	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type BulkActionResponse struct {
	Message string `json:"message"`
}

func ThreatDataListBulkAction(endpoint string, payload any) (*common.APIResponse, error) {
	bulk_resp := BulkActionResponse{}
	resp, err := CTIX_CLIENT.MakeRequest("POST", endpoint, nil, &bulk_resp, payload, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(bulk_resp),
		RawResponse:     resp,
	}, err
}

// This function uses an action map and registers tools for all the bulk actions of threat data
func ThreatDataListBulkActionTools(s *server.MCPServer) {
	mp := helpers.GetThreatDataBulkActionsMapping()
	for _, v := range mp {
		tool := mcp.NewToolWithRawSchema(v["tool_name"], v["tool_description"], []byte(v["schema"]))

		s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := ThreatDataListBulkAction(v["endpoint"], request.Params.Arguments)
			return common.MCPToolResponse(resp, []int{200}, err)
		})
	}
}
