package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightLimitsProfile = `{
  "block": {
    "attributes": {
      "account_id": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "profile_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_limits": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "unit": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "map"
        },
        "required": true
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::QuickSight::LimitsProfile Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightLimitsProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightLimitsProfile), &result)
	return &result
}
