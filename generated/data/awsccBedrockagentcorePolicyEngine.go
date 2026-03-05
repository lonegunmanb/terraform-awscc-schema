package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcorePolicyEngine = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the policy engine was created",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A human-readable description of the policy engine's purpose and scope",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_key_arn": {
        "computed": true,
        "description": "The ARN of the KMS key used to encrypt the policy engine data",
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
        "description": "The customer-assigned immutable name for the policy engine",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_engine_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the policy engine",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_engine_id": {
        "computed": true,
        "description": "The unique identifier for the policy engine",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the policy engine",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reasons": {
        "computed": true,
        "description": "Additional information about the policy engine status",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to assign to the policy engine.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the policy engine was last updated",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::BedrockAgentCore::PolicyEngine",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcorePolicyEngineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcorePolicyEngine), &result)
	return &result
}
