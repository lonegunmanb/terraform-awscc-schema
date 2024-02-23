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
                    "description": "The device name (for example, /dev/sdh or xvdh).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ebs": {
                    "computed": true,
                    "description": "Parameters used to automatically set up EBS volumes when the instance is launched.",
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
                          "description": "The number of I/O operations per second (IOPS). For ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + ` volumes, this represents the number of IOPS that are provisioned for the volume. For ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + ` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.\n The following are the supported values for each volume type:\n  +   ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + `: 3,000 - 16,000 IOPS\n  +   ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `: 100 - 64,000 IOPS\n  +   ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + `: 100 - 256,000 IOPS\n  \n For ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + ` volumes, you can achieve up to 256,000 IOPS on [instances built on the Nitro System](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances). On other instances, you can achieve performance up to 32,000 IOPS.\n This parameter is supported for ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + ` volumes only.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "kms_key_id": {
                          "computed": true,
                          "description": "The ARN of the symmetric KMSlong (KMS) CMK used for encryption.",
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
                          "description": "The throughput to provision for a ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + ` volume, with a maximum of 1,000 MiB/s.\n Valid Range: Minimum value of 125. Maximum value of 1000.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "volume_size": {
                          "computed": true,
                          "description": "The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size. The following are the supported volumes sizes for each volume type:\n  +   ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + `: 1 - 16,384 GiB\n  +   ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `: 4 - 16,384 GiB\n  +   ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + `: 4 - 65,536 GiB\n  +   ` + "`" + `` + "`" + `st1` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `sc1` + "`" + `` + "`" + `: 125 - 16,384 GiB\n  +   ` + "`" + `` + "`" + `standard` + "`" + `` + "`" + `: 1 - 1024 GiB",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "volume_type": {
                          "computed": true,
                          "description": "The volume type. For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the *Amazon Elastic Compute Cloud User Guide*.",
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
                    "description": "The virtual device name (ephemeralN). Instance store volumes are numbered starting from 0. An instance type with 2 available instance store volumes can specify mappings for ephemeral0 and ephemeral1. The number of available instance store volumes depends on the instance type. After you connect to the instance, you must mount the volume.",
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
              "description": "The Capacity Reservation targeting option. If you do not specify this parameter, the instance's Capacity Reservation preference defaults to ` + "`" + `` + "`" + `open` + "`" + `` + "`" + `, which enables it to run in any open Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity_reservation_preference": {
                    "computed": true,
                    "description": "Indicates the instance's Capacity Reservation preferences. Possible preferences include:\n  +   ` + "`" + `` + "`" + `open` + "`" + `` + "`" + ` - The instance can run in any ` + "`" + `` + "`" + `open` + "`" + `` + "`" + ` Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).\n  +   ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` - The instance avoids running in a Capacity Reservation even if one is available. The instance runs in On-Demand capacity.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "capacity_reservation_target": {
                    "computed": true,
                    "description": "Information about the target Capacity Reservation or Capacity Reservation group.",
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
              "description": "The CPU options for the instance. For more information, see [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the *Amazon Elastic Compute Cloud User Guide*.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "amd_sev_snp": {
                    "computed": true,
                    "description": "Indicates whether to enable the instance for AMD SEV-SNP. AMD SEV-SNP is supported with M6a, R6a, and C6a instance types only. For more information, see [AMD SEV-SNP](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/sev-snp.html).",
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
                    "description": "The number of threads per CPU core. To disable multithreading for the instance, specify a value of ` + "`" + `` + "`" + `1` + "`" + `` + "`" + `. Otherwise, specify the default value of ` + "`" + `` + "`" + `2` + "`" + `` + "`" + `.",
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
              "description": "The credit option for CPU usage of the instance. Valid only for T instances.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cpu_credits": {
                    "computed": true,
                    "description": "The credit option for CPU usage of a T instance.\n Valid values: ` + "`" + `` + "`" + `standard` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `unlimited` + "`" + `` + "`" + `",
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
              "description": "Indicates whether to enable the instance for stop protection. For more information, see [Stop protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html#Using_StopProtection) in the *Amazon Elastic Compute Cloud User Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disable_api_termination": {
              "computed": true,
              "description": "If you set this parameter to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can. To change this attribute after launch, use [ModifyInstanceAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html). Alternatively, if you set ` + "`" + `` + "`" + `InstanceInitiatedShutdownBehavior` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `terminate` + "`" + `` + "`" + `, you can terminate the instance by running the shutdown command from the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ebs_optimized": {
              "computed": true,
              "description": "Indicates whether the instance is optimized for Amazon EBS I/O. This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "elastic_gpu_specifications": {
              "computed": true,
              "description": "Deprecated.\n  Amazon Elastic Graphics reached end of life on January 8, 2024. For workloads that require graphics acceleration, we recommend that you use Amazon EC2 G4ad, G4dn, or G5 instances.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description": "The type of Elastic Graphics accelerator. For more information about the values to specify for ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + `, see [Elastic Graphics Basics](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html#elastic-graphics-basics), specifically the Elastic Graphics accelerator column, in the *Amazon Elastic Compute Cloud User Guide for Windows Instances*.",
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
              "description": "An elastic inference accelerator to associate with the instance. Elastic inference accelerators are a resource you can attach to your Amazon EC2 instances to accelerate your Deep Learning (DL) inference workloads.\n You cannot specify accelerators from different generations in the same request.\n  Starting April 15, 2023, AWS will not onboard new customers to Amazon Elastic Inference (EI), and will help current customers migrate their workloads to options that offer better price and performance. After April 15, 2023, new customers will not be able to launch instances with Amazon EI accelerators in Amazon SageMaker, Amazon ECS, or Amazon EC2. However, customers who have used Amazon EI at least once during the past 30-day period are considered current customers and will be able to continue using the service.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "count": {
                    "computed": true,
                    "description": "The number of elastic inference accelerators to attach to the instance. \n Default: 1",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description": "The type of elastic inference accelerator. The possible values are eia1.medium, eia1.large, and eia1.xlarge.",
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
              "description": "Indicates whether the instance is enabled for AWS Nitro Enclaves. For more information, see [What is Nitro Enclaves?](https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html) in the *Nitro Enclaves User Guide*.\n You can't enable AWS Nitro Enclaves and hibernation on the same instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "If this parameter is set to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, the instance is enabled for AWS Nitro Enclaves; otherwise, it is not enabled for AWS Nitro Enclaves.",
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
              "description": "Indicates whether an instance is enabled for hibernation. This parameter is valid only if the instance meets the [hibernation prerequisites](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/hibernating-prerequisites.html). For more information, see [Hibernate your instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the *Amazon Elastic Compute Cloud User Guide*.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "configured": {
                    "computed": true,
                    "description": "If you set this parameter to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, the instance is enabled for hibernation.\n Default: ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `",
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
              "description": "The name or Amazon Resource Name (ARN) of an IAM instance profile.",
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
              "description": "The ID of the AMI. Alternatively, you can specify a Systems Manager parameter, which will resolve to an AMI ID on launch.\n Valid formats:\n  +   ` + "`" + `` + "`" + `ami-17characters00000` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `resolve:ssm:parameter-name` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `resolve:ssm:parameter-name:version-number` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `resolve:ssm:parameter-name:label` + "`" + `` + "`" + ` \n  \n For more information, see [Use a Systems Manager parameter to find an AMI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html#using-systems-manager-parameter-to-find-AMI) in the *Amazon Elastic Compute Cloud User Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_initiated_shutdown_behavior": {
              "computed": true,
              "description": "Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).\n Default: ` + "`" + `` + "`" + `stop` + "`" + `` + "`" + `",
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
                    "description": "The options for Spot Instances.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "block_duration_minutes": {
                          "computed": true,
                          "description": "Deprecated.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "instance_interruption_behavior": {
                          "computed": true,
                          "description": "The behavior when a Spot Instance is interrupted. The default is ` + "`" + `` + "`" + `terminate` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "max_price": {
                          "computed": true,
                          "description": "The maximum hourly price you're willing to pay for the Spot Instances. We do not recommend using this parameter because it can lead to increased interruptions. If you do not specify this parameter, you will pay the current Spot price.\n  If you specify a maximum price, your Spot Instances will be interrupted more frequently than if you do not specify this parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "spot_instance_type": {
                          "computed": true,
                          "description": "The Spot Instance request type.\n If you are using Spot Instances with an Auto Scaling group, use ` + "`" + `` + "`" + `one-time` + "`" + `` + "`" + ` requests, as the Amazon EC2 Auto Scaling service handles requesting new Spot Instances whenever the group is below its desired capacity.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "valid_until": {
                          "computed": true,
                          "description": "The end date of the request, in UTC format (*YYYY-MM-DD*T*HH:MM:SS*Z). Supported only for persistent requests.\n  +  For a persistent request, the request remains active until the ` + "`" + `` + "`" + `ValidUntil` + "`" + `` + "`" + ` date and time is reached. Otherwise, the request remains active until you cancel it.\n  +  For a one-time request, ` + "`" + `` + "`" + `ValidUntil` + "`" + `` + "`" + ` is not supported. The request remains active until all instances launch or you cancel the request.\n  \n Default: 7 days from the current date",
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
              "description": "The attributes for the instance types. When you specify instance attributes, Amazon EC2 will identify instance types with these attributes.\n You must specify ` + "`" + `` + "`" + `VCpuCount` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `MemoryMiB` + "`" + `` + "`" + `. All other attributes are optional. Any unspecified optional attribute is set to its default.\n When you specify multiple attributes, you get instance types that satisfy all of the specified attributes. If you specify multiple values for an attribute, you get instance types that satisfy any of the specified values.\n To limit the list of instance types from which Amazon EC2 can identify matching instance types, you can use one of the following parameters, but not both in the same request:\n  +   ` + "`" + `` + "`" + `AllowedInstanceTypes` + "`" + `` + "`" + ` - The instance types to include in the list. All other instance types are ignored, even if they match your specified attributes.\n  +   ` + "`" + `` + "`" + `ExcludedInstanceTypes` + "`" + `` + "`" + ` - The instance types to exclude from the list, even if they match your specified attributes.\n  \n  If you specify ` + "`" + `` + "`" + `InstanceReq",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "accelerator_count": {
                    "computed": true,
                    "description": "The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) on an instance.\n To exclude accelerator-enabled instance types, set ` + "`" + `` + "`" + `Max` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.\n Default: No minimum or maximum limits",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum number of accelerators. To specify no maximum limit, omit this parameter. To exclude accelerator-enabled instance types, set ` + "`" + `` + "`" + `Max` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum number of accelerators. To specify no minimum limit, omit this parameter.",
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
                    "description": "Indicates whether instance types must have accelerators by specific manufacturers.\n  +  For instance types with AWS devices, specify ` + "`" + `` + "`" + `amazon-web-services` + "`" + `` + "`" + `.\n  +  For instance types with AMD devices, specify ` + "`" + `` + "`" + `amd` + "`" + `` + "`" + `.\n  +  For instance types with Habana devices, specify ` + "`" + `` + "`" + `habana` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA devices, specify ` + "`" + `` + "`" + `nvidia` + "`" + `` + "`" + `.\n  +  For instance types with Xilinx devices, specify ` + "`" + `` + "`" + `xilinx` + "`" + `` + "`" + `.\n  \n Default: Any manufacturer",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "accelerator_names": {
                    "computed": true,
                    "description": "The accelerators that must be on the instance type.\n  +  For instance types with NVIDIA A10G GPUs, specify ` + "`" + `` + "`" + `a10g` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA A100 GPUs, specify ` + "`" + `` + "`" + `a100` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA H100 GPUs, specify ` + "`" + `` + "`" + `h100` + "`" + `` + "`" + `.\n  +  For instance types with AWS Inferentia chips, specify ` + "`" + `` + "`" + `inferentia` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA GRID K520 GPUs, specify ` + "`" + `` + "`" + `k520` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA K80 GPUs, specify ` + "`" + `` + "`" + `k80` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA M60 GPUs, specify ` + "`" + `` + "`" + `m60` + "`" + `` + "`" + `.\n  +  For instance types with AMD Radeon Pro V520 GPUs, specify ` + "`" + `` + "`" + `radeon-pro-v520` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA T4 GPUs, specify ` + "`" + `` + "`" + `t4` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA T4G GPUs, specify ` + "`" + `` + "`" + `t4g` + "`" + `` + "`" + `.\n  +  For instance types with Xilinx VU9P FPGAs, specify ` + "`" + `` + "`" + `vu9p` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA V100 GPUs, specify ` + "`" + `` + "`" + `v100` + "`" + `` + "`" + `.\n  \n Default: Any accelerator",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "accelerator_total_memory_mi_b": {
                    "computed": true,
                    "description": "The minimum and maximum amount of total accelerator memory, in MiB.\n Default: No minimum or maximum limits",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum amount of accelerator memory, in MiB. To specify no maximum limit, omit this parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum amount of accelerator memory, in MiB. To specify no minimum limit, omit this parameter.",
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
                    "description": "The accelerator types that must be on the instance type.\n  +  For instance types with GPU accelerators, specify ` + "`" + `` + "`" + `gpu` + "`" + `` + "`" + `.\n  +  For instance types with FPGA accelerators, specify ` + "`" + `` + "`" + `fpga` + "`" + `` + "`" + `.\n  +  For instance types with inference accelerators, specify ` + "`" + `` + "`" + `inference` + "`" + `` + "`" + `.\n  \n Default: Any accelerator type",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_instance_types": {
                    "computed": true,
                    "description": "The instance types to apply your specified attributes against. All other instance types are ignored, even if they match your specified attributes.\n You can use strings with one or more wild cards, represented by an asterisk (` + "`" + `` + "`" + `*` + "`" + `` + "`" + `), to allow an instance type, size, or generation. The following are examples: ` + "`" + `` + "`" + `m5.8xlarge` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `c5*.*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `m5a.*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `r*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `*3*` + "`" + `` + "`" + `.\n For example, if you specify ` + "`" + `` + "`" + `c5*` + "`" + `` + "`" + `,Amazon EC2 will allow the entire C5 instance family, which includes all C5a and C5n instance types. If you specify ` + "`" + `` + "`" + `m5a.*` + "`" + `` + "`" + `, Amazon EC2 will allow all the M5a instance types, but not the M5n instance types.\n  If you specify ` + "`" + `` + "`" + `AllowedInstanceTypes` + "`" + `` + "`" + `, you can't specify ` + "`" + `` + "`" + `ExcludedInstanceTypes` + "`" + `` + "`" + `.\n  Default: All instance types",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "bare_metal": {
                    "computed": true,
                    "description": "Indicates whether bare metal instance types must be included, excluded, or required.\n  +  To include bare metal instance types, specify ` + "`" + `` + "`" + `included` + "`" + `` + "`" + `.\n  +  To require only bare metal instance types, specify ` + "`" + `` + "`" + `required` + "`" + `` + "`" + `.\n  +  To exclude bare metal instance types, specify ` + "`" + `` + "`" + `excluded` + "`" + `` + "`" + `.\n  \n Default: ` + "`" + `` + "`" + `excluded` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "baseline_ebs_bandwidth_mbps": {
                    "computed": true,
                    "description": "The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps. For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html) in the *Amazon EC2 User Guide*.\n Default: No minimum or maximum limits",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum baseline bandwidth, in Mbps. To specify no maximum limit, omit this parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum baseline bandwidth, in Mbps. To specify no minimum limit, omit this parameter.",
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
                    "description": "Indicates whether burstable performance T instance types are included, excluded, or required. For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html).\n  +  To include burstable performance instance types, specify ` + "`" + `` + "`" + `included` + "`" + `` + "`" + `.\n  +  To require only burstable performance instance types, specify ` + "`" + `` + "`" + `required` + "`" + `` + "`" + `.\n  +  To exclude burstable performance instance types, specify ` + "`" + `` + "`" + `excluded` + "`" + `` + "`" + `.\n  \n Default: ` + "`" + `` + "`" + `excluded` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cpu_manufacturers": {
                    "computed": true,
                    "description": "The CPU manufacturers to include.\n  +  For instance types with Intel CPUs, specify ` + "`" + `` + "`" + `intel` + "`" + `` + "`" + `.\n  +  For instance types with AMD CPUs, specify ` + "`" + `` + "`" + `amd` + "`" + `` + "`" + `.\n  +  For instance types with AWS CPUs, specify ` + "`" + `` + "`" + `amazon-web-services` + "`" + `` + "`" + `.\n  \n  Don't confuse the CPU manufacturer with the CPU architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template.\n  Default: Any manufacturer",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "excluded_instance_types": {
                    "computed": true,
                    "description": "The instance types to exclude.\n You can use strings with one or more wild cards, represented by an asterisk (` + "`" + `` + "`" + `*` + "`" + `` + "`" + `), to exclude an instance type, size, or generation. The following are examples: ` + "`" + `` + "`" + `m5.8xlarge` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `c5*.*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `m5a.*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `r*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `*3*` + "`" + `` + "`" + `.\n For example, if you specify ` + "`" + `` + "`" + `c5*` + "`" + `` + "`" + `,Amazon EC2 will exclude the entire C5 instance family, which includes all C5a and C5n instance types. If you specify ` + "`" + `` + "`" + `m5a.*` + "`" + `` + "`" + `, Amazon EC2 will exclude all the M5a instance types, but not the M5n instance types.\n  If you specify ` + "`" + `` + "`" + `ExcludedInstanceTypes` + "`" + `` + "`" + `, you can't specify ` + "`" + `` + "`" + `AllowedInstanceTypes` + "`" + `` + "`" + `.\n  Default: No excluded instance types",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "instance_generations": {
                    "computed": true,
                    "description": "Indicates whether current or previous generation instance types are included. The current generation instance types are recommended for use. Current generation instance types are typically the latest two to three generations in each instance family. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide*.\n For current generation instance types, specify ` + "`" + `` + "`" + `current` + "`" + `` + "`" + `.\n For previous generation instance types, specify ` + "`" + `` + "`" + `previous` + "`" + `` + "`" + `.\n Default: Current and previous generation instance types",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "local_storage": {
                    "computed": true,
                    "description": "Indicates whether instance types with instance store volumes are included, excluded, or required. For more information, [Amazon EC2 instance store](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html) in the *Amazon EC2 User Guide*.\n  +  To include instance types with instance store volumes, specify ` + "`" + `` + "`" + `included` + "`" + `` + "`" + `.\n  +  To require only instance types with instance store volumes, specify ` + "`" + `` + "`" + `required` + "`" + `` + "`" + `.\n  +  To exclude instance types with instance store volumes, specify ` + "`" + `` + "`" + `excluded` + "`" + `` + "`" + `.\n  \n Default: ` + "`" + `` + "`" + `included` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "local_storage_types": {
                    "computed": true,
                    "description": "The type of local storage that is required.\n  +  For instance types with hard disk drive (HDD) storage, specify ` + "`" + `` + "`" + `hdd` + "`" + `` + "`" + `.\n  +  For instance types with solid state drive (SSD) storage, specify ` + "`" + `` + "`" + `ssd` + "`" + `` + "`" + `.\n  \n Default: ` + "`" + `` + "`" + `hdd` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `ssd` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "max_spot_price_as_percentage_of_optimal_on_demand_price": {
                    "computed": true,
                    "description": "[Price protection] The price protection threshold for Spot Instances, as a percentage of an identified On-Demand price. The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from the lowest priced current generation instance types, and failing that, from the lowest priced previous generation instance types that match your attributes. When Amazon EC2 selects instance types with your attributes, it will exclude instance types whose price exceeds your specified threshold.\n The parameter accepts an integer, which Amazon EC2 interprets as a percentage.\n To indicate no price protection threshold, specify a high value, such as ` + "`" + `` + "`" + `999999` + "`" + `` + "`" + `.\n If you set ` + "`" + `` + "`" + `DesiredCapacityType` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `vcpu` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `memory-mib` + "`" + `` + "`" + `, the price protection threshold is based on the per vCPU or per memory price instead of the per instanc",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "memory_gi_b_per_v_cpu": {
                    "computed": true,
                    "description": "The minimum and maximum amount of memory per vCPU, in GiB.\n Default: No minimum or maximum limits",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum amount of memory per vCPU, in GiB. To specify no maximum limit, omit this parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum amount of memory per vCPU, in GiB. To specify no minimum limit, omit this parameter.",
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
                          "description": "The maximum amount of memory, in MiB. To specify no maximum limit, omit this parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum amount of memory, in MiB. To specify no minimum limit, specify ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.",
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
                    "description": "The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).\n Default: No minimum or maximum limits",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum amount of network bandwidth, in Gbps. To specify no maximum limit, omit this parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum amount of network bandwidth, in Gbps. If this parameter is not specified, there is no minimum limit.",
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
                    "description": "The minimum and maximum number of network interfaces.\n Default: No minimum or maximum limits",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum number of network interfaces. To specify no maximum limit, omit this parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum number of network interfaces. To specify no minimum limit, omit this parameter.",
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
                    "description": "[Price protection] The price protection threshold for On-Demand Instances, as a percentage higher than an identified On-Demand price. The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. When Amazon EC2 selects instance types with your attributes, it will exclude instance types whose price exceeds your specified threshold.\n The parameter accepts an integer, which Amazon EC2 interprets as a percentage.\n To turn off price protection, specify a high value, such as ` + "`" + `` + "`" + `999999` + "`" + `` + "`" + `.\n This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html) and [GetInstanceTypesFromInstanceRequirements](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceTypesFromInstanceRequirements.html).\n  If you set ` + "`" + `` + "`" + `TargetCapacityUnitType` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `vcpu` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `memory-mib` + "`" + `` + "`" + `, the price protection threshold is applied based on the per-",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "require_hibernate_support": {
                    "computed": true,
                    "description": "Indicates whether instance types must support hibernation for On-Demand Instances.\n This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html).\n Default: ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "spot_max_price_percentage_over_lowest_price": {
                    "computed": true,
                    "description": "[Price protection] The price protection threshold for Spot Instances, as a percentage higher than an identified Spot price. The identified Spot price is the Spot price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified Spot price is from the lowest priced current generation instance types, and failing that, from the lowest priced previous generation instance types that match your attributes. When Amazon EC2 selects instance types with your attributes, it will exclude instance types whose Spot price exceeds your specified threshold.\n The parameter accepts an integer, which Amazon EC2 interprets as a percentage.\n To indicate no price protection threshold, specify a high value, such as ` + "`" + `` + "`" + `999999` + "`" + `` + "`" + `.\n If you set ` + "`" + `` + "`" + `TargetCapacityUnitType` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `vcpu` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `memory-mib` + "`" + `` + "`" + `, the price protection threshold is applied based on the per-vCPU or per-memory price i",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "total_local_storage_gb": {
                    "computed": true,
                    "description": "The minimum and maximum amount of total local storage, in GB.\n Default: No minimum or maximum limits",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max": {
                          "computed": true,
                          "description": "The maximum amount of total local storage, in GB. To specify no maximum limit, omit this parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum amount of total local storage, in GB. To specify no minimum limit, omit this parameter.",
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
                          "description": "The maximum number of vCPUs. To specify no maximum limit, omit this parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "computed": true,
                          "description": "The minimum number of vCPUs. To specify no minimum limit, specify ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.",
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
              "description": "The instance type. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon Elastic Compute Cloud User Guide*.\n If you specify ` + "`" + `` + "`" + `InstanceType` + "`" + `` + "`" + `, you can't specify ` + "`" + `` + "`" + `InstanceRequirements` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kernel_id": {
              "computed": true,
              "description": "The ID of the kernel.\n We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User Provided Kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon EC2 User Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_name": {
              "computed": true,
              "description": "The name of the key pair. You can create a key pair using [CreateKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html) or [ImportKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html).\n  If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.",
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
              "description": "The metadata options for the instance. For more information, see [Instance metadata and user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) in the *Amazon Elastic Compute Cloud User Guide*.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "http_endpoint": {
                    "computed": true,
                    "description": "Enables or disables the HTTP metadata endpoint on your instances. If the parameter is not specified, the default state is ` + "`" + `` + "`" + `enabled` + "`" + `` + "`" + `.\n  If you specify a value of ` + "`" + `` + "`" + `disabled` + "`" + `` + "`" + `, you will not be able to access your instance metadata.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "http_protocol_ipv_6": {
                    "computed": true,
                    "description": "Enables or disables the IPv6 endpoint for the instance metadata service.\n Default: ` + "`" + `` + "`" + `disabled` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "http_put_response_hop_limit": {
                    "computed": true,
                    "description": "The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel.\n Default: ` + "`" + `` + "`" + `1` + "`" + `` + "`" + ` \n Possible values: Integers from 1 to 64",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "http_tokens": {
                    "computed": true,
                    "description": "Indicates whether IMDSv2 is required.\n  +   ` + "`" + `` + "`" + `optional` + "`" + `` + "`" + ` - IMDSv2 is optional. You can choose whether to send a session token in your instance metadata retrieval requests. If you retrieve IAM role credentials without a session token, you receive the IMDSv1 role credentials. If you retrieve IAM role credentials using a valid session token, you receive the IMDSv2 role credentials.\n  +   ` + "`" + `` + "`" + `required` + "`" + `` + "`" + ` - IMDSv2 is required. You must send a session token in your instance metadata retrieval requests. With this option, retrieving the IAM role credentials always returns IMDSv2 credentials; IMDSv1 credentials are not available.\n  \n Default: If the value of ` + "`" + `` + "`" + `ImdsSupport` + "`" + `` + "`" + ` for the Amazon Machine Image (AMI) for your instance is ` + "`" + `` + "`" + `v2.0` + "`" + `` + "`" + `, the default is ` + "`" + `` + "`" + `required` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_metadata_tags": {
                    "computed": true,
                    "description": "Set to ` + "`" + `` + "`" + `enabled` + "`" + `` + "`" + ` to allow access to instance tags from the instance metadata. Set to ` + "`" + `` + "`" + `disabled` + "`" + `` + "`" + ` to turn off access to instance tags from the instance metadata. For more information, see [Work with instance tags using the instance metadata](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#work-with-tags-in-IMDS).\n Default: ` + "`" + `` + "`" + `disabled` + "`" + `` + "`" + `",
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
              "description": "The monitoring for the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "Specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` to enable detailed monitoring. Otherwise, basic monitoring is enabled.",
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
              "description": "One or more network interfaces. If you specify a network interface, you must specify any security groups and subnets as part of the network interface.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "associate_carrier_ip_address": {
                    "computed": true,
                    "description": "Associates a Carrier IP address with eth0 for a new network interface.\n Use this option when you launch an instance in a Wavelength Zone and want to associate a Carrier IP address with the network interface. For more information about Carrier IP addresses, see [Carrier IP addresses](https://docs.aws.amazon.com/wavelength/latest/developerguide/how-wavelengths-work.html#provider-owned-ip) in the *Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "associate_public_ip_address": {
                    "computed": true,
                    "description": "Associates a public IPv4 address with eth0 for a new network interface.\n  AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [Amazon VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "connection_tracking_specification": {
                    "computed": true,
                    "description": "A connection tracking specification for the network interface.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "tcp_established_timeout": {
                          "computed": true,
                          "description": "Timeout (in seconds) for idle TCP connections in an established state. Min: 60 seconds. Max: 432000 seconds (5 days). Default: 432000 seconds. Recommended: Less than 432000 seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "udp_stream_timeout": {
                          "computed": true,
                          "description": "Timeout (in seconds) for idle UDP flows classified as streams which have seen more than one request-response transaction. Min: 60 seconds. Max: 180 seconds (3 minutes). Default: 180 seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "udp_timeout": {
                          "computed": true,
                          "description": "Timeout (in seconds) for idle UDP flows that have seen traffic only in a single direction or a single request-response transaction. Min: 30 seconds. Max: 60 seconds. Default: 30 seconds.",
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
                    "description": "The ENA Express configuration for the network interface.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ena_srd_enabled": {
                          "computed": true,
                          "description": "Indicates whether ENA Express is enabled for the network interface.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "ena_srd_udp_specification": {
                          "computed": true,
                          "description": "Configures ENA Express for UDP network traffic.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "ena_srd_udp_enabled": {
                                "computed": true,
                                "description": "Indicates whether UDP traffic to and from the instance uses ENA Express. To specify this setting, you must first enable ENA Express.",
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
                    "description": "The type of network interface. To create an Elastic Fabric Adapter (EFA), specify ` + "`" + `` + "`" + `efa` + "`" + `` + "`" + `. For more information, see [Elastic Fabric Adapter](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/efa.html) in the *Amazon Elastic Compute Cloud User Guide*.\n If you are not creating an EFA, specify ` + "`" + `` + "`" + `interface` + "`" + `` + "`" + ` or omit this parameter.\n Valid values: ` + "`" + `` + "`" + `interface` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `efa` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ipv_4_prefix_count": {
                    "computed": true,
                    "description": "The number of IPv4 prefixes to be automatically assigned to the network interface. You cannot use this option if you use the ` + "`" + `` + "`" + `Ipv4Prefix` + "`" + `` + "`" + ` option.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ipv_4_prefixes": {
                    "computed": true,
                    "description": "One or more IPv4 prefixes to be assigned to the network interface. You cannot use this option if you use the ` + "`" + `` + "`" + `Ipv4PrefixCount` + "`" + `` + "`" + ` option.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ipv_4_prefix": {
                          "computed": true,
                          "description": "The IPv4 prefix. For information, see [Assigning prefixes to Amazon EC2 network interfaces](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html) in the *Amazon Elastic Compute Cloud User Guide*.",
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
                    "description": "The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. You can't use this option if specifying specific IPv6 addresses.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ipv_6_addresses": {
                    "computed": true,
                    "description": "One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying a number of IPv6 addresses.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ipv_6_address": {
                          "computed": true,
                          "description": "One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying a number of IPv6 addresses.",
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
                    "description": "The number of IPv6 prefixes to be automatically assigned to the network interface. You cannot use this option if you use the ` + "`" + `` + "`" + `Ipv6Prefix` + "`" + `` + "`" + ` option.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ipv_6_prefixes": {
                    "computed": true,
                    "description": "One or more IPv6 prefixes to be assigned to the network interface. You cannot use this option if you use the ` + "`" + `` + "`" + `Ipv6PrefixCount` + "`" + `` + "`" + ` option.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ipv_6_prefix": {
                          "computed": true,
                          "description": "The IPv6 prefix.",
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
                    "description": "The index of the network card. Some instance types support multiple network cards. The primary network interface must be assigned to network card index 0. The default is network card index 0.",
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
                    "description": "The primary IPv6 address of the network interface. When you enable an IPv6 GUA address to be a primary IPv6, the first IPv6 GUA will be made the primary IPv6 address until the instance is terminated or the network interface is detached. For more information about primary IPv6 addresses, see [RunInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html).",
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
              "description": "The placement for the instance.",
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
                    "description": "The Group Id of a placement group. You must specify the Placement Group *Group Id* to launch an instance in a shared placement group.",
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
                    "description": "The ARN of the host resource group in which to launch the instances. If you specify a host resource group ARN, omit the *Tenancy* parameter or set it to ` + "`" + `` + "`" + `host` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partition_number": {
                    "computed": true,
                    "description": "The number of the partition the instance should launch in. Valid only if the placement group strategy is set to ` + "`" + `` + "`" + `partition` + "`" + `` + "`" + `.",
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
                    "description": "The tenancy of the instance. An instance with a tenancy of dedicated runs on single-tenant hardware.",
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
              "description": "The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries should be handled. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.",
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
                    "description": "The type of hostname for EC2 instances. For IPv4 only subnets, an instance DNS name must be based on the instance IPv4 address. For IPv6 only subnets, an instance DNS name must be based on the instance ID. For dual-stack subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.",
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
              "description": "The ID of the RAM disk.\n  We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User provided kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon Elastic Compute Cloud User Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_group_ids": {
              "computed": true,
              "description": "The IDs of the security groups. You can specify the IDs of existing security groups and references to resources created by the stack template.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "security_groups": {
              "computed": true,
              "description": "One or more security group names. For a nondefault VPC, you must use security group IDs instead.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tag_specifications": {
              "computed": true,
              "description": "The tags to apply to the resources that are created during instance launch.\n To tag a resource after it has been created, see [CreateTags](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html).\n To tag the launch template itself, use [TagSpecifications](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-tagspecifications).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "resource_type": {
                    "computed": true,
                    "description": "The type of resource to tag.\n Valid Values lists all resource types for Amazon EC2 that can be tagged. When you create a launch template, you can specify tags for the following resource types only: ` + "`" + `` + "`" + `instance` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `volume` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `network-interface` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `spot-instances-request` + "`" + `` + "`" + `. If the instance does not include the resource type that you specify, the instance launch fails. For example, not all instance types include a volume.\n To tag a resource after it has been created, see [CreateTags](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
                    "description": "The tags to apply to the resource.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "description": "The tag key.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "The tag value.",
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
              "description": "The user data to make available to the instance. You must provide base64-encoded text. User data is limited to 16 KB. For more information, see [Run commands on your Linux instance at launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) (Linux) or [Work with instance user data](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/instancedata-add-user-data.html) (Windows) in the *Amazon Elastic Compute Cloud User Guide*.\n If you are creating the launch template for use with BATCH, the user data must be provided in the [MIME multi-part archive format](https://docs.aws.amazon.com/https://cloudinit.readthedocs.io/en/latest/topics/format.html#mime-multi-part-archive). For more information, see [Amazon EC2 user data in launch templates](https://docs.aws.amazon.com/batch/latest/userguide/launch-templates.html) in the *User Guide*.",
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
        "description": "The tags to apply to the launch template on creation. To tag the launch template, the resource type must be ` + "`" + `` + "`" + `launch-template` + "`" + `` + "`" + `.\n To specify the tags for the resources that are created when an instance is launched, you must use [TagSpecifications](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-tagspecifications).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_type": {
              "computed": true,
              "description": "The type of resource. To tag the launch template, ` + "`" + `` + "`" + `ResourceType` + "`" + `` + "`" + ` must be ` + "`" + `` + "`" + `launch-template` + "`" + `` + "`" + `.",
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
                    "description": "The tag key.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The tag value.",
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
    "description": "Specifies the properties for creating a launch template.\n The minimum required properties for specifying a launch template are as follows:\n  +  You must specify at least one property for the launch template data.\n  +  You do not need to specify a name for the launch template. If you do not specify a name, CFN creates the name for you.\n  \n A launch template can contain some or all of the configuration information to launch an instance. When you launch an instance using a launch template, instance properties that are not specified in the launch template use default values, except the ` + "`" + `` + "`" + `ImageId` + "`" + `` + "`" + ` property, which has no default value. If you do not specify an AMI ID for the launch template ` + "`" + `` + "`" + `ImageId` + "`" + `` + "`" + ` property, you must specify an AMI ID for the instance ` + "`" + `` + "`" + `ImageId` + "`" + `` + "`" + ` property.\n For more information, see [Launch an instance from a launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html) in the *Amazon EC2 User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2LaunchTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2LaunchTemplate), &result)
	return &result
}
