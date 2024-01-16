package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailDisk = `{
  "block": {
    "attributes": {
      "add_ons": {
        "computed": true,
        "description": "An array of objects representing the add-ons to enable for the new instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "add_on_type": {
              "description": "The add-on type",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "auto_snapshot_add_on_request": {
              "computed": true,
              "description": "An object that represents additional parameters when enabling or modifying the automatic snapshot add-on",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "snapshot_time_of_day": {
                    "computed": true,
                    "description": "The daily time when an automatic snapshot will be created.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "status": {
              "computed": true,
              "description": "Status of the Addon",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "attached_to": {
        "computed": true,
        "description": "Name of the attached Lightsail Instance",
        "description_kind": "plain",
        "type": "string"
      },
      "attachment_state": {
        "computed": true,
        "description": "Attachment State of the Lightsail disk",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk_name": {
        "description": "The names to use for your new Lightsail disk.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "iops": {
        "computed": true,
        "description": "Iops of the Lightsail disk",
        "description_kind": "plain",
        "type": "number"
      },
      "is_attached": {
        "computed": true,
        "description": "Check is Disk is attached state",
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description": "Location of a resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description": "The Availability Zone in which to create your disk. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
              "description_kind": "plain",
              "type": "string"
            },
            "region_name": {
              "computed": true,
              "description": "The Region Name in which to create your disk.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "path": {
        "computed": true,
        "description": "Path of the  attached Disk",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "Resource type of Lightsail instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "size_in_gb": {
        "description": "Size of the Lightsail disk",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "State of the Lightsail disk",
        "description_kind": "plain",
        "type": "string"
      },
      "support_code": {
        "computed": true,
        "description": "Support code to help identify any issues",
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
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Lightsail::Disk",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLightsailDiskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailDisk), &result)
	return &result
}
