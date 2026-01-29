package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserverlessCollectionGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the collection group.",
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_limits": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_indexing_capacity_in_ocu": {
              "computed": true,
              "description": "The maximum indexing capacity for collections in the group.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_search_capacity_in_ocu": {
              "computed": true,
              "description": "The maximum search capacity for collections in the group.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_indexing_capacity_in_ocu": {
              "computed": true,
              "description": "The minimum indexing capacity for collections in the group.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_search_capacity_in_ocu": {
              "computed": true,
              "description": "The minimum search capacity for collections in the group.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "collection_group_id": {
        "computed": true,
        "description": "The unique identifier of the collection group.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the collection group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the collection group.",
        "description_kind": "plain",
        "type": "string"
      },
      "standby_replicas": {
        "computed": true,
        "description": "Indicates whether standby replicas are used for the collection group.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key in the key-value pair",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value in the key-value pair",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::OpenSearchServerless::CollectionGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchserverlessCollectionGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserverlessCollectionGroup), &result)
	return &result
}
