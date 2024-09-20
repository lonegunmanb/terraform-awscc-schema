package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccB2BiProfile = `{
  "block": {
    "attributes": {
      "business_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "email": {
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
      "log_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "logging": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "phone": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::B2BI::Profile Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccB2BiProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccB2BiProfile), &result)
	return &result
}
