package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockAutomatedReasoningPolicyVersion = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "Time this policy version was created",
        "description_kind": "plain",
        "type": "string"
      },
      "definition_hash": {
        "computed": true,
        "description": "The hash for this version",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description inherited from the policy",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_definition_hash": {
        "computed": true,
        "description": "The hash for this version",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name inherited from the policy",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_arn": {
        "description": "Arn of the policy ",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description": "The id of the associated policy",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag Key",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag Value",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "Time this policy was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version of the policy",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::Bedrock::AutomatedReasoningPolicyVersion Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockAutomatedReasoningPolicyVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockAutomatedReasoningPolicyVersion), &result)
	return &result
}
