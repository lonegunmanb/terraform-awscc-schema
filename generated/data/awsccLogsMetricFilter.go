package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsMetricFilter = `{
  "block": {
    "attributes": {
      "filter_name": {
        "computed": true,
        "description": "A name for the metric filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "filter_pattern": {
        "computed": true,
        "description": "Pattern that Logs follows to interpret each entry in a log.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_group_name": {
        "computed": true,
        "description": "Existing log group that you want to associate with this filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "metric_transformations": {
        "computed": true,
        "description": "A collection of information that defines how metric data gets emitted.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_value": {
              "computed": true,
              "description": "The value to emit when a filter pattern does not match a log event. This value can be null.",
              "description_kind": "plain",
              "type": "number"
            },
            "dimensions": {
              "computed": true,
              "description": "Dimensions are the key-value pairs that further define a metric",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The key of the dimension. Maximum length of 255.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value of the dimension. Maximum length of 255.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "metric_name": {
              "computed": true,
              "description": "The name of the CloudWatch metric. Metric name must be in ASCII format.",
              "description_kind": "plain",
              "type": "string"
            },
            "metric_namespace": {
              "computed": true,
              "description": "The namespace of the CloudWatch metric.",
              "description_kind": "plain",
              "type": "string"
            },
            "metric_value": {
              "computed": true,
              "description": "The value to publish to the CloudWatch metric when a filter pattern matches a log event.",
              "description_kind": "plain",
              "type": "string"
            },
            "unit": {
              "computed": true,
              "description": "The unit to assign to the metric. If you omit this, the unit is set as None.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Logs::MetricFilter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsMetricFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsMetricFilter), &result)
	return &result
}
