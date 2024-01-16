package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticacheSubnetGroup = `{
  "block": {
    "attributes": {
      "cache_subnet_group_name": {
        "computed": true,
        "description": "The name for the cache subnet group. This value is stored as a lowercase string.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description for the cache subnet group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "The EC2 subnet IDs for the cache subnet group.",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::ElastiCache::SubnetGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticacheSubnetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticacheSubnetGroup), &result)
	return &result
}
