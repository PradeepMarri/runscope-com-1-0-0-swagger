package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/runscope-api/mcp-server/config"
	"github.com/runscope-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Put_buckets_bucketkey_tests_testidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		bucketKeyVal, ok := args["bucketKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: bucketKey"), nil
		}
		bucketKey, ok := bucketKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: bucketKey"), nil
		}
		testIdVal, ok := args["testId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: testId"), nil
		}
		testId, ok := testIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: testId"), nil
		}
		url := fmt.Sprintf("%s/buckets/%s/tests/%s", cfg.BaseURL, bucketKey, testId)
		req, err := http.NewRequest("PUT", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.TestDetail
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreatePut_buckets_bucketkey_tests_testidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_buckets_bucketKey_tests_testId",
		mcp.WithDescription("Modify a test's name, description, default environment and its steps. To modify other individual properties of a test, make requests to the steps, environments, and schedules subresources of the test."),
		mcp.WithString("bucketKey", mcp.Required(), mcp.Description("Unique identifier for a bucket")),
		mcp.WithString("testId", mcp.Required(), mcp.Description("Unique identifier for a test")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Put_buckets_bucketkey_tests_testidHandler(cfg),
	}
}
