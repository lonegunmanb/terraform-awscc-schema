package data

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
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description": "Arn of the resource to which the policy statement is being attached.",
        "description_kind": "plain",
        "type": "string"
      },
      "condition": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "effect": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "statement_id": {
        "computed": true,
        "description": "The Statement Id of the policy statement that is being attached.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EntityResolution::PolicyStatement",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEntityresolutionPolicyStatementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEntityresolutionPolicyStatement), &result)
	return &result
}
