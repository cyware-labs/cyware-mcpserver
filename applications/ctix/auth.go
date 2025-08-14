package ctix

import (
	"log"

	"github.com/cyware-labs/cyware-mcpserver/common"
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

// AuthParams holds the authentication parameters
type AuthParams struct {
	AccessID  string
	Signature string
	Expires   string
}

func GenerateAuthHeaders() string {
	login_resp := LoginResponse{}
	login_payload := LoginPayload{
		Email:    CTIX_CONFIG.Auth.Username,
		Password: CTIX_CONFIG.Auth.Password,
	}
	CTIX_CLIENT.MakeRequest("POST", Login_endpoint, nil, &login_resp, login_payload, nil)
	return common.FormatCywareToken(login_resp.Token)
}

func Login() {

	// based on the auth type generate the auth header and update the client.
	switch CTIX_CONFIG.Auth.Type {
	case "basic":
		auth_token := GenerateAuthHeaders()
		CTIX_CLIENT.Client.SetHeader("Authorization", auth_token)

	case "token":
		token := common.FormatCywareToken(CTIX_CONFIG.Auth.Token)
		CTIX_CLIENT.Client.SetHeader("Authorization", token)

	case "openapicreds":
		params := common.GenerateAuthParams(CTIX_CONFIG.Auth.AccessID, CTIX_CONFIG.Auth.SecretKey)
		CTIX_CLIENT.Client.SetQueryParams(params)

	default:
		log.Printf("unsupported auth_type: %s", CTIX_CONFIG.Auth.Type)
	}
}

// func LoginTool(s *server.MCPServer) {
// 	loginTool := mcp.NewTool("login-to-ctix",
// 		mcp.WithDescription("This tool will login into CTIX and set the auth token."),
// 	)

// 	s.AddTool(loginTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
// 		resp, err := Login()
// 		return common.MCPToolResponse(resp, []int{200}, err)
// 	})
// }
