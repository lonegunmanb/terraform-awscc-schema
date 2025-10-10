package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailDiskSnapshot = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the disk snapshot was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "disk_name": {
        "description": "The name of the source disk from which the snapshot was created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disk_snapshot_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the disk snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "disk_snapshot_name": {
        "description": "The name of the disk snapshot (e.g., my-disk-snapshot).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "from_disk_name": {
        "computed": true,
        "description": "The name of the source disk from which the disk snapshot was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_from_auto_snapshot": {
        "computed": true,
        "description": "A Boolean value indicating whether the snapshot was created from an automatic snapshot.",
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description": "The AWS Region and Availability Zone where the disk snapshot was created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description": "The Availability Zone where the disk snapshot was created.",
              "description_kind": "plain",
              "type": "string"
            },
            "region_name": {
              "computed": true,
              "description": "The AWS Region where the disk snapshot was created.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "progress": {
        "computed": true,
        "description": "The progress of the disk snapshot creation operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The Lightsail resource type (DiskSnapshot).",
        "description_kind": "plain",
        "type": "string"
      },
      "size_in_gb": {
        "computed": true,
        "description": "The size of the disk snapshot in GB.",
        "description_kind": "plain",
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "The status of the disk snapshot operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "support_code": {
        "computed": true,
        "description": "The support code. Include this code in your email to support when you have questions about an instance or another resource in Lightsail.",
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
              "optional": true,
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
    "description": "Resource Type definition for AWS::Lightsail::DiskSnapshot",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLightsailDiskSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailDiskSnapshot), &result)
	return &result
}
