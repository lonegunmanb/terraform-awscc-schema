package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNeptuneDbSubnetGroup = `{
  "block": {
    "attributes": {
      "db_subnet_group_description": {
        "description": "The description for the DB subnet group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description": "The name for the DB subnet group. This value is stored as a lowercase string.\n\nConstraints: Must contain no more than 255 lowercase alphanumeric characters or hyphens. Must not be \"Default\".\n\nExample: mysubnetgroup\n\n",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "description": "The Amazon EC2 subnet IDs for the DB subnet group.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "An optional array of key-value pairs to apply to this DB subnet group.",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "The AWS::Neptune::DBSubnetGroup type creates an Amazon Neptune DB subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same AWS Region.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNeptuneDbSubnetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNeptuneDbSubnetGroup), &result)
	return &result
}
