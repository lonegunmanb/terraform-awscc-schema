package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupFramework = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The date and time that a framework is created, in ISO 8601 representation. The value of CreationTime is accurate to milliseconds. For example, 2020-07-10T15:00:00.000-08:00 represents the 10th of July 2020 at 3:00 PM 8 hours behind UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_status": {
        "computed": true,
        "description": "The deployment status of a framework. The statuses are: ` + "`" + `CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED | FAILED` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "framework_arn": {
        "computed": true,
        "description": "An Amazon Resource Name (ARN) that uniquely identifies Framework as a resource",
        "description_kind": "plain",
        "type": "string"
      },
      "framework_controls": {
        "description": "Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "control_input_parameters": {
              "computed": true,
              "description": "A list of ParameterName and ParameterValue pairs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "parameter_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parameter_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "control_name": {
              "description": "The name of a control. This name is between 1 and 256 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "control_scope": {
              "computed": true,
              "description": "The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "compliance_resource_ids": {
                    "computed": true,
                    "description": "The ID of the only AWS resource that you want your control scope to contain.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "compliance_resource_types": {
                    "computed": true,
                    "description": "Describes whether the control scope includes one or more types of resources, such as ` + "`" + `EFS` + "`" + ` or ` + "`" + `RDS` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "tags": {
                    "computed": true,
                    "description": "Describes whether the control scope includes resources with one or more tags. Each tag is a key-value pair.",
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
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "framework_description": {
        "computed": true,
        "description": "An optional description of the framework with a maximum 1,024 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "framework_name": {
        "computed": true,
        "description": "The unique name of a framework. This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "framework_status": {
        "computed": true,
        "description": "A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. The statuses are:\n\n` + "`" + `ACTIVE` + "`" + ` when recording is turned on for all resources governed by the framework.\n\n` + "`" + `PARTIALLY_ACTIVE` + "`" + ` when recording is turned off for at least one resource governed by the framework.\n\n` + "`" + `INACTIVE` + "`" + ` when recording is turned off for all resources governed by the framework.\n\n` + "`" + `UNAVAILABLE` + "`" + ` when AWS Backup is unable to validate recording status at this time.",
        "description_kind": "plain",
        "type": "string"
      },
      "framework_tags": {
        "computed": true,
        "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Contains detailed information about a framework. Frameworks contain controls, which evaluate and report on your backup events and resources. Frameworks generate daily compliance results.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBackupFrameworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupFramework), &result)
	return &result
}
