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

// DateStringToEpoch converts a date string in format "dd-mm-yyyy-hh-min-sec" to epoch time
func DateStringToEpoch(dateStr string) (int64, error) {
	// Parse the date string using the custom layout
	layout := "02-01-2006-15-04-05"

	// Get local time zone
	loc := time.Now().Location()

	// Parse using the local time zone
	t, err := time.ParseInLocation(layout, dateStr, loc)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %v", err)
	}

	// Return Unix timestamp (epoch time)
	return t.Unix(), nil
}

func ConvertDateStringToEpochTool(s *server.MCPServer) {
	convertDateStringToEpochTool := mcp.NewTool("convert-date-string-to-epoch",
		mcp.WithDescription(`This tool will convert the give date string of format "dd-mm-yyyy-hh-min-sec" to epoch. 
		!!Important!! You must always use this tool to convert datetime into epoch.`),
		mcp.WithString("date",
			mcp.Description(`Represents the date in the format "dd-mm-yyyy-hh-min-sec"`)),
	)

	s.AddTool(convertDateStringToEpochTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		date := request.Params.Arguments["date"].(string)
		resp, _ := DateStringToEpoch(date)

		return mcp.NewToolResultText(fmt.Sprintf(`{"timestamp":%v}`, resp)), nil
	})

}
