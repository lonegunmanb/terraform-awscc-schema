package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAutoscalingLaunchConfiguration = `{
  "block": {
    "attributes": {
      "associate_public_ip_address": {
        "computed": true,
        "description": "For Auto Scaling groups that are running in a virtual private cloud (VPC), specifies whether to assign a public IP address to the group's instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "block_device_mappings": {
        "computed": true,
        "description": "Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "device_name": {
              "description": "The device name exposed to the EC2 instance (for example, /dev/sdh or xvdh). ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ebs": {
              "computed": true,
              "description": "Parameters used to automatically set up EBS volumes when an instance is launched.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delete_on_termination": {
                    "computed": true,
                    "description": "Indicates whether the volume is deleted on instance termination. ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "encrypted": {
                    "computed": true,
                    "description": "Specifies whether the volume should be encrypted. ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "iops": {
                    "computed": true,
                    "description": "The number of input/output (I/O) operations per second (IOPS) to provision for the volume. ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "snapshot_id": {
                    "computed": true,
                    "description": "The snapshot ID of the volume to use.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "throughput": {
                    "computed": true,
                    "description": "The throughput (MiBps) to provision for a gp3 volume.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "volume_size": {
                    "computed": true,
                    "description": "The volume size, in GiBs.",
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
              "description": "Setting this value to true suppresses the specified device included in the block device mapping of the AMI.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "virtual_name": {
              "computed": true,
              "description": "The name of the virtual device.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "classic_link_vpc_id": {
        "computed": true,
        "description": "The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances to.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "classic_link_vpc_security_groups": {
        "computed": true,
        "description": "The IDs of one or more security groups for the VPC that you specified in the ClassicLinkVPCId property.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ebs_optimized": {
        "computed": true,
        "description": "Specifies whether the launch configuration is optimized for EBS I/O (true) or not (false).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "iam_instance_profile": {
        "computed": true,
        "description": "Provides the name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance. The instance profile contains the IAM role.",
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
        "description": "Provides the unique ID of the Amazon Machine Image (AMI) that was assigned during registration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The ID of the Amazon EC2 instance you want to use to create the launch configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_monitoring": {
        "computed": true,
        "description": "Controls whether instances in this group are launched with detailed (true) or basic (false) monitoring.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "instance_type": {
        "description": "Specifies the instance type of the EC2 instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kernel_id": {
        "computed": true,
        "description": "Provides the ID of the kernel associated with the EC2 AMI.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_name": {
        "computed": true,
        "description": "Provides the name of the EC2 key pair.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "launch_configuration_name": {
        "computed": true,
        "description": "The name of the launch configuration. This name must be unique per Region per account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metadata_options": {
        "computed": true,
        "description": "The metadata options for the instances.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "http_endpoint": {
              "computed": true,
              "description": "This parameter enables or disables the HTTP metadata endpoint on your instances.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_put_response_hop_limit": {
              "computed": true,
              "description": "The desired HTTP PUT response hop limit for instance metadata requests.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "http_tokens": {
              "computed": true,
              "description": "The state of token usage for your instance metadata requests.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "placement_tenancy": {
        "computed": true,
        "description": "The tenancy of the instance, either default or dedicated.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ram_disk_id": {
        "computed": true,
        "description": "The ID of the RAM disk to select.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_groups": {
        "computed": true,
        "description": "A list that contains the security groups to assign to the instances in the Auto Scaling group.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "spot_price": {
        "computed": true,
        "description": "The maximum hourly price you are willing to pay for any Spot Instances launched to fulfill the request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_data": {
        "computed": true,
        "description": "The Base64-encoded user data to make available to the launched EC2 instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The AWS::AutoScaling::LaunchConfiguration resource specifies the launch configuration that can be used by an Auto Scaling group to configure Amazon EC2 instances.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAutoscalingLaunchConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAutoscalingLaunchConfiguration), &result)
	return &result
}
