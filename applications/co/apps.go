package co

import (
	"context"
	"fmt"
	"log"

	"github.com/cyware-labs/cyware-mcpserver/applications/co/helpers"
	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	list_apps_endpoint   = "integrations/apps/"
	app_actions_endpoint = "integrations/app-action/"
	list_app_instance    = "integrations/app-instance/"
	execute_action       = "/soarapi/integrations/sync-exec-action/"
)

type AppsListingResponse struct {
	Next     string `json:"next"`
	PageSize int    `json:"page_size"`
	Previous string `json:"previous"`
	Results  []struct {
		ActionsCount    int    `json:"actions_count"`
		Apphash         string `json:"apphash"`
		Appid           string `json:"appid"`
		Created         string `json:"created"`
		Createdby       string `json:"createdby"`
		Description     string `json:"description"`
		Docurl          string `json:"docurl"`
		Islatestversion bool   `json:"islatestversion"`
		Meta            struct {
			Dummy int `json:"_dummy"`
		} `json:"meta"`
		Modifiedby         string `json:"modifiedby"`
		Packagehash        string `json:"packagehash"`
		Published          string `json:"published"`
		Publishedby        string `json:"publishedby"`
		Publisherlogo      string `json:"publisherlogo"`
		Status             int    `json:"status"`
		Supportedversion   string `json:"supportedversion"`
		Title              string `json:"title"`
		TotalInstanceCount int    `json:"total_instance_count"`
		Version            string `json:"version"`
	} `json:"results"`
	Total int `json:"total"`
}

type AppActionsResponse struct {
	Next     string `json:"next"`
	PageSize int    `json:"page_size"`
	Previous string `json:"previous"`
	Results  []struct {
		Actionid string `json:"actionid"`
		// Apphash      string `json:"apphash"`
		// Appid        string `json:"appid"`
		// Created      string `json:"created"`
		// Createdby    string `json:"createdby"`
		Description string `json:"description"`
		ID          string `json:"id"`
		// IsActionSupp bool   `json:"is_action_supp"`
		// MinPltfVer   string `json:"min_pltf_ver"`
		// Modified     string `json:"modified"`
		// Modifiedby   string `json:"modifiedby"`
		Title string `json:"title"`
	} `json:"results"`
	Total int `json:"total"`
}

func GetCOAppsListing(params map[string]string) (*common.APIResponse, error) {
	app_listing_resp := AppsListingResponse{}
	resp, err := CO_CLIENT.MakeRequest("GET", GetSoarEndpoint(list_apps_endpoint), params, &app_listing_resp, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(app_listing_resp),
		RawResponse:     resp,
	}, err
}

