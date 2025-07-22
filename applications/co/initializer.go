package co

import (
	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/server"
	"resty.dev/v3"
)

// CO_CLIENT is the shared HTTP client for all CO-related API requests.
var CO_CLIENT common.APIClient

// CO_CONFIG holds the loaded configuration for the CO application,
// including its base URL and authentication credentials.
var CO_CONFIG common.Application

// Default CO workspace of the user
var USER_WS string

// InitClient initializes the global CO client using the application
// configuration from the main config.
//
// It sets the base URL to the domain of the configured endpoint and appends
// The initialized client is stored in CO_CLIENT and is used by all CO tools.
func InitClient(cfg *common.Config) {

	CO_CONFIG = cfg.Applications["co"]
	CO_CONFIG.BASE_URL = common.GetDomain(CO_CONFIG.BASE_URL)

	// initializing global httpclient which will be used for all the CO related APIs
	CO_CLIENT = common.APIClient{
		BASE_URL: CO_CONFIG.BASE_URL,
		Client:   resty.New(),
	}
}

// Initialize sets up all CO tools and the API client within the MCP server context.
//
// It first initializes the client configuration, then registers all CO-specific tools
// to the server instance.
func Initialize(cfg *common.Config, s *server.MCPServer) {
	InitClient(cfg)
	InitTools(s)
}

// InitTools performs login and registers all CO-specific tools with the MCP server.
// It ensures a valid session token is set via Login(), and then exposes all relevant
// CO tools such as getting playbook list, executing the playbook, executing actions capabilities.
func InitTools(s *server.MCPServer) {
	Login()
	SetUpWorkspace()

	LoginTool(s)
	GetPlayBookListTool(s)
	GetPlaybookDetailsTool(s)
	ExecutePlaybookTool(s)
	GetCOAppsListingTool(s)
	GetCOAppDetailsTool(s)
	COAppActionsListingTool(s)
	GetCOAppActionDetailsTool(s)
	GetConfiguredInstancesOfCOAppTool(s)
	ExecuteActionOfCOAppTool(s)
}
