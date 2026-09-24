package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccStoragegatewayVolume = `{
  "block": {
    "attributes": {
      "created_date": {
        "computed": true,
        "description": "The date the volume was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the gateway on which to create the volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_encrypted": {
        "computed": true,
        "description": "Set to true to use Amazon S3 server-side encryption with your own KMS key, or false to use a key managed by Amazon S3.",
        "description_kind": "plain",
        "type": "bool"
      },
      "kms_key": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for Amazon S3 server-side encryption.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description": "The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_id": {
        "computed": true,
        "description": "The snapshot ID of the snapshot to restore as the new cached volume (e.g., snap-1122aabb).",
        "description_kind": "plain",
        "type": "string"
      },
      "source_volume_arn": {
        "computed": true,
        "description": "The ARN of an existing volume from which to create the new volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of up to 50 tags to assign to the volume.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the volume target, which includes the iSCSI target name.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_name": {
        "computed": true,
        "description": "The name of the iSCSI target used by an initiator to connect to a volume and used as a suffix for the target ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the storage volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_attachment_status": {
        "computed": true,
        "description": "Indicates whether the storage volume is attached to or detached from the gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_id": {
        "computed": true,
        "description": "The unique identifier of the volume, extracted from the ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_size_in_bytes": {
        "computed": true,
        "description": "The size of the volume in bytes.",
        "description_kind": "plain",
        "type": "number"
      },
      "volume_status": {
        "computed": true,
        "description": "The status of the storage volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_type": {
        "computed": true,
        "description": "The type of the volume (CACHED iSCSI).",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_used_in_bytes": {
        "computed": true,
        "description": "The size of the data stored on the volume in bytes.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::StorageGateway::Volume",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccStoragegatewayVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccStoragegatewayVolume), &result)
	return &result
}
