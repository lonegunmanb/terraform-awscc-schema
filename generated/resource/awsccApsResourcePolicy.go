package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApsResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description": "The JSON to use as the Resource-based Policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workspace_arn": {
        "description": "The Arn of an APS Workspace that the PolicyDocument will be attached to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::APS::ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApsResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApsResourcePolicy), &result)
	return &result
}
