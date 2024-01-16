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
        "description": "Indicates the name of the engine that this option group can be applied to.",
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
        "description": "Indicates the major engine version associated with this option group.",
        "description_kind": "plain",
        "type": "string"
      },
      "option_configurations": {
        "computed": true,
        "description": "Indicates what options are available in the option group.",
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
        "description": "Provides a description of the option group.",
        "description_kind": "plain",
        "type": "string"
      },
      "option_group_name": {
        "computed": true,
        "description": "Specifies the name of the option group.",
        "description_kind": "plain",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
