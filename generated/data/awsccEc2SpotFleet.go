package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SpotFleet = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "spot_fleet_request_config_data": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allocation_strategy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
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
            "iam_fleet_role": {
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
                                "type": "bool"
                              },
                              "encrypted": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "iops": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "snapshot_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "volume_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "volume_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "no_device": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "virtual_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "ebs_optimized": {
                    "computed": true,
                    "description_kind": "plain",
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
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "image_id": {
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
                  "kernel_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_name": {
                    "computed": true,
                    "description_kind": "plain",
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
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "network_interfaces": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "associate_public_ip_address": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "delete_on_termination": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "description": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "device_index": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "groups": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "ipv_6_address_count": {
                          "computed": true,
                          "description_kind": "plain",
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
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "network_interface_id": {
                          "computed": true,
                          "description_kind": "plain",
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
                                "type": "bool"
                              },
                              "private_ip_address": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "secondary_private_ip_address_count": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "subnet_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "placement": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
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
                        "tenancy": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "ramdisk_id": {
                    "computed": true,
                    "description_kind": "plain",
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
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "spot_price": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subnet_id": {
                    "computed": true,
                    "description_kind": "plain",
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
                  "user_data": {
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
                        "priority": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "spot_price": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
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
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
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
            "on_demand_allocation_strategy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "on_demand_max_total_price": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "on_demand_target_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "replace_unhealthy_instances": {
              "computed": true,
              "description_kind": "plain",
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
            "spot_max_total_price": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "spot_price": {
              "computed": true,
              "description_kind": "plain",
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
            "target_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "target_capacity_unit_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::EC2::SpotFleet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2SpotFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SpotFleet), &result)
	return &result
}
