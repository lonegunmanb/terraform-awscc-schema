package resource

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
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "The description for the cache subnet group.",
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
      "subnet_ids": {
        "description": "The EC2 subnet IDs for the cache subnet group.",
        "description_kind": "plain",
        "required": true,
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::ElastiCache::SubnetGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccElasticacheSubnetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticacheSubnetGroup), &result)
	return &result
}
