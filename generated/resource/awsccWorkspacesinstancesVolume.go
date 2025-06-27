package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspacesinstancesVolume = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "description": "The Availability Zone in which to create the volume",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encrypted": {
        "computed": true,
        "description": "Indicates whether the volume should be encrypted",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "iops": {
        "computed": true,
        "description": "The number of I/O operations per second (IOPS)",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for Amazon EBS encryption",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "size_in_gb": {
        "computed": true,
        "description": "The size of the volume, in GiBs",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "snapshot_id": {
        "computed": true,
        "description": "The snapshot from which to create the volume",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tag_specifications": {
        "computed": true,
        "description": "The tags passed to EBS volume",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tags": {
              "computed": true,
              "description": "The tags to apply to the resource",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The key name of the tag",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value for the tag",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "throughput": {
        "computed": true,
        "description": "The throughput to provision for a volume, with a maximum of 1,000 MiB/s",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "volume_id": {
        "computed": true,
        "description": "Unique identifier for the volume",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_type": {
        "computed": true,
        "description": "The volume type",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::WorkspacesInstances::Volume - Manages WorkSpaces Volume resources",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWorkspacesinstancesVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspacesinstancesVolume), &result)
	return &result
}
