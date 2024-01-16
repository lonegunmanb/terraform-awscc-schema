package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMemorydbCluster = `{
  "block": {
    "attributes": {
      "acl_name": {
        "description": "The name of the Access Control List to associate with the cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description": "A flag that enables automatic minor version upgrade when set to true.\n\nYou cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cluster_endpoint": {
        "computed": true,
        "description": "The cluster endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "The DNS address of the primary read-write node.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port number that the engine is listening on. ",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "cluster_name": {
        "description": "The name of the cluster. This value must be unique as it also serves as the cluster identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_tiering": {
        "computed": true,
        "description": "Enables data tiering. Data tiering is only supported for clusters using the r6gd node type. This parameter must be set when using r6gd nodes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The Redis engine version used by the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "final_snapshot_name": {
        "computed": true,
        "description": "The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The ID of the KMS key used to encrypt the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maintenance_window": {
        "computed": true,
        "description": "Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_type": {
        "description": "The compute and memory capacity of the nodes in the cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "num_replicas_per_shard": {
        "computed": true,
        "description": "The number of replicas to apply to each shard. The limit is 5.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "num_shards": {
        "computed": true,
        "description": "The number of shards the cluster will contain.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "parameter_group_name": {
        "computed": true,
        "description": "The name of the parameter group associated with the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameter_group_status": {
        "computed": true,
        "description": "The status of the parameter group used by the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port number on which each member of the cluster accepts connections.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "security_group_ids": {
        "computed": true,
        "description": "One or more Amazon VPC security groups associated with this cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "snapshot_arns": {
        "computed": true,
        "description": "A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "snapshot_name": {
        "computed": true,
        "description": "The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_retention_limit": {
        "computed": true,
        "description": "The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "snapshot_window": {
        "computed": true,
        "description": "The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sns_topic_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sns_topic_status": {
        "computed": true,
        "description": "The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the cluster. For example, Available, Updating, Creating.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_group_name": {
        "computed": true,
        "description": "The name of the subnet group to be used for the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key for the tag. May not be null.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The tag's value. May be null.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "tls_enabled": {
        "computed": true,
        "description": "A flag that enables in-transit encryption when set to true.\n\nYou cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "The AWS::MemoryDB::Cluster resource creates an Amazon MemoryDB Cluster.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMemorydbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMemorydbCluster), &result)
	return &result
}
