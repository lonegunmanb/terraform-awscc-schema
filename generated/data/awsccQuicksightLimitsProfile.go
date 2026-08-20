package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightLimitsProfile = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "profile_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_limits": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "unit": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "map"
        }
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QuickSight::LimitsProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightLimitsProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightLimitsProfile), &result)
	return &result
}
