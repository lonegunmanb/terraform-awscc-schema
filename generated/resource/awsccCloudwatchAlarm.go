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
        "description": "The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *API Reference*.",
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
        "description": "The name of the alarm. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the alarm name. \n  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "comparison_operator": {
        "description": "The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "datapoints_to_alarm": {
        "computed": true,
        "description": "The number of datapoints that must be breaching to trigger the alarm. This is used only if you are setting an \"M out of N\" alarm. In that case, this value is the M, and the value that you set for ` + "`" + `` + "`" + `EvaluationPeriods` + "`" + `` + "`" + ` is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.\n If you omit this parameter, CW uses the same value here that you set for ` + "`" + `` + "`" + `EvaluationPeriods` + "`" + `` + "`" + `, and the alarm goes to alarm state if that many consecutive periods are breaching.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "dimensions": {
        "computed": true,
        "description": "The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify ` + "`" + `` + "`" + `Dimensions` + "`" + `` + "`" + `. Instead, you use ` + "`" + `` + "`" + `Metrics` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description": "The name of the dimension, from 1?255 characters in length. This dimension name must have been included when the metric was published.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the dimension, from 1?255 characters in length.",
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
        "description": "Used only for alarms based on percentiles. If ` + "`" + `` + "`" + `ignore` + "`" + `` + "`" + `, the alarm state does not change during periods with too few data points to be statistically significant. If ` + "`" + `` + "`" + `evaluate` + "`" + `` + "`" + ` or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "evaluation_periods": {
        "description": "The number of periods over which data is compared to the specified threshold. If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an \"M out of N\" alarm, this value is the N, and ` + "`" + `` + "`" + `DatapointsToAlarm` + "`" + `` + "`" + ` is the M.\n For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "extended_statistic": {
        "computed": true,
        "description": "The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.\n For an alarm based on a metric, you must specify either ` + "`" + `` + "`" + `Statistic` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ExtendedStatistic` + "`" + `` + "`" + ` but not both.\n For an alarm based on a math expression, you can't specify ` + "`" + `` + "`" + `ExtendedStatistic` + "`" + `` + "`" + `. Instead, you use ` + "`" + `` + "`" + `Metrics` + "`" + `` + "`" + `.",
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
        "description": "The actions to execute when this alarm transitions to the ` + "`" + `` + "`" + `INSUFFICIENT_DATA` + "`" + `` + "`" + ` state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "metric_name": {
        "computed": true,
        "description": "The name of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you use ` + "`" + `` + "`" + `Metrics` + "`" + `` + "`" + ` instead and you can't specify ` + "`" + `` + "`" + `MetricName` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metrics": {
        "computed": true,
        "description": "An array that enables you to create an alarm based on the result of a metric math expression. Each item in the array either retrieves a metric or performs a math expression.\n If you specify the ` + "`" + `` + "`" + `Metrics` + "`" + `` + "`" + ` parameter, you cannot specify ` + "`" + `` + "`" + `MetricName` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Dimensions` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Period` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Namespace` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Statistic` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `ExtendedStatistic` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `Unit` + "`" + `` + "`" + `.",
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
              "description": "The math expression to be performed on the returned data, if this object is performing a math expression. This expression can use the ` + "`" + `` + "`" + `Id` + "`" + `` + "`" + ` of the other metrics to refer to those metrics, and can also use the ` + "`" + `` + "`" + `Id` + "`" + `` + "`" + ` of other expressions to use the result of those expressions. For more information about metric math expressions, see [Metric Math Syntax and Functions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax) in the *User Guide*.\n Within each MetricDataQuery object, you must specify either ` + "`" + `` + "`" + `Expression` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MetricStat` + "`" + `` + "`" + ` but not both.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "description": "A short name used to tie this object to the results in the response. This name must be unique within a single call to ` + "`" + `` + "`" + `GetMetricData` + "`" + `` + "`" + `. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "label": {
              "computed": true,
              "description": "A human-readable label for this metric or expression. This is especially useful if this is an expression, so that you know what the value represents. If the metric or expression is shown in a CW dashboard widget, the label is shown. If ` + "`" + `` + "`" + `Label` + "`" + `` + "`" + ` is omitted, CW generates a default.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_stat": {
              "computed": true,
              "description": "The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.\n Within one MetricDataQuery object, you must specify either ` + "`" + `` + "`" + `Expression` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MetricStat` + "`" + `` + "`" + ` but not both.",
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
                          "description": "The metric dimensions that you want to be used for the metric that the alarm will watch.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "description": "The name of the dimension, from 1?255 characters in length. This dimension name must have been included when the metric was published.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "The value for the dimension, from 1?255 characters in length.",
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
                          "description": "The name of the metric that you want the alarm to watch. This is a required field.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "namespace": {
                          "computed": true,
                          "description": "The namespace of the metric that the alarm will watch.",
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
                    "description": "The granularity, in seconds, of the returned data points. For metrics with regular resolution, a period can be as short as one minute (60 seconds) and must be a multiple of 60. For high-resolution metrics that are collected at intervals of less than one minute, the period can be 1, 5, 10, 30, 60, or any multiple of 60. High-resolution metrics are those metrics stored by a ` + "`" + `` + "`" + `PutMetricData` + "`" + `` + "`" + ` call that includes a ` + "`" + `` + "`" + `StorageResolution` + "`" + `` + "`" + ` of 1 second.\n If the ` + "`" + `` + "`" + `StartTime` + "`" + `` + "`" + ` parameter specifies a time stamp that is greater than 3 hours ago, you must specify the period as follows or no data points in that time range is returned:\n  +  Start time between 3 hours and 15 days ago - Use a multiple of 60 seconds (1 minute).\n  +  Start time between 15 and 63 days ago - Use a multiple of 300 seconds (5 minutes).\n  +  Start time greater than 63 days ago - Use a multiple of 3600 seconds (1 hour).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "stat": {
                    "description": "The statistic to return. It can include any CW statistic or extended statistic. For a list of valid values, see the table in [Statistics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Statistic) in the *User Guide*.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "unit": {
                    "computed": true,
                    "description": "The unit to use for the returned data points. \n Valid values are: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.",
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
              "description": "The granularity, in seconds, of the returned data points. For metrics with regular resolution, a period can be as short as one minute (60 seconds) and must be a multiple of 60. For high-resolution metrics that are collected at intervals of less than one minute, the period can be 1, 5, 10, 30, 60, or any multiple of 60. High-resolution metrics are those metrics stored by a ` + "`" + `` + "`" + `PutMetricData` + "`" + `` + "`" + ` operation that includes a ` + "`" + `` + "`" + `StorageResolution of 1 second` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "return_data": {
              "computed": true,
              "description": "This option indicates whether to return the timestamps and raw data values of this metric.\n When you create an alarm based on a metric math expression, specify ` + "`" + `` + "`" + `True` + "`" + `` + "`" + ` for this value for only the one math expression that the alarm is based on. You must specify ` + "`" + `` + "`" + `False` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `ReturnData` + "`" + `` + "`" + ` for all the other metrics and expressions used in the alarm.\n This field is required.",
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
        "description": "The namespace of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify ` + "`" + `` + "`" + `Namespace` + "`" + `` + "`" + ` and you use ` + "`" + `` + "`" + `Metrics` + "`" + `` + "`" + ` instead.\n For a list of namespaces for metrics from AWS services, see [Services That Publish Metrics.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ok_actions": {
        "computed": true,
        "description": "The actions to execute when this alarm transitions to the ` + "`" + `` + "`" + `OK` + "`" + `` + "`" + ` state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "period": {
        "computed": true,
        "description": "The period, in seconds, over which the statistic is applied. This is required for an alarm based on a metric. Valid values are 10, 30, 60, and any multiple of 60.\n For an alarm based on a math expression, you can't specify ` + "`" + `` + "`" + `Period` + "`" + `` + "`" + `, and instead you use the ` + "`" + `` + "`" + `Metrics` + "`" + `` + "`" + ` parameter.\n  *Minimum:* 10",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "statistic": {
        "computed": true,
        "description": "The statistic for the metric associated with the alarm, other than percentile. For percentile statistics, use ` + "`" + `` + "`" + `ExtendedStatistic` + "`" + `` + "`" + `.\n For an alarm based on a metric, you must specify either ` + "`" + `` + "`" + `Statistic` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ExtendedStatistic` + "`" + `` + "`" + ` but not both.\n For an alarm based on a math expression, you can't specify ` + "`" + `` + "`" + `Statistic` + "`" + `` + "`" + `. Instead, you use ` + "`" + `` + "`" + `Metrics` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate tags with the alarm when you create the alarm, you must have the ` + "`" + `` + "`" + `cloudwatch:TagResource` + "`" + `` + "`" + ` permission.\n Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A string that you can use to assign a value. The combination of tag keys and values can help you organize and categorize your resources.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the specified tag key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "threshold": {
        "computed": true,
        "description": "The value to compare with the specified statistic.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "threshold_metric_id": {
        "computed": true,
        "description": "In an alarm based on an anomaly detection model, this is the ID of the ` + "`" + `` + "`" + `ANOMALY_DETECTION_BAND` + "`" + `` + "`" + ` function used as the threshold for the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "treat_missing_data": {
        "computed": true,
        "description": "Sets how this alarm is to handle missing data points. Valid values are ` + "`" + `` + "`" + `breaching` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `notBreaching` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `ignore` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `missing` + "`" + `` + "`" + `. For more information, see [Configuring How Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon User Guide*.\n If you omit this parameter, the default behavior of ` + "`" + `` + "`" + `missing` + "`" + `` + "`" + ` is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "unit": {
        "computed": true,
        "description": "The unit of the metric associated with the alarm. Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a ` + "`" + `` + "`" + `Metrics` + "`" + `` + "`" + ` array.\n  You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::CloudWatch::Alarm` + "`" + `` + "`" + ` type specifies an alarm and associates it with the specified metric or metric math expression.\n When this operation creates an alarm, the alarm state is immediately set to ` + "`" + `` + "`" + `INSUFFICIENT_DATA` + "`" + `` + "`" + `. The alarm is then evaluated and its state is set appropriately. Any actions associated with the new state are then executed.\n When you update an existing alarm, its state is left unchanged, but the update completely overwrites the previous configuration of the alarm.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchAlarmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchAlarm), &result)
	return &result
}
