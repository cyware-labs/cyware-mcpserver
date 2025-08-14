package ctix

import (
	"context"
	"fmt"
	"strings"

	"github.com/cyware-labs/cyware-mcpserver/applications/ctix/helpers"
	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	relation_type_list = "ingestion/threat-data/relationship-types/"
	threat_data_list   = "ingestion/threat-data/list/"
)

type ThreatDataDetailsResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ThreatDataListResp struct {
	Next     string                  `json:"next"`
	PageSize int                     `json:"page_size"`
	Previous any                     `json:"previous"`
	Results  []ThreatDataDetailsResp `json:"results"`
	Total    int                     `json:"total"`
}

func CQLCTIXSearchGrammarTool(s *server.MCPServer) {
	content := helpers.CQL_grammar_rule
	cqlCtixSearchGrammarTool := mcp.NewTool("cql-ctix-grammar-rules",
		mcp.WithDescription(
			"This tool will return the complete CQL grammar details which is used to generate CQL queries. Always understand the grammar before making the CQL query."),
	)
	s.AddTool(cqlCtixSearchGrammarTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText(content), nil
	})
}

func GetCQLQuerySearchResult(sort string, query string, page string, page_size string) (*common.APIResponse, error) {
	query = strings.ReplaceAll(query, "\"", "\\\"")
	payload := strings.NewReader(fmt.Sprintf(`{"query": "%s"}`, query))

	params := map[string]string{
		"sort":      sort,
		"page":      page,
		"page_size": page_size,
	}

	threat_data_list_resp := ThreatDataListResp{}

	resp, err := CTIX_CLIENT.MakeRequest("POST", threat_data_list, params, &threat_data_list_resp, payload, nil)

	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(threat_data_list_resp),
		RawResponse:     resp,
	}, err
}

func GetCQLQuerySearchResultTool(s *server.MCPServer) {
	getCQLQuerySearchResultTool := mcp.NewTool("get-cql-query-search-result",
		mcp.WithDescription("This tool will give result for the CQL"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("This is the query which is used as CQL to make a search and return the result"),
		),
		mcp.WithString("page",
			mcp.Required(),
			mcp.Description("This is the page number for the paginated query. Used to get the result of specific page number"),
		),
		mcp.WithString("page_size",
			mcp.Required(),
			mcp.Description("This is the page size number of result per page. Used to get the specified number of result per page. Please note here if you are making paginated call then keep the page_size same in all the pages otherwise you will get duplicate entries in two different pages."),
		),
		mcp.WithString("sort",
			mcp.Required(),
			mcp.Description(`This is 'sort' params used to get the result in either descending/ascending order based on the value. Supported values are: confidence_score, ctix_modified, ctix_created only. Pass the value prefixed with '-' for descending order or as it for ascending order.
			If nothing is specified then pass "-ctix_modified"`),
		),
	)

	s.AddTool(getCQLQuerySearchResultTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := request.Params.Arguments["query"].(string)
		page := request.Params.Arguments["page"].(string)
		page_size := request.Params.Arguments["page_size"].(string)
		sort := request.Params.Arguments["sort"].(string)

		resp, err := GetCQLQuerySearchResult(sort, query, page, page_size)

		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetAvailableRelationTypeListing(params map[string]string) (*common.APIResponse, error) {
	relation_listing_resp := RelationTypeListingResponse{}

	resp, err := CTIX_CLIENT.MakeRequest("GET", relation_type_list, params, &relation_listing_resp, nil, nil)

	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(relation_listing_resp),
		RawResponse:     resp,
	}, err
}

func GetAvailableRelationTypeListingTool(s *server.MCPServer) {
	getThreatDataObjectRelations := mcp.NewTool("get-available-relation-type",
		mcp.WithDescription("This tool will give all the available relation type. eg indicator is related-to malware"),
	)

	s.AddTool(getThreatDataObjectRelations, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := map[string]string{
			"page":      "1",
			"page_size": "100",
			"nominal":   "true",
			"sort":      "name",
		}

		resp, err := GetAvailableRelationTypeListing(params)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}
