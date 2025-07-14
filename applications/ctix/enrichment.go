package ctix

import (
	"context"

	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	list_enrichment_tool    = "integration/apps/"
	enrichment_tool_details = "integration/apps/detail/"
)

type EnrichmentToolsList struct {
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	PageSize int         `json:"page_size"`
	Total    int         `json:"total"`
	Results  []struct {
		ID             string      `json:"id"`
		Title          string      `json:"title"`
		Slug           string      `json:"slug"`
		Category       string      `json:"category"`
		IsActive       bool        `json:"is_active"`
		Bulk           bool        `json:"bulk"`
		BatchSize      interface{} `json:"batch_size"`
		Timer          interface{} `json:"timer"`
		ReportKey      interface{} `json:"report_key"`
		ConfiguredOnce bool        `json:"configured_once"`
		// EnrichmentPolicy bool        `json:"enrichment_policy"`
		// ThirdPartyLogo   string      `json:"third_party_logo"`
		DefaultAPIURL string `json:"default_api_url,omitempty"`
		// VisualizerTool            bool        `json:"visualizer_tool,omitempty"`
		// LitePlusVersion           bool        `json:"lite_plus_version,omitempty"`
		// ConnectorVersion          string      `json:"connector_version,omitempty"`
		// CtixCodeVersion           string      `json:"ctix_code_version,omitempty"`
		// SupportedProducts         []string    `json:"supported_products,omitempty"`
		// MinimumCtixDropinVersions []string    `json:"minimum_ctix_dropin_versions,omitempty"`
		EnrichmentTool struct {
			Vulncheck struct {
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"vulncheck"`
		} `json:"enrichment_tool,omitempty"`
	} `json:"results"`
}

type EnrichmentToolDetails struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Category string `json:"category"`
	IsActive bool   `json:"is_active"`
	Metadata struct {
		Fields []struct {
			Key   string `json:"key"`
			Rules struct {
				Required  bool   `json:"required"`
				Validator string `json:"validator"`
			} `json:"rules,omitempty"`
			FieldType    string `json:"field_type"`
			DefaultValue string `json:"default_value,omitempty"`
			Label        string `json:"label,omitempty"`
		} `json:"fields"`
	} `json:"metadata"`
	UploadedPackage interface{} `json:"uploaded_package"`
	Bulk            bool        `json:"bulk"`
	BatchSize       interface{} `json:"batch_size"`
	Timer           interface{} `json:"timer"`
	ReportKey       interface{} `json:"report_key"`
	ThirdPartyLogo  string      `json:"third_party_logo"`
}

type EnrichmentToolActionConfig struct {
	Results []struct {
		Name               string `json:"name"`
		ThirdPartyConfigID string `json:"third_party_config_id"`
		IsActive           bool   `json:"is_active"`
		Actions            []struct {
			PollingType    string `json:"polling_type,omitempty"`
			ActionConfigID string `json:"action_config_id,omitempty"`
			ConnectorName  string `json:"connector_name,omitempty"`
			Name           string `json:"name,omitempty"`
			// PollFrequency  int    `json:"poll_frequency,omitempty"`
			// LastPolled         interface{} `json:"last_polled,omitempty"`
			// LastConnected      interface{} `json:"last_connected,omitempty"`
			ActionSlug         string `json:"action_slug"`
			Component          string `json:"component,omitempty"`
			IsActive           bool   `json:"is_active,omitempty"`
			ThirdPartyActionID string `json:"third_party_action_id"`
			CollectionID       string `json:"collection_id,omitempty"`
			IsConfigured       bool   `json:"is_configured"`
			// Created            time.Time   `json:"created,omitempty"`
			// Modified           time.Time   `json:"modified,omitempty"`
			WorkingStatus string `json:"working_status,omitempty"`
		} `json:"actions"`
		Quota struct {
			Total       string `json:"total"`
			Available   string `json:"available"`
			Used        string `json:"used"`
			Duration    string `json:"duration"`
			LastUpdated int    `json:"last_updated"`
		} `json:"quota"`
	} `json:"results"`
}

type SupportedEnrichmentToolForSDO struct {
	IP []struct {
		ID             string `json:"id"`
		AppName        string `json:"app_name"`
		AppSlug        string `json:"app_slug"`
		ThirdPartyLogo string `json:"third_party_logo"`
		IsActive       bool   `json:"is_active"`
	} `json:"ip"`

	CVE []struct {
		ID             string `json:"id"`
		AppName        string `json:"app_name"`
		AppSlug        string `json:"app_slug"`
		ThirdPartyLogo string `json:"third_party_logo"`
		IsActive       bool   `json:"is_active"`
	} `json:"cve"`
}

type EnrichmentResponse struct {
	RawData    any `json:"raw_data"`
	ParsedData struct {
		Reputation             int           `json:"reputation"`
		PulseCount             int           `json:"pulse_count"`
		PulseAggregatedDetails []interface{} `json:"pulse_aggregated_details"`
		PulseReferences        []interface{} `json:"pulse_references"`
		Geo                    struct {
			City        interface{} `json:"city"`
			Region      interface{} `json:"region"`
			CountryCode string      `json:"country_code"`
			Asn         string      `json:"asn"`
			CountryName string      `json:"country_name"`
		} `json:"geo"`
	} `json:"parsed_data"`
	Status     int         `json:"status"`
	Timestamp  interface{} `json:"timestamp"`
	EnrichedOn int         `json:"enriched_on"`
	Verdict    string      `json:"verdict"`
}

func GetEnrichmenToolsList(params map[string]string) (*common.APIResponse, error) {
	enrichmet_tool_list_resp := EnrichmentToolsList{}
	resp, err := CTIX_CLIENT.MakeRequest("GET", list_enrichment_tool, params, &enrichmet_tool_list_resp, nil, nil)

	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(enrichmet_tool_list_resp),
		RawResponse:     resp,
	}, err
}

func GetEnrichmenToolsListTool(s *server.MCPServer) {
	getEnrichmenToolsList := mcp.NewTool("get-enrichment-tools-list",
		mcp.WithDescription("This tool will give list of all the enrichment tools. To use any any tool for enrichment, the tool must be configured."),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query params information. Query params which can be send
			1. "page":  This is the page number for the paginated query. Used to get the result of specific page number
			2. "page_size" : This is the page size number of result per page. Used to get the specified number of result per page. Please note here if you are making paginated call then keep the page_size same in all the pages otherwise you will get duplicate entries in two different pages.
			3. "is_active" : This represents query params which filters only the tools which are active. It has a string value either true or false.
			4. "q" : This represent name tool that is specifically needs to searched in the erichment tools list.
			`),
			// mcp.AdditionalProperties(false),
		),
	)

	s.AddTool(getEnrichmenToolsList, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mp := request.Params.Arguments["params"].(map[string]interface{})

		params_list := []string{"page", "page_size", "is_active", "q"}

		params := map[string]string{}
		params["category"] = "threat_intelligence_enrichment"
		for _, v := range params_list {
			if _, ok := mp[v]; ok {
				params[v] = mp[v].(string)
			}
		}
		resp, err := GetEnrichmenToolsList(params)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetEnrichmentToolsDetails(app_id string) (*common.APIResponse, error) {
	enrichment_tool_details_resp := EnrichmentToolDetails{}
	endpoint := enrichment_tool_details + app_id + "/"
	resp, err := CTIX_CLIENT.MakeRequest("GET", endpoint, nil, &enrichment_tool_details_resp, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(enrichment_tool_details_resp),
		RawResponse:     resp,
	}, err
}

func GetEnrichmentToolsDetailsTool(s *server.MCPServer) {
	getEnrichmenToolsDetail := mcp.NewTool("get-enrichment-tool-details",
		mcp.WithDescription("This tool will give details about the enrichment tool. Details like title, logo, metadata etc.s"),
		mcp.WithString(
			"app_id",
			mcp.Description("This is the app_id of the enrichment tool which is required to hit the endpoint to get the details"),
		),
	)

	s.AddTool(getEnrichmenToolsDetail, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		app_id := request.Params.Arguments["app_id"].(string)

		resp, err := GetEnrichmentToolsDetails(app_id)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetEnrichmentToolActionConfigs(app_id string, params map[string]string) (*common.APIResponse, error) {
	endpoint := "/integration/apps/" + app_id + "/action_configs/"
	enrichment_tool_action_config_resp := EnrichmentToolActionConfig{}

	resp, err := CTIX_CLIENT.MakeRequest("GET", endpoint, params, &enrichment_tool_action_config_resp, nil, nil)

	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(enrichment_tool_action_config_resp),
		RawResponse:     resp,
	}, err
}

func GetEnrichmentToolActionConfigsTool(s *server.MCPServer) {
	getEnrichmentToolActionConfigs := mcp.NewTool("get-enrichment-tool-action-configs",
		mcp.WithDescription("This tool provides detailed information about all supported actions for a specific enrichment tool. For example, if the tool is AbuseIPDB, it will return actions related to enriching IP addresses, as that's the only type it supports. It also includes quota usage for each action."),
		mcp.WithString(
			"app_id",
			mcp.Description("The unique app_id of the enrichment tool. This is required to fetch supported actions and quota usage."),
		),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query params information. Query params which can be send
			1. "page":  This is the page number for the paginated query. Used to get the result of specific page number
			2. "page_size" : This is the page size number of result per page. Used to get the specified number of result per page. Please note here if you are making paginated call then keep the page_size same in all the pages otherwise you will get duplicate entries in two different pages.
			`),
			// mcp.AdditionalProperties(false),
		),
	)

	s.AddTool(getEnrichmentToolActionConfigs, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		app_id := request.Params.Arguments["app_id"].(string)

		mp, _ := request.Params.Arguments["params"].(map[string]interface{})

		params_list := []string{"page", "page_size"}
		params := map[string]string{}
		if mp != nil {
			for _, v := range params_list {
				if _, ok := mp[v]; ok {
					params[v] = mp[v].(string)
				}
			}
		}
		resp, err := GetEnrichmentToolActionConfigs(app_id, params)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func GetAllEnrichmentToolSupportedForThreatDataObject(params map[string]string) (*common.APIResponse, error) {
	endpoint := "integration/apps/actions/"
	supported_enrichmenttool_resp := SupportedEnrichmentToolForSDO{}
	resp, err := CTIX_CLIENT.MakeRequest("GET", endpoint, params, &supported_enrichmenttool_resp, nil, nil)

	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(supported_enrichmenttool_resp),
		RawResponse:     resp,
	}, err

}

func GetAllEnrichmentToolSupportedForThreatDataObjectTool(s *server.MCPServer) {
	getAllEnrichmentToolSupportedForThreatDataObject := mcp.NewTool("enrichment-tool-supported-for-threat-data-object",
		mcp.WithDescription(`This tool returns a list of enrichment tools that support a specific Threat Data Object (SDO).

		📌 Important: Use this tool only when the user explicitly requests enrichment for an object. It helps determine which enrichment tools are applicable to the specified object type.

		❗ Note: This tool does not indicate whether an object has already been enriched. It only lists potential enrichment tools available for the object type provided.
`),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query params information. Query params which can be send
			1. "action_name": Specifies the type of enrichment action to match tools against. One of below value should be passed based on the ioc type:
   				- "get_ip" for IPv4/IPv6 indicators
				- "get_domain" for domain indicators
				- "get_url" for URL indicators
				- "get_hash" for hash indicators like MD5, SHA1, SHA256, etc.
				- "get_cve" for vulnerabilities
			2. "is_active": A boolean value ("true" or "false") indicating whether to return only active tools. If set to "true", only currently active enrichment tools are included.
			3. "full_list" : Always set to "true"
			`),
			// mcp.AdditionalProperties(false),
		),
	)

	s.AddTool(getAllEnrichmentToolSupportedForThreatDataObject, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mp := request.Params.Arguments["params"].(map[string]interface{})

		params_list := []string{"action_name", "is_active", "full_list"}
		params := map[string]string{}

		params["full_list"] = "true"
		for _, v := range params_list {
			if _, ok := mp[v]; ok {
				params[v] = mp[v].(string)
			}
		}
		resp, err := GetAllEnrichmentToolSupportedForThreatDataObject(params)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}

func EnrichThreatDataObject(params map[string]string) (*common.APIResponse, error) {
	endpoint := "integration/apps/update/threatdata/"
	enrichment_resp := EnrichmentResponse{}
	resp, err := CTIX_CLIENT.MakeRequest("GET", endpoint, params, &enrichment_resp, nil, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(enrichment_resp),
		RawResponse:     resp,
	}, err
}

func EnrichThreatDataObjectTool(s *server.MCPServer) {
	enrichThreatDataObjectTool := mcp.NewTool("enrich-threat-data-object",
		mcp.WithDescription(`
			PURPOSE: This tool enriches a Threat Data Object using data from enrichment tools.
			Before using this tool always use GetEnrichmentToolActionConfigsTool to get the action config which is then used in the enrich payload
			⚠️ IMPORTANT: This tool performs an active operation that consumes API quota. Only use when explicitly requested by the user. Never use as part of a general information gathering process.
			If user requests information about an threat data object BUT doesn't specifically request enrichment:
   				- First provide basic information using get-threat-data-object-details
   				- Then ASK if they want to enrich the object before proceeding		
			`),
		mcp.WithObject(
			"params",
			mcp.Description(`Key-value pairs for query parameters. These parameters help identify the enrichment tool and the threat data to be enriched. The following fields are must be present in the payload:

				1. "app_slug": (Required) The unique identifier (slug) of the enrichment tool. You can obtain this from the list of supported enrichment tools for that object.

				2. "value": (Required) The actual value to enrich. Example: "12.1.2.1" for an IP address.

				3. "action_slug": (Required) Specifies the enrichment action to perform. It must be chosen based on the type of IOC (Indicator of Compromise):
				- Use "get_ip" for IPv4 or IPv6 addresses
				- Use "get_domain" for domain names
				- Use "get_url" for URLs
				- Use "get_hash" for file hashes such as MD5, SHA1, SHA256, etc.

				4. "object_id": (Required) The unique identifier of the threat data object to be enriched. You can retrieve this using a CQL (Cyware Query Language) query.

				5. "object_type": (Required) The type of the threat data object (e.g., indicator, SDO). This helps in identifying how the enrichment should be applied.

				6. "ioc_type": (Required only if enriching an indicator object) Specifies the type of indicator being enriched (e.g., "ipv4", "domain", "sha256").

				Make sure all required fields are present in the payload for successful enrichment.`),
			// mcp.AdditionalProperties(false),
		),
	)

	s.AddTool(enrichThreatDataObjectTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mp := request.Params.Arguments["params"].(map[string]interface{})
		params_list := []string{"app_slug", "value", "action_slug", "object_id", "object_type", "ioc_type"}
		params := map[string]string{}

		for _, v := range params_list {
			if _, ok := mp[v]; ok {
				params[v] = mp[v].(string)
			}
		}
		resp, err := EnrichThreatDataObject(params)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}
