package general

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func Current_time() time.Time {
	currentTime := time.Now()
	return currentTime
}

// GetEpochWithDeltaDays returns the epoch time in milliseconds after subtracting delta days if negative
func GetEpochWithDeltaFromNowDays(deltaDays int) int64 {
	// Add (or subtract) delta days from current time
	targetTime := time.Now().AddDate(0, 0, deltaDays)
	// Return epoch time in seconds
	return targetTime.Unix()
}

func GetEpochWithDeltaFromNowDaysTool(s *server.MCPServer) {
	getEpochWithDeltaFromNowDaysTool := mcp.NewTool("get-epoch-with-delta-from-now",
		mcp.WithDescription(`This tool will give you the epoch in milliseconds after subtracting delta days from current time
		Always use this tool to calculate the timestamps.
		`),
		mcp.WithNumber("delta",
			mcp.Description("Represents the delta, which means current time minus number of days. It must be a int value. Also for past timestamp use negative value and for future use postive value")),
	)

	s.AddTool(getEpochWithDeltaFromNowDaysTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		delta := request.Params.Arguments["delta"].(float64)
		resp := GetEpochWithDeltaFromNowDays(int(delta))

		return mcp.NewToolResultText(fmt.Sprintf(`{"timestamp":%v}`, resp)), nil
	})

}
