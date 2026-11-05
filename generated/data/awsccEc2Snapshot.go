package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Snapshot = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "encrypted": {
        "computed": true,
        "description": "Indicates whether the snapshot is encrypted.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the KMS key that was used to protect the volume encryption key for the parent volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "The ID of the Amazon Web Services account that owns the EBS snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "progress": {
        "computed": true,
        "description": "The progress of the snapshot, as a percentage.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_id": {
        "computed": true,
        "description": "The ID of the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description": "The time stamp when the snapshot was initiated.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The snapshot state.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_tier": {
        "computed": true,
        "description": "The storage tier in which the snapshot is stored.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the snapshot.",
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
      "volume_id": {
        "computed": true,
        "description": "The ID of the Amazon EBS volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_size": {
        "computed": true,
        "description": "The size of the volume, in GiB.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::EC2::Snapshot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2SnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Snapshot), &result)
	return &result
}
