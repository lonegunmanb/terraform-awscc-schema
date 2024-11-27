package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAutoscalingScalingPolicy = `{
  "block": {
    "attributes": {
      "adjustment_type": {
        "computed": true,
        "description": "Specifies how the scaling adjustment is interpreted. The valid values are ChangeInCapacity, ExactCapacity, and PercentChangeInCapacity.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the AutoScaling scaling policy",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_scaling_group_name": {
        "computed": true,
        "description": "The name of the Auto Scaling group.",
        "description_kind": "plain",
        "type": "string"
      },
      "cooldown": {
        "computed": true,
        "description": "The duration of the policy's cooldown period, in seconds. When a cooldown period is specified here, it overrides the default cooldown period defined for the Auto Scaling group.",
        "description_kind": "plain",
        "type": "string"
      },
      "estimated_instance_warmup": {
        "computed": true,
        "description": "The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics. If not provided, the default is to use the value from the default cooldown period for the Auto Scaling group. Valid only if the policy type is TargetTrackingScaling or StepScaling.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metric_aggregation_type": {
        "computed": true,
        "description": "The aggregation type for the CloudWatch metrics. The valid values are Minimum, Maximum, and Average. If the aggregation type is null, the value is treated as Average. Valid only if the policy type is StepScaling.",
        "description_kind": "plain",
        "type": "string"
      },
      "min_adjustment_magnitude": {
        "computed": true,
        "description": "The minimum value to scale by when the adjustment type is PercentChangeInCapacity. For example, suppose that you create a step scaling policy to scale out an Auto Scaling group by 25 percent and you specify a MinAdjustmentMagnitude of 2. If the group has 4 instances and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a MinAdjustmentMagnitude of 2, Amazon EC2 Auto Scaling scales out the group by 2 instances.",
        "description_kind": "plain",
        "type": "number"
      },
      "policy_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_type": {
        "computed": true,
        "description": "One of the following policy types: TargetTrackingScaling, StepScaling, SimpleScaling (default), PredictiveScaling",
        "description_kind": "plain",
        "type": "string"
      },
      "predictive_scaling_configuration": {
        "computed": true,
        "description": "A predictive scaling policy. Includes support for predefined metrics only.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_capacity_breach_behavior": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "max_capacity_buffer": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "metric_specifications": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "customized_capacity_metric_specification": {
                    "computed": true,
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
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "label": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "metric_stat": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "metric": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimensions": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "name": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "set"
                                            }
                                          },
                                          "metric_name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "namespace": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "stat": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "unit": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "return_data": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "set"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "customized_load_metric_specification": {
                    "computed": true,
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
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "label": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "metric_stat": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "metric": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimensions": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "name": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "set"
                                            }
                                          },
                                          "metric_name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "namespace": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "stat": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "unit": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "return_data": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "set"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "customized_scaling_metric_specification": {
                    "computed": true,
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
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "label": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "metric_stat": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "metric": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimensions": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "name": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "set"
                                            }
                                          },
                                          "metric_name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "namespace": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "stat": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "unit": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "return_data": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "set"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "predefined_load_metric_specification": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "predefined_metric_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "resource_label": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "predefined_metric_pair_specification": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "predefined_metric_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "resource_label": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "predefined_scaling_metric_specification": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "predefined_metric_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "resource_label": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "target_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "scheduling_buffer_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "scaling_adjustment": {
        "computed": true,
        "description": "The amount by which to scale, based on the specified adjustment type. A positive value adds to the current capacity while a negative number removes from the current capacity. For exact capacity, you must specify a positive value. Required if the policy type is SimpleScaling. (Not used with any other policy type.)",
        "description_kind": "plain",
        "type": "number"
      },
      "step_adjustments": {
        "computed": true,
        "description": "A set of adjustments that enable you to scale based on the size of the alarm breach. Required if the policy type is StepScaling. (Not used with any other policy type.)",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "metric_interval_lower_bound": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "metric_interval_upper_bound": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "scaling_adjustment": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "set"
        }
      },
      "target_tracking_configuration": {
        "computed": true,
        "description": "A target tracking scaling policy. Includes support for predefined or customized metrics.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "customized_metric_specification": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dimensions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    }
                  },
                  "metric_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "metrics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "expression": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "label": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_stat": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "metric": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimensions": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "set"
                                      }
                                    },
                                    "metric_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "namespace": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "period": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "stat": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "unit": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "period": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "return_data": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "set"
                    }
                  },
                  "namespace": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "period": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "statistic": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "unit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "disable_scale_in": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "predefined_metric_specification": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "predefined_metric_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_label": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "target_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::AutoScaling::ScalingPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAutoscalingScalingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAutoscalingScalingPolicy), &result)
	return &result
}
