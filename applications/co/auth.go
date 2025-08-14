package co

import (
	"fmt"
	"log"

	"github.com/cyware-labs/cyware-mcpserver/common"
	"resty.dev/v3"
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

func GenerateAuthHeaders() string {
	login_resp := LoginResponse{}
	login_payload := LoginPayload{
		Email:    CO_CONFIG.Auth.Username,
		Password: common.Base64Encode(CO_CONFIG.Auth.Password),
	}
	client := common.APIClient{
		BASE_URL: CO_CONFIG.BASE_URL,
		Client:   resty.New(),
	}
	client.MakeRequest("POST", Login_endpoint, nil, &login_resp, login_payload, nil)

	return common.FormatCywareToken(login_resp.Token)
}

func Login() {

	// based on the auth type generate the auth header and update the client.
	switch CO_CONFIG.Auth.Type {
	case "basic":
		token := GenerateAuthHeaders()
		CO_CLIENT.Client.SetHeader("Authorization", token)
	case "token":
		token := common.FormatCywareToken(CO_CONFIG.Auth.Token)
		CO_CLIENT.Client.SetHeader("Authorization", token)

	case "openapicreds":
		params := common.GenerateAuthParams(CO_CONFIG.Auth.AccessID, CO_CONFIG.Auth.SecretKey)
		CO_CLIENT.Client.SetQueryParams(params)

	default:
		log.Printf("unsupported auth_type: %s", CO_CONFIG.Auth.Type)
	}
}

func SetUpWorkspace() {
	resp := GetLoggedInUserDetails()
	USER_WS = resp.PreferredWorkspace.Code
}

func GetSoarEndpoint(endpoint string) string {
	return fmt.Sprintf("/soarapi/%v/%v", USER_WS, endpoint)
}
