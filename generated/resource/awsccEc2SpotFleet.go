package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SpotFleet = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "spot_fleet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "spot_fleet_request_config_data": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allocation_strategy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "context": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "excess_capacity_termination_policy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "iam_fleet_role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "instance_interruption_behavior": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_pools_to_use_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "launch_specifications": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "block_device_mappings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "device_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "ebs": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "delete_on_termination": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "encrypted": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "iops": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "snapshot_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "volume_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "volume_type": {
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
                        "no_device": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "virtual_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "ebs_optimized": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "iam_instance_profile": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "arn": {
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
                  "image_id": {
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
                        "baseline_performance_factors": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "cpu": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "references": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "instance_family": {
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
                                  "nesting_mode": "single"
                                },
                                "optional": true
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
                  "instance_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kernel_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "monitoring": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "network_interfaces": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "associate_public_ip_address": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "delete_on_termination": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "description": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "device_index": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "groups": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "ipv_6_address_count": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "ipv_6_addresses": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "ipv_6_address": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "network_interface_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "private_ip_addresses": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "primary": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "private_ip_address": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "secondary_private_ip_address_count": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "subnet_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "placement": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "availability_zone": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "group_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "tenancy": {
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
                  "ramdisk_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "security_groups": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "group_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "spot_price": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnet_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tag_specifications": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "resource_type": {
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
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "user_data": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "weighted_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
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
                          "optional": true,
                          "type": "string"
                        },
                        "launch_template_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "version": {
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
                  "overrides": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "availability_zone": {
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
                              "baseline_performance_factors": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "cpu": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "references": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "instance_family": {
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
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
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
                        "instance_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "priority": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "spot_price": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "subnet_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "weighted_capacity": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "load_balancers_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "classic_load_balancers_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "classic_load_balancers": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
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
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "target_groups_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "target_groups": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "arn": {
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
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "on_demand_allocation_strategy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "on_demand_max_total_price": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "on_demand_target_capacity": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "replace_unhealthy_instances": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "spot_maintenance_strategies": {
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
                          "optional": true,
                          "type": "string"
                        },
                        "termination_delay": {
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
            "spot_max_total_price": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "spot_price": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_specifications": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "resource_type": {
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
                "nesting_mode": "list"
              },
              "optional": true
            },
            "target_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "target_capacity_unit_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "terminate_instances_with_expiration": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "valid_from": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "valid_until": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Resource Type definition for AWS::EC2::SpotFleet",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SpotFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SpotFleet), &result)
	return &result
}
