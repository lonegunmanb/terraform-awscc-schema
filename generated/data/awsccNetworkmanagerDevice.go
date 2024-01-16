package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerDevice = `{
  "block": {
    "attributes": {
      "aws_location": {
        "computed": true,
        "description": "The Amazon Web Services location of the device, if applicable.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "subnet_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the subnet that the device is located in.",
              "description_kind": "plain",
              "type": "string"
            },
            "zone": {
              "computed": true,
              "description": "The Zone that the device is located in. Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "The date and time that the device was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the device.",
        "description_kind": "plain",
        "type": "string"
      },
      "device_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the device.",
        "description_kind": "plain",
        "type": "string"
      },
      "device_id": {
        "computed": true,
        "description": "The ID of the device.",
        "description_kind": "plain",
        "type": "string"
      },
      "global_network_id": {
        "computed": true,
        "description": "The ID of the global network.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "computed": true,
        "description": "The site location.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "The physical address.",
              "description_kind": "plain",
              "type": "string"
            },
            "latitude": {
              "computed": true,
              "description": "The latitude.",
              "description_kind": "plain",
              "type": "string"
            },
            "longitude": {
              "computed": true,
              "description": "The longitude.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "model": {
        "computed": true,
        "description": "The device model",
        "description_kind": "plain",
        "type": "string"
      },
      "serial_number": {
        "computed": true,
        "description": "The device serial number.",
        "description_kind": "plain",
        "type": "string"
      },
      "site_id": {
        "computed": true,
        "description": "The site ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the device.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the device.",
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
      "type": {
        "computed": true,
        "description": "The device type.",
        "description_kind": "plain",
        "type": "string"
      },
      "vendor": {
        "computed": true,
        "description": "The device vendor.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkManager::Device",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkmanagerDeviceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerDevice), &result)
	return &result
}
