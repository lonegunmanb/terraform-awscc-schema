package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVpclatticeAuthPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates or updates the auth policy.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccVpclatticeAuthPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVpclatticeAuthPolicy), &result)
	return &result
}
