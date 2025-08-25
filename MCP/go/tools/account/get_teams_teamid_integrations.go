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

func Get_teams_teamid_integrationsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		teamIdVal, ok := args["teamId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: teamId"), nil
		}
		teamId, ok := teamIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: teamId"), nil
		}
		url := fmt.Sprintf("%s/teams/%s/integrations", cfg.BaseURL, teamId)
		req, err := http.NewRequest("GET", url, nil)
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
		var result interface{}
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

func CreateGet_teams_teamid_integrationsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_teams_teamId_integrations",
		mcp.WithDescription("Team integrations list"),
		mcp.WithString("teamId", mcp.Required(), mcp.Description("Unique identifier for team")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_teams_teamid_integrationsHandler(cfg),
	}
}
