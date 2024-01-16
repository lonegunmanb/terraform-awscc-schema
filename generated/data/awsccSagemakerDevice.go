package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerDevice = `{
  "block": {
    "attributes": {
      "device": {
        "computed": true,
        "description": "The Edge Device you want to register against a device fleet",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "Description of the device",
              "description_kind": "plain",
              "type": "string"
            },
            "device_name": {
              "computed": true,
              "description": "The name of the device",
              "description_kind": "plain",
              "type": "string"
            },
            "iot_thing_name": {
              "computed": true,
              "description": "AWS Internet of Things (IoT) object name.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "device_fleet_name": {
        "computed": true,
        "description": "The name of the edge device fleet",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Associate tags with the resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::Device",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerDeviceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerDevice), &result)
	return &result
}
