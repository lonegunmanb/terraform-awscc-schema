package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description": "The policy document",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_name": {
        "description": "A name for resource policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The resource schema for AWSLogs ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsResourcePolicy), &result)
	return &result
}
