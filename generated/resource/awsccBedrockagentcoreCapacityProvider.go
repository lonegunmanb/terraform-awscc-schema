package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreCapacityProvider = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the capacity provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_provider_id": {
        "computed": true,
        "description": "The unique identifier of the capacity provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "compute_configuration": {
        "description": "The capacity configuration for the capacity provider. Defines the compute resources for this capacity provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ec_2_configuration": {
              "description": "Configuration for EC2-based capacity.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "launch_template_source": {
                    "description": "How the launch template is specified.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "launch_parameters": {
                          "description": "Parameters for launching EC2 instances.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "capacity_reservation_specification": {
                                "computed": true,
                                "description": "The Capacity Reservation targeting option.",
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
                              "ephemeral_volumes": {
                                "computed": true,
                                "description": "The block device mapping for ephemeral (instance store) volumes.",
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
                                          "ebs_card_index": {
                                            "computed": true,
                                            "description": "The index of the EBS card. Applies to instances with multiple EBS cards.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "encrypted": {
                                            "computed": true,
                                            "description": "Indicates whether the EBS volume is encrypted.",
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
                                            "description": "Identifier of the customer managed KMS key to use for EBS encryption.",
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
                                            "description": "The throughput to provision for a gp3 volume, in MiB/s.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "volume_initialization_rate": {
                                            "computed": true,
                                            "description": "The rate at which the volume is initialized after creation, in MiB/s. Supported only for volumes created from snapshots. If the snapshot is enabled for fast snapshot restore and a volume initialization rate is also specified, the volume is initialized at the specified rate instead of by fast snapshot restore. Valid range: 100-300 MiB/s.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "volume_size": {
                                            "computed": true,
                                            "description": "The size of the volume, in GiBs.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "volume_type": {
                                            "computed": true,
                                            "description": "The volume type. Defaults to gp3 if not specified.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
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
                              "instance_profile_arn": {
                                "computed": true,
                                "description": "The ARN of the IAM instance profile to associate with launched instances.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "instance_requirements": {
                                "description": "Requirements for EC2 instance types.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "allowed_instance_types": {
                                      "description": "List of allowed instance types.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "required": true
                              },
                              "license_specifications": {
                                "computed": true,
                                "description": "The license configurations.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "license_configuration_arn": {
                                      "computed": true,
                                      "description": "The ARN of the license configuration.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "optional": true
                              },
                              "monitoring": {
                                "computed": true,
                                "description": "The monitoring level for the instance.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "operating_system": {
                                "description": "The operating system and CPU architecture for the instances.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "propagated_tags": {
                                "computed": true,
                                "description": "Tags to apply to all EC2 resources (instances, volumes, and network interfaces) created by this capacity provider.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "ssh_key_name": {
                                "computed": true,
                                "description": "The name of the SSH key pair to configure on instances for SSH connectivity.",
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
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "lifecycle_configuration": {
                    "computed": true,
                    "description": "Configuration for managing the lifecycle of instances in a capacity provider.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "idle_instance_timeout": {
                          "computed": true,
                          "description": "The number of seconds an instance can remain idle before it is stopped.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_lifetime": {
                          "computed": true,
                          "description": "Maximum lifetime for the instance in seconds. Once reached, instances will be automatically terminated regardless of activity. Default: 28800 seconds (8 hours). Maximum: 1209600 seconds (14 days).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "root_volume": {
                    "computed": true,
                    "description": "Customer-facing configuration for the (service-managed) root volume. The service provisions the root volume at its own AMI size estimate plus FreeSpaceGiB, and pins the visible free space to FreeSpaceGiB with a filler file, so the space you are guaranteed does not change as the underlying AMI grows. The device name and the delete-on-termination behavior are service-owned and are not configurable.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "encrypted": {
                          "computed": true,
                          "description": "Indicates whether the EBS volume is encrypted. Encrypted volumes can only be attached to instances that support Amazon EBS encryption. If you are creating a volume from a snapshot, you can't specify an encryption value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "free_space_gi_b": {
                          "computed": true,
                          "description": "The free space guaranteed on the root volume, in GiB. The service adds the operating system overhead on top of this value. Defaults to 8 GiB. The maximum is below the 65,536 GiB gp3 ceiling because the service adds the AMI size bucket on top of this value, and the resulting total must still be a provisionable gp3 volume.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "iops": {
                          "computed": true,
                          "description": "The number of IOPS to provision. Only valid for gp3, io1, and io2 volumes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "kms_key_id": {
                          "computed": true,
                          "description": "Identifier of the customer managed KMS key to use for EBS encryption.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "throughput": {
                          "computed": true,
                          "description": "The throughput to provision for a gp3 volume, in MiB/s. Valid range: 125-2000 MiB/s.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "volume_type": {
                          "computed": true,
                          "description": "The EBS volume type. Defaults to gp3 if not specified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "volumes": {
                    "computed": true,
                    "description": "Named persistent EBS volumes for this capacity provider.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ebs_configuration": {
                          "computed": true,
                          "description": "Configuration for an EBS-backed persistent volume.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "encrypted": {
                                "computed": true,
                                "description": "Whether to encrypt the volume. Defaults to true.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "iops": {
                                "computed": true,
                                "description": "The number of IOPS to provision. Only valid for gp3, io1, and io2 volumes.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "kms_key_id": {
                                "computed": true,
                                "description": "Identifier of the KMS key to use for encryption.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The logical name of the volume, used to reference it when mounting.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "size_gi_b": {
                                "computed": true,
                                "description": "The size of the volume in GiB.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "snapshot_id": {
                                "computed": true,
                                "description": "Optional EBS snapshot ID to initialize the volume from.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "throughput": {
                                "computed": true,
                                "description": "The throughput in MiB/s. Only valid for gp3 volumes.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "volume_type": {
                                "computed": true,
                                "description": "The EBS volume type. Defaults to gp3 if not specified.",
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
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "vpc_configuration": {
                    "description": "VPC configuration for launching EC2 instances.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "security_groups": {
                          "description": "The IDs of the security groups to associate with the instances.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "subnets": {
                          "description": "The IDs of the subnets in which to launch instances.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the capacity provider was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of the capacity provider.",
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
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp when the capacity provider was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the capacity provider.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permissions_configuration": {
        "description": "Configuration for permissions associated with a capacity provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_provider_operator_role_arn": {
              "description": "The ARN of the IAM role that operators use to manage the capacity provider.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "status": {
        "computed": true,
        "description": "The current status of the capacity provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to the capacity provider.",
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
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::BedrockAgentCore::CapacityProvider. A capacity provider defines the compute resources (EC2) used to run Amazon Bedrock AgentCore agent runtimes.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreCapacityProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreCapacityProvider), &result)
	return &result
}
