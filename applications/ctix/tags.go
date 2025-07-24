package ctix

import (
	"context"

	"github.com/cyware-labs/cyware-mcpserver/applications/ctix/helpers"
	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	tag_listing_endpoint = "ingestion/tags/"
)

type TagCreationResponse struct {
	Result struct {
		Details string `json:"details"`
		Count   int    `json:"count"`
	} `json:"result"`
}

type TagListingResponse struct {
	Next     interface{} `json:"next"`
	PageSize int         `json:"page_size"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		// ColourCode string `json:"colour_code"`
		ID       string `json:"id"`
		IsActive bool   `json:"is_active"`
		Name     string `json:"name"`
		TagType  struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"tag_type"`
		Theme string `json:"theme"`
	} `json:"results"`
	Total int `json:"total"`
}

func GetCTIXTagListing(params map[string]string) (*common.APIResponse, error) {
	tag_listing_resp := TagListingResponse{}
	resp, err := CTIX_CLIENT.MakeRequest("GET", tag_listing_endpoint, params, &tag_listing_resp, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(tag_listing_resp),
		RawResponse:     resp,
	}, err
}

func GetCTIXTagListingTool(s *server.MCPServer) {

	getCTIXTagListing := mcp.NewTool("get-ctix-tags-list",
		mcp.WithDescription("This tool will give the list of all the tags created in CTIX. Always use params with lower limit if its not required. Also use 'q' params whenever specific tags details are required."),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query params information with value as strings. Query params which can be send
			1. "page":  This is the page number for the paginated query. Used to get the result of specific page number
			2. "page_size" : This is the page size number of result per page. Used to get the specified number of result per page. Please note here if you are making paginated call then keep the page_size same in all the pages otherwise you will get duplicate entries in two different pages.
			3. "tag_type" : This represents a the type of tags, which must be any of these "source" (created by source during ingestion), "system" (internal tag), "user" (created by user), "privileged" (tags having restriction and allowed specific user groups). If we don't this params then we will all the tags which are in the platform.
			4. "q" : This represent if there is any specific tag value to be searched. Note ->❗❗❗ This must be used if there is a tag name specified to reduce the search space.`),
			// mcp.AdditionalProperties(false),
		),
	)
	s.AddTool(getCTIXTagListing, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mp := request.Params.Arguments["params"].(map[string]interface{})

		params_list := []string{"page", "page_size", "tag_type", "q"}
		params := map[string]string{}

		for _, v := range params_list {
			if _, ok := mp[v]; ok {
				params[v] = mp[v].(string)
			}
		}
		resp, err := GetCTIXTagListing(params)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func CreateTaginCTIX(payload any) (*common.APIResponse, error) {
	endpoint := "ingestion/tags/bulk-actions/"
	param := map[string]string{
		"component": "tag",
	}
	tag_resp := TagCreationResponse{}
	resp, err := CTIX_CLIENT.MakeRequest("POST", endpoint, param, &tag_resp, payload, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(tag_resp),
		RawResponse:     resp,
	}, err
}

func CreateTaginCTIXTool(s *server.MCPServer) {
	schema := helpers.Create_tag_schema

	createTaginCTIXTool := mcp.NewToolWithRawSchema("create-tag-in-ctix",
		`This tool creates tags with the specified category in CTIX.
		Tags creation can be partially successfull in case we try to add the same tag again.
    `,
		[]byte(schema),
	)

	s.AddTool(createTaginCTIXTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := CreateTaginCTIX(request.Params.Arguments)
		return common.MCPToolResponse(resp, []int{200}, err)
	})

}
