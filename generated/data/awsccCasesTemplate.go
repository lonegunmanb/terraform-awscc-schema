package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCasesTemplate = `{
  "block": {
    "attributes": {
      "created_time": {
        "computed": true,
        "description": "The time at which the template was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description explaining the purpose and use case for this template. Should indicate what types of cases this template is designed for and any specific workflow it supports.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The unique identifier of the Cases domain.",
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
        "description": "The time at which the template was created or last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "layout_configuration": {
        "computed": true,
        "description": "Specifies the default layout to use when displaying cases created from this template. The layout determines which fields are visible and their arrangement in the agent interface.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_layout": {
              "computed": true,
              "description": "The unique identifier of a layout.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "A name for the template. It must be unique per domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "required_fields": {
        "computed": true,
        "description": "A list of fields that must contain a value for a case to be successfully created with this template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "field_id": {
              "computed": true,
              "description": "The unique identifier of a field.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "rules": {
        "computed": true,
        "description": "A list of case rules (also known as case field conditions) on a template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "case_rule_id": {
              "computed": true,
              "description": "The unique identifier of a case rule.",
              "description_kind": "plain",
              "type": "string"
            },
            "field_id": {
              "computed": true,
              "description": "The ID of the field that this rule applies to.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "status": {
        "computed": true,
        "description": "The current status of the template. Active templates can be used to create new cases, while Inactive templates are disabled but preserved for existing cases.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags that you attach to this template.",
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
      "template_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the template.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_id": {
        "computed": true,
        "description": "The unique identifier of a template.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cases::Template",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCasesTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCasesTemplate), &result)
	return &result
}
