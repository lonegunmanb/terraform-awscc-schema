package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbSubnetGroup = `{
  "block": {
    "attributes": {
      "db_subnet_group_description": {
        "computed": true,
        "description": "The description for the DB subnet group.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description": "The name for the DB subnet group. This value is stored as a lowercase string.\n Constraints: Must contain no more than 255 lowercase alphanumeric characters or hyphens. Must not be \"Default\".\n Example: ` + "`" + `` + "`" + `mysubnetgroup` + "`" + `` + "`" + `",
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
        "description": "The EC2 Subnet IDs for the DB subnet group.",
        "description_kind": "plain",
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
              "description": "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::RDS::DBSubnetGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsDbSubnetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbSubnetGroup), &result)
	return &result
}
