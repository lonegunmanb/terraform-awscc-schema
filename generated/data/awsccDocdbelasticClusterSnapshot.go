package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDocdbelasticClusterSnapshot = `{
  "block": {
    "attributes": {
      "admin_user_name": {
        "computed": true,
        "description": "The name of the elastic cluster administrator.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_arn": {
        "computed": true,
        "description": "The ARN of the elastic cluster of which to create a snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_creation_time": {
        "computed": true,
        "description": "The time when the source elastic cluster was created, in Universal Coordinated Time (UTC).",
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
        "description": "The KMS key identifier used to encrypt the source elastic cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the elastic cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_creation_time": {
        "computed": true,
        "description": "The time when the elastic cluster snapshot was created, in Universal Coordinated Time (UTC).",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_name": {
        "computed": true,
        "description": "The name of the elastic cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_type": {
        "computed": true,
        "description": "The type of the elastic cluster snapshot (manual or automated).",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the elastic cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "The Amazon EC2 subnet IDs associated with the source elastic cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to the elastic cluster snapshot.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "The Amazon EC2 security group IDs associated with the source elastic cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::DocDBElastic::ClusterSnapshot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDocdbelasticClusterSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDocdbelasticClusterSnapshot), &result)
	return &result
}
