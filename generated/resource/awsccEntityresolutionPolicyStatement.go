package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEntityresolutionPolicyStatement = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "description": "Arn of the resource to which the policy statement is being attached.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "condition": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effect": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "statement_id": {
        "description": "The Statement Id of the policy statement that is being attached.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Policy Statement defined in AWS Entity Resolution Service",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEntityresolutionPolicyStatementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEntityresolutionPolicyStatement), &result)
	return &result
}
