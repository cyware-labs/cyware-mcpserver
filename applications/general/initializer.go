package general

import "github.com/mark3labs/mcp-go/server"

func Initialize(s *server.MCPServer) {
	InitTools(s)
}

func InitTools(s *server.MCPServer) {
	// GetEpochWithDeltaFromNowDaysTool(s)
	ConvertDateStringToEpochTool(s)
}
