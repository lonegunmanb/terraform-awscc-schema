package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectSecurityKey = `{
  "block": {
    "attributes": {
      "association_id": {
        "computed": true,
        "description": "An associationID is automatically generated when a storage config is associated with an instance",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "Amazon Connect instance identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "key": {
        "computed": true,
        "description": "A valid security key in PEM format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::SecurityKey",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectSecurityKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectSecurityKey), &result)
	return &result
}
