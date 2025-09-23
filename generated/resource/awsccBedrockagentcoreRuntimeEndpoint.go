package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreRuntimeEndpoint = `{
  "block": {
    "attributes": {
      "agent_runtime_arn": {
        "computed": true,
        "description": "The ARN of the Agent Runtime",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_runtime_endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AgentCore Runtime.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_runtime_id": {
        "description": "The ID of the parent Agent Runtime (required for creation)",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "agent_runtime_version": {
        "computed": true,
        "description": "The version of the AgentCore Runtime to use for the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the Agent Runtime Endpoint was created",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the AgentCore Runtime endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_reason": {
        "computed": true,
        "description": "The reason for failure if the endpoint is in a failed state",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp when the Agent Runtime Endpoint was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "live_version": {
        "computed": true,
        "description": "The Live version of the Agent Runtime",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the Agent Runtime Endpoint",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "runtime_endpoint_id": {
        "computed": true,
        "description": "The unique identifier of the AgentCore Runtime endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Agent Runtime Endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of tag keys and values",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target_version": {
        "computed": true,
        "description": "The target version of the AgentCore Runtime for the endpoint.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource definition for AWS::BedrockAgentCore::RuntimeEndpoint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreRuntimeEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreRuntimeEndpoint), &result)
	return &result
}
