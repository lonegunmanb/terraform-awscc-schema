package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsCapacityProvider = `{
  "block": {
    "attributes": {
      "auto_scaling_group_provider": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_scaling_group_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "managed_draining": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "managed_scaling": {
              "computed": true,
              "description": "The managed scaling settings for the Auto Scaling group capacity provider.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_warmup_period": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_scaling_step_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "minimum_scaling_step_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "managed_termination_protection": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "cluster_name": {
        "computed": true,
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
      "managed_instances_provider": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "infrastructure_optimization": {
              "computed": true,
              "description": "Defines how Amazon ECS Managed Instances optimizes the infrastructure in your capacity provider. Configure it to turn on or off the infrastructure optimization in your capacity provider, and to control the idle EC2 instances optimization delay.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "scale_in_after": {
                    "computed": true,
                    "description": "This parameter defines the number of seconds Amazon ECS Managed Instances waits before optimizing EC2 instances that have become idle or underutilized. A longer delay increases the likelihood of placing new tasks on idle instances, reducing startup time. A shorter delay helps reduce infrastructure costs by optimizing idle instances more quickly. Valid values are: Not set (null) - Uses the default optimization behavior, ` + "`" + `-1` + "`" + ` - Disables automatic infrastructure optimization, ` + "`" + `0` + "`" + ` to ` + "`" + `3600` + "`" + ` (inclusive) - Specifies the number of seconds to wait before optimizing instances.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "infrastructure_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_launch_template": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ec_2_instance_profile_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
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
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "accelerator_manufacturers": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "accelerator_names": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
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
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "accelerator_types": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allowed_instance_types": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "bare_metal": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
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
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "burstable_performance": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cpu_manufacturers": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "excluded_instance_types": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "instance_generations": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "local_storage": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "local_storage_types": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "max_spot_price_as_percentage_of_optimal_on_demand_price": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "memory_gi_b_per_v_cpu": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "memory_mi_b": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "network_bandwidth_gbps": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "network_interface_count": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "on_demand_max_price_percentage_over_lowest_price": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "require_hibernate_support": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "spot_max_price_percentage_over_lowest_price": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
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
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "v_cpu_count": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "computed": true,
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
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "monitoring": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "network_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "security_groups": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "subnets": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "storage_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "storage_size_gi_b": {
                          "computed": true,
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "propagate_tags": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::ECS::CapacityProvider.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcsCapacityProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsCapacityProvider), &result)
	return &result
}
