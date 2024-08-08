package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbSubnetGroup = `{
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
        "description": "The name for the DB subnet group. This value is stored as a lowercase string.\n Constraints:\n  +  Must contain no more than 255 letters, numbers, periods, underscores, spaces, or hyphens.\n  +  Must not be default.\n  +  First character must be a letter.\n  \n Example: ` + "`" + `` + "`" + `mydbsubnetgroup` + "`" + `` + "`" + `",
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
        "description": "The EC2 Subnet IDs for the DB subnet group.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the DB subnet group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
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
    "description": "The ` + "`" + `` + "`" + `AWS::RDS::DBSubnetGroup` + "`" + `` + "`" + ` resource creates a database subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same region. \n For more information, see [Working with DB subnet groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html#USER_VPC.Subnets) in the *Amazon RDS User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsDbSubnetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbSubnetGroup), &result)
	return &result
}
