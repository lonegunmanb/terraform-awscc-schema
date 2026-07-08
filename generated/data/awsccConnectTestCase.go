package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectTestCase = `{
  "block": {
    "attributes": {
      "content": {
        "computed": true,
        "description": "The content of the test case.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the test case.",
        "description_kind": "plain",
        "type": "string"
      },
      "entry_point": {
        "computed": true,
        "description": "Entry point for Testcase.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "chat_entry_point_parameters": {
              "computed": true,
              "description": "The chat entry point parameters for the test case",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "flow_id": {
                    "computed": true,
                    "description": "The flow id used for the TestCase",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "type": {
              "computed": true,
              "description": "The type of the Entry Point",
              "description_kind": "plain",
              "type": "string"
            },
            "voice_call_entry_point_parameters": {
              "computed": true,
              "description": "The voice call entry point parameters for the test case",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_phone_number": {
                    "computed": true,
                    "description": "The destination phonenumber of the EntryPoint",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "flow_id": {
                    "computed": true,
                    "description": "The flow id used for the TestCase",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source_phone_number": {
                    "computed": true,
                    "description": "The source phonenumber of the EntryPoint",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "initialization_data": {
        "computed": true,
        "description": "The initialization data of the test case.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_region": {
        "computed": true,
        "description": "Last modified region.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "Last modified time.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "The name of the test case.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the test case.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
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
              "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "test_case_arn": {
        "computed": true,
        "description": "The identifier of the test case.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::TestCase",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectTestCaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectTestCase), &result)
	return &result
}
