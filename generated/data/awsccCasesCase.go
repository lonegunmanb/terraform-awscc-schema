package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCasesCase = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the case.",
        "description_kind": "plain",
        "type": "string"
      },
      "case_id": {
        "computed": true,
        "description": "A unique identifier of the case.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_id": {
        "computed": true,
        "description": "The full customer profile ARN for the case.",
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
      "tags": {
        "computed": true,
        "description": "A list of tags for the case.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "template_id": {
        "computed": true,
        "description": "A unique identifier of a template.",
        "description_kind": "plain",
        "type": "string"
      },
      "title": {
        "computed": true,
        "description": "The title of the case.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cases::Case",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCasesCaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCasesCase), &result)
	return &result
}
