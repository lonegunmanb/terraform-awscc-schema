package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticacheSnapshot = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN (Amazon Resource Name) of the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description": "If true, opt-in to the next auto minor version upgrade campaign.",
        "description_kind": "plain",
        "type": "bool"
      },
      "automatic_failover": {
        "computed": true,
        "description": "Indicates the status of automatic failover for the source replication group.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_cluster_create_time": {
        "computed": true,
        "description": "The date and time when the source cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_cluster_id": {
        "computed": true,
        "description": "The identifier of an existing cluster to create a snapshot from. The snapshot is created from this cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_node_type": {
        "computed": true,
        "description": "The name of the compute and memory capacity node type for the source cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_parameter_group_name": {
        "computed": true,
        "description": "The cache parameter group that is associated with the source cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_subnet_group_name": {
        "computed": true,
        "description": "The name of the cache subnet group associated with the source cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_tiering": {
        "computed": true,
        "description": "Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "computed": true,
        "description": "The name of the cache engine used by the source cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version of the cache engine version used by the source cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The ID of the KMS key used to encrypt the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "node_snapshots": {
        "computed": true,
        "description": "A list of the cache nodes in the source cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cache_cluster_id": {
              "computed": true,
              "description": "A unique identifier for the source cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "cache_node_create_time": {
              "computed": true,
              "description": "The date and time when the cache node was created in the source cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "cache_node_id": {
              "computed": true,
              "description": "The cache node identifier for the node in the source cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "cache_size": {
              "computed": true,
              "description": "The size of the cache on the source cache node.",
              "description_kind": "plain",
              "type": "string"
            },
            "node_group_id": {
              "computed": true,
              "description": "A unique identifier for the source node group (shard).",
              "description_kind": "plain",
              "type": "string"
            },
            "snapshot_create_time": {
              "computed": true,
              "description": "The date and time when the source node's metadata and cache data set was obtained for the snapshot.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "num_cache_nodes": {
        "computed": true,
        "description": "The number of cache nodes in the source cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "num_node_groups": {
        "computed": true,
        "description": "The number of node groups (shards) in this snapshot.",
        "description_kind": "plain",
        "type": "number"
      },
      "port": {
        "computed": true,
        "description": "The port number used by each cache nodes in the source cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "preferred_availability_zone": {
        "computed": true,
        "description": "The name of the Availability Zone in which the source cluster is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "Specifies the weekly time range during which maintenance on the cluster is performed.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_group_description": {
        "computed": true,
        "description": "A description of the source replication group.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_group_id": {
        "computed": true,
        "description": "The identifier of an existing replication group to create a snapshot from. The snapshot is created from this replication group.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_name": {
        "computed": true,
        "description": "The name of a snapshot. Must be unique within the customer account.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_retention_limit": {
        "computed": true,
        "description": "For an automatic snapshot, the number of days for which ElastiCache retains the snapshot before deleting it.",
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_source": {
        "computed": true,
        "description": "Indicates whether the snapshot is from an automatic backup (automated) or was created manually (manual).",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_status": {
        "computed": true,
        "description": "The status of the snapshot. Valid values: creating | available | restoring | copying | deleting | failed | deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_window": {
        "computed": true,
        "description": "The daily time range during which ElastiCache takes daily snapshots of the source cluster.",
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
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "topic_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the topic used by the source cluster for publishing notifications.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The Amazon Virtual Private Cloud identifier (VPC ID) of the cache subnet group for the source cluster.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ElastiCache::Snapshot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticacheSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticacheSnapshot), &result)
	return &result
}
