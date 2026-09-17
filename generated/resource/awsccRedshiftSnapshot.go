package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftSnapshot = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone in which the cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_create_time": {
        "computed": true,
        "description": "The time (UTC) when the cluster was originally created.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_identifier": {
        "description": "The cluster identifier for which you want a snapshot.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_version": {
        "computed": true,
        "description": "The version ID of the Amazon Redshift engine that is running on the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_name": {
        "computed": true,
        "description": "The name of the database that was created when the cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "encrypted": {
        "computed": true,
        "description": "If true, the data in the snapshot is encrypted at rest.",
        "description_kind": "plain",
        "type": "bool"
      },
      "encrypted_with_hsm": {
        "computed": true,
        "description": "A boolean that indicates whether the snapshot data is encrypted using the HSM keys of the source cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "engine_full_version": {
        "computed": true,
        "description": "The engine full version of the cluster at the time the snapshot was taken.",
        "description_kind": "plain",
        "type": "string"
      },
      "enhanced_vpc_routing": {
        "computed": true,
        "description": "An option that specifies whether to create the cluster with enhanced VPC routing enabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Key Management Service (KMS) key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.",
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_track_name": {
        "computed": true,
        "description": "The name of the maintenance track for the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "manual_snapshot_retention_period": {
        "computed": true,
        "description": "The number of days that a manual snapshot is retained. If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3653.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "master_username": {
        "computed": true,
        "description": "The admin user name for the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "node_type": {
        "computed": true,
        "description": "The node type that the cluster is provisioned with.",
        "description_kind": "plain",
        "type": "string"
      },
      "number_of_nodes": {
        "computed": true,
        "description": "The number of nodes in the cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "owner_account": {
        "computed": true,
        "description": "The Amazon Web Services account that owns the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port that the cluster is listening on.",
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_create_time": {
        "computed": true,
        "description": "The time (in UTC format) when Amazon Redshift began the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_identifier": {
        "description": "A unique identifier for the snapshot that you are requesting. This identifier must be unique for all snapshots within the Amazon Web Services account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "snapshot_type": {
        "computed": true,
        "description": "The snapshot type. Snapshots created using CreateClusterSnapshot and CopyClusterSnapshot are of type manual.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The snapshot status. The value of the status depends on the API operation used: CreateClusterSnapshot and CopyClusterSnapshot returns status as creating. DescribeClusterSnapshots returns status as creating, available, final snapshot, or failed. DeleteClusterSnapshot returns status as deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tag instances.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key, or name, for the resource tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the resource tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_id": {
        "computed": true,
        "description": "The VPC identifier of the cluster if the snapshot is from a cluster in a VPC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Redshift::Snapshot. Creates a manual snapshot of the specified cluster. The cluster must be in the available state.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftSnapshot), &result)
	return &result
}
