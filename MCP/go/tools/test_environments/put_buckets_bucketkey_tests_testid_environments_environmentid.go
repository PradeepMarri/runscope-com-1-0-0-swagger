package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/runscope-api/mcp-server/config"
	"github.com/runscope-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Put_buckets_bucketkey_tests_testid_environments_environmentidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		environmentIdVal, ok := args["environmentId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: environmentId"), nil
		}
		environmentId, ok := environmentIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: environmentId"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Environment
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/buckets/%s/tests/%s/environments/%s", cfg.BaseURL, bucketKey, testId, environmentId)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result map[string]interface{}
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

func CreatePut_buckets_bucketkey_tests_testid_environments_environmentidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_buckets_bucketKey_tests_testId_environments_environmentId",
		mcp.WithDescription("Update the details of a test environment."),
		mcp.WithString("bucketKey", mcp.Required(), mcp.Description("Unique identifier for a bucket")),
		mcp.WithString("testId", mcp.Required(), mcp.Description("Unique identifier for a test")),
		mcp.WithString("environmentId", mcp.Required(), mcp.Description("Unique identifier for a test environment")),
		mcp.WithArray("remote_agents", mcp.Description("")),
		mcp.WithBoolean("retry_on_failure", mcp.Description("")),
		mcp.WithObject("headers", mcp.Description("")),
		mcp.WithObject("initial_variables", mcp.Description("")),
		mcp.WithArray("script_library", mcp.Description("Input parameter: The list of ids for scripts, part of the script libraries, being used for this environment.")),
		mcp.WithString("auth", mcp.Description("")),
		mcp.WithArray("integrations", mcp.Description("Input parameter: The list of integrations for this environment.")),
		mcp.WithString("parent_environment_id", mcp.Description("")),
		mcp.WithString("initial_script_hash", mcp.Description("")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: Name of this environment.")),
		mcp.WithArray("regions", mcp.Description("Input parameter: An array of the region codes that this environment is using.")),
		mcp.WithString("webhooks", mcp.Description("")),
		mcp.WithString("client_certificate", mcp.Description("")),
		mcp.WithString("test_id", mcp.Description("Input parameter: The unique identifier for this test.")),
		mcp.WithString("script", mcp.Description("")),
		mcp.WithString("id", mcp.Description("Input parameter: The unique identifier for this environment.")),
		mcp.WithBoolean("stop_on_failure", mcp.Description("Input parameter: Stop executing the test after the first failed step.")),
		mcp.WithString("version", mcp.Description("")),
		mcp.WithBoolean("preserve_cookies", mcp.Description("")),
		mcp.WithObject("emails", mcp.Description("")),
		mcp.WithNumber("exported_at", mcp.Description("")),
		mcp.WithBoolean("verify_ssl", mcp.Description("Input parameter: Validate all SSL certificates on any HTTPS connections.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Put_buckets_bucketkey_tests_testid_environments_environmentidHandler(cfg),
	}
}
