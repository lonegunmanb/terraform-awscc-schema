package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApplicationautoscalingScalingPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "ARN is a read only property for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_name": {
        "computed": true,
        "description": "The name of the scaling policy.\n\nUpdates to the name of a target tracking scaling policy are not supported, unless you also update the metric used for scaling. To change only a target tracking scaling policy's name, first delete the policy by removing the existing AWS::ApplicationAutoScaling::ScalingPolicy resource from the template and updating the stack. Then, recreate the resource with the same settings and a different name.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_type": {
        "computed": true,
        "description": "The scaling policy type.\n\nThe following policy types are supported:\n\nTargetTrackingScaling Not supported for Amazon EMR\n\nStepScaling Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "scalable_dimension": {
        "computed": true,
        "description": "The scalable dimension. This string consists of the service namespace, resource type, and scaling property.",
        "description_kind": "plain",
        "type": "string"
      },
      "scaling_target_id": {
        "computed": true,
        "description": "The CloudFormation-generated ID of an Application Auto Scaling scalable target. For more information about the ID, see the Return Value section of the AWS::ApplicationAutoScaling::ScalableTarget resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_namespace": {
        "computed": true,
        "description": "The namespace of the AWS service that provides the resource, or a custom-resource.",
        "description_kind": "plain",
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
              "description": "Specifies how the ScalingAdjustment value in a StepAdjustment is interpreted.",
              "description_kind": "plain",
              "type": "string"
            },
            "cooldown": {
              "computed": true,
              "description": "The amount of time, in seconds, to wait for a previous scaling activity to take effect.",
              "description_kind": "plain",
              "type": "number"
            },
            "metric_aggregation_type": {
              "computed": true,
              "description": "The aggregation type for the CloudWatch metrics. Valid values are Minimum, Maximum, and Average. If the aggregation type is null, the value is treated as Average",
              "description_kind": "plain",
              "type": "string"
            },
            "min_adjustment_magnitude": {
              "computed": true,
              "description": "The minimum value to scale by when the adjustment type is PercentChangeInCapacity.",
              "description_kind": "plain",
              "type": "number"
            },
            "step_adjustments": {
              "computed": true,
              "description": "A set of adjustments that enable you to scale based on the size of the alarm breach.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "metric_interval_lower_bound": {
                    "computed": true,
                    "description": "The lower bound for the difference between the alarm threshold and the CloudWatch metric. If the metric value is above the breach threshold, the lower bound is inclusive (the metric must be greater than or equal to the threshold plus the lower bound). Otherwise, it is exclusive (the metric must be greater than the threshold plus the lower bound). A null value indicates negative infinity.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "metric_interval_upper_bound": {
                    "computed": true,
                    "description": "The upper bound for the difference between the alarm threshold and the CloudWatch metric. If the metric value is above the breach threshold, the upper bound is exclusive (the metric must be less than the threshold plus the upper bound). Otherwise, it is inclusive (the metric must be less than or equal to the threshold plus the upper bound). A null value indicates positive infinity.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "scaling_adjustment": {
                    "computed": true,
                    "description": "The amount by which to scale, based on the specified adjustment type. A positive value adds to the current capacity while a negative number removes from the current capacity. For exact capacity, you must specify a positive value.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "set"
              }
            }
          },
          "nesting_mode": "single"
        }
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
                    "description": "The dimensions of the metric.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the dimension.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value of the dimension.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "metric_name": {
                    "computed": true,
                    "description": "The name of the metric. To get the exact metric name, namespace, and dimensions, inspect the Metric object that is returned by a call to ListMetrics.",
                    "description_kind": "plain",
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
                          "description": "The math expression to perform on the returned data, if this object is performing a math expression.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description": "A short name that identifies the object's results in the response.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "label": {
                          "computed": true,
                          "description": "A human-readable label for this metric or expression. This is especially useful if this is a math expression, so that you know what the value represents.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_stat": {
                          "computed": true,
                          "description": "Information about the metric data to return.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "metric": {
                                "computed": true,
                                "description": "The CloudWatch metric to return, including the metric name, namespace, and dimensions. ",
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
                                            "computed": true,
                                            "description": "The name of the dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description": "The value of the dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      }
                                    },
                                    "metric_name": {
                                      "computed": true,
                                      "description": "The name of the metric.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "namespace": {
                                      "computed": true,
                                      "description": "The namespace of the metric.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "stat": {
                                "computed": true,
                                "description": "The statistic to return. It can include any CloudWatch statistic or extended statistic.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "unit": {
                                "computed": true,
                                "description": "The unit to use for the returned data points.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "return_data": {
                          "computed": true,
                          "description": "Indicates whether to return the timestamps and raw data values of this metric.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "namespace": {
                    "computed": true,
                    "description": "The namespace of the metric.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "statistic": {
                    "computed": true,
                    "description": "The statistic of the metric.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "unit": {
                    "computed": true,
                    "description": "The unit of the metric. For a complete list of the units that CloudWatch supports, see the MetricDatum data type in the Amazon CloudWatch API Reference.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "disable_scale_in": {
              "computed": true,
              "description": "Indicates whether scale in by the target tracking scaling policy is disabled. If the value is true, scale in is disabled and the target tracking scaling policy won't remove capacity from the scalable target. Otherwise, scale in is enabled and the target tracking scaling policy can remove capacity from the scalable target. The default value is false.",
              "description_kind": "plain",
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
                    "description": "The metric type. The ALBRequestCountPerTarget metric type applies only to Spot Fleets and ECS services.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_label": {
                    "computed": true,
                    "description": "Identifies the resource associated with the metric type. You can't specify a resource label unless the metric type is ALBRequestCountPerTarget and there is a target group attached to the Spot Fleet or ECS service.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "scale_in_cooldown": {
              "computed": true,
              "description": "The amount of time, in seconds, after a scale-in activity completes before another scale-in activity can start.",
              "description_kind": "plain",
              "type": "number"
            },
            "scale_out_cooldown": {
              "computed": true,
              "description": "The amount of time, in seconds, to wait for a previous scale-out activity to take effect.",
              "description_kind": "plain",
              "type": "number"
            },
            "target_value": {
              "computed": true,
              "description": "The target value for the metric. Although this property accepts numbers of type Double, it won't accept values that are either too small or too large. Values must be in the range of -2^360 to 2^360. The value must be a valid number based on the choice of metric. For example, if the metric is CPU utilization, then the target value is a percent value that represents how much of the CPU can be used before scaling out.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::ApplicationAutoScaling::ScalingPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApplicationautoscalingScalingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApplicationautoscalingScalingPolicy), &result)
	return &result
}
