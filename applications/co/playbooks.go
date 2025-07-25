package co

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cyware-labs/cyware-mcpserver/applications/co/helpers"
	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	list_playbook_endpoint    = "playbooks/"
	execute_playbook_endpoint = "execute/"
)

type ListPlaybookResponse struct {
	Next     string `json:"next"`
	PageSize int    `json:"page_size"`
	Previous string `json:"previous"`
	Results  []struct {
		Actionlist   []string      `json:"actionlist"`
		Applist      []string      `json:"applist"`
		Busers       []interface{} `json:"busers"`
		Created      string        `json:"created"`
		Createdby    string        `json:"createdby"`
		Description  string        `json:"description"`
		Hash         string        `json:"hash"`
		ID           string        `json:"id"`
		IsFollowed   int           `json:"is_followed"`
		Labels       []interface{} `json:"labels"`
		Modified     string        `json:"modified"`
		Modifiedby   string        `json:"modifiedby"`
		Readableid   string        `json:"readableid"`
		ScheduleInfo string        `json:"schedule_info"`
		Status       int           `json:"status"`
		Tags         []interface{} `json:"tags"`
		Title        string        `json:"title"`
		Version      int           `json:"version"`
		Workspaceid  string        `json:"workspaceid"`
	} `json:"results"`
	Total int `json:"total"`
}

type ExecuteAPIResponse struct {
	ResultID string `json:"result_id"`
}

func GetPlayBookList(params map[string]string) (*common.APIResponse, error) {
	playbook_listing_resp := ListPlaybookResponse{}
	resp, err := CO_CLIENT.MakeRequest("GET", GetSoarEndpoint(list_playbook_endpoint), params, &playbook_listing_resp, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(playbook_listing_resp),
		RawResponse:     resp,
	}, err
}

func GetPlayBookListTool(s *server.MCPServer) {

	getPlayBookListTool := mcp.NewTool("get-co-playbooks-list",
		mcp.WithDescription("This tool will give the list of all the playbooks created in CO. Always use params with lower limit if its not required. Also use 'q' params whenever specific playbook details are required."),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query params information with value as strings. Query params which can be send
			1. "page":  This is the page number for the paginated query. Used to get the result of specific page number
			2. "page_size" : This is the page size number of result per page. Used to get the specified number of result per page. Please note here if you are making paginated call then keep the page_size same in all the pages otherwise you will get duplicate entries in two different pages.
			3. "status" : This represents the status of the playbook whether its active or not. Use 10 for active and 11 for inactive.
			4. "q" : This represent if there is any specific playbook value to be searched. Note ->❗❗❗ This must be used if there is a playbook name specified to reduce the search space.`),
		),
	)
	s.AddTool(getPlayBookListTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mp := request.Params.Arguments["params"].(map[string]interface{})

		params_list := []string{"page", "page_size", "status", "q"}
		params := map[string]string{}

		for _, v := range params_list {
			if _, ok := mp[v]; ok {
				params[v] = mp[v].(string)
			}
		}
		resp, err := GetPlayBookList(params)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetPlaybookDetails(playbook_id string) (*common.APIResponse, error) {
	endpoint := fmt.Sprintf("%v%v/", GetSoarEndpoint(list_playbook_endpoint), playbook_id)
	resp, err := CO_CLIENT.MakeRequest("GET", endpoint, nil, nil, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(resp.String()),
		RawResponse:     resp,
	}, err
}

func GetPlaybookDetailsTool(s *server.MCPServer) {
	getPlaybookDetailsTool := mcp.NewTool("get-co-playbook-details",
		mcp.WithDescription(`This tool retrieves the details of a specific playbook using the playbook id.`),
		mcp.WithString("playbook_id",
			mcp.Required(),
			mcp.Description("This is id of the playbook which is used to hit the API to get the details"),
		),
	)

	s.AddTool(getPlaybookDetailsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		playbook_id := request.Params.Arguments["playbook_id"].(string)

		resp, err := GetPlaybookDetails(playbook_id)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func ExecutePlaybook(payload any) (*common.APIResponse, error) {
	hash := payload.(map[string]any)["pbhash"]
	endpoint := GetSoarEndpoint(fmt.Sprintf("playbooks/%v/execute/", hash))
	exec_resp := ExecuteAPIResponse{}
	resp, err := CO_CLIENT.MakeRequest("POST", endpoint, nil, nil, payload, nil)
	json.Unmarshal([]byte(resp.String()), &exec_resp)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(exec_resp),
		RawResponse:     resp,
	}, err
}

func ExecutePlaybookTool(s *server.MCPServer) {
	schema := helpers.Execute_playbook_schema

	executePlaybookTool := mcp.NewToolWithRawSchema("execute-playbook-in-co",
		`This tool executes the specified playbook in CO.
		If the playbook requires input then generate it based on the playbook details.
    `,
		[]byte(schema),
	)

	s.AddTool(executePlaybookTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := ExecutePlaybook(request.Params.Arguments)
		return common.MCPToolResponse(resp, []int{201}, err)
	})

}
