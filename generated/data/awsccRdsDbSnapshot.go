package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbSnapshot = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description": "The allocated storage size in gibibytes (GiB).",
        "description_kind": "plain",
        "type": "number"
      },
      "availability_zone": {
        "computed": true,
        "description": "The name of the Availability Zone the DB instance was located in at the time of the DB snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_identifier": {
        "computed": true,
        "description": "The identifier of the DB instance that you want to create the snapshot of.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_snapshot_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the DB snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_snapshot_identifier": {
        "computed": true,
        "description": "The identifier for the DB snapshot. Must contain from 1 to 255 letters, numbers, or hyphens. First character must be a letter. Can't end with a hyphen or contain two consecutive hyphens.",
        "description_kind": "plain",
        "type": "string"
      },
      "dbi_resource_id": {
        "computed": true,
        "description": "The identifier for the source DB instance, which is unique to an AWS Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "encrypted": {
        "computed": true,
        "description": "Indicates whether the DB snapshot is encrypted.",
        "description_kind": "plain",
        "type": "bool"
      },
      "engine": {
        "computed": true,
        "description": "The name of the database engine.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version of the database engine.",
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
      "instance_create_time": {
        "computed": true,
        "description": "The time when the DB instance was created, in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "iops": {
        "computed": true,
        "description": "The Provisioned IOPS value of the DB instance at the time of the snapshot.",
        "description_kind": "plain",
        "type": "number"
      },
      "kms_key_id": {
        "computed": true,
        "description": "If Encrypted is true, the AWS KMS key identifier for the encrypted DB snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description": "License model information for the restored DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "master_username": {
        "computed": true,
        "description": "The master username for the DB snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "option_group_name": {
        "computed": true,
        "description": "The option group name for the DB snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "original_snapshot_create_time": {
        "computed": true,
        "description": "The time of the CreateDBSnapshot operation in UTC. Doesn't change when the snapshot is copied.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port that the database engine was listening on at the time of the snapshot.",
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_create_time": {
        "computed": true,
        "description": "The time when the snapshot was taken, in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_type": {
        "computed": true,
        "description": "The type of the DB snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of this DB snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_throughput": {
        "computed": true,
        "description": "The storage throughput for the DB snapshot.",
        "description_kind": "plain",
        "type": "number"
      },
      "storage_type": {
        "computed": true,
        "description": "The storage type associated with the DB snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to be assigned to the DB snapshot.",
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
        "description": "The VPC ID associated with the DB snapshot.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::RDS::DBSnapshot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsDbSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbSnapshot), &result)
	return &result
}
