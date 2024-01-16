package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessWirelessDeviceImportTask = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Arn for Wireless Device Import Task, Returned upon successful start.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date": {
        "computed": true,
        "description": "CreationDate for import task",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_name": {
        "description": "Destination Name for import task",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "failed_imported_devices_count": {
        "computed": true,
        "description": "Failed Imported Devices Count",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Id for Wireless Device Import Task, Returned upon successful start.",
        "description_kind": "plain",
        "type": "string"
      },
      "initialized_imported_devices_count": {
        "computed": true,
        "description": "Initialized Imported Devices Count",
        "description_kind": "plain",
        "type": "number"
      },
      "onboarded_imported_devices_count": {
        "computed": true,
        "description": "Onboarded Imported Devices Count",
        "description_kind": "plain",
        "type": "number"
      },
      "pending_imported_devices_count": {
        "computed": true,
        "description": "Pending Imported Devices Count",
        "description_kind": "plain",
        "type": "number"
      },
      "sidewalk": {
        "description": "sidewalk contain file for created device and role",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "device_creation_file": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_creation_file_list": {
              "computed": true,
              "description": "sidewalk create device's file path",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "role": {
              "computed": true,
              "description": "sidewalk role",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sidewalk_manufacturing_sn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "status": {
        "computed": true,
        "description": "Status for import task",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "StatusReason for import task",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Wireless Device Import Tasks",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotwirelessWirelessDeviceImportTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessWirelessDeviceImportTask), &result)
	return &result
}
