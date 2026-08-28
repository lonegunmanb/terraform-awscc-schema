package resource

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
        "description": "The full customer profile ARN for the case.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_id": {
        "description": "The unique identifier of the Cases domain.",
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "template_id": {
        "description": "A unique identifier of a template.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "title": {
        "description": "The title of the case.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Creates a case in the specified Cases domain.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCasesCaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCasesCase), &result)
	return &result
}
