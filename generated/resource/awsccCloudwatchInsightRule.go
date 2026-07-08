package resource

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
        "optional": true,
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_body": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_state": {
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::CloudWatch::InsightRule. Creates a Contributor Insights rule that analyzes log data to identify top contributors and usage patterns.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchInsightRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchInsightRule), &result)
	return &result
}
