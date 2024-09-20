package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDocdbelasticCluster = `{
  "block": {
    "attributes": {
      "admin_user_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "admin_user_password": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auth_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "backup_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cluster_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "description_kind": "plain",
        "required": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_backup_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "shard_capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "shard_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "shard_instance_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "The AWS::DocDBElastic::Cluster Amazon DocumentDB (with MongoDB compatibility) Elastic Scale resource describes a Cluster",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDocdbelasticClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDocdbelasticCluster), &result)
	return &result
}
