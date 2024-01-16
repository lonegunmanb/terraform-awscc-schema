package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2LaunchTemplate = `{
  "block": {
    "attributes": {
      "default_version_number": {
        "computed": true,
        "description": "The default version of the launch template",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "latest_version_number": {
        "computed": true,
        "description": "The latest version of the launch template",
        "description_kind": "plain",
        "type": "string"
      },
      "launch_template_data": {
        "description": "The information for the launch template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "block_device_mappings": {
              "computed": true,
              "description": "The block device mapping.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "device_name": {
                    "computed": true,
                    "description": "The user data to make available to the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ebs": {
                    "computed": true,
                    "description": "Parameters for a block device for an EBS volume in an Amazon EC2 launch template.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delete_on_termination": {
                          "computed": true,
                          "description": "Indicates whether the EBS volume is deleted on instance termination.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "encrypted": {
                          "computed": true,
                          "description": "Indicates whether the EBS volume is encrypted. Encrypted volumes can only be attached to instances that support Amazon EBS encryption. If you are creating a volume from a snapshot, you can't specify an encryption value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "iops": {
                          "computed": true,
                          "description": "The number of I/O operations per second (IOPS).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "kms_key_id": {
                          "computed": true,
                          "description": "The ARN of the symmetric AWS Key Management Service (AWS KMS) CMK used for encryption.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "snapshot_id": {
                          "computed": true,
                          "description": "The ID of the snapshot.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "throughput": {
                          "computed": true,
                          "description": "The throughput to provision for a gp3 volume, with a maximum of 1,000 MiB/s.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "volume_size": {
                          "computed": true,
                          "description": "The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "volume_type": {
                          "computed": true,
                          "description": "The volume type.",
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
                    "description": "To omit the device from the block device mapping, specify an empty string.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "virtual_name": {
                    "computed": true,
                    "description": "The virtual device name (ephemeralN).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "capacity_reservation_specification": {
              "computed": true,
              "description": "Specifies an instance's Capacity Reservation targeting option.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity_reservation_preference": {
                    "computed": true,
                    "description": "Indicates the instance's Capacity Reservation preferences.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "capacity_reservation_target": {
                    "computed": true,
                    "description": "Specifies a target Capacity Reservation.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "capacity_reservation_id": {
                          "computed": true,
                          "description": "The ID of the Capacity Reservation in which to run the instance.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "capacity_reservation_resource_group_arn": {
                          "computed": true,
                          "description": "The ARN of the Capacity Reservation resource group in which to run the instance.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
            "cpu_options": {
              "computed": true,
              "description": "specifies the CPU options for an instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "amd_sev_snp": {
                    "computed": true,
                    "description": "Indicates whether to enable the instance for AMD SEV-SNP. AMD SEV-SNP is supported with M6a, R6a, and C6a instance types only.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "core_count": {
                    "computed": true,
                    "description": "The number of CPU cores for the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "threads_per_core": {
                    "computed": true,
                    "description": "The number of threads per CPU core. To disable multithreading for the instance, specify a value of 1. Otherwise, specify the default value of 2.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "credit_specification": {
              "computed": true,
              "description": "The user data to make available to the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cpu_credits": {
                    "computed": true,
                    "description": "The user data to make available to the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "disable_api_stop": {
              "computed": true,
              "description": "Indicates whether to enable the instance for stop protection.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disable_api_termination": {
              "computed": true,
              "description": "If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ebs_optimized": {
              "computed": true,
              "description": "Indicates whether the instance is optimized for Amazon EBS I/O.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "elastic_gpu_specifications": {
              "computed": true,
              "description": "An elastic GPU to associate with the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description": "The type of Elastic Graphics accelerator.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "elastic_inference_accelerators": {
              "computed": true,
              "description": "The elastic inference accelerator for the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "count": {
                    "computed": true,
                    "description": "The number of elastic inference accelerators to attach to the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description": "The type of elastic inference accelerator.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "enclave_options": {
              "computed": true,
              "description": "Indicates whether the instance is enabled for AWS Nitro Enclaves.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "If this parameter is set to true, the instance is enabled for AWS Nitro Enclaves; otherwise, it is not enabled for AWS Nitro Enclaves.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "hibernation_options": {
              "computed": true,
              "description": "Specifies whether your instance is configured for hibernation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "configured": {
                    "computed": true,
                    "description": "TIf you set this parameter to true, the instance is enabled for hibernation.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "iam_instance_profile": {
              "computed": true,
              "description": "Specifies an IAM instance profile, which is a container for an IAM role for your instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the instance profile.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the instance profile.",
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
              "description": "The ID of the AMI. Alternatively, you can specify a Systems Manager parameter, which will resolve to an AMI ID on launch.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_initiated_shutdown_behavior": {
              "computed": true,
              "description": "Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_market_options": {
              "computed": true,
              "description": "The market (purchasing) option for the instances.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "market_type": {
                    "computed": true,
                    "description": "The market type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spot_options": {
                    "computed": true,
                    "description": "Specifies options for Spot Instances.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "block_duration_minutes": {
                          "computed": true,
                          "description": "Deprecated",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "instance_interruption_behavior": {
                          "computed": true,
                          "description": "The behavior when a Spot Instance is interrupted. The default is terminate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "max_price": {
                          "computed": true,
                          "description": "The maximum hourly price you're willing to pay for the Spot Instances.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "spot_instance_type": {
                          "computed": true,
                          "description": "The Spot Instance request type.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "valid_until": {
                          "computed": true,
                          "description": "The end date of the request, in UTC format (YYYY-MM-DDTHH:MM:SSZ). Supported only for persistent requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
            "instance_requirements": {
              "computed": true,
              "description": "The attributes for the instance types.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "accelerator_count": {
                    "computed": true,
                    "description": "The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferential chips) on an instance.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum number of accelerators.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum number of accelerators.",
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
                    "description": "Indicates whether instance types must have accelerators by specific manufacturers.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "accelerator_names": {
                    "computed": true,
                    "description": "The accelerators that must be on the instance type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "accelerator_total_memory_mi_b": {
                    "computed": true,
                    "description": "The minimum and maximum amount of total accelerator memory, in MiB.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum amount of accelerator memory, in MiB.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum amount of accelerator memory, in MiB.",
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
                    "description": "The accelerator types that must be on the instance type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_instance_types": {
                    "computed": true,
                    "description": "The instance types to apply your specified attributes against.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "bare_metal": {
                    "computed": true,
                    "description": "Indicates whether bare metal instance types must be included, excluded, or required.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "baseline_ebs_bandwidth_mbps": {
                    "computed": true,
                    "description": "The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum baseline bandwidth, in Mbps.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum baseline bandwidth, in Mbps.",
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
                    "description": "The CPU manufacturers to include.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "excluded_instance_types": {
                    "computed": true,
                    "description": "The instance types to exclude.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "instance_generations": {
                    "computed": true,
                    "description": "Indicates whether current or previous generation instance types are included.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "local_storage": {
                    "computed": true,
                    "description": "The user data to make available to the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "local_storage_types": {
                    "computed": true,
                    "description": "The type of local storage that is required.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "memory_gi_b_per_v_cpu": {
                    "computed": true,
                    "description": "The minimum and maximum amount of memory per vCPU, in GiB.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum amount of memory per vCPU, in GiB.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "TThe minimum amount of memory per vCPU, in GiB.",
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
                    "description": "The minimum and maximum amount of memory, in MiB.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum amount of memory, in MiB.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum amount of memory, in MiB.",
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
                    "description": "The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum amount of network bandwidth, in Gbps.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum amount of network bandwidth, in Gbps.",
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
                    "description": "TThe minimum and maximum number of network interfaces.",
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
                    "description": "The price protection threshold for On-Demand Instances.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "require_hibernate_support": {
                    "computed": true,
                    "description": "Indicates whether instance types must support hibernation for On-Demand Instances.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "spot_max_price_percentage_over_lowest_price": {
                    "computed": true,
                    "description": "The price protection threshold for Spot Instances.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "total_local_storage_gb": {
                    "computed": true,
                    "description": "The minimum and maximum amount of total local storage, in GB.",
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
                    "description": "The minimum and maximum number of vCPUs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum number of vCPUs.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum number of vCPUs.",
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
              "description": "The ID of the kernel.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_name": {
              "computed": true,
              "description": "The name of the EC2 key pair",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "license_specifications": {
              "computed": true,
              "description": "The license configurations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "license_configuration_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the license configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "maintenance_options": {
              "computed": true,
              "description": "The maintenance options of your instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auto_recovery": {
                    "computed": true,
                    "description": "Disables the automatic recovery behavior of your instance or sets it to default.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "reboot_migration": {
                    "computed": true,
                    "description": "Disables the automatic reboot-migration behavior of your instance or sets it to default.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "metadata_options": {
              "computed": true,
              "description": "The metadata options for the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "http_endpoint": {
                    "computed": true,
                    "description": "Enables or disables the HTTP metadata endpoint on your instances. If the parameter is not specified, the default state is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "http_protocol_ipv_6": {
                    "computed": true,
                    "description": "Enables or disables the IPv6 endpoint for the instance metadata service.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "http_put_response_hop_limit": {
                    "computed": true,
                    "description": "The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "http_tokens": {
                    "computed": true,
                    "description": "IMDSv2 uses token-backed sessions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_metadata_tags": {
                    "computed": true,
                    "description": "Set to enabled to allow access to instance tags from the instance metadata.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "monitoring": {
              "computed": true,
              "description": "Specifies whether detailed monitoring is enabled for an instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "Specify true to enable detailed monitoring.",
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
              "description": "If you specify a network interface, you must specify any security groups and subnets as part of the network interface.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "associate_carrier_ip_address": {
                    "computed": true,
                    "description": "Indicates whether to associate a Carrier IP address with eth0 for a new network interface.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "associate_public_ip_address": {
                    "computed": true,
                    "description": "Associates a public IPv4 address with eth0 for a new network interface.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "connection_tracking_specification": {
                    "computed": true,
                    "description": "Allows customer to specify Connection Tracking options",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "tcp_established_timeout": {
                          "computed": true,
                          "description": "Integer value for TCP Established Timeout",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "udp_stream_timeout": {
                          "computed": true,
                          "description": "Integer value for UDP Stream Timeout",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "udp_timeout": {
                          "computed": true,
                          "description": "Integer value for UDP Timeout",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "delete_on_termination": {
                    "computed": true,
                    "description": "Indicates whether the network interface is deleted when the instance is terminated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "description": {
                    "computed": true,
                    "description": "A description for the network interface.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "device_index": {
                    "computed": true,
                    "description": "The device index for the network interface attachment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ena_srd_specification": {
                    "computed": true,
                    "description": "Allows customer to specify ENA-SRD options",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ena_srd_enabled": {
                          "computed": true,
                          "description": "Enables TCP ENA-SRD",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "ena_srd_udp_specification": {
                          "computed": true,
                          "description": "Allows customer to specify ENA-SRD (UDP) options",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "ena_srd_udp_enabled": {
                                "computed": true,
                                "description": "Enables UDP ENA-SRD",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
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
                  "groups": {
                    "computed": true,
                    "description": "The IDs of one or more security groups.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "interface_type": {
                    "computed": true,
                    "description": "The type of network interface.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ipv_4_prefix_count": {
                    "computed": true,
                    "description": "The number of IPv4 prefixes to be automatically assigned to the network interface.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ipv_4_prefixes": {
                    "computed": true,
                    "description": "One or more IPv4 prefixes to be assigned to the network interface.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ipv_4_prefix": {
                          "computed": true,
                          "description": "The IPv4 prefix.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "ipv_6_address_count": {
                    "computed": true,
                    "description": "The number of IPv6 addresses to assign to a network interface.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ipv_6_addresses": {
                    "computed": true,
                    "description": "One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet.",
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
                  "ipv_6_prefix_count": {
                    "computed": true,
                    "description": "The number of IPv6 prefixes to be automatically assigned to the network interface.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ipv_6_prefixes": {
                    "computed": true,
                    "description": "One or more IPv6 prefixes to be assigned to the network interface.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ipv_6_prefix": {
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
                  "network_card_index": {
                    "computed": true,
                    "description": "The index of the network card.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "network_interface_id": {
                    "computed": true,
                    "description": "The ID of the network interface.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "primary_ipv_6": {
                    "computed": true,
                    "description": "Enables the first IPv6 global unique address (GUA) on a dual stack or IPv6-only ENI immutable.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "private_ip_address": {
                    "computed": true,
                    "description": "The primary private IPv4 address of the network interface.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "private_ip_addresses": {
                    "computed": true,
                    "description": "One or more private IPv4 addresses.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "primary": {
                          "computed": true,
                          "description": "Indicates whether the private IPv4 address is the primary private IPv4 address. Only one IPv4 address can be designated as primary.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "private_ip_address": {
                          "computed": true,
                          "description": "The private IPv4 address.",
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
                    "description": "The number of secondary private IPv4 addresses to assign to a network interface.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "subnet_id": {
                    "computed": true,
                    "description": "The ID of the subnet for the network interface.",
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
              "description": "Specifies the placement of an instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "affinity": {
                    "computed": true,
                    "description": "The affinity setting for an instance on a Dedicated Host.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "availability_zone": {
                    "computed": true,
                    "description": "The Availability Zone for the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "group_id": {
                    "computed": true,
                    "description": "The Group Id of a placement group. You must specify the Placement Group Group Id to launch an instance in a shared placement group.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "group_name": {
                    "computed": true,
                    "description": "The name of the placement group for the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "host_id": {
                    "computed": true,
                    "description": "The ID of the Dedicated Host for the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "host_resource_group_arn": {
                    "computed": true,
                    "description": "The ARN of the host resource group in which to launch the instances. If you specify a host resource group ARN, omit the Tenancy parameter or set it to host.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partition_number": {
                    "computed": true,
                    "description": "The number of the partition the instance should launch in. Valid only if the placement group strategy is set to partition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "spread_domain": {
                    "computed": true,
                    "description": "Reserved for future use.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tenancy": {
                    "computed": true,
                    "description": "The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "private_dns_name_options": {
              "computed": true,
              "description": "Describes the options for instance hostnames.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_resource_name_dns_a_record": {
                    "computed": true,
                    "description": "Indicates whether to respond to DNS queries for instance hostnames with DNS A records.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_resource_name_dns_aaaa_record": {
                    "computed": true,
                    "description": "Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "hostname_type": {
                    "computed": true,
                    "description": "The type of hostname for EC2 instances.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "ram_disk_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_group_ids": {
              "computed": true,
              "description": "One or more security group IDs. ",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "security_groups": {
              "computed": true,
              "description": "One or more security group names.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tag_specifications": {
              "computed": true,
              "description": "The tags to apply to the resources that are created during instance launch.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "resource_type": {
                    "computed": true,
                    "description": "The type of resource to tag.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
                    "description": "The tags for the resource.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
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
              "description": "The user data to make available to the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "launch_template_id": {
        "computed": true,
        "description": "LaunchTemplate ID generated by service",
        "description_kind": "plain",
        "type": "string"
      },
      "launch_template_name": {
        "computed": true,
        "description": "A name for the launch template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tag_specifications": {
        "computed": true,
        "description": "The tags to apply to the launch template on creation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_type": {
              "computed": true,
              "description": "The type of resource to tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tags": {
              "computed": true,
              "description": "The tags for the resource.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
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
      "version_description": {
        "computed": true,
        "description": "A description for the first version of the launch template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::LaunchTemplate",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2LaunchTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2LaunchTemplate), &result)
	return &result
}
