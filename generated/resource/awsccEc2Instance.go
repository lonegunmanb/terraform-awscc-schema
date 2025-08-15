package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Instance = `{
  "block": {
    "attributes": {
      "additional_info": {
        "computed": true,
        "description": "This property is reserved for internal use. If you use it, the stack fails with this error: Bad property set: [Testing this property] (Service: AmazonEC2; Status Code: 400; Error Code: InvalidParameterCombination; Request ID: 0XXXXXX-49c7-4b40-8bcc-76885dcXXXXX).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "affinity": {
        "computed": true,
        "description": "Indicates whether the instance is associated with a dedicated host. If you want the instance to always restart on the same host on which it was launched, specify host. If you want the instance to restart on any available host, but try to launch onto the last host it ran on (on a best-effort basis), specify default.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "block_device_mappings": {
        "computed": true,
        "description": "The block device mapping entries that defines the block devices to attach to the instance at launch.",
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
                    "description": "Indicates whether the volume should be encrypted.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "iops": {
                    "computed": true,
                    "description": "The number of I/O operations per second (IOPS). For gp3, io1, and io2 volumes, this represents the number of IOPS that are provisioned for the volume. For gp2 volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description": "The identifier of the AWS Key Management Service (AWS KMS) customer managed CMK to use for Amazon EBS encryption. If KmsKeyId is specified, the encrypted state must be true. If the encrypted state is true but you do not specify KmsKeyId, your AWS managed CMK for EBS is used.",
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
                  "volume_size": {
                    "computed": true,
                    "description": "The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size. If you specify a snapshot, the default is the snapshot size. You can specify a volume size that is equal to or larger than the snapshot size.",
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
      "cpu_options": {
        "computed": true,
        "description": "The CPU options for the instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "core_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "threads_per_core": {
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
      "credit_specification": {
        "computed": true,
        "description": "The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cpu_credits": {
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
      "disable_api_termination": {
        "computed": true,
        "description": "If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can.",
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
              "description": "The type of Elastic Graphics accelerator. Amazon Elastic Graphics is no longer available.",
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
        "description": "An elastic inference accelerator to associate with the instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "count": {
              "computed": true,
              "description": "The number of elastic inference accelerators to attach to the instance. Amazon Elastic Inference is no longer available.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "computed": true,
              "description": "The type of elastic inference accelerator. Amazon Elastic Inference is no longer available.",
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
        "description": "Indicates whether an instance is enabled for hibernation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configured": {
              "computed": true,
              "description": "If you set this parameter to true, your instance is enabled for hibernation.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "host_id": {
        "computed": true,
        "description": "If you specify host for the Affinity property, the ID of a dedicated host that the instance is associated with. If you don't specify an ID, Amazon EC2 launches the instance onto any available, compatible dedicated host in your account.",
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
      "iam_instance_profile": {
        "computed": true,
        "description": "The IAM instance profile.",
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
      "image_id": {
        "computed": true,
        "description": "The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The EC2 Instance ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_initiated_shutdown_behavior": {
        "computed": true,
        "description": "Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description": "The instance type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_6_address_count": {
        "computed": true,
        "description": "[EC2-VPC] The number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ipv_6_addresses": {
        "computed": true,
        "description": "[EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ipv_6_address": {
              "computed": true,
              "description": "The IPv6 address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
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
        "description": "The name of the key pair.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "launch_template": {
        "computed": true,
        "description": "The launch template to use to launch the instances.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "launch_template_id": {
              "computed": true,
              "description": "The ID of the launch template. You must specify the LaunchTemplateName or the LaunchTemplateId, but not both.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "launch_template_name": {
              "computed": true,
              "description": "The name of the launch template. You must specify the LaunchTemplateName or the LaunchTemplateId, but not both.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version number of the launch template.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
      "metadata_options": {
        "computed": true,
        "description": "The metadata options for the instance",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "http_endpoint": {
              "computed": true,
              "description": "Enables or disables the HTTP metadata endpoint on your instances. If you specify a value of disabled, you cannot access your instance metadata.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_protocol_ipv_6": {
              "computed": true,
              "description": "Enables or disables the IPv6 endpoint for the instance metadata service. To use this option, the instance must be a Nitro-based instance launched in a subnet that supports IPv6.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_put_response_hop_limit": {
              "computed": true,
              "description": "The number of network hops that the metadata token can travel. Maximum is 64.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "http_tokens": {
              "computed": true,
              "description": "Indicates whether IMDSv2 is required.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_metadata_tags": {
              "computed": true,
              "description": "Indicates whether tags from the instance are propagated to the EBS volumes.",
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
        "description": "Specifies whether detailed monitoring is enabled for the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "network_interfaces": {
        "computed": true,
        "description": "The network interfaces to associate with the instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "associate_carrier_ip_address": {
              "computed": true,
              "description": "Not currently supported by AWS CloudFormation.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "associate_public_ip_address": {
              "computed": true,
              "description": "Indicates whether to assign a public IPv4 address to an instance you launch in a VPC.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "delete_on_termination": {
              "computed": true,
              "description": "If set to true, the interface is deleted when the instance is terminated.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "description": {
              "computed": true,
              "description": "The description of the network interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_index": {
              "computed": true,
              "description": "The position of the network interface in the attachment order. A primary network interface has a device index of 0.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ena_srd_specification": {
              "computed": true,
              "description": "Specifies the ENA Express settings for the network interface that's attached to the instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ena_srd_enabled": {
                    "computed": true,
                    "description": "Specifies whether ENA Express is enabled for the network interface when you launch an instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "ena_srd_udp_specification": {
                    "computed": true,
                    "description": "Contains ENA Express settings for UDP network traffic for the network interface that's attached to the instance.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ena_srd_udp_enabled": {
                          "computed": true,
                          "description": "Indicates whether UDP traffic uses ENA Express for your instance.",
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
            "group_set": {
              "computed": true,
              "description": "The IDs of the security groups for the network interface.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "ipv_6_address_count": {
              "computed": true,
              "description": "A number of IPv6 addresses to assign to the network interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipv_6_addresses": {
              "computed": true,
              "description": "The IPv6 addresses associated with the network interface.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ipv_6_address": {
                    "computed": true,
                    "description": "The IPv6 address.",
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
              "description": "The ID of the network interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_ip_address": {
              "computed": true,
              "description": "The private IPv4 address of the network interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_ip_addresses": {
              "computed": true,
              "description": "One or more private IPv4 addresses to assign to the network interface.",
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
                    "description": "The private IPv4 addresses.",
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
              "description": "The number of secondary private IPv4 addresses.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "subnet_id": {
              "computed": true,
              "description": "The ID of the subnet.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "placement_group_name": {
        "computed": true,
        "description": "The name of an existing placement group that you want to launch the instance into (cluster | partition | spread).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_dns_name": {
        "computed": true,
        "description": "The private DNS name of the specified instance. For example: ip-10-24-34-0.ec2.internal.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns_name_options": {
        "computed": true,
        "description": "The options for the instance hostname.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_resource_name_dns_a_record": {
              "computed": true,
              "description": "Indicates whether to respond to DNS queries for instance hostnames with DNS A records. For more information, see Amazon EC2 instance hostname types in the Amazon Elastic Compute Cloud User Guide.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_resource_name_dns_aaaa_record": {
              "computed": true,
              "description": "Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. For more information, see Amazon EC2 instance hostname types in the Amazon Elastic Compute Cloud User Guide.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hostname_type": {
              "computed": true,
              "description": "The type of hostnames to assign to instances in the subnet at launch. For IPv4 only subnets, an instance DNS name must be based on the instance IPv4 address. For IPv6 only subnets, an instance DNS name must be based on the instance ID. For dual-stack subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. For more information, see Amazon EC2 instance hostname types in the Amazon Elastic Compute Cloud User Guide.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "private_ip": {
        "computed": true,
        "description": "The private IP address of the specified instance. For example: 10.24.34.0.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description": "[EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "propagate_tags_to_volume_on_creation": {
        "computed": true,
        "description": "Indicates whether to assign the tags from the instance to all of the volumes attached to the instance at launch. If you specify true and you assign tags to the instance, those tags are automatically assigned to all of the volumes that you attach to the instance at launch. If you specify false, those tags are not assigned to the attached volumes.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "public_dns_name": {
        "computed": true,
        "description": "The public DNS name of the specified instance. For example: ec2-107-20-50-45.compute-1.amazonaws.com.",
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip": {
        "computed": true,
        "description": "The public IP address of the specified instance. For example: 192.0.2.0.",
        "description_kind": "plain",
        "type": "string"
      },
      "ramdisk_id": {
        "computed": true,
        "description": "The ID of the RAM disk to select.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description": "The IDs of the security groups.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "security_groups": {
        "computed": true,
        "description": "the names of the security groups. For a nondefault VPC, you must use security group IDs instead.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "source_dest_check": {
        "computed": true,
        "description": "Specifies whether to enable an instance launched in a VPC to perform NAT.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ssm_associations": {
        "computed": true,
        "description": "The SSM document and parameter values in AWS Systems Manager to associate with this instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "association_parameters": {
              "computed": true,
              "description": "The input parameter values to use with the associated SSM document.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The name of an input parameter that is in the associated SSM document.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value of an input parameter.",
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
            },
            "document_name": {
              "computed": true,
              "description": "The name of an SSM document to associate with the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "state": {
        "computed": true,
        "description": "The current state of the instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code": {
              "computed": true,
              "description": "The state of the instance as a 16-bit unsigned integer.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The current state of the instance.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "subnet_id": {
        "computed": true,
        "description": "[EC2-VPC] The ID of the subnet to launch the instance into.\n\n",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to add to the instance.",
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
      },
      "tenancy": {
        "computed": true,
        "description": "The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_data": {
        "computed": true,
        "description": "The user data to make available to the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volumes": {
        "computed": true,
        "description": "The volumes to attach to the instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "device": {
              "computed": true,
              "description": "The device name (for example, /dev/sdh or xvdh).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_id": {
              "computed": true,
              "description": "The ID of the EBS volume. The volume and instance must be within the same Availability Zone.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC that the instance is running in.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::Instance",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2InstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Instance), &result)
	return &result
}
