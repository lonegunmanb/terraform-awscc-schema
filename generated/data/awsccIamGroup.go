package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Arn of the group to create",
        "description_kind": "plain",
        "type": "string"
      },
      "group_name": {
        "computed": true,
        "description": "The name of the group to create",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_policy_arns": {
        "computed": true,
        "description": "A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role. ",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "path": {
        "computed": true,
        "description": "The path to the group",
        "description_kind": "plain",
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description": "Adds or updates an inline policy document that is embedded in the specified IAM group",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy_document": {
              "computed": true,
              "description": "The policy document.",
              "description_kind": "plain",
              "type": "string"
            },
            "policy_name": {
              "computed": true,
              "description": "The friendly name (not ARN) identifying the policy.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::IAM::Group",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIamGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamGroup), &result)
	return &result
}
