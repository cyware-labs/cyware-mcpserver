package ctix

import (
	"context"

	"github.com/cyware-labs/cyware-mcpserver/applications/ctix/helpers"
	"github.com/cyware-labs/cyware-mcpserver/common"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const Quick_add_intel_create = "conversion/quick-intel/create-stix/"

type QuickAddIntelPayload struct {
	Context          string `json:"context"`
	ParsedIndicators struct {
	} `json:"parsed_indicators"`
	Metadata struct {
		Tlp                      string        `json:"tlp"`
		DefaultMarkingDefinition string        `json:"default_marking_definition"`
		MarkingConfig            string        `json:"marking_config"`
		Tags                     []interface{} `json:"tags"`
		IsApplyAll               bool          `json:"is_apply_all"`
		AdditionalDefaults       []interface{} `json:"additional_defaults"`
		CustomScores             struct {
		} `json:"custom_scores"`
		Confidence  int    `json:"confidence"`
		Description string `json:"description"`
	} `json:"metadata"`
	Import struct {
	} `json:"import"`
	Indicators struct {
		Ipv4Addr           string `json:"ipv4-addr"`
		Ipv6Addr           string `json:"ipv6-addr"`
		Domain             string `json:"domain"`
		URL                string `json:"url"`
		Email              string `json:"email"`
		Md5                string `json:"md5"`
		Sha1               string `json:"sha1"`
		Sha224             string `json:"sha224"`
		Sha256             string `json:"sha256"`
		Sha512             string `json:"sha512"`
		Sha384             string `json:"sha384"`
		Ssdeep             string `json:"ssdeep"`
		AutonomousSystem   string `json:"autonomous-system"`
		WindowsRegistryKey string `json:"windows-registry-key"`
	} `json:"indicators"`
	Title string `json:"title"`
	Sdos  struct {
		Vulnerability  string `json:"vulnerability"`
		IntrusionSet   string `json:"intrusion-set"`
		Malware        string `json:"malware"`
		Campaign       string `json:"campaign"`
		ThreatActor    string `json:"threat-actor"`
		AttackPattern  string `json:"attack-pattern"`
		Incident       string `json:"incident"`
		CourseOfAction string `json:"course-of-action"`
		Identity       string `json:"identity"`
		Tool           string `json:"tool"`
		Infrastructure string `json:"infrastructure"`
		Location       struct {
			Type   string `json:"type"`
			Values []struct {
				Value string `json:"value"`
				Label string `json:"label"`
			} `json:"values"`
		} `json:"location"`
		MalwareAnalysis string `json:"malware-analysis"`
	} `json:"sdos"`
	Observables struct {
		Artifact        string `json:"artifact"`
		Directory       string `json:"directory"`
		MacAddr         string `json:"mac-addr"`
		EmailMessage    string `json:"email-message"`
		Mutex           string `json:"mutex"`
		NetworkTraffic  string `json:"network-traffic"`
		Process         string `json:"process"`
		Software        string `json:"software"`
		UserAccount     string `json:"user-account"`
		X509Certificate string `json:"x509-certificate"`
		File            string `json:"file"`
	} `json:"observables"`
	Relations       []interface{} `json:"relations"`
	CreateIntelFeed bool          `json:"create_intel_feed"`
}

type QuickAddIntelResponse struct {
	Details string `json:"details"`
	TaskID  string `json:"task_id"`
}

func CreateQuickAddIntel(payload any) (*common.APIResponse, error) {
	// implementation of quick add intel
	quick_add_resp := &QuickAddIntelResponse{}
	resp, err := CTIX_CLIENT.MakeRequest("POST", Quick_add_intel_create, nil, quick_add_resp, payload, nil)
	return &common.APIResponse{
		FilteredReponse: common.JsonifyResponse(quick_add_resp),
		RawResponse:     resp,
	}, err
}

func CreateQuickAddIntelTool(s *server.MCPServer) {

	schema := helpers.Quick_add_intel_schema
	quickAddIntelCreateTool := mcp.NewToolWithRawSchema("quick-add-intel-create",
		`This tool creates Threat Intel objects using the Quick Add Intel module in CTIX.
    A Report object is always created alongside, with the same name as title in the payload. The report will have relation with all the objects created with it.

    Please note: The creation process runs in the background, so if you're running a CQL query to search for the newly created Intel, you may need to:

      -Wait a few seconds before running the query, or

      -Retry the search 2–3 times.
    `,
		[]byte(schema),
	)

	s.AddTool(quickAddIntelCreateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := CreateQuickAddIntel(request.Params.Arguments)
		return common.MCPToolResponse(resp, []int{200}, err)
	})
}
