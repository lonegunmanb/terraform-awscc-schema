package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApsResourcePolicy = `{
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
        "description": "The JSON to use as the Resource-based Policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_arn": {
        "computed": true,
        "description": "The Arn of an APS Workspace that the PolicyDocument will be attached to.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::APS::ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApsResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApsResourcePolicy), &result)
	return &result
}
