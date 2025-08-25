package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// NewMessage represents the NewMessage schema from the OpenAPI specification
type NewMessage struct {
	Response map[string]interface{} `json:"response,omitempty"`
	Request map[string]interface{} `json:"request,omitempty"`
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Name string `json:"name,omitempty"` // The name of this team.
	Id string `json:"id,omitempty"` // The unique identifier for this team.
}

// Account represents the Account schema from the OpenAPI specification
type Account struct {
	Email string `json:"email,omitempty"` // The email address for this account. Only present if authorized with the account:email scope.
	Id string `json:"id,omitempty"` // The unique identifier for this account.
	Name string `json:"name,omitempty"` // The name of the person for this account.
	Teams []Team `json:"teams,omitempty"`
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code int `json:"code,omitempty"`
	Fields string `json:"fields,omitempty"`
	Message string `json:"message,omitempty"`
}

// StandardError represents the StandardError schema from the OpenAPI specification
type StandardError struct {
	Meta Meta `json:"meta,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
	ErrorField Error400 `json:"error,omitempty"`
}

// Test represents the Test schema from the OpenAPI specification
type Test struct {
	Description string `json:"description,omitempty"` // The description for the test.
	Id string `json:"id,omitempty"`
	Last_run map[string]interface{} `json:"last_run,omitempty"`
	Name string `json:"name"` // The name for the test.
	Trigger_url string `json:"trigger_url,omitempty"`
	Created_at int `json:"created_at,omitempty"` // The date the test was created in seconds (Unix time stamp format).
	Created_by map[string]interface{} `json:"created_by,omitempty"`
	Default_environment_id string `json:"default_environment_id,omitempty"`
}

// Variable represents the Variable schema from the OpenAPI specification
type Variable struct {
	Name string `json:"name,omitempty"`
	Property string `json:"property,omitempty"`
	Source string `json:"source,omitempty"`
}

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Script string `json:"script,omitempty"`
	Id string `json:"id,omitempty"` // The unique identifier for this environment.
	Stop_on_failure bool `json:"stop_on_failure,omitempty"` // Stop executing the test after the first failed step.
	Version string `json:"version,omitempty"`
	Preserve_cookies bool `json:"preserve_cookies,omitempty"`
	Emails map[string]interface{} `json:"emails,omitempty"`
	Exported_at int `json:"exported_at,omitempty"`
	Verify_ssl bool `json:"verify_ssl,omitempty"` // Validate all SSL certificates on any HTTPS connections.
	Remote_agents []Agent `json:"remote_agents,omitempty"`
	Retry_on_failure bool `json:"retry_on_failure,omitempty"`
	Headers map[string]interface{} `json:"headers,omitempty"`
	Initial_variables map[string]interface{} `json:"initial_variables,omitempty"`
	Script_library []string `json:"script_library,omitempty"` // The list of ids for scripts, part of the script libraries, being used for this environment.
	Auth string `json:"auth,omitempty"`
	Integrations []Integration `json:"integrations,omitempty"` // The list of integrations for this environment.
	Parent_environment_id string `json:"parent_environment_id,omitempty"`
	Initial_script_hash string `json:"initial_script_hash,omitempty"`
	Name string `json:"name"` // Name of this environment.
	Regions []string `json:"regions,omitempty"` // An array of the region codes that this environment is using.
	Webhooks string `json:"webhooks,omitempty"`
	Client_certificate string `json:"client_certificate,omitempty"`
	Test_id string `json:"test_id,omitempty"` // The unique identifier for this test.
}

// TestStep represents the TestStep schema from the OpenAPI specification
type TestStep struct {
	Step_type string `json:"step_type,omitempty"` // Type of test step -- request, pause, condition, ghost-inspector, or subtest.
}

// Bucket represents the Bucket schema from the OpenAPI specification
type Bucket struct {
	Trigger_url string `json:"trigger_url,omitempty"`
	Auth_token string `json:"auth_token,omitempty"` // Bucket auth token if set, otherwise this value is null.
	Messages_url string `json:"messages_url,omitempty"`
	Tests_url string `json:"tests_url,omitempty"`
	Verify_ssl bool `json:"verify_ssl,omitempty"` // True if this bucket is configured to verify ssl for requests made to it.
	Team Team `json:"team,omitempty"`
	Collections_url string `json:"collections_url,omitempty"`
	DefaultField bool `json:"default,omitempty"` // True if this bucket is the 'default' for a team. Default buckets cannot be deleted.
	Key string `json:"key,omitempty"` // The unique identifier used to address a bucket.
	Name string `json:"name,omitempty"` // The name of this bucket as displayed in your dashboard.
}

// Schedule represents the Schedule schema from the OpenAPI specification
type Schedule struct {
	Id string `json:"id,omitempty"`
	Interval string `json:"interval,omitempty"`
	Note string `json:"note,omitempty"`
	Version string `json:"version,omitempty"`
	Environment_id string `json:"environment_id,omitempty"`
	Exported_at int `json:"exported_at,omitempty"`
}

// TestStepRequest represents the TestStepRequest schema from the OpenAPI specification
type TestStepRequest struct {
	Method string `json:"method,omitempty"` // The HTTP method for this request step. E.g. GET, POST, PUT, DELETE, etc.
	Scripts []string `json:"scripts,omitempty"` // A list of post-response scripts to run after this request.
	Assertions []Assertion `json:"assertions,omitempty"` // A list of assertions to apply to the HTTP response from this request.
	Headers map[string]interface{} `json:"headers,omitempty"` // An object with keys as header names matched to their values. Values can either be a single string or an array of strings.
	Url string `json:"url,omitempty"` // The URL to make a request to for this step. This may contain both query string parameters and variables.
	Variables []Variable `json:"variables,omitempty"` // A list of variables to extract out of the HTTP response from this request.
	Auth map[string]interface{} `json:"auth,omitempty"` // An authentication object with either basic, oauth1, or client_certificate credentials for authenticating this request.
	Before_scripts []string `json:"before_scripts,omitempty"` // A list of pre-request scripts to run before this request.
	Note string `json:"note,omitempty"` // A description or note for this request step.
	Body string `json:"body,omitempty"` // A string to use as the body of the request.
	Form map[string]interface{} `json:"form,omitempty"` // An object with keys as form post parameter names matched to their values. Values can either be a single string or an array of strings.
	Step_type string `json:"step_type,omitempty"` // Type of test step -- request, pause, condition, ghost-inspector, or subtest.
}

// Agent represents the Agent schema from the OpenAPI specification
type Agent struct {
	Version string `json:"version,omitempty"` // The version for this agent.
	Agent_id string `json:"agent_id,omitempty"` // The unique identifier for this agent.
	Name string `json:"name,omitempty"` // The name of the agent set in the configuration file or with the command line flag.
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	TypeField string `json:"type,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
}

