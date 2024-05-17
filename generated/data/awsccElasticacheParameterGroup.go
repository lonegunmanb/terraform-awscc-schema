package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticacheParameterGroup = `{
  "block": {
    "attributes": {
      "cache_parameter_group_family": {
        "computed": true,
        "description": "The name of the cache parameter group family that this cache parameter group is compatible with.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_parameter_group_name": {
        "computed": true,
        "description": "The name of the Cache Parameter Group.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description for this cache parameter group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "properties": {
        "computed": true,
        "description": "A comma-delimited list of parameter name/value pairs. For more information see ModifyCacheParameterGroup in the Amazon ElastiCache API Reference Guide.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted.",
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
    "description": "Data Source schema for AWS::ElastiCache::ParameterGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticacheParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticacheParameterGroup), &result)
	return &result
}
