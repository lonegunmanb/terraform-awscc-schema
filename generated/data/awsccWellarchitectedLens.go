package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedLens = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the lens.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "json_string": {
        "computed": true,
        "description": "The JSON representation of a lens.",
        "description_kind": "plain",
        "type": "string"
      },
      "lens_arn": {
        "computed": true,
        "description": "The ARN of the lens.",
        "description_kind": "plain",
        "type": "string"
      },
      "lens_id": {
        "computed": true,
        "description": "The unique identifier of the lens.",
        "description_kind": "plain",
        "type": "string"
      },
      "lens_version": {
        "computed": true,
        "description": "The version of the lens.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The full name of the lens.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description": "The Amazon Web Services account ID that owns the lens.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the lens.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::WellArchitected::Lens",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWellarchitectedLensSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedLens), &result)
	return &result
}
