package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDocdbelasticCluster = `{
  "block": {
    "attributes": {
      "admin_user_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "admin_user_password": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auth_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
        "computed": true,
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
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "shard_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "shard_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::DocDBElastic::Cluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDocdbelasticClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDocdbelasticCluster), &result)
	return &result
}
