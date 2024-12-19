package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQbusinessPermission = `{
  "block": {
    "attributes": {
      "actions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "application_id": {
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
      "principal": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "statement_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QBusiness::Permission",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQbusinessPermissionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQbusinessPermission), &result)
	return &result
}
