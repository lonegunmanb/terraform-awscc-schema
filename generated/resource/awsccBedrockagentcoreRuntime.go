package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreRuntime = `{
  "block": {
    "attributes": {
      "agent_runtime_arn": {
        "computed": true,
        "description": "The Amazon Resource Name(ARN) that uniquely identifies the Agent",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_runtime_artifact": {
        "description": "The artifact of the agent",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_uri": {
                    "computed": true,
                    "description": "The ECR URI of the container",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "agent_runtime_id": {
        "computed": true,
        "description": "Identifier for a resource",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_runtime_name": {
        "description": "Name for a resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "agent_runtime_version": {
        "computed": true,
        "description": "Version of the Agent",
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_configuration": {
        "computed": true,
        "description": "Authorizer configuration for the agent runtime",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_jwt_authorizer": {
              "computed": true,
              "description": "Configuration for custom JWT authorizer",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_audience": {
                    "computed": true,
                    "description": "List of allowed audiences",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_clients": {
                    "computed": true,
                    "description": "List of allowed clients",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "discovery_url": {
                    "computed": true,
                    "description": "OpenID Connect discovery URL",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "created_at": {
        "computed": true,
        "description": "Timestamp when the Agent was created",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_variables": {
        "computed": true,
        "description": "Environment variables for the agent runtime",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "When resource was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "network_configuration": {
        "description": "Network access configuration for the Agent",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_mode": {
              "description": "Network mode configuration type",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "protocol_configuration": {
        "computed": true,
        "description": "Protocol configuration for the agent runtime",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "description": "Amazon Resource Name (ARN) of an IAM role",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Current status of the agent",
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
      "workload_identity_details": {
        "computed": true,
        "description": "Workload identity details for the agent",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "workload_identity_arn": {
              "computed": true,
              "description": "ARN of the workload identity",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Resource Type definition for AWS::BedrockAgentCore::Runtime",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreRuntimeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreRuntime), &result)
	return &result
}
