package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsOptionGroup = `{
  "block": {
    "attributes": {
      "engine_name": {
        "computed": true,
        "description": "Specifies the name of the engine that this option group should be associated with.\n Valid Values: \n  +   ` + "`" + `` + "`" + `mariadb` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `mysql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-ee` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-ee-cdb` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-se2` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-se2-cdb` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `postgres` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-ee` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-se` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-ex` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-web` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "major_engine_version": {
        "computed": true,
        "description": "Specifies the major version of the engine that this option group should be associated with.",
        "description_kind": "plain",
        "type": "string"
      },
      "option_configurations": {
        "computed": true,
        "description": "A list of options and the settings for each option.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "db_security_group_memberships": {
              "computed": true,
              "description": "A list of DBSecurityGroupMembership name strings used for this option.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "option_name": {
              "computed": true,
              "description": "The configuration of options to include in a group.",
              "description_kind": "plain",
              "type": "string"
            },
            "option_settings": {
              "computed": true,
              "description": "The option settings to include in an option group.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the option that has settings that you can set.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The current value of the option setting.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "option_version": {
              "computed": true,
              "description": "The version for the option.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The optional port for the option.",
              "description_kind": "plain",
              "type": "number"
            },
            "vpc_security_group_memberships": {
              "computed": true,
              "description": "A list of VpcSecurityGroupMembership name strings used for this option.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "option_group_description": {
        "computed": true,
        "description": "The description of the option group.",
        "description_kind": "plain",
        "type": "string"
      },
      "option_group_name": {
        "computed": true,
        "description": "The name of the option group to be created.\n Constraints:\n  +  Must be 1 to 255 letters, numbers, or hyphens\n  +  First character must be a letter\n  +  Can't end with a hyphen or contain two consecutive hyphens\n  \n Example: ` + "`" + `` + "`" + `myoptiongroup` + "`" + `` + "`" + ` \n If you don't specify a value for ` + "`" + `` + "`" + `OptionGroupName` + "`" + `` + "`" + ` property, a name is automatically created for the option group.\n  This value is stored as a lowercase string.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An optional array of key-value pairs to apply to this option group.",
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
    "description": "Data Source schema for AWS::RDS::OptionGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsOptionGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsOptionGroup), &result)
	return &result
}
