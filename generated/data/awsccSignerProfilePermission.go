package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSignerProfilePermission = `{
  "block": {
    "attributes": {
      "action": {
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
      "profile_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "profile_version": {
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
    "description": "Data Source schema for AWS::Signer::ProfilePermission",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSignerProfilePermissionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSignerProfilePermission), &result)
	return &result
}
