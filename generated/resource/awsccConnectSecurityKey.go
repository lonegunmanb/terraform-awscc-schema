package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "Amazon Connect instance identifier",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key": {
        "description": "A valid security key in PEM format.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::SecurityKey",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectSecurityKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectSecurityKey), &result)
	return &result
}
