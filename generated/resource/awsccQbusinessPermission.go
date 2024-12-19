package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQbusinessPermission = `{
  "block": {
    "attributes": {
      "actions": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "application_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "statement_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::QBusiness::Permission Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQbusinessPermissionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQbusinessPermission), &result)
	return &result
}
