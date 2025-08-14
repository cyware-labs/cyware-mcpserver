package ctix

import (
	"log"

	"github.com/cyware-labs/cyware-mcpserver/common"
	"resty.dev/v3"
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
	client := common.APIClient{
		BASE_URL: CTIX_CONFIG.BASE_URL,
		Client:   resty.New(),
	}
	client.MakeRequest("POST", Login_endpoint, nil, &login_resp, login_payload, nil)
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
