package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCasesField = `{
  "block": {
    "attributes": {
      "created_time": {
        "computed": true,
        "description": "The time at which the field was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description explaining the purpose and usage of this field in cases. Helps agents and administrators understand what information should be captured in this field.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The unique identifier of the Cases domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "field_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the field.",
        "description_kind": "plain",
        "type": "string"
      },
      "field_id": {
        "computed": true,
        "description": "The unique identifier of a field.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time at which the field was created or last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The display name of the field as it appears to agents in the case interface. Should be descriptive and user-friendly (e.g., 'Customer Priority Level', 'Issue Category').",
        "description_kind": "plain",
        "type": "string"
      },
      "namespace": {
        "computed": true,
        "description": "Indicates whether this is a System field (predefined by AWS) or a Custom field (created by your organization). System fields cannot be modified or deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags that you attach to this field.",
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
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "The data type of the field, which determines validation rules, input constraints, and display format. Each type has specific constraints: Text (string input), Number (numeric values), Boolean (true/false), DateTime (date/time picker), SingleSelect (dropdown options), Url (URL validation), User (Amazon Connect user selection).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cases::Field",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCasesFieldSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCasesField), &result)
	return &result
}
