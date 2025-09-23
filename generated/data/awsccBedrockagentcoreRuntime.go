package data

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
        "computed": true,
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
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "agent_runtime_id": {
        "computed": true,
        "description": "Identifier for a resource",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_runtime_name": {
        "computed": true,
        "description": "Name for a resource",
        "description_kind": "plain",
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
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_clients": {
                    "computed": true,
                    "description": "List of allowed clients",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "discovery_url": {
                    "computed": true,
                    "description": "OpenID Connect discovery URL",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
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
        "type": "string"
      },
      "environment_variables": {
        "computed": true,
        "description": "Environment variables for the agent runtime",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "When resource was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "network_configuration": {
        "computed": true,
        "description": "Network access configuration for the Agent",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_mode": {
              "computed": true,
              "description": "Network mode configuration type",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "protocol_configuration": {
        "computed": true,
        "description": "Protocol configuration for the agent runtime",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of an IAM role",
        "description_kind": "plain",
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
    "description": "Data Source schema for AWS::BedrockAgentCore::Runtime",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcoreRuntimeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreRuntime), &result)
	return &result
}
