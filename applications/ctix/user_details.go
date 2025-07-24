package ctix

import (
	"context"

	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const user_details_endpoint = "rest-auth/user-details/"

type LoggedInUserDetailsResponse struct {
	ID            string      `json:"id"`
	FirstName     string      `json:"first_name"`
	LastName      string      `json:"last_name"`
	Email         string      `json:"email"`
	IsActive      bool        `json:"is_active"`
	ContactNumber string      `json:"contact_number"`
	CountryCode   string      `json:"country_code"`
	IsReadOnly    bool        `json:"is_read_only"`
	UserID        string      `json:"user_id"`
	Created       int         `json:"created"`
	InvitedBy     interface{} `json:"invited_by"`
	InviteStatus  string      `json:"invite_status"`
	EmailAlerts   bool        `json:"email_alerts"`
	SmsAlerts     bool        `json:"sms_alerts"`
	Groups        []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
	AllowedTlp        []interface{} `json:"allowed_tlp"`
	Username          string        `json:"username"`
	ImageURL          interface{}   `json:"image_url"`
	DateJoined        int           `json:"date_joined"`
	LastActiveSession int           `json:"last_active_session"`
	AdminComponent    []string      `json:"admin_component"`
	Component         []string      `json:"component"`
	Permission        []string      `json:"permission"`
	DefaultPageSize   int           `json:"default_page_size"`
	LastLoginIP       string        `json:"last_login_ip"`
	LastLoginLocation string        `json:"last_login_location"`
	AllowedTagTypes   []string      `json:"allowed_tag_types"`
	CpConfiguration   struct {
		Error string `json:"error"`
	} `json:"cp_configuration"`
	MetaData struct {
		Theme string `json:"theme"`
	} `json:"meta_data"`
}

func GetLoggedInUserDetails() (*common.APIResponse, error) {
	user_details_resp := LoggedInUserDetailsResponse{}
	resp, err := CTIX_CLIENT.MakeRequest("GET", user_details_endpoint, nil, &user_details_resp, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(user_details_resp),
		RawResponse:     resp,
	}, err
}

func GetLoggedInUserDetailsTool(s *server.MCPServer) {
	getLoggedInUserDetailsTool := mcp.NewTool("logged-in-user-details",
		mcp.WithDescription(`This tool will give the current logged-in user details of the CTIX`),
	)
	s.AddTool(getLoggedInUserDetailsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := GetLoggedInUserDetails()
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}
