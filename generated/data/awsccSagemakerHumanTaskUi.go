package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerHumanTaskUi = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The timestamp when the human task user interface was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "human_task_ui_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the human task user interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "human_task_ui_name": {
        "computed": true,
        "description": "The name of the human task user interface.",
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
        "description": "An array of key-value pairs that contain metadata to help you categorize and organize a human review workflow user interface.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "ui_template": {
        "computed": true,
        "description": "The Liquid template for the worker user interface.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content": {
              "computed": true,
              "description": "The content of the Liquid template for the worker user interface.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::HumanTaskUi",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerHumanTaskUiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerHumanTaskUi), &result)
	return &result
}
