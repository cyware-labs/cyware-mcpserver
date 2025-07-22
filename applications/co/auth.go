package co

import (
	"context"
	"fmt"
	"log"

	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const Login_endpoint = "/cpapi/rest-auth/login/"

type LoginPayload struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	Email   string `json:"email"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func Login() (*common.APIResponse, error) {

	// based on the auth type generate the auth header and update the client.
	login_resp := LoginResponse{}
	log.Println("AUTHTYPE", CO_CONFIG.Auth.Type, " BASE URL", CO_CONFIG.BASE_URL)
	switch CO_CONFIG.Auth.Type {
	case "basic":
		login_payload := LoginPayload{
			Email:    CO_CONFIG.Auth.Username,
			Password: common.Base64Encode(CO_CONFIG.Auth.Password),
		}
		resp, err := CO_CLIENT.MakeRequest("POST", Login_endpoint, nil, &login_resp, login_payload, nil)
		CO_CLIENT.Client.SetHeader("Authorization", common.FormatCywareToken(login_resp.Token))

		return &common.APIResponse{
			FilteredReponse: common.JsonifyResponse(login_resp),
			RawResponse:     resp,
		}, err

	case "token":
		token := common.FormatCywareToken(CO_CONFIG.Auth.Token)
		CO_CLIENT.Client.SetHeader("Authorization", token)
		login_resp.Token = token
	default:
		return nil, fmt.Errorf("unsupported auth_type: %s", CO_CONFIG.Auth.Type)
	}

	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(login_resp),
		RawResponse:     nil,
	}, nil
}

func SetUpWorkspace() {
	resp := GetLoggedInUserDetails()
	USER_WS = resp.PreferredWorkspace.Code
}

func LoginTool(s *server.MCPServer) {
	loginTool := mcp.NewTool("login-to-co",
		mcp.WithDescription("This tool will login into CO and set the auth token."),
	)
	s.AddTool(loginTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := Login()
		SetUpWorkspace()
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetSoarEndpoint(endpoint string) string {
	return fmt.Sprintf("/soarapi/%v/%v", USER_WS, endpoint)
}
