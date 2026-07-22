package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"resty.dev/v3"
)

const retry = 4

var failed_status = []int{400, 401}

func FormatCywareToken(rawToken string) string {
	const prefix = "CYW "

	if rawToken == "" {
		return ""
	}

	if strings.HasPrefix(rawToken, prefix) {
		return rawToken
	}

	return prefix + rawToken
}

// Base64Encode encodes the input string to Base64 format
func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// GenerateAuthParams generates authentication parameters
func GenerateAuthParams(accessID, secretKey string) map[string]string {
	// Generating unix timestamp
	unixTimestamp := time.Now().Unix()

	// Adding 20 seconds for expires
	expires := unixTimestamp + 20

	// Creating the string to sign
	toSign := accessID + "\n" + strconv.FormatInt(expires, 10)

	// Generating HMAC-SHA1 hash
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(toSign))
	hash := h.Sum(nil)

	// Converting to base64
	hashInBase64 := base64.StdEncoding.EncodeToString(hash)

	params := map[string]string{
		"Expires":   strconv.FormatInt(expires, 10),
		"AccessID":  accessID,
		"Signature": hashInBase64,
	}
	return params
}

// ExtractParams extracts params key from the tool call request and convert them into a map
func ExtractParams(request mcp.CallToolRequest, params_list []string) map[string]string {
	params := map[string]string{}
	mp, ok := request.Params.Arguments["params"].(map[string]interface{})
	if !ok {
		return params
	}

	for _, v := range params_list {
		if val, ok := mp[v]; ok {
			params[v] = anyToString(val)
		}
	}
	return params
}

func ParseTimeoutDuration(timeoutArg any) (time.Duration, error) {
	var timeoutSeconds float64

	switch v := timeoutArg.(type) {
	case float64:
		timeoutSeconds = v
	case float32:
		timeoutSeconds = float64(v)
	case int:
		timeoutSeconds = float64(v)
	case int8:
		timeoutSeconds = float64(v)
	case int16:
		timeoutSeconds = float64(v)
	case int32:
		timeoutSeconds = float64(v)
	case int64:
		timeoutSeconds = float64(v)
	case uint:
		timeoutSeconds = float64(v)
	case uint8:
		timeoutSeconds = float64(v)
	case uint16:
		timeoutSeconds = float64(v)
	case uint32:
		timeoutSeconds = float64(v)
	case uint64:
		timeoutSeconds = float64(v)
	case json.Number:
		parsedTimeoutSeconds, err := v.Float64()
		if err != nil {
			return 0, fmt.Errorf("timeout_seconds must be a number")
		}
		timeoutSeconds = parsedTimeoutSeconds
	default:
		return 0, fmt.Errorf("timeout_seconds must be a number")
	}

	if timeoutSeconds <= 0 || math.IsNaN(timeoutSeconds) || math.IsInf(timeoutSeconds, 0) {
		return 0, fmt.Errorf("timeout_seconds must be greater than 0")
	}

	return time.Duration(timeoutSeconds * float64(time.Second)), nil
}

// anyToString coerces a JSON-decoded value to a string. JSON numbers decode to
// float64, so callers that pass numeric params (e.g. {"page": 1}) must not be
// assumed to be strings — an unchecked type assertion panics the tool handler.
func anyToString(val any) string {
	switch t := val.(type) {
	case string:
		return t
	case float64:
		// Render integral values without a trailing ".0" (page=1 -> "1").
		return strconv.FormatFloat(t, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(t)
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", t)
	}
}

func GetRestyClient(retryHook func(r *resty.Response, err error)) *resty.Client {
	c := resty.New()
	c.SetAllowNonIdempotentRetry(true)
	c.SetRetryCount(retry)
	c.SetRetryWaitTime(1 * time.Second)

	// Retry condition
	c.AddRetryConditions(func(r *resty.Response, err error) bool {
		return r != nil && ContainsStatusCode(failed_status, r.StatusCode())
	})
	c.AddRetryHooks(retryHook)
	return c
}
