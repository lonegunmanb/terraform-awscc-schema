package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserverlessSecurityPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the policy",
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
        "description": "The name of the policy",
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description": "The JSON policy document that is the content for the policy",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The possible types for the network policy",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::OpenSearchServerless::SecurityPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchserverlessSecurityPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserverlessSecurityPolicy), &result)
	return &result
}
