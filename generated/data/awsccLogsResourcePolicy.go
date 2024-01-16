package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "The policy document",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_name": {
        "computed": true,
        "description": "A name for resource policy",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Logs::ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsResourcePolicy), &result)
	return &result
}
