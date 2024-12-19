package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsMetricFilter = `{
  "block": {
    "attributes": {
      "apply_on_transformed_logs": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "filter_name": {
        "computed": true,
        "description": "The name of the metric filter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_pattern": {
        "description": "A filter pattern for extracting metric data out of ingested log events. For more information, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).",
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
      "log_group_name": {
        "description": "The name of an existing log group that you want to associate with this metric filter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metric_transformations": {
        "description": "The metric transformations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_value": {
              "computed": true,
              "description": "(Optional) The value to emit when a filter pattern does not match a log event. This value can be null.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dimensions": {
              "computed": true,
              "description": "The fields to use as dimensions for the metric. One metric filter can include as many as three dimensions.\n  Metrics extracted from log events are charged as custom metrics. To prevent unexpected high charges, do not specify high-cardinality fields such as ` + "`" + `` + "`" + `IPAddress` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `requestID` + "`" + `` + "`" + ` as dimensions. Each different value found for a dimension is treated as a separate metric and accrues charges as a separate custom metric. \n CloudWatch Logs disables a metric filter if it generates 1000 different name/value pairs for your specified dimensions within a certain amount of time. This helps to prevent accidental high charges.\n You can also set up a billing alarm to alert you if your charges are higher than expected. For more information, see [Creating a Billing Alarm to Monitor Your Estimated Charges](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/monitor_estimated_charges_with_cloudwatch.html).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The name for the CW metric dimension that the metric filter creates.\n Dimension names must contain only ASCII characters, must include at least one non-whitespace character, and cannot start with a colon (:).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The log event field that will contain the value for this dimension. This dimension will only be published for a metric if the value is found in the log event. For example, ` + "`" + `` + "`" + `$.eventType` + "`" + `` + "`" + ` for JSON log events, or ` + "`" + `` + "`" + `$server` + "`" + `` + "`" + ` for space-delimited log events.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "metric_name": {
              "description": "The name of the CloudWatch metric.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "metric_namespace": {
              "description": "A custom namespace to contain your metric in CloudWatch. Use namespaces to group together metrics that are similar. For more information, see [Namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Namespace).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "metric_value": {
              "description": "The value that is published to the CloudWatch metric. For example, if you're counting the occurrences of a particular term like ` + "`" + `` + "`" + `Error` + "`" + `` + "`" + `, specify 1 for the metric value. If you're counting the number of bytes transferred, reference the value that is in the log event by using $. followed by the name of the field that you specified in the filter pattern, such as ` + "`" + `` + "`" + `$.size` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "unit": {
              "computed": true,
              "description": "The unit to assign to the metric. If you omit this, the unit is set as ` + "`" + `` + "`" + `None` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::Logs::MetricFilter` + "`" + `` + "`" + ` resource specifies a metric filter that describes how CWL extracts information from logs and transforms it into Amazon CloudWatch metrics. If you have multiple metric filters that are associated with a log group, all the filters are applied to the log streams in that group.\n The maximum number of metric filters that can be associated with a log group is 100.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsMetricFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsMetricFilter), &result)
	return &result
}
