package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2PlacementGroup = `{
  "block": {
    "attributes": {
      "group_name": {
        "computed": true,
        "description": "The Group Name of Placement Group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "partition_count": {
        "computed": true,
        "description": "The number of partitions. Valid only when **Strategy** is set to ` + "`" + `partition` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "spread_level": {
        "computed": true,
        "description": "The Spread Level of Placement Group is an enum where it accepts either host or rack when strategy is spread",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "strategy": {
        "computed": true,
        "description": "The placement strategy.",
        "description_kind": "plain",
        "optional": true,
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "Resource Type definition for AWS::EC2::PlacementGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2PlacementGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2PlacementGroup), &result)
	return &result
}
