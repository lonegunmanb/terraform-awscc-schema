package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApplicationautoscalingScalingPolicy = `{
  "block": {
    "attributes": {
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
      "policy_name": {
        "description": "The name of the scaling policy.\n Updates to the name of a target tracking scaling policy are not supported, unless you also update the metric used for scaling. To change only a target tracking scaling policy's name, first delete the policy by removing the existing ` + "`" + `` + "`" + `AWS::ApplicationAutoScaling::ScalingPolicy` + "`" + `` + "`" + ` resource from the template and updating the stack. Then, recreate the resource with the same settings and a different name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_type": {
        "description": "The scaling policy type.\n The following policy types are supported: \n ` + "`" + `` + "`" + `TargetTrackingScaling` + "`" + `` + "`" + `—Not supported for Amazon EMR\n ` + "`" + `` + "`" + `StepScaling` + "`" + `` + "`" + `—Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.\n ` + "`" + `` + "`" + `PredictiveScaling` + "`" + `` + "`" + `—Only supported for Amazon ECS",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "predictive_scaling_policy_configuration": {
        "computed": true,
        "description": "The predictive scaling policy configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_capacity_breach_behavior": {
              "computed": true,
              "description": "Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity. Defaults to ` + "`" + `` + "`" + `HonorMaxCapacity` + "`" + `` + "`" + ` if not specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_capacity_buffer": {
              "computed": true,
              "description": "The size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity. The value is specified as a percentage relative to the forecast capacity. For example, if the buffer is 10, this means a 10 percent buffer, such that if the forecast capacity is 50, and the maximum capacity is 40, then the effective maximum capacity is 55. \n Required if the ` + "`" + `` + "`" + `MaxCapacityBreachBehavior` + "`" + `` + "`" + ` property is set to ` + "`" + `` + "`" + `IncreaseMaxCapacity` + "`" + `` + "`" + `, and cannot be used otherwise.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "metric_specifications": {
              "computed": true,
              "description": "This structure includes the metrics and target utilization to use for predictive scaling. \n This is an array, but we currently only support a single metric specification. That is, you can specify a target value and a single metric pair, or a target value and one scaling metric and one load metric.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "customized_capacity_metric_specification": {
                    "computed": true,
                    "description": "The customized capacity metric specification.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metric_data_queries": {
                          "computed": true,
                          "description": "One or more metric data queries to provide data points for a metric specification.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "expression": {
                                "computed": true,
                                "description": "The math expression to perform on the returned data, if this object is performing a math expression. This expression can use the ` + "`" + `` + "`" + `Id` + "`" + `` + "`" + ` of the other metrics to refer to those metrics, and can also use the ` + "`" + `` + "`" + `Id` + "`" + `` + "`" + ` of other expressions to use the result of those expressions. \n Conditional: Within each ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` object, you must specify either ` + "`" + `` + "`" + `Expression` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MetricStat` + "`" + `` + "`" + `, but not both.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "id": {
                                "computed": true,
                                "description": "A short name that identifies the object's results in the response. This name must be unique among all ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` objects specified for a single scaling policy. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscores. The first character must be a lowercase letter.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "label": {
                                "computed": true,
                                "description": "A human-readable label for this metric or expression. This is especially useful if this is a math expression, so that you know what the value represents.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "metric_stat": {
                                "computed": true,
                                "description": "Information about the metric data to return. \n Conditional: Within each ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` object, you must specify either ` + "`" + `` + "`" + `Expression` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MetricStat` + "`" + `` + "`" + `, but not both.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "metric": {
                                      "computed": true,
                                      "description": "The CloudWatch metric to return, including the metric name, namespace, and dimensions. To get the exact metric name, namespace, and dimensions, inspect the [Metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html) object that is returned by a call to [ListMetrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html).",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimensions": {
                                            "computed": true,
                                            "description": "Describes the dimensions of the metric.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "name": {
                                                  "computed": true,
                                                  "description": "The name of the dimension.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "The value of the dimension.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "optional": true
                                    },
                                    "stat": {
                                      "computed": true,
                                      "description": "The statistic to return. It can include any CloudWatch statistic or extended statistic. For a list of valid values, see the table in [Statistics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Statistic) in the *Amazon CloudWatch User Guide*. \n The most commonly used metrics for predictive scaling are ` + "`" + `` + "`" + `Average` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `Sum` + "`" + `` + "`" + `.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "unit": {
                                      "computed": true,
                                      "description": "The unit to use for the returned data points. For a complete list of the units that CloudWatch supports, see the [MetricDatum](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html) data type in the *Amazon CloudWatch API Reference*.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "return_data": {
                                "computed": true,
                                "description": "Indicates whether to return the timestamps and raw data values of this metric. \n If you use any math expressions, specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` for this value for only the final math expression that the metric specification is based on. You must specify ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `ReturnData` + "`" + `` + "`" + ` for all the other metrics and expressions used in the metric specification.\n If you are only retrieving metrics and not performing any math expressions, do not specify anything for ` + "`" + `` + "`" + `ReturnData` + "`" + `` + "`" + `. This sets it to its default (` + "`" + `` + "`" + `true` + "`" + `` + "`" + `).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "set"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "customized_load_metric_specification": {
                    "computed": true,
                    "description": "The customized load metric specification.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metric_data_queries": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "expression": {
                                "computed": true,
                                "description": "The math expression to perform on the returned data, if this object is performing a math expression. This expression can use the ` + "`" + `` + "`" + `Id` + "`" + `` + "`" + ` of the other metrics to refer to those metrics, and can also use the ` + "`" + `` + "`" + `Id` + "`" + `` + "`" + ` of other expressions to use the result of those expressions. \n Conditional: Within each ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` object, you must specify either ` + "`" + `` + "`" + `Expression` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MetricStat` + "`" + `` + "`" + `, but not both.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "id": {
                                "computed": true,
                                "description": "A short name that identifies the object's results in the response. This name must be unique among all ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` objects specified for a single scaling policy. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscores. The first character must be a lowercase letter.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "label": {
                                "computed": true,
                                "description": "A human-readable label for this metric or expression. This is especially useful if this is a math expression, so that you know what the value represents.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "metric_stat": {
                                "computed": true,
                                "description": "Information about the metric data to return. \n Conditional: Within each ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` object, you must specify either ` + "`" + `` + "`" + `Expression` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MetricStat` + "`" + `` + "`" + `, but not both.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "metric": {
                                      "computed": true,
                                      "description": "The CloudWatch metric to return, including the metric name, namespace, and dimensions. To get the exact metric name, namespace, and dimensions, inspect the [Metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html) object that is returned by a call to [ListMetrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html).",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimensions": {
                                            "computed": true,
                                            "description": "Describes the dimensions of the metric.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "name": {
                                                  "computed": true,
                                                  "description": "The name of the dimension.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "The value of the dimension.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "optional": true
                                    },
                                    "stat": {
                                      "computed": true,
                                      "description": "The statistic to return. It can include any CloudWatch statistic or extended statistic. For a list of valid values, see the table in [Statistics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Statistic) in the *Amazon CloudWatch User Guide*. \n The most commonly used metrics for predictive scaling are ` + "`" + `` + "`" + `Average` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `Sum` + "`" + `` + "`" + `.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "unit": {
                                      "computed": true,
                                      "description": "The unit to use for the returned data points. For a complete list of the units that CloudWatch supports, see the [MetricDatum](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html) data type in the *Amazon CloudWatch API Reference*.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "return_data": {
                                "computed": true,
                                "description": "Indicates whether to return the timestamps and raw data values of this metric. \n If you use any math expressions, specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` for this value for only the final math expression that the metric specification is based on. You must specify ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `ReturnData` + "`" + `` + "`" + ` for all the other metrics and expressions used in the metric specification.\n If you are only retrieving metrics and not performing any math expressions, do not specify anything for ` + "`" + `` + "`" + `ReturnData` + "`" + `` + "`" + `. This sets it to its default (` + "`" + `` + "`" + `true` + "`" + `` + "`" + `).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "set"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "customized_scaling_metric_specification": {
                    "computed": true,
                    "description": "The customized scaling metric specification.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metric_data_queries": {
                          "computed": true,
                          "description": "One or more metric data queries to provide data points for a metric specification.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "expression": {
                                "computed": true,
                                "description": "The math expression to perform on the returned data, if this object is performing a math expression. This expression can use the ` + "`" + `` + "`" + `Id` + "`" + `` + "`" + ` of the other metrics to refer to those metrics, and can also use the ` + "`" + `` + "`" + `Id` + "`" + `` + "`" + ` of other expressions to use the result of those expressions. \n Conditional: Within each ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` object, you must specify either ` + "`" + `` + "`" + `Expression` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MetricStat` + "`" + `` + "`" + `, but not both.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "id": {
                                "computed": true,
                                "description": "A short name that identifies the object's results in the response. This name must be unique among all ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` objects specified for a single scaling policy. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscores. The first character must be a lowercase letter.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "label": {
                                "computed": true,
                                "description": "A human-readable label for this metric or expression. This is especially useful if this is a math expression, so that you know what the value represents.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "metric_stat": {
                                "computed": true,
                                "description": "Information about the metric data to return. \n Conditional: Within each ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` object, you must specify either ` + "`" + `` + "`" + `Expression` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MetricStat` + "`" + `` + "`" + `, but not both.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "metric": {
                                      "computed": true,
                                      "description": "The CloudWatch metric to return, including the metric name, namespace, and dimensions. To get the exact metric name, namespace, and dimensions, inspect the [Metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html) object that is returned by a call to [ListMetrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html).",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimensions": {
                                            "computed": true,
                                            "description": "Describes the dimensions of the metric.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "name": {
                                                  "computed": true,
                                                  "description": "The name of the dimension.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "The value of the dimension.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "optional": true
                                    },
                                    "stat": {
                                      "computed": true,
                                      "description": "The statistic to return. It can include any CloudWatch statistic or extended statistic. For a list of valid values, see the table in [Statistics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Statistic) in the *Amazon CloudWatch User Guide*. \n The most commonly used metrics for predictive scaling are ` + "`" + `` + "`" + `Average` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `Sum` + "`" + `` + "`" + `.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "unit": {
                                      "computed": true,
                                      "description": "The unit to use for the returned data points. For a complete list of the units that CloudWatch supports, see the [MetricDatum](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html) data type in the *Amazon CloudWatch API Reference*.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "return_data": {
                                "computed": true,
                                "description": "Indicates whether to return the timestamps and raw data values of this metric. \n If you use any math expressions, specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` for this value for only the final math expression that the metric specification is based on. You must specify ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `ReturnData` + "`" + `` + "`" + ` for all the other metrics and expressions used in the metric specification.\n If you are only retrieving metrics and not performing any math expressions, do not specify anything for ` + "`" + `` + "`" + `ReturnData` + "`" + `` + "`" + `. This sets it to its default (` + "`" + `` + "`" + `true` + "`" + `` + "`" + `).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "set"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "predefined_load_metric_specification": {
                    "computed": true,
                    "description": "The predefined load metric specification.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "predefined_metric_type": {
                          "computed": true,
                          "description": "The metric type.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_label": {
                          "computed": true,
                          "description": "A label that uniquely identifies a target group.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "predefined_metric_pair_specification": {
                    "computed": true,
                    "description": "The predefined metric pair specification that determines the appropriate scaling metric and load metric to use.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "predefined_metric_type": {
                          "computed": true,
                          "description": "Indicates which metrics to use. There are two different types of metrics for each metric type: one is a load metric and one is a scaling metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_label": {
                          "computed": true,
                          "description": "A label that uniquely identifies a specific target group from which to determine the total and average request count.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "predefined_scaling_metric_specification": {
                    "computed": true,
                    "description": "The predefined scaling metric specification.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "predefined_metric_type": {
                          "computed": true,
                          "description": "The metric type.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_label": {
                          "computed": true,
                          "description": "A label that uniquely identifies a specific target group from which to determine the average request count.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "target_value": {
                    "computed": true,
                    "description": "Specifies the target utilization.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "mode": {
              "computed": true,
              "description": "The predictive scaling mode. Defaults to ` + "`" + `` + "`" + `ForecastOnly` + "`" + `` + "`" + ` if not specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scheduling_buffer_time": {
              "computed": true,
              "description": "The amount of time, in seconds, that the start time can be advanced. \n The value must be less than the forecast interval duration of 3600 seconds (60 minutes). Defaults to 300 seconds if not specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "resource_id": {
        "computed": true,
        "description": "The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.\n  +  ECS service - The resource type is ` + "`" + `` + "`" + `service` + "`" + `` + "`" + ` and the unique identifier is the cluster name and service name. Example: ` + "`" + `` + "`" + `service/my-cluster/my-service` + "`" + `` + "`" + `.\n  +  Spot Fleet - The resource type is ` + "`" + `` + "`" + `spot-fleet-request` + "`" + `` + "`" + ` and the unique identifier is the Spot Fleet request ID. Example: ` + "`" + `` + "`" + `spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE` + "`" + `` + "`" + `.\n  +  EMR cluster - The resource type is ` + "`" + `` + "`" + `instancegroup` + "`" + `` + "`" + ` and the unique identifier is the cluster ID and instance group ID. Example: ` + "`" + `` + "`" + `instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0` + "`" + `` + "`" + `.\n  +  AppStream 2.0 fleet - The resource type is ` + "`" + `` + "`" + `fleet` + "`" + `` + "`" + ` and the unique identifier is the fleet name. Example: ` + "`" + `` + "`" + `fleet/sample-fleet` + "`" + `` + "`" + `.\n  +  DynamoDB table - The resource type is ` + "`" + `` + "`" + `table` + "`" + `` + "`" + ` and the unique identifier is the table name. Example: ` + "`" + `` + "`" + `table/my-table` + "`" + `` + "`" + `.\n  +  DynamoDB global secondary index - The resource type is ` + "`" + `` + "`" + `index` + "`" + `` + "`" + ` and the unique identifier is the index name. Example: ` + "`" + `` + "`" + `table/my-table/index/my-table-index` + "`" + `` + "`" + `.\n  +  Aurora DB cluster - The resource type is ` + "`" + `` + "`" + `cluster` + "`" + `` + "`" + ` and the unique identifier is the cluster name. Example: ` + "`" + `` + "`" + `cluster:my-db-cluster` + "`" + `` + "`" + `.\n  +  SageMaker endpoint variant - The resource type is ` + "`" + `` + "`" + `variant` + "`" + `` + "`" + ` and the unique identifier is the resource ID. Example: ` + "`" + `` + "`" + `endpoint/my-end-point/variant/KMeansClustering` + "`" + `` + "`" + `.\n  +  Custom resources are not supported with a resource type. This parameter must specify the ` + "`" + `` + "`" + `OutputValue` + "`" + `` + "`" + ` from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our [GitHub repository](https://docs.aws.amazon.com/https://github.com/aws/aws-auto-scaling-custom-resource).\n  +  Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: ` + "`" + `` + "`" + `arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE` + "`" + `` + "`" + `.\n  +  Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: ` + "`" + `` + "`" + `arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE` + "`" + `` + "`" + `.\n  +  Lambda provisioned concurrency - The resource type is ` + "`" + `` + "`" + `function` + "`" + `` + "`" + ` and the unique identifier is the function name with a function version or alias name suffix that is not ` + "`" + `` + "`" + `$LATEST` + "`" + `` + "`" + `. Example: ` + "`" + `` + "`" + `function:my-function:prod` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `function:my-function:1` + "`" + `` + "`" + `.\n  +  Amazon Keyspaces table - The resource type is ` + "`" + `` + "`" + `table` + "`" + `` + "`" + ` and the unique identifier is the table name. Example: ` + "`" + `` + "`" + `keyspace/mykeyspace/table/mytable` + "`" + `` + "`" + `.\n  +  Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: ` + "`" + `` + "`" + `arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5` + "`" + `` + "`" + `.\n  +  Amazon ElastiCache replication group - The resource type is ` + "`" + `` + "`" + `replication-group` + "`" + `` + "`" + ` and the unique identifier is the replication group name. Example: ` + "`" + `` + "`" + `replication-group/mycluster` + "`" + `` + "`" + `.\n  +  Amazon ElastiCache cache cluster - The resource type is ` + "`" + `` + "`" + `cache-cluster` + "`" + `` + "`" + ` and the unique identifier is the cache cluster name. Example: ` + "`" + `` + "`" + `cache-cluster/mycluster` + "`" + `` + "`" + `.\n  +  Neptune cluster - The resource type is ` + "`" + `` + "`" + `cluster` + "`" + `` + "`" + ` and the unique identifier is the cluster name. Example: ` + "`" + `` + "`" + `cluster:mycluster` + "`" + `` + "`" + `.\n  +  SageMaker serverless endpoint - The resource type is ` + "`" + `` + "`" + `variant` + "`" + `` + "`" + ` and the unique identifier is the resource ID. Example: ` + "`" + `` + "`" + `endpoint/my-end-point/variant/KMeansClustering` + "`" + `` + "`" + `.\n  +  SageMaker inference component - The resource type is ` + "`" + `` + "`" + `inference-component` + "`" + `` + "`" + ` and the unique identifier is the resource ID. Example: ` + "`" + `` + "`" + `inference-component/my-inference-component` + "`" + `` + "`" + `.\n  +  Pool of WorkSpaces - The resource type is ` + "`" + `` + "`" + `workspacespool` + "`" + `` + "`" + ` and the unique identifier is the pool ID. Example: ` + "`" + `` + "`" + `workspacespool/wspool-123456` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scalable_dimension": {
        "computed": true,
        "description": "The scalable dimension. This string consists of the service namespace, resource type, and scaling property.\n  +  ` + "`" + `` + "`" + `ecs:service:DesiredCount` + "`" + `` + "`" + ` - The task count of an ECS service.\n  +  ` + "`" + `` + "`" + `elasticmapreduce:instancegroup:InstanceCount` + "`" + `` + "`" + ` - The instance count of an EMR Instance Group.\n  +  ` + "`" + `` + "`" + `ec2:spot-fleet-request:TargetCapacity` + "`" + `` + "`" + ` - The target capacity of a Spot Fleet.\n  +  ` + "`" + `` + "`" + `appstream:fleet:DesiredCapacity` + "`" + `` + "`" + ` - The capacity of an AppStream 2.0 fleet.\n  +  ` + "`" + `` + "`" + `dynamodb:table:ReadCapacityUnits` + "`" + `` + "`" + ` - The provisioned read capacity for a DynamoDB table.\n  +  ` + "`" + `` + "`" + `dynamodb:table:WriteCapacityUnits` + "`" + `` + "`" + ` - The provisioned write capacity for a DynamoDB table.\n  +  ` + "`" + `` + "`" + `dynamodb:index:ReadCapacityUnits` + "`" + `` + "`" + ` - The provisioned read capacity for a DynamoDB global secondary index.\n  +  ` + "`" + `` + "`" + `dynamodb:index:WriteCapacityUnits` + "`" + `` + "`" + ` - The provisioned write capacity for a DynamoDB global secondary index.\n  +  ` + "`" + `` + "`" + `rds:cluster:ReadReplicaCount` + "`" + `` + "`" + ` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.\n  +  ` + "`" + `` + "`" + `sagemaker:variant:DesiredInstanceCount` + "`" + `` + "`" + ` - The number of EC2 instances for a SageMaker model endpoint variant.\n  +  ` + "`" + `` + "`" + `custom-resource:ResourceType:Property` + "`" + `` + "`" + ` - The scalable dimension for a custom resource provided by your own application or service.\n  +  ` + "`" + `` + "`" + `comprehend:document-classifier-endpoint:DesiredInferenceUnits` + "`" + `` + "`" + ` - The number of inference units for an Amazon Comprehend document classification endpoint.\n  +  ` + "`" + `` + "`" + `comprehend:entity-recognizer-endpoint:DesiredInferenceUnits` + "`" + `` + "`" + ` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.\n  +  ` + "`" + `` + "`" + `lambda:function:ProvisionedConcurrency` + "`" + `` + "`" + ` - The provisioned concurrency for a Lambda function.\n  +  ` + "`" + `` + "`" + `cassandra:table:ReadCapacityUnits` + "`" + `` + "`" + ` - The provisioned read capacity for an Amazon Keyspaces table.\n  +  ` + "`" + `` + "`" + `cassandra:table:WriteCapacityUnits` + "`" + `` + "`" + ` - The provisioned write capacity for an Amazon Keyspaces table.\n  +  ` + "`" + `` + "`" + `kafka:broker-storage:VolumeSize` + "`" + `` + "`" + ` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.\n  +  ` + "`" + `` + "`" + `elasticache:cache-cluster:Nodes` + "`" + `` + "`" + ` - The number of nodes for an Amazon ElastiCache cache cluster.\n  +  ` + "`" + `` + "`" + `elasticache:replication-group:NodeGroups` + "`" + `` + "`" + ` - The number of node groups for an Amazon ElastiCache replication group.\n  +  ` + "`" + `` + "`" + `elasticache:replication-group:Replicas` + "`" + `` + "`" + ` - The number of replicas per node group for an Amazon ElastiCache replication group.\n  +  ` + "`" + `` + "`" + `neptune:cluster:ReadReplicaCount` + "`" + `` + "`" + ` - The count of read replicas in an Amazon Neptune DB cluster.\n  +  ` + "`" + `` + "`" + `sagemaker:variant:DesiredProvisionedConcurrency` + "`" + `` + "`" + ` - The provisioned concurrency for a SageMaker serverless endpoint.\n  +  ` + "`" + `` + "`" + `sagemaker:inference-component:DesiredCopyCount` + "`" + `` + "`" + ` - The number of copies across an endpoint for a SageMaker inference component.\n  +  ` + "`" + `` + "`" + `workspaces:workspacespool:DesiredUserSessions` + "`" + `` + "`" + ` - The number of user sessions for the WorkSpaces in the pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_target_id": {
        "computed": true,
        "description": "The CloudFormation-generated ID of an Application Auto Scaling scalable target. For more information about the ID, see the Return Value section of the ` + "`" + `` + "`" + `AWS::ApplicationAutoScaling::ScalableTarget` + "`" + `` + "`" + ` resource.\n  You must specify either the ` + "`" + `` + "`" + `ScalingTargetId` + "`" + `` + "`" + ` property, or the ` + "`" + `` + "`" + `ResourceId` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `ScalableDimension` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `ServiceNamespace` + "`" + `` + "`" + ` properties, but not both.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_namespace": {
        "computed": true,
        "description": "The namespace of the AWS service that provides the resource, or a ` + "`" + `` + "`" + `custom-resource` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "step_scaling_policy_configuration": {
        "computed": true,
        "description": "A step scaling policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "adjustment_type": {
              "computed": true,
              "description": "Specifies whether the ` + "`" + `` + "`" + `ScalingAdjustment` + "`" + `` + "`" + ` value in the ` + "`" + `` + "`" + `StepAdjustment` + "`" + `` + "`" + ` property is an absolute number or a percentage of the current capacity.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cooldown": {
              "computed": true,
              "description": "The amount of time, in seconds, to wait for a previous scaling activity to take effect. If not specified, the default value is 300. For more information, see [Cooldown period](https://docs.aws.amazon.com/autoscaling/application/userguide/step-scaling-policy-overview.html#step-scaling-cooldown) in the *Application Auto Scaling User Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "metric_aggregation_type": {
              "computed": true,
              "description": "The aggregation type for the CloudWatch metrics. Valid values are ` + "`" + `` + "`" + `Minimum` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Maximum` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `Average` + "`" + `` + "`" + `. If the aggregation type is null, the value is treated as ` + "`" + `` + "`" + `Average` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_adjustment_magnitude": {
              "computed": true,
              "description": "The minimum value to scale by when the adjustment type is ` + "`" + `` + "`" + `PercentChangeInCapacity` + "`" + `` + "`" + `. For example, suppose that you create a step scaling policy to scale out an Amazon ECS service by 25 percent and you specify a ` + "`" + `` + "`" + `MinAdjustmentMagnitude` + "`" + `` + "`" + ` of 2. If the service has 4 tasks and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a ` + "`" + `` + "`" + `MinAdjustmentMagnitude` + "`" + `` + "`" + ` of 2, Application Auto Scaling scales out the service by 2 tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "step_adjustments": {
              "computed": true,
              "description": "A set of adjustments that enable you to scale based on the size of the alarm breach.\n At least one step adjustment is required if you are adding a new step scaling policy configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "metric_interval_lower_bound": {
                    "computed": true,
                    "description": "The lower bound for the difference between the alarm threshold and the CloudWatch metric. If the metric value is above the breach threshold, the lower bound is inclusive (the metric must be greater than or equal to the threshold plus the lower bound). Otherwise, it is exclusive (the metric must be greater than the threshold plus the lower bound). A null value indicates negative infinity.\n You must specify at least one upper or lower bound.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "metric_interval_upper_bound": {
                    "computed": true,
                    "description": "The upper bound for the difference between the alarm threshold and the CloudWatch metric. If the metric value is above the breach threshold, the upper bound is exclusive (the metric must be less than the threshold plus the upper bound). Otherwise, it is inclusive (the metric must be less than or equal to the threshold plus the upper bound). A null value indicates positive infinity.\n You must specify at least one upper or lower bound.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "scaling_adjustment": {
                    "computed": true,
                    "description": "The amount by which to scale. The adjustment is based on the value that you specified in the ` + "`" + `` + "`" + `AdjustmentType` + "`" + `` + "`" + ` property (either an absolute number or a percentage). A positive value adds to the current capacity and a negative number subtracts from the current capacity.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "target_tracking_scaling_policy_configuration": {
        "computed": true,
        "description": "A target tracking scaling policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "customized_metric_specification": {
              "computed": true,
              "description": "A customized metric. You can specify either a predefined metric or a customized metric.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dimensions": {
                    "computed": true,
                    "description": "The dimensions of the metric. \n Conditional: If you published your metric with dimensions, you must specify the same dimensions in your scaling policy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the dimension.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value of the dimension.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "metric_name": {
                    "computed": true,
                    "description": "The name of the metric. To get the exact metric name, namespace, and dimensions, inspect the [Metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html) object that's returned by a call to [ListMetrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "metrics": {
                    "computed": true,
                    "description": "The metrics to include in the target tracking scaling policy, as a metric data query. This can include both raw metric and metric math expressions.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "expression": {
                          "computed": true,
                          "description": "The math expression to perform on the returned data, if this object is performing a math expression. This expression can use the ` + "`" + `` + "`" + `Id` + "`" + `` + "`" + ` of the other metrics to refer to those metrics, and can also use the ` + "`" + `` + "`" + `Id` + "`" + `` + "`" + ` of other expressions to use the result of those expressions. \n Conditional: Within each ` + "`" + `` + "`" + `TargetTrackingMetricDataQuery` + "`" + `` + "`" + ` object, you must specify either ` + "`" + `` + "`" + `Expression` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MetricStat` + "`" + `` + "`" + `, but not both.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description": "A short name that identifies the object's results in the response. This name must be unique among all ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` objects specified for a single scaling policy. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscores. The first character must be a lowercase letter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "label": {
                          "computed": true,
                          "description": "A human-readable label for this metric or expression. This is especially useful if this is a math expression, so that you know what the value represents.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "metric_stat": {
                          "computed": true,
                          "description": "Information about the metric data to return.\n Conditional: Within each ` + "`" + `` + "`" + `MetricDataQuery` + "`" + `` + "`" + ` object, you must specify either ` + "`" + `` + "`" + `Expression` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MetricStat` + "`" + `` + "`" + `, but not both.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "metric": {
                                "computed": true,
                                "description": "The CloudWatch metric to return, including the metric name, namespace, and dimensions. To get the exact metric name, namespace, and dimensions, inspect the [Metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html) object that is returned by a call to [ListMetrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html).",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimensions": {
                                      "computed": true,
                                      "description": "The dimensions for the metric. For the list of available dimensions, see the AWS documentation available from the table in [services that publish CloudWatch metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) in the *Amazon CloudWatch User Guide*. \n Conditional: If you published your metric with dimensions, you must specify the same dimensions in your scaling policy.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "The name of the dimension.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description": "The value of the dimension.",
                                            "description_kind": "plain",
                                            "optional": true,
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
                                      "description": "The namespace of the metric. For more information, see the table in [services that publish CloudWatch metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) in the *Amazon CloudWatch User Guide*.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "stat": {
                                "computed": true,
                                "description": "The statistic to return. It can include any CloudWatch statistic or extended statistic. For a list of valid values, see the table in [Statistics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Statistic) in the *Amazon CloudWatch User Guide*.\n The most commonly used metric for scaling is ` + "`" + `` + "`" + `Average` + "`" + `` + "`" + `.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "unit": {
                                "computed": true,
                                "description": "The unit to use for the returned data points. For a complete list of the units that CloudWatch supports, see the [MetricDatum](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html) data type in the *Amazon CloudWatch API Reference*.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "return_data": {
                          "computed": true,
                          "description": "Indicates whether to return the timestamps and raw data values of this metric. \n If you use any math expressions, specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` for this value for only the final math expression that the metric specification is based on. You must specify ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `ReturnData` + "`" + `` + "`" + ` for all the other metrics and expressions used in the metric specification.\n If you are only retrieving metrics and not performing any math expressions, do not specify anything for ` + "`" + `` + "`" + `ReturnData` + "`" + `` + "`" + `. This sets it to its default (` + "`" + `` + "`" + `true` + "`" + `` + "`" + `).",
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
                    "description": "The namespace of the metric.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "statistic": {
                    "computed": true,
                    "description": "The statistic of the metric.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "unit": {
                    "computed": true,
                    "description": "The unit of the metric. For a complete list of the units that CloudWatch supports, see the [MetricDatum](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html) data type in the *Amazon CloudWatch API Reference*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "disable_scale_in": {
              "computed": true,
              "description": "Indicates whether scale in by the target tracking scaling policy is disabled. If the value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, scale in is disabled and the target tracking scaling policy won't remove capacity from the scalable target. Otherwise, scale in is enabled and the target tracking scaling policy can remove capacity from the scalable target. The default value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "predefined_metric_specification": {
              "computed": true,
              "description": "A predefined metric. You can specify either a predefined metric or a customized metric.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "predefined_metric_type": {
                    "computed": true,
                    "description": "The metric type. The ` + "`" + `` + "`" + `ALBRequestCountPerTarget` + "`" + `` + "`" + ` metric type applies only to Spot fleet requests and ECS services.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resource_label": {
                    "computed": true,
                    "description": "Identifies the resource associated with the metric type. You can't specify a resource label unless the metric type is ` + "`" + `` + "`" + `ALBRequestCountPerTarget` + "`" + `` + "`" + ` and there is a target group attached to the Spot Fleet or ECS service.\n You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). The format of the resource label is:\n ` + "`" + `` + "`" + `app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff` + "`" + `` + "`" + `.\n Where:\n  +  app/\u003cload-balancer-name\u003e/\u003cload-balancer-id\u003e is the final portion of the load balancer ARN\n  +  targetgroup/\u003ctarget-group-name\u003e/\u003ctarget-group-id\u003e is the final portion of the target group ARN.\n  \n To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operation. To find the ARN for the target group, use the [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) API operation.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "scale_in_cooldown": {
              "computed": true,
              "description": "The amount of time, in seconds, after a scale-in activity completes before another scale-in activity can start. For more information and for default values, see [Define cooldown periods](https://docs.aws.amazon.com/autoscaling/application/userguide/target-tracking-scaling-policy-overview.html#target-tracking-cooldown) in the *Application Auto Scaling User Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "scale_out_cooldown": {
              "computed": true,
              "description": "The amount of time, in seconds, to wait for a previous scale-out activity to take effect. For more information and for default values, see [Define cooldown periods](https://docs.aws.amazon.com/autoscaling/application/userguide/target-tracking-scaling-policy-overview.html#target-tracking-cooldown) in the *Application Auto Scaling User Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_value": {
              "computed": true,
              "description": "The target value for the metric. Although this property accepts numbers of type Double, it won't accept values that are either too small or too large. Values must be in the range of -2^360 to 2^360. The value must be a valid number based on the choice of metric. For example, if the metric is CPU utilization, then the target value is a percent value that represents how much of the CPU can be used before scaling out.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApplicationAutoScaling::ScalingPolicy` + "`" + `` + "`" + ` resource defines a scaling policy that Application Auto Scaling uses to adjust the capacity of a scalable target. \n For more information, see [Target tracking scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the *Application Auto Scaling User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApplicationautoscalingScalingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApplicationautoscalingScalingPolicy), &result)
	return &result
}
