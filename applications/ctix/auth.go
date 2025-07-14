package ctix

import (
	"context"
	"fmt"
	"strings"

	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const Login_endpoint = "rest-auth/login/user-pass/"

type LoginPayload struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	User_id string `json:"user_id"`
	Token   string `json:"token"`
	Email   string `json:"email"`
}

func formatCywareToken(rawToken string) string {
	const prefix = "CYW "

	if rawToken == "" {
		return ""
	}

	if strings.HasPrefix(rawToken, prefix) {
		return rawToken
	}

	return prefix + rawToken
}

func Login() (*common.APIResponse, error) {

	// based on the auth type generate the auth header and update the client.
	login_resp := LoginResponse{}
	switch CTIX_CONFIG.Auth.Type {
	case "basic":
		login_payload := LoginPayload{
			Email:    CTIX_CONFIG.Auth.Username,
			Password: CTIX_CONFIG.Auth.Password,
		}
		resp, err := CTIX_CLIENT.MakeRequest("POST", Login_endpoint, nil, &login_resp, login_payload, nil)
		CTIX_CLIENT.Client.SetHeader("Authorization", formatCywareToken(login_resp.Token))

		return &common.APIResponse{
			FilteredReponse: common.JsonifyResponse(login_resp),
			RawResponse:     resp,
		}, err

	case "token":
		token := formatCywareToken(CTIX_CONFIG.Auth.Token)
		CTIX_CLIENT.Client.SetHeader("Authorization", token)
		login_resp.Token = token
	default:
		return nil, fmt.Errorf("unsupported auth_type: %s", CTIX_CONFIG.Auth.Type)
	}

	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(login_resp),
		RawResponse:     nil,
	}, nil
}

func LoginTool(s *server.MCPServer) {
	loginTool := mcp.NewTool("login-to-ctix",
		mcp.WithDescription("This tool will login into CTIX and set the auth token."),
	)

	s.AddTool(loginTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := Login()
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}
