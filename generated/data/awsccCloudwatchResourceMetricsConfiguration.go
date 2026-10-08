package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchResourceMetricsConfiguration = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time at which the resource metrics configuration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metric_selections": {
        "computed": true,
        "description": "The metric selections that define which metrics are enabled for detailed monitoring on the resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "include_metrics": {
              "computed": true,
              "description": "The list of metric names to include in detailed monitoring for the resource.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "resource_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resource for which the detailed monitoring metrics configuration is managed.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The time at which the resource metrics configuration was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudWatch::ResourceMetricsConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudwatchResourceMetricsConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchResourceMetricsConfiguration), &result)
	return &result
}
