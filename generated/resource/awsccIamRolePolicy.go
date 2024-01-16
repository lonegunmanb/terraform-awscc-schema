package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamRolePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "The policy document.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_name": {
        "description": "The friendly name (not ARN) identifying the policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_name": {
        "description": "The name of the policy document.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Schema for IAM Role Policy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIamRolePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamRolePolicy), &result)
	return &result
}
