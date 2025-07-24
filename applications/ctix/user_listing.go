package ctix

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	user_listing_endpoint       = "rest-auth/users/"
	user_group_listing_endpoint = "rest-auth/groups/"
)

type UserListingResponse struct {
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	PageSize int         `json:"page_size"`
	Total    int         `json:"total"`
	Results  []struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		// ContactNumber     string `json:"contact_number"`
		// CountryCode       string `json:"country_code"`
		IsReadOnly bool   `json:"is_read_only"`
		UserID     string `json:"user_id"`
		Username   string `json:"username"`
		// InviteStatus      string `json:"invite_status"`
		// ImageURL          string `json:"image_url"`
		IsActive  bool `json:"is_active"`
		IsBlocked bool `json:"is_blocked"`
		// LastActiveSession int    `json:"last_active_session"`
		// Created           int    `json:"created"`
		// InvitedBy         struct {
		// 	ID            string `json:"id"`
		// 	FirstName     string `json:"first_name"`
		// 	LastName      string `json:"last_name"`
		// 	Email         string `json:"email"`
		// 	IsActive      bool   `json:"is_active"`
		// 	ContactNumber string `json:"contact_number"`
		// 	CountryCode   string `json:"country_code"`
		// 	IsReadOnly    bool   `json:"is_read_only"`
		// } `json:"invited_by"`
		UserGroups []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"user_groups"`
	} `json:"results"`
	ActiveCount int `json:"active_count"`
}

type UserGroupListingResponse struct {
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	PageSize int         `json:"page_size"`
	Total    int         `json:"total"`
	Results  []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		// CreatedBy   struct {
		// 	ID            string `json:"id"`
		// 	FirstName     string `json:"first_name"`
		// 	LastName      string `json:"last_name"`
		// 	Email         string `json:"email"`
		// 	IsActive      bool   `json:"is_active"`
		// 	ContactNumber string `json:"contact_number"`
		// 	CountryCode   string `json:"country_code"`
		// 	IsReadOnly    bool   `json:"is_read_only"`
		// } `json:"created_by"`
		IsEditable bool `json:"is_editable"`
		IsActive   bool `json:"is_active"`
		// Created              int      `json:"created"`
		PermissionCount int  `json:"permission_count"`
		UserCount       int  `json:"user_count"`
		IsReadOnly      bool `json:"is_read_only"`
		IsDefault       bool `json:"is_default"`
		// SamlAssociatedGroups []string `json:"saml_associated_groups"`
	} `json:"results"`
	PermissionGranted int `json:"permission_granted"`
}

func GetCTIXUserListing(params map[string]string) UserListingResponse {

	resp := UserListingResponse{}
	CTIX_CLIENT.MakeRequest("GET", user_listing_endpoint, params, &resp, nil, nil)

	return resp
}

func GetCTIXUserListingTool(s *server.MCPServer) {

	getCTIXUserListingTool := mcp.NewTool("get-ctix-user-list",
		mcp.WithDescription("This tool will give the list of all the users added in CTIX"),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query params information. Query params which can be send
			1. "page":  This is the page number for the paginated query. Used to get the result of specific page number
			2. "page_size" : This is the page size number of result per page. Used to get the specified number of result per page. Please note here if you are making paginated call then keep the page_size same in all the pages otherwise you will get duplicate entries in two different pages.
			3. "q" :  This represents a specific string value which must be used to reduce the search space. Always use this if you have a specific value.
			4. "is_active" : This represent if the user is active or not. It can have only two values either "true" or "false".
			5. "is_blocked" : This represents if the user is blocked or not.  It can have only two values either "true" or "false".
			6. "is_read_only": This represents if the user is only read only. It can have only two values either "true" or "false".
			`),
			mcp.AdditionalProperties(false),
		),
	)
	s.AddTool(getCTIXUserListingTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mp := request.Params.Arguments["params"].(map[string]interface{})

		params_list := []string{"page", "page_size", "q", "is_active", "is_blocked", "is_read_only"}
		params := map[string]string{}

		for _, v := range params_list {
			if _, ok := mp[v]; ok {
				params[v] = mp[v].(string)
			}
		}
		resp := GetCTIXUserListing(params)
		result, _ := json.Marshal(resp)

		return mcp.NewToolResultText(fmt.Sprintf("Successfully got the list of users in CTIX %v", string(result))), nil
	})
}

func GetCTIXUserGroupList(params map[string]string) UserGroupListingResponse {
	resp := UserGroupListingResponse{}
	CTIX_CLIENT.MakeRequest("GET", user_group_listing_endpoint, params, &resp, nil, nil)
	return resp
}

func GetCTIXUserGroupListTool(s *server.MCPServer) {

	getCTIXUserGroupListTool := mcp.NewTool("get-ctix-user-group-list",
		mcp.WithDescription("This tool will give the list of all the users groups added in CTIX"),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query params information. Query params which can be send
			1. "page":  This is the page number for the paginated query. Used to get the result of specific page number
			2. "page_size" : This is the page size number of result per page. Used to get the specified number of result per page. Please note here if you are making paginated call then keep the page_size same in all the pages otherwise you will get duplicate entries in two different pages.
			3. "q" : This represents a string value, which is supposed to be searched.
			`),
			mcp.AdditionalProperties(false),
		),
	)
	s.AddTool(getCTIXUserGroupListTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mp := request.Params.Arguments["params"].(map[string]interface{})

		params_list := []string{"page", "page_size", "q"}
		params := map[string]string{}

		for _, v := range params_list {
			if _, ok := mp[v]; ok {
				params[v] = mp[v].(string)
			}
		}
		resp := GetCTIXUserGroupList(params)
		result, _ := json.Marshal(resp)

		return mcp.NewToolResultText(fmt.Sprintf("Successfully got the list of users group in CTIX %v", string(result))), nil
	})
}
