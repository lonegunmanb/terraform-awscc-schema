package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAutoscalingplansScalingPlan = `{
  "block": {
    "attributes": {
      "application_source": {
        "description": "A CloudFormation stack or a set of tags. You can create one scaling plan per application source. The ApplicationSource property must be present to ensure interoperability with the AWS Auto Scaling console.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudformation_stack_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of a CloudFormation stack. The stack must contain at least one scalable resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_filters": {
              "computed": true,
              "description": "A set of tag filters (keys and values). Each tag filter specified must contain a key with values as optional. Each scaling plan can include up to 50 keys, and each key can include up to 20 values.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The tag key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description": "The tag values (0 to 20).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "scaling_instructions": {
        "description": "The scaling instructions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "customized_load_metric_specification": {
              "computed": true,
              "description": "The customized load metric to use for predictive scaling. This property or a PredefinedLoadMetricSpecification is required when configuring predictive scaling, and cannot be used otherwise.",
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
                    "description": "The unit of the metric.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "disable_dynamic_scaling": {
              "computed": true,
              "description": "Controls whether dynamic scaling is disabled. When dynamic scaling is enabled, AWS Auto Scaling creates target tracking scaling policies based on the specified target tracking configurations. The default is enabled (false).",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_capacity": {
              "description": "The maximum capacity of the resource. The exception to this upper limit is if you specify a non-default setting for PredictiveScalingMaxCapacityBehavior.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_capacity": {
              "description": "The minimum capacity of the resource.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "predefined_load_metric_specification": {
              "computed": true,
              "description": "The predefined load metric to use for predictive scaling. This property or a CustomizedLoadMetricSpecification is required when configuring predictive scaling, and cannot be used otherwise.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "predefined_load_metric_type": {
                    "computed": true,
                    "description": "The metric type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resource_label": {
                    "computed": true,
                    "description": "Identifies the resource associated with the metric type. You can't specify a resource label unless the metric type is ALBTargetGroupRequestCount and there is a target group for an Application Load Balancer attached to the Auto Scaling group.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "predictive_scaling_max_capacity_behavior": {
              "computed": true,
              "description": "Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity specified for the resource. The default value is SetForecastCapacityToMaxCapacity.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "predictive_scaling_max_capacity_buffer": {
              "computed": true,
              "description": "The size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity. The value is specified as a percentage relative to the forecast capacity. For example, if the buffer is 10, this means a 10 percent buffer. The range is 1-100. Required if PredictiveScalingMaxCapacityBehavior is set to SetMaxCapacityAboveForecastCapacity, and cannot be used otherwise.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "predictive_scaling_mode": {
              "computed": true,
              "description": "The predictive scaling mode. The default value is ForecastAndScale. Otherwise, AWS Auto Scaling forecasts capacity but does not apply any scheduled scaling actions based on the capacity forecast.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_id": {
              "description": "The ID of the resource. This string consists of the resource type and unique identifier.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scalable_dimension": {
              "description": "The scalable dimension associated with the resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scaling_policy_update_behavior": {
              "computed": true,
              "description": "Controls whether a resource's externally created scaling policies are deleted and new target tracking scaling policies created. The default value is KeepExternalPolicies. Valid only when configuring dynamic scaling.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scheduled_action_buffer_time": {
              "computed": true,
              "description": "The amount of time, in seconds, to buffer the run time of scheduled scaling actions when scaling out. The value must be less than the forecast interval duration of 3600 seconds (60 minutes). The default is 300 seconds. Valid only when configuring predictive scaling.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "service_namespace": {
              "description": "The namespace of the AWS service.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_tracking_configurations": {
              "description": "The target tracking configurations (up to 10). Each of these structures must specify a unique scaling metric and a target value for the metric.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "customized_scaling_metric_specification": {
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
                          "description": "The unit of the metric.",
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
                    "description": "Indicates whether scale in by the target tracking scaling policy is disabled. If the value is true, scale in is disabled and the target tracking scaling policy doesn't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking scaling policy can remove capacity from the scalable resource. The default value is false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "estimated_instance_warmup": {
                    "computed": true,
                    "description": "The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics. This value is used only if the resource is an Auto Scaling group.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "predefined_scaling_metric_specification": {
                    "computed": true,
                    "description": "A predefined metric. You can specify either a predefined metric or a customized metric.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "predefined_scaling_metric_type": {
                          "computed": true,
                          "description": "The metric type. The ALBRequestCountPerTarget metric type applies only to Auto Scaling groups, Spot Fleet requests, and ECS services.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_label": {
                          "computed": true,
                          "description": "Identifies the resource associated with the metric type. You can't specify a resource label unless the metric type is ALBRequestCountPerTarget and there is a target group for an Application Load Balancer attached to the Auto Scaling group, Spot Fleet request, or ECS service.",
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
                    "description": "The amount of time, in seconds, after a scale-in activity completes before another scale in activity can start. This value is not used if the scalable resource is an Auto Scaling group.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "scale_out_cooldown": {
                    "computed": true,
                    "description": "The amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start. This value is not used if the scalable resource is an Auto Scaling group.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_value": {
                    "description": "The target value for the metric.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "scaling_plan_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the scaling plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "scaling_plan_name": {
        "computed": true,
        "description": "The name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.",
        "description_kind": "plain",
        "type": "string"
      },
      "scaling_plan_version": {
        "computed": true,
        "description": "The version number of the scaling plan. This value is always 1.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AutoScalingPlans::ScalingPlan",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAutoscalingplansScalingPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAutoscalingplansScalingPlan), &result)
	return &result
}
