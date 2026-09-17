package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMemorydbSnapshot = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN (Amazon Resource Name) of the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_configuration": {
        "computed": true,
        "description": "The configuration of the cluster from which the snapshot was taken.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "The description of the cluster configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "engine": {
              "computed": true,
              "description": "The name of the engine used by the cluster configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "engine_version": {
              "computed": true,
              "description": "The Redis OSS engine version used by the cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "maintenance_window": {
              "computed": true,
              "description": "The specified maintenance window for the cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "node_type": {
              "computed": true,
              "description": "The node type used for the cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "num_shards": {
              "computed": true,
              "description": "The number of shards in the cluster.",
              "description_kind": "plain",
              "type": "number"
            },
            "parameter_group_name": {
              "computed": true,
              "description": "The name of parameter group used by the cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port used by the cluster.",
              "description_kind": "plain",
              "type": "number"
            },
            "snapshot_retention_limit": {
              "computed": true,
              "description": "The snapshot retention limit set by the cluster.",
              "description_kind": "plain",
              "type": "number"
            },
            "snapshot_window": {
              "computed": true,
              "description": "The snapshot window set by the cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "subnet_group_name": {
              "computed": true,
              "description": "The name of the subnet group used by the cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "topic_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the SNS notification topic for the cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_id": {
              "computed": true,
              "description": "The ID of the VPC the cluster belongs to.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "cluster_name": {
        "description": "The name of the cluster from which the snapshot was taken.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_tiering": {
        "computed": true,
        "description": "Enables data tiering. Data tiering is only supported for clusters using the r6gd node type.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The ID of the KMS key used to encrypt the snapshot.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_name": {
        "description": "The name of the snapshot.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source": {
        "computed": true,
        "description": "Indicates whether the snapshot is from an automatic backup (automated) or was created manually (manual).",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the snapshot. Valid values: creating | available | restoring | copying | deleting.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to be added to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
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
    "description": "Resource Type definition for AWS::MemoryDB::Snapshot. Represents a copy of an entire cluster as of the time when the snapshot was taken.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMemorydbSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMemorydbSnapshot), &result)
	return &result
}
