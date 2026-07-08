package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchInsightRule = `{
  "block": {
    "attributes": {
      "apply_on_transformed_logs": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_body": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rule_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rule_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::CloudWatch::InsightRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudwatchInsightRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchInsightRule), &result)
	return &result
}
