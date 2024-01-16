package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchAlarm = `{
  "block": {
    "attributes": {
      "actions_enabled": {
        "computed": true,
        "description": "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "alarm_actions": {
        "computed": true,
        "description": "The list of actions to execute when this alarm transitions into an ALARM state from any other state.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "alarm_description": {
        "computed": true,
        "description": "The description of the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alarm_name": {
        "computed": true,
        "description": "The name of the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name is a unique name for each resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "comparison_operator": {
        "description": "The arithmetic operation to use when comparing the specified statistic and threshold.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "datapoints_to_alarm": {
        "computed": true,
        "description": "The number of datapoints that must be breaching to trigger the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "dimensions": {
        "computed": true,
        "description": "The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify Dimensions. Instead, you use Metrics.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description": "The name of the dimension.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the dimension.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "evaluate_low_sample_count_percentile": {
        "computed": true,
        "description": "Used only for alarms based on percentiles.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "evaluation_periods": {
        "description": "The number of periods over which data is compared to the specified threshold.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "extended_statistic": {
        "computed": true,
        "description": "The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.",
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
      "insufficient_data_actions": {
        "computed": true,
        "description": "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "metric_name": {
        "computed": true,
        "description": "The name of the metric associated with the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metrics": {
        "computed": true,
        "description": "An array that enables you to create an alarm based on the result of a metric math expression.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_id": {
              "computed": true,
              "description": "The ID of the account where the metrics are located, if this is a cross-account alarm.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expression": {
              "computed": true,
              "description": "The math expression to be performed on the returned data.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "description": "A short name used to tie this object to the results in the response.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "label": {
              "computed": true,
              "description": "A human-readable label for this metric or expression.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_stat": {
              "computed": true,
              "description": "The metric to be returned, along with statistics, period, and units.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "metric": {
                    "description": "The metric to return, including the metric name, namespace, and dimensions.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dimensions": {
                          "computed": true,
                          "description": "The dimensions for the metric.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "description": "The name of the dimension.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "The value for the dimension.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "metric_name": {
                          "computed": true,
                          "description": "The name of the metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "namespace": {
                          "computed": true,
                          "description": "The namespace of the metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "period": {
                    "description": "The granularity, in seconds, of the returned data points.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "stat": {
                    "description": "The statistic to return.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "unit": {
                    "computed": true,
                    "description": "The unit to use for the returned data points.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "period": {
              "computed": true,
              "description": "The period in seconds, over which the statistic is applied.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "return_data": {
              "computed": true,
              "description": "This option indicates whether to return the timestamps and raw data values of this metric.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "namespace": {
        "computed": true,
        "description": "The namespace of the metric associated with the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ok_actions": {
        "computed": true,
        "description": "The actions to execute when this alarm transitions to the OK state from any other state.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "period": {
        "computed": true,
        "description": "The period in seconds, over which the statistic is applied.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "statistic": {
        "computed": true,
        "description": "The statistic for the metric associated with the alarm, other than percentile.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "threshold": {
        "computed": true,
        "description": "In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "threshold_metric_id": {
        "computed": true,
        "description": "In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "treat_missing_data": {
        "computed": true,
        "description": "Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "unit": {
        "computed": true,
        "description": "The unit of the metric associated with the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudWatch::Alarm",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchAlarmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchAlarm), &result)
	return &result
}
