package ctix

import (
	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/server"
	"resty.dev/v3"
)

// CTIX_CLIENT is the shared HTTP client for all CTIX-related API requests.
var CTIX_CLIENT common.APIClient

// CTIX_CONFIG holds the loaded configuration for the CTIX application,
// including its base URL and authentication credentials.
var CTIX_CONFIG common.Application

// InitClient initializes the global CTIX client using the application
// configuration from the main config.
//
// It sets the base URL to the domain of the configured endpoint and appends
// "/ctixapi/" as the API base path. The initialized client is stored in
// CTIX_CLIENT and is used by all CTIX tools.
func InitClient(cfg *common.Config) {

	CTIX_CONFIG = cfg.Applications["ctix"]
	CTIX_CONFIG.BASE_URL = common.GetDomain(CTIX_CONFIG.BASE_URL) + "/ctixapi/"

	// initializing global httpclient which will be used for all the CTIX related APIs
	CTIX_CLIENT = common.APIClient{
		BASE_URL: CTIX_CONFIG.BASE_URL,
		Client:   resty.New(),
	}
}

// Initialize sets up all CTIX tools and the API client within the MCP server context.
//
// It first initializes the client configuration, then registers all CTIX-specific tools
// to the server instance.
func Initialize(cfg *common.Config, s *server.MCPServer) {

	InitClient(cfg)
	InitTools(s)
}

// InitTools performs login and registers all CTIX-specific tools with the MCP server.
//
// It ensures a valid session token is set via Login(), and then exposes all relevant
// CTIX tools such as user info, threat data actions, tagging, and CQL search capabilities.
func InitTools(s *server.MCPServer) {

	// login and setting the token for all the subsequent requests
	Login()

	LoginTool(s)
	GetLoggedInUserDetailsTool(s)

	// cql and search
	CQLCTIXSearchGrammarTool(s)
	GetCQLQuerySearchResultTool(s)
	GetThreatDataObjectDetailsTool(s)
	GetThreatDataObjectRelationsTool(s)
	GetAvailableRelationTypeListingTool(s)

	// bulk action threat data
	ThreatDataListBulkActionTools(s)

	// tag management
	CreateTaginCTIXTool(s)
	GetCTIXTagListingTool(s)

	// enrichment
	GetEnrichmenToolsListTool(s)
	GetEnrichmentToolsDetailsTool(s)
	GetEnrichmentToolActionConfigsTool(s)
	GetAllEnrichmentToolSupportedForThreatDataObjectTool(s)
	EnrichThreatDataObjectTool(s)

	// intel creation
	CreateQuickAddIntelTool(s)

}
