package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbClusterParameterGroup = `{
  "block": {
    "attributes": {
      "db_cluster_parameter_group_name": {
        "computed": true,
        "description": "The name of the DB cluster parameter group.\n Constraints:\n  +  Must not match the name of an existing DB cluster parameter group.\n  \n  This value is stored as a lowercase string.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "The description for the DB cluster parameter group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "family": {
        "description": "The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.\n  *Aurora MySQL* \n Example: ` + "`" + `` + "`" + `aurora-mysql5.7` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `aurora-mysql8.0` + "`" + `` + "`" + ` \n  *Aurora PostgreSQL* \n Example: ` + "`" + `` + "`" + `aurora-postgresql14` + "`" + `` + "`" + ` \n  *RDS for MySQL* \n Example: ` + "`" + `` + "`" + `mysql8.0` + "`" + `` + "`" + ` \n  *RDS for PostgreSQL* \n Example: ` + "`" + `` + "`" + `postgres13` + "`" + `` + "`" + ` \n To list all of the available parameter group families for a DB engine, use the following command:\n  ` + "`" + `` + "`" + `aws rds describe-db-engine-versions --query \"DBEngineVersions[].DBParameterGroupFamily\" --engine \u003cengine\u003e` + "`" + `` + "`" + ` \n For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:\n  ` + "`" + `` + "`" + `aws rds describe-db-engine-versions --query \"DBEngineVersions[].DBParameterGroupFamily\" --engine aurora-postgresql` + "`" + `` + "`" + ` \n  The output contains duplicates.\n  The following are the valid DB engine values:\n  +   ` + "`" + `` + "`" + `aurora-mysql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `aurora-postgresql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `mysql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `postgres` + "`" + `` + "`" + `",
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
      "parameters": {
        "description": "Provides a list of parameters for the DB cluster parameter group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the DB cluster parameter group.",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::RDS::DBClusterParameterGroup` + "`" + `` + "`" + ` resource creates a new Amazon RDS DB cluster parameter group.\n For information about configuring parameters for Amazon Aurora DB clusters, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.\n  If you apply a parameter group to a DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.\n If you apply a change to parameter group associated with a stopped DB cluster, then the updated stack waits until the DB cluster is started.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsDbClusterParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbClusterParameterGroup), &result)
	return &result
}
