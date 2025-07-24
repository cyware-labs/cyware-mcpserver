package ctix

import (
	"context"
	"fmt"

	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type ThreadDataDetailsBasicResponse struct {
	Aliases                  any    `json:"aliases"`
	AnalystDescription       any    `json:"analyst_description"`
	AnalystMarkings          []any  `json:"analyst_markings"`
	AnalystScore             any    `json:"analyst_score"`
	AnalystTlp               any    `json:"analyst_tlp"`
	Asn                      any    `json:"asn"`
	AttributeField           any    `json:"attribute_field"`
	AttributeValue           string `json:"attribute_value"`
	BaseType                 string `json:"base_type"`
	ConfidenceScore          int    `json:"confidence_score"`
	ConfidenceType           string `json:"confidence_type"`
	Country                  string `json:"country"`
	Created                  int    `json:"created"`
	CtixCreated              int    `json:"ctix_created"`
	CtixModified             int    `json:"ctix_modified"`
	CtixScore                int    `json:"ctix_score"`
	CtixTlp                  any    `json:"ctix_tlp"`
	DefangAnalystDescription any    `json:"defang_analyst_description"`
	Description              any    `json:"description"`
	FangAnalystDescription   any    `json:"fang_analyst_description"`
	FirstSeen                any    `json:"first_seen"`
	LastSeen                 any    `json:"last_seen"`
	MarkingDefinitions       []any  `json:"marking_definitions"`
	Modified                 int    `json:"modified"`
	Name                     string `json:"name"`
	Pattern                  string `json:"pattern"`
	PatternType              string `json:"pattern_type"`
	PatternVersion           string `json:"pattern_version"`
	Sources                  []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		SourceType string `json:"source_type"`
	} `json:"sources"`
	SubType string `json:"sub_type"`
	Tags    []struct {
		ColourCode string `json:"colour_code"`
		ID         string `json:"id"`
		Name       string `json:"name"`
		TagType    string `json:"tag_type"`
		Theme      string `json:"theme"`
	} `json:"tags"`
	Tld        string   `json:"tld"`
	Tlp        string   `json:"tlp"`
	Type       string   `json:"type"`
	Types      []string `json:"types"`
	ValidFrom  int      `json:"valid_from"`
	ValidUntil any      `json:"valid_until"`
}

type ThreatDataRelationsResponse struct {
	Next     any `json:"next"`
	PageSize int `json:"page_size"`
	Previous any `json:"previous"`
	Results  []struct {
		CreatedByRef           any    `json:"created_by_ref"`
		ID                     string `json:"id"`
		IsForward              bool   `json:"is_forward"`
		IsRedacted             bool   `json:"is_redacted"`
		RelationEnd            any    `json:"relation_end"`
		RelationStart          any    `json:"relation_start"`
		RelationshipConfidence any    `json:"relationship_confidence"`
		RelationshipType       string `json:"relationship_type"`
		Sources                []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			SourceType string `json:"source_type"`
		} `json:"sources"`
		TargetRef struct {
			Created    int    `json:"created"`
			ID         string `json:"id"`
			Modified   int    `json:"modified"`
			Name       string `json:"name"`
			ObjectType string `json:"object_type"`
			SubType    any    `json:"sub_type"`
			Tlp        string `json:"tlp"`
		} `json:"target_ref"`
	} `json:"results"`
	Total int `json:"total"`
}

type RelationTypeListingResponse struct {
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Total    int    `json:"total"`
	Results  []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"results"`
	PageSize int `json:"page_size"`
}

func GetThreatDataObjectDetails(object_id, object_type string) (*common.APIResponse, error) {
	endpoint := fmt.Sprintf("ingestion/threat-data/%s/%s/basic/", object_type, object_id)
	threat_data_details_basic_resp := ThreadDataDetailsBasicResponse{}

	resp, err := CTIX_CLIENT.MakeRequest("GET", endpoint, nil, &threat_data_details_basic_resp, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(threat_data_details_basic_resp),
		RawResponse:     resp,
	}, err
}

func GetThreatDataObjectDetailsTool(s *server.MCPServer) {
	getThreatDataObjectDetailsTool := mcp.NewTool("get-threat-data-object-details",
		mcp.WithDescription(`This tool retrieves all details of a specific Threat Data Object based on the object type and object ID. 
		⚠️ IMPORTANT: It must not perform any enrichment unless explicitly requested.`),
		mcp.WithString("object_type",
			mcp.Required(),
			mcp.Description("This is type of the threat data object which is used to hit the API"),
		),
		mcp.WithString("object_id",
			mcp.Required(),
			mcp.Description("This is id of the threat data object which is used to hit the API to get the details"),
		),
	)

	s.AddTool(getThreatDataObjectDetailsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		object_type := request.Params.Arguments["object_type"].(string)
		object_id := request.Params.Arguments["object_id"].(string)

		resp, err := GetThreatDataObjectDetails(object_id, object_type)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetThreatDataObjectRelations(params map[string]string, object_id, object_type string) (*common.APIResponse, error) {
	endpoint := fmt.Sprintf("ingestion/threat-data/%s/%s/relations/", object_type, object_id)
	threatDataRelation_resp := ThreatDataRelationsResponse{}
	resp, err := CTIX_CLIENT.MakeRequest("GET", endpoint, params, &threatDataRelation_resp, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(threatDataRelation_resp),
		RawResponse:     resp,
	}, err
}

func GetThreatDataObjectRelationsTool(s *server.MCPServer) {
	getThreatDataObjectRelations := mcp.NewTool("get-threat-data-object-relations",
		mcp.WithDescription("This tool will give all the relations of a threat data objects, which mean it will return all the other threat data objects which are related to this object"),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query params information. Query params which can be send
			1. "direction":  with value "all", "forward", "backward" which represents the direction of the relation
			2. "page":  This is the page number for the paginated query. Used to get the result of specific page number
			3. "page_size" : This is the page size number of result per page. Used to get the specified number of result per page. Please note here if you are making paginated call then keep the page_size same in all the pages otherwise you will get duplicate entries in two different pages.
			4. "object_type" : This represent the type of the related objects. It will filter the relations based on the provided object type.
			5. "q" : This represent the specific object value. This must be used if requested for any specific object value.
			`),
		),
		mcp.WithString("object_type",
			mcp.Required(),
			mcp.Description("This is type of the threat data object which is used to hit the API"),
		),
		mcp.WithString("object_id",
			mcp.Required(),
			mcp.Description("This is id of the threat data object which is used to hit the API to get the details"),
		),
	)

	s.AddTool(getThreatDataObjectRelations, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		object_type := request.Params.Arguments["object_type"].(string)
		object_id := request.Params.Arguments["object_id"].(string)

		mp := request.Params.Arguments["params"].(map[string]any)

		params_list := []string{"direction", "page", "page_size"}

		params := map[string]string{}
		for _, v := range params_list {
			if _, ok := mp[v]; ok {
				params[v] = mp[v].(string)
			}
		}
		resp, err := GetThreatDataObjectRelations(params, object_id, object_type)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}
