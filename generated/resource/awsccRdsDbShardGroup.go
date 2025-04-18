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
        "description": "Specifies whether to create standby DB shard groups for the DB shard group. Valid values are the following:\n  +  0 - Creates a DB shard group without a standby DB shard group. This is the default value.\n  +  1 - Creates a DB shard group with a standby DB shard group in a different Availability Zone (AZ).\n  +  2 - Creates a DB shard group with two standby DB shard groups in two different AZs.",
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
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
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
        "description": "Specifies whether the DB shard group is publicly accessible.\n When the DB shard group is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB shard group's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB shard group's VPC. Access to the DB shard group is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB shard group doesn't permit it.\n When the DB shard group isn't publicly accessible, it is an internal DB shard group with a DNS name that resolves to a private IP address.\n Default: The default behavior varies depending on whether ` + "`" + `` + "`" + `DBSubnetGroupName` + "`" + `` + "`" + ` is specified.\n If ` + "`" + `` + "`" + `DBSubnetGroupName` + "`" + `` + "`" + ` isn't specified, and ` + "`" + `` + "`" + `PubliclyAccessible` + "`" + `` + "`" + ` isn't specified, the following applies:\n  +  If the default VPC in the target Region doesn?t have an internet gateway attached to it, the DB shard group is private.\n  +  If the default VPC in the target Region has an internet gateway attached to it, the DB shard group is public.\n  \n If ` + "`" + `` + "`" + `DBSubnetGroupName` + "`" + `` + "`" + ` is specified, and ` + "`" + `` + "`" + `PubliclyAccessible` + "`" + `` + "`" + ` isn't specified, the following applies:\n  +  If the subnets are part of a VPC that doesn?t have an internet gateway attached to it, the DB shard group is private.\n  +  If the subnets are part of a VPC that has an internet gateway attached to it, the DB shard group is public.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "An optional set of key-value pairs to associate arbitrary data of your choosing with the DB shard group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "optional": true,
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Creates a new DB shard group for Aurora Limitless Database. You must enable Aurora Limitless Database to create a DB shard group.\n Valid for: Aurora DB clusters only",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsDbShardGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbShardGroup), &result)
	return &result
}
