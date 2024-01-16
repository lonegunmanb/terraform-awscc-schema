package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFrauddetectorVariable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the variable.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The time when the variable was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_source": {
        "description": "The source of the data.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_type": {
        "description": "The data type.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_value": {
        "description": "The default value for the variable when no value is received.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The time when the variable was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the variable.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with this variable.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "variable_type": {
        "computed": true,
        "description": "The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "A resource schema for a Variable in Amazon Fraud Detector.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccFrauddetectorVariableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFrauddetectorVariable), &result)
	return &result
}
