package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotcoredeviceadvisorSuiteDefinition = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "suite_definition_arn": {
        "computed": true,
        "description": "The Amazon Resource name for the suite definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "suite_definition_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "device_permission_role_arn": {
              "computed": true,
              "description": "The device permission role arn of the test suite.",
              "description_kind": "plain",
              "type": "string"
            },
            "devices": {
              "computed": true,
              "description": "The devices being tested in the test suite",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "thing_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "intended_for_qualification": {
              "computed": true,
              "description": "Whether the tests are intended for qualification in a suite.",
              "description_kind": "plain",
              "type": "bool"
            },
            "root_group": {
              "computed": true,
              "description": "The root group of the test suite.",
              "description_kind": "plain",
              "type": "string"
            },
            "suite_definition_name": {
              "computed": true,
              "description": "The Name of the suite definition.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "suite_definition_id": {
        "computed": true,
        "description": "The unique identifier for the suite definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "suite_definition_version": {
        "computed": true,
        "description": "The suite definition version of a test suite.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::IoTCoreDeviceAdvisor::SuiteDefinition",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotcoredeviceadvisorSuiteDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotcoredeviceadvisorSuiteDefinition), &result)
	return &result
}
