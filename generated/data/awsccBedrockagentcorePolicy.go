package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcorePolicy = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the policy was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "definition": {
        "computed": true,
        "description": "The definition structure for policies. Encapsulates different policy formats.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cedar": {
              "computed": true,
              "description": "A Cedar policy statement within the AgentCore Policy system.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "statement": {
                    "computed": true,
                    "description": "The Cedar policy statement that defines the authorization logic.",
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
      "description": {
        "computed": true,
        "description": "A human-readable description of the policy's purpose and functionality.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The customer-assigned immutable name for the policy. Must be unique within the policy engine.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_engine_id": {
        "computed": true,
        "description": "The identifier of the policy engine which contains this policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description": "The unique identifier for the policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reasons": {
        "computed": true,
        "description": "Additional information about the policy status.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the policy was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "validation_mode": {
        "computed": true,
        "description": "The validation mode for the policy. Determines how Cedar analyzer validation results are handled.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::BedrockAgentCore::Policy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcorePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcorePolicy), &result)
	return &result
}
