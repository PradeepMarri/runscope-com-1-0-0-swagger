package main

import (
	"github.com/runscope-api/mcp-server/config"
	"github.com/runscope-api/mcp-server/models"
	tools_messages "github.com/runscope-api/mcp-server/tools/messages"
	tools_account "github.com/runscope-api/mcp-server/tools/account"
	tools_tests "github.com/runscope-api/mcp-server/tools/tests"
	tools_test_environments "github.com/runscope-api/mcp-server/tools/test_environments"
	tools_shared_environments "github.com/runscope-api/mcp-server/tools/shared_environments"
	tools_test_steps "github.com/runscope-api/mcp-server/tools/test_steps"
	tools_buckets "github.com/runscope-api/mcp-server/tools/buckets"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_messages.CreatePost_buckets_bucketkey_messagesTool(cfg),
		tools_messages.CreateDelete_buckets_bucketkey_messagesTool(cfg),
		tools_messages.CreateGet_buckets_bucketkey_messagesTool(cfg),
		tools_account.CreateGet_teams_teamid_peopleTool(cfg),
		tools_account.CreateGet_accountTool(cfg),
		tools_tests.CreateDelete_buckets_bucketkey_tests_testidTool(cfg),
		tools_tests.CreateGet_buckets_bucketkey_tests_testidTool(cfg),
		tools_tests.CreatePut_buckets_bucketkey_tests_testidTool(cfg),
		tools_test_environments.CreateGet_buckets_bucketkey_tests_testid_environmentsTool(cfg),
		tools_test_environments.CreatePost_buckets_bucketkey_tests_testid_environmentsTool(cfg),
		tools_shared_environments.CreatePut_buckets_bucketkey_environments_environmentidTool(cfg),
		tools_account.CreateGet_teams_teamid_integrationsTool(cfg),
		tools_shared_environments.CreateGet_buckets_bucketkey_environmentsTool(cfg),
		tools_shared_environments.CreatePost_buckets_bucketkey_environmentsTool(cfg),
		tools_test_steps.CreateGet_buckets_bucketkey_tests_testid_stepsTool(cfg),
		tools_test_steps.CreatePost_buckets_bucketkey_tests_testid_stepsTool(cfg),
		tools_messages.CreateGet_buckets_bucketkey_errorsTool(cfg),
		tools_buckets.CreateGet_bucketsTool(cfg),
		tools_buckets.CreatePost_bucketsTool(cfg),
		tools_tests.CreateGet_buckets_bucketkey_tests_testid_metricsTool(cfg),
		tools_tests.CreateGet_buckets_bucketkey_testsTool(cfg),
		tools_tests.CreatePost_buckets_bucketkey_testsTool(cfg),
		tools_test_steps.CreateDelete_buckets_bucketkey_tests_testid_steps_stepidTool(cfg),
		tools_test_steps.CreatePut_buckets_bucketkey_tests_testid_steps_stepidTool(cfg),
		tools_account.CreateGet_teams_teamid_agentsTool(cfg),
		tools_messages.CreateGet_buckets_bucketkey_messages_messageidTool(cfg),
		tools_buckets.CreateGet_buckets_bucketkeyTool(cfg),
		tools_buckets.CreateDelete_buckets_bucketkeyTool(cfg),
		tools_test_environments.CreatePut_buckets_bucketkey_tests_testid_environments_environmentidTool(cfg),
	}
}