func GetCOAppsListingTool(s *server.MCPServer) {

	getCOAppsListingTool := mcp.NewTool("get-co-apps-list",
		mcp.WithDescription("This tool will give the list of all the apps in CO(Cyware Orchestrate). Always use params with lower limit if its not required. Also use 'q' params whenever specific app details are required."),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query params information with value as strings. Query params which can be send
			1. "page":  This is the page number for the paginated query. Used to get the result of specific page number
			2. "page_size" : This is the page size number of result per page. Used to get the specified number of result per page. Please note here if you are making paginated call then keep the page_size same in all the pages otherwise you will get duplicate entries in two different pages.
			3. "configured" : This represents whether the app is configured or not. Pass 1 for configured and 0 for not configured.
			4. "q" : This represent if there is any specific app name value to be searched. Note ->❗❗❗ This must be used if there is a app name specified to reduce the search space.`),
			// mcp.AdditionalProperties(false),
		),
	)
	s.AddTool(getCOAppsListingTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mp := request.Params.Arguments["params"].(map[string]interface{})

		params_list := []string{"page", "page_size", "configured", "q"}
		params := map[string]string{}

		for _, v := range params_list {
			if _, ok := mp[v]; ok {
				params[v] = mp[v].(string)
			}
		}
		resp, err := GetCOAppsListing(params)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetCOAppDetails(apphash string) (*common.APIResponse, error) {
	endpoint := GetSoarEndpoint(fmt.Sprintf("%v%v/", list_apps_endpoint, apphash))
	resp, err := CO_CLIENT.MakeRequest("GET", endpoint, nil, nil, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(resp.String()),
		RawResponse:     resp,
	}, err
}

func GetCOAppDetailsTool(s *server.MCPServer) {
	getCOAppDetailsTool := mcp.NewTool("get-co-app-details",
		mcp.WithDescription("This tool provides detailed information of an app in the CO(Cyware orchestrate)."),
		mcp.WithString(
			"apphash",
			mcp.Description("The unique apphash of app requested. This is required to fetch the details of the app."),
		),
	)
	s.AddTool(getCOAppDetailsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		apphash := request.Params.Arguments["apphash"].(string)
		resp, err := GetCOAppDetails(apphash)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetCOAppActionsListing(params map[string]string) (*common.APIResponse, error) {
	app_action_resp := AppActionsResponse{}
	resp, err := CO_CLIENT.MakeRequest("GET", GetSoarEndpoint(app_actions_endpoint), params, &app_action_resp, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(app_action_resp),
		RawResponse:     resp,
	}, err
}

func COAppActionsListingTool(s *server.MCPServer) {

	getCOAppActionsListing := mcp.NewTool("get-co-actions-of-app",
		mcp.WithDescription("This tool will give the list of all the actions of the specified app. Always use params with lower limit if its not required. Also use 'q' params whenever specific actions details are required."),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query params information with value as strings. Query params which can be send
			1. "page":  This is the page number for the paginated query. Used to get the result of specific page number
			2. "page_size" : This is the page size number of result per page. Used to get the specified number of result per page. Please note here if you are making paginated call then keep the page_size same in all the pages otherwise you will get duplicate entries in two different pages.
			3. "app_unique_id" : The unique apphash of app requested. This is required to fetch the actions of the app..
			4. "q" : This represent if there is any specific action name to be searched. Note ->❗❗❗ This must be used if there is any action name specified to reduce the search space.`),
			// mcp.AdditionalProperties(false),
		),
	)
	s.AddTool(getCOAppActionsListing, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mp := request.Params.Arguments["params"].(map[string]interface{})

		params_list := []string{"page", "page_size", "app_unique_id", "q"}
		params := map[string]string{}
		for _, v := range params_list {
			if _, ok := mp[v]; ok {
				params[v] = mp[v].(string)
			}
		}

		resp, err := GetCOAppActionsListing(params)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetCOAppActionDetails(id string) (*common.APIResponse, error) {
	endpoint := GetSoarEndpoint(fmt.Sprintf("%v%v/", app_actions_endpoint, id))
	resp, err := CO_CLIENT.MakeRequest("GET", endpoint, nil, nil, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(resp.String()),
		RawResponse:     resp,
	}, err
}

func GetCOAppActionDetailsTool(s *server.MCPServer) {
	getCOAppActionDetailsTool := mcp.NewTool("get-co-app-action-details",
		mcp.WithDescription("This tool provides detailed information of a specific action of an app in the CO(Cyware orchestrate)."),
		mcp.WithString(
			"id",
			mcp.Description("The id of the specific action, used to fetch the details of the action. Details also consist of input structure for specific action."),
		),
	)
	s.AddTool(getCOAppActionDetailsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id := request.Params.Arguments["id"].(string)
		resp, err := GetCOAppActionDetails(id)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetConfiguredInstancesOfCOApp(apphash string) (*common.APIResponse, error) {
	params := map[string]string{"app_unique_id": apphash}
	resp, err := CO_CLIENT.MakeRequest("GET", GetSoarEndpoint(list_app_instance), params, nil, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(resp.String()),
		RawResponse:     resp,
	}, err
}

func GetConfiguredInstancesOfCOAppTool(s *server.MCPServer) {
	getConfiguredInstancesOfCOAppTool := mcp.NewTool("get-instances-of-co-app",
		mcp.WithDescription("This tool provides the list of all the instances(account) configured in the specific app which can be used to execute the action."),
		mcp.WithString(
			"apphash",
			mcp.Description("The unique apphash of app. This is required to fetch the instances(accounts) configured in the specific app."),
		),
	)
	s.AddTool(getConfiguredInstancesOfCOAppTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		apphash := request.Params.Arguments["apphash"].(string)
		resp, err := GetConfiguredInstancesOfCOApp(apphash)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func ExecuteActionOfCOApp(payload any) (*common.APIResponse, error) {
	resp, err := CO_CLIENT.MakeRequest("POST", execute_action, nil, nil, payload, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(resp.String()),
		RawResponse:     resp,
	}, err
}

func ExecuteActionOfCOAppTool(s *server.MCPServer) {
	schema := helpers.Execute_actions_of_app_schema
	executePlaybookTool := mcp.NewToolWithRawSchema("execute-action-of-co-app",
		`This tool executes the action of the specified app. Some action requires input which can be generated using the details of the action.`,
		[]byte(schema),
	)

	s.AddTool(executePlaybookTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mp := request.Params.Arguments

		// calling to make the full payload
		logged_in_user_details := GetLoggedInUserDetails()
		mp["tenantid"] = logged_in_user_details.TenantID
		mp["workspaceid"] = logged_in_user_details.PreferredWorkspaceID
		mp["sku"] = 1

		log.Println(mp)
		resp, err := ExecuteActionOfCOApp(mp)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}
