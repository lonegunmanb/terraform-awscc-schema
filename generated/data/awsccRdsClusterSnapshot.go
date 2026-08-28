package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsClusterSnapshot = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description": "The allocated storage size of the DB cluster snapshot in gibibytes (GiB).",
        "description_kind": "plain",
        "type": "number"
      },
      "availability_zones": {
        "computed": true,
        "description": "The list of Availability Zones where instances in the DB cluster snapshot can be restored.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cluster_create_time": {
        "computed": true,
        "description": "The time when the DB cluster was created, in Universal Coordinated Time (UTC).",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_identifier": {
        "computed": true,
        "description": "The identifier of the DB cluster to create a snapshot for.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_resource_id": {
        "computed": true,
        "description": "The resource ID of the DB cluster that this DB cluster snapshot was created from.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_snapshot_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_snapshot_identifier": {
        "computed": true,
        "description": "The identifier for the DB cluster snapshot. Must contain from 1 to 63 letters, numbers, or hyphens. First character must be a letter. Can't end with a hyphen or contain two consecutive hyphens.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "computed": true,
        "description": "The name of the database engine for this DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_mode": {
        "computed": true,
        "description": "The engine mode of the database engine for this DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version of the database engine for this DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "iam_database_authentication_enabled": {
        "computed": true,
        "description": "Indicates whether mapping of AWS IAM accounts to database accounts is enabled.",
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
        "description": "If StorageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description": "The license model information for this DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "master_username": {
        "computed": true,
        "description": "The master username for this DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port that the DB cluster was listening on at the time of the snapshot.",
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_create_time": {
        "computed": true,
        "description": "The time when the snapshot was taken, in Universal Coordinated Time (UTC).",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_type": {
        "computed": true,
        "description": "The type of the DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of this DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description": "Indicates whether the DB cluster snapshot is encrypted.",
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "The tags to be assigned to the DB cluster snapshot.",
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
          "nesting_mode": "list"
        }
      },
      "vpc_id": {
        "computed": true,
        "description": "The VPC ID associated with the DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::RDS::ClusterSnapshot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsClusterSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsClusterSnapshot), &result)
	return &result
}
