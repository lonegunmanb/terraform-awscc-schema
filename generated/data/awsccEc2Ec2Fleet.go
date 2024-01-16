package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Ec2Fleet = `{
  "block": {
    "attributes": {
      "context": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "excess_capacity_termination_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fleet_id": {
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
      "launch_template_configs": {
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
                  "availability_zone": {
                    "computed": true,
                    "description_kind": "plain",
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
                            "list",
                            "string"
                          ]
                        },
                        "accelerator_names": {
                          "computed": true,
                          "description_kind": "plain",
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
                            "list",
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
                            "list",
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
                            "list",
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
                            "list",
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
                  "max_price": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "placement": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "affinity": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "availability_zone": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "group_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "host_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "host_resource_group_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "partition_number": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "spread_domain": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "tenancy": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "priority": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "subnet_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "weighted_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "on_demand_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allocation_strategy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "capacity_reservation_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "usage_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "max_total_price": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "min_target_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "single_availability_zone": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "single_instance_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "replace_unhealthy_instances": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "spot_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allocation_strategy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_interruption_behavior": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_pools_to_use_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "maintenance_strategies": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity_rebalance": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "replacement_strategy": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "termination_delay": {
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
            "max_total_price": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "min_target_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "single_availability_zone": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "single_instance_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tag_specifications": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_type": {
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
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_capacity_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_target_capacity_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "on_demand_target_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "spot_target_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "target_capacity_unit_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "total_target_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "terminate_instances_with_expiration": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "valid_from": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "valid_until": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::EC2Fleet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2Ec2FleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Ec2Fleet), &result)
	return &result
}
