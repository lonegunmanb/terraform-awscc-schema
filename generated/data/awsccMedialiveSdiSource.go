package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveSdiSource = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The unique arn of the SdiSource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "inputs": {
        "computed": true,
        "description": "The list of inputs currently using this SDI source.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "mode": {
        "computed": true,
        "description": "The current state of the SdiSource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the SdiSource.",
        "description_kind": "plain",
        "type": "string"
      },
      "sdi_source_id": {
        "computed": true,
        "description": "The unique identifier of the SdiSource.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the SdiSource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of key-value pairs.",
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
      },
      "type": {
        "computed": true,
        "description": "The interface mode of the SdiSource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::MediaLive::SdiSource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMedialiveSdiSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveSdiSource), &result)
	return &result
}
