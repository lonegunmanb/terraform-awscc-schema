package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoCimdPolicy = `{
  "block": {
    "attributes": {
      "cache_duration": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "cimd_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_id": {
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
      "policy_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "url_patterns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::CIMDPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoCimdPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoCimdPolicy), &result)
	return &result
}
