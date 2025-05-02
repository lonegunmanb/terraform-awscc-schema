package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbParameterGroup = `{
  "block": {
    "attributes": {
      "db_parameter_group_name": {
        "computed": true,
        "description": "The name of the DB parameter group.\n Constraints:\n  +  Must be 1 to 255 letters, numbers, or hyphens.\n  +  First character must be a letter\n  +  Can't end with a hyphen or contain two consecutive hyphens\n  \n If you don't specify a value for ` + "`" + `` + "`" + `DBParameterGroupName` + "`" + `` + "`" + ` property, a name is automatically created for the DB parameter group.\n  This value is stored as a lowercase string.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Provides the customer-specified description for this DB parameter group.",
        "description_kind": "plain",
        "type": "string"
      },
      "family": {
        "computed": true,
        "description": "The DB parameter group family name. A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a database engine and engine version compatible with that DB parameter group family.\n To list all of the available parameter group families for a DB engine, use the following command:\n  ` + "`" + `` + "`" + `aws rds describe-db-engine-versions --query \"DBEngineVersions[].DBParameterGroupFamily\" --engine \u003cengine\u003e` + "`" + `` + "`" + ` \n For example, to list all of the available parameter group families for the MySQL DB engine, use the following command:\n  ` + "`" + `` + "`" + `aws rds describe-db-engine-versions --query \"DBEngineVersions[].DBParameterGroupFamily\" --engine mysql` + "`" + `` + "`" + ` \n  The output contains duplicates.\n  The following are the valid DB engine values:\n  +   ` + "`" + `` + "`" + `aurora-mysql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `aurora-postgresql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `db2-ae` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `db2-se` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `mysql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-ee` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-ee-cdb` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-se2` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-se2-cdb` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `postgres` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-ee` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-se` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-ex` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-web` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description": "A mapping of parameter names and values for the parameter update. You must specify at least one parameter name and value.\n For more information about parameter groups, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide*, or [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.\n  AWS CloudFormation doesn't support specifying an apply method for each individual parameter. The default apply method for each parameter is used.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the DB parameter group.",
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
    "description": "Data Source schema for AWS::RDS::DBParameterGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsDbParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbParameterGroup), &result)
	return &result
}
