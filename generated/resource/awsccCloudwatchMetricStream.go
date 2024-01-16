package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchMetricStream = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name of the metric stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date": {
        "computed": true,
        "description": "The date of creation of the metric stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "exclude_filters": {
        "computed": true,
        "description": "Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "metric_names": {
              "computed": true,
              "description": "Only metrics with MetricNames matching these values will be streamed. Must be set together with Namespace.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "namespace": {
              "description": "Only metrics with Namespace matching this value will be streamed.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "firehose_arn": {
        "description": "The ARN of the Kinesis Firehose where to stream the data.",
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
      "include_filters": {
        "computed": true,
        "description": "Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "metric_names": {
              "computed": true,
              "description": "Only metrics with MetricNames matching these values will be streamed. Must be set together with Namespace.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "namespace": {
              "description": "Only metrics with Namespace matching this value will be streamed.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "include_linked_accounts_metrics": {
        "computed": true,
        "description": "If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "last_update_date": {
        "computed": true,
        "description": "The date of the last update of the metric stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the metric stream.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_format": {
        "description": "The output format of the data streamed to the Kinesis Firehose.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The ARN of the role that provides access to the Kinesis Firehose.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Displays the state of the Metric Stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "statistics_configurations": {
        "computed": true,
        "description": "By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "additional_statistics": {
              "description": "The additional statistics to stream for the metrics listed in IncludeMetrics.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "include_metrics": {
              "description": "An array that defines the metrics that are to have additional statistics streamed.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "metric_name": {
                    "description": "The name of the metric.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "namespace": {
                    "description": "The namespace of the metric.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A set of tags to assign to the delivery stream.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A unique identifier for the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "String which you can use to describe or define the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for Metric Stream",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchMetricStreamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchMetricStream), &result)
	return &result
}
