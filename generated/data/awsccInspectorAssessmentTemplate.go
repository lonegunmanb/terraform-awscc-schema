package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccInspectorAssessmentTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assessment_target_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assessment_template_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "duration_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rules_package_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "user_attributes_for_findings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Inspector::AssessmentTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccInspectorAssessmentTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInspectorAssessmentTemplate), &result)
	return &result
}
