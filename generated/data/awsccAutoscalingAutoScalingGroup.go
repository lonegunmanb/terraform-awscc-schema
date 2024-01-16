package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAutoscalingAutoScalingGroup = `{
  "block": {
    "attributes": {
      "auto_scaling_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "capacity_rebalance": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "context": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cooldown": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_instance_warmup": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "desired_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "desired_capacity_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "health_check_grace_period": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "health_check_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_maintenance_policy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_healthy_percentage": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "min_healthy_percentage": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "launch_configuration_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "launch_template": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "launch_template_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "launch_template_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "lifecycle_hook_specification_list": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_result": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "heartbeat_timeout": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "lifecycle_hook_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "lifecycle_transition": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "notification_metadata": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "notification_target_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "load_balancer_names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "max_instance_lifetime": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metrics_collection": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "granularity": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "metrics": {
              "computed": true,
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
      "min_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mixed_instances_policy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "instances_distribution": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "on_demand_allocation_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "on_demand_base_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "on_demand_percentage_above_base_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "spot_allocation_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "spot_instance_pools": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "spot_max_price": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "launch_template": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "launch_template_specification": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "launch_template_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "launch_template_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "version": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "overrides": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_requirements": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "accelerator_count": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "accelerator_manufacturers": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "accelerator_names": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "accelerator_total_memory_mi_b": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "accelerator_types": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "allowed_instance_types": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "bare_metal": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "baseline_ebs_bandwidth_mbps": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "burstable_performance": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "cpu_manufacturers": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "excluded_instance_types": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "instance_generations": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "local_storage": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "local_storage_types": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "memory_gi_b_per_v_cpu": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "memory_mi_b": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "network_bandwidth_gbps": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "network_interface_count": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "on_demand_max_price_percentage_over_lowest_price": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "require_hibernate_support": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "spot_max_price_percentage_over_lowest_price": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "total_local_storage_gb": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "v_cpu_count": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "instance_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "launch_template_specification": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "launch_template_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "launch_template_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "version": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "weighted_capacity": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "new_instances_protected_from_scale_in": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "notification_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "notification_types": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "topic_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "notification_configurations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "notification_types": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "topic_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "placement_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_linked_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "propagate_at_launch": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_group_ar_ns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "termination_policies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_zone_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::AutoScaling::AutoScalingGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAutoscalingAutoScalingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAutoscalingAutoScalingGroup), &result)
	return &result
}
