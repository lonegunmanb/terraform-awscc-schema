package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbShardGroup = `{
  "block": {
    "attributes": {
      "compute_redundancy": {
        "computed": true,
        "description": "Specifies whether to create standby instances for the DB shard group.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "db_cluster_identifier": {
        "description": "The name of the primary DB cluster for the DB shard group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_shard_group_identifier": {
        "computed": true,
        "description": "The name of the DB shard group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_shard_group_resource_id": {
        "computed": true,
        "description": "The Amazon Web Services Region-unique, immutable identifier for the DB shard group.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "The connection endpoint for the DB shard group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_acu": {
        "description": "The maximum capacity of the DB shard group in Aurora capacity units (ACUs).",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "min_acu": {
        "computed": true,
        "description": "The minimum capacity of the DB shard group in Aurora capacity units (ACUs).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Indicates whether the DB shard group is publicly accessible.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The AWS::RDS::DBShardGroup resource creates an Amazon Aurora Limitless DB Shard Group.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsDbShardGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbShardGroup), &result)
	return &result
}
