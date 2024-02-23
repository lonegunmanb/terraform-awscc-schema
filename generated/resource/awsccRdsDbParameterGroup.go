package resource

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
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "Provides the customer-specified description for this DB parameter group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "family": {
        "description": "The DB parameter group family name. A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a DB engine and engine version compatible with that DB parameter group family.\n  The DB parameter group family can't be changed when updating a DB parameter group.\n  To list all of the available parameter group families, use the following command:\n ` + "`" + `` + "`" + `aws rds describe-db-engine-versions --query \"DBEngineVersions[].DBParameterGroupFamily\"` + "`" + `` + "`" + `\n The output contains duplicates.\n For more information, see ` + "`" + `` + "`" + `CreateDBParameterGroup` + "`" + `` + "`" + `.",
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
        "computed": true,
        "description": "An array of parameter names and values for the parameter update. At least one parameter name and value must be supplied. Subsequent arguments are optional.\n RDS for Db2 requires you to bring your own Db2 license. You must enter your IBM customer ID (` + "`" + `` + "`" + `rds.ibm_customer_id` + "`" + `` + "`" + `) and site number (` + "`" + `` + "`" + `rds.ibm_site_id` + "`" + `` + "`" + `) before starting a Db2 instance.\n For more information about DB parameters and DB parameter groups for Amazon RDS DB engines, see [Working with DB Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide*.\n For more information about DB cluster and DB instance parameters and parameter groups for Amazon Aurora DB engines, see [Working with DB Parameter Groups and DB Cluster Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.\n  AWS CloudFormation doesn't support specifying an apply method for each individual ",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An optional array of key-value pairs to apply to this DB parameter group.\n  Currently, this is the only property that supports drift detection.",
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
    "description": "The ` + "`" + `` + "`" + `AWS::RDS::DBParameterGroup` + "`" + `` + "`" + ` resource creates a custom parameter group for an RDS database family.\n This type can be declared in a template and referenced in the ` + "`" + `` + "`" + `DBParameterGroupName` + "`" + `` + "`" + ` property of an ` + "`" + `` + "`" + `AWS::RDS::DBInstance` + "`" + `` + "`" + ` resource.\n For information about configuring parameters for Amazon RDS DB instances, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide*.\n For information about configuring parameters for Amazon Aurora DB instances, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.\n  Applying a parameter group to a DB instance may require the DB instance to reboot, resulting in a database outage for the duration of the reboot.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsDbParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbParameterGroup), &result)
	return &result
}
