package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailInstanceSnapshot = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "from_instance_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the instance from which the snapshot was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "from_instance_name": {
        "computed": true,
        "description": "The instance from which the snapshot was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_name": {
        "computed": true,
        "description": "The instance from which the snapshot was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_snapshot_name": {
        "computed": true,
        "description": "The name of the snapshot.",
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
        "description": "The region name and Availability Zone where you created the snapshot.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description": "The Availability Zone. Follows the format us-east-2a (case-sensitive).",
              "description_kind": "plain",
              "type": "string"
            },
            "region_name": {
              "computed": true,
              "description": "The AWS Region name.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_type": {
        "computed": true,
        "description": "The type of resource (usually InstanceSnapshot).",
        "description_kind": "plain",
        "type": "string"
      },
      "size_in_gb": {
        "computed": true,
        "description": "The size in GB of the SSD",
        "description_kind": "plain",
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "The state the snapshot is in.",
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
      }
    },
    "description": "Data Source schema for AWS::Lightsail::InstanceSnapshot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLightsailInstanceSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailInstanceSnapshot), &result)
	return &result
}
