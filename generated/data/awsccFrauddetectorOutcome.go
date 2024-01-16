package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFrauddetectorOutcome = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The outcome ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The timestamp when the outcome was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The outcome description.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The timestamp when the outcome was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the outcome.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with this outcome.",
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
    "description": "Data Source schema for AWS::FraudDetector::Outcome",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFrauddetectorOutcomeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFrauddetectorOutcome), &result)
	return &result
}