// Assertion represents the Assertion schema from the OpenAPI specification
type Assertion struct {
	Source string `json:"source,omitempty"`
	Value string `json:"value,omitempty"`
	Comparison string `json:"comparison,omitempty"`
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Status string `json:"status,omitempty"` // Success or failure status of call.
}

// Metrics represents the Metrics schema from the OpenAPI specification
type Metrics struct {
	This_time_period map[string]interface{} `json:"this_time_period,omitempty"` // The average response time for different percentiles for the request in the requested timeframe.
	Timeframe string `json:"timeframe,omitempty"` // The timeframe that filters this request.
	Changes_from_last_period map[string]interface{} `json:"changes_from_last_period,omitempty"` // The changes in average response time compared to the last period.
	Environment_uuid string `json:"environment_uuid,omitempty"` // The environment_uuid that filters this request.
	Region string `json:"region,omitempty"` // The region that filters this request.
	Response_times []interface{} `json:"response_times,omitempty"` // The list of response times based on the timeframe of the request.
}

// NewBucket represents the NewBucket schema from the OpenAPI specification
type NewBucket struct {
	Team_id string `json:"team_id"` // Unique identifier for the team to create this bucket for.
	Name string `json:"name"` // Name of this bucket
}

// TestDetail represents the TestDetail schema from the OpenAPI specification
type TestDetail struct {
	Version string `json:"version,omitempty"`
	Environments Environment `json:"environments,omitempty"`
	Exported_at int `json:"exported_at,omitempty"`
	Last_run map[string]interface{} `json:"last_run,omitempty"`
	Schedules []Schedule `json:"schedules,omitempty"`
	Steps []map[string]interface{} `json:"steps,omitempty"`
	Id string `json:"id,omitempty"`
	Last_run map[string]interface{} `json:"last_run,omitempty"`
	Name string `json:"name"` // The name for the test.
	Trigger_url string `json:"trigger_url,omitempty"`
	Created_at int `json:"created_at,omitempty"` // The date the test was created in seconds (Unix time stamp format).
	Created_by map[string]interface{} `json:"created_by,omitempty"`
	Default_environment_id string `json:"default_environment_id,omitempty"`
	Description string `json:"description,omitempty"` // The description for the test.
}

// Error400 represents the Error400 schema from the OpenAPI specification
type Error400 struct {
	ErrorField string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	More_info string `json:"more_info,omitempty"`
	Status int `json:"status,omitempty"`
}
