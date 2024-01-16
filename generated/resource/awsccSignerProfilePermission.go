package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSignerProfilePermission = `{
  "block": {
    "attributes": {
      "action": {
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
      "profile_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "statement_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "An example resource schema demonstrating some basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSignerProfilePermissionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSignerProfilePermission), &result)
	return &result
}
