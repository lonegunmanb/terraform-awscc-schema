package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessNetworkAnalyzerConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Arn for network analyzer configuration, Returned upon successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the new resource",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the network analyzer configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "trace_content": {
        "computed": true,
        "description": "Trace content for your wireless gateway and wireless device resources",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_level": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "wireless_device_frame_info": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "wireless_devices": {
        "computed": true,
        "description": "List of wireless gateway resources that have been added to the network analyzer configuration",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "wireless_gateways": {
        "computed": true,
        "description": "List of wireless gateway resources that have been added to the network analyzer configuration",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::IoTWireless::NetworkAnalyzerConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotwirelessNetworkAnalyzerConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessNetworkAnalyzerConfiguration), &result)
	return &result
}
