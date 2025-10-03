package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcsComputeNodeGroup = `{
  "block": {
    "attributes": {
      "ami_id": {
        "computed": true,
        "description": "The ID of the Amazon Machine Image (AMI) that AWS PCS uses to launch instances. If not provided, AWS PCS uses the AMI ID specified in the custom launch template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The unique Amazon Resource Name (ARN) of the compute node group.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_id": {
        "description": "The ID of the cluster of the compute node group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compute_node_group_id": {
        "computed": true,
        "description": "The generated unique ID of the compute node group.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_launch_template": {
        "description": "An Amazon EC2 launch template AWS PCS uses to launch compute nodes.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "template_id": {
              "computed": true,
              "description": "The ID of the EC2 launch template to use to provision instances.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "description": "The version of the EC2 launch template to use to provision instances.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "error_info": {
        "computed": true,
        "description": "The list of errors that occurred during compute node group provisioning.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code": {
              "computed": true,
              "description": "The short-form error code.",
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description": "The detailed error information.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "iam_instance_profile_arn": {
        "description": "The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when launching EC2 instances. The role contained in your instance profile must have pcs:RegisterComputeNodeGroupInstance permissions attached to provision instances correctly.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_configs": {
        "description": "A list of EC2 instance configurations that AWS PCS can provision in the compute node group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "instance_type": {
              "computed": true,
              "description": "The EC2 instance type that AWS PCS can provision in the compute node group.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "name": {
        "computed": true,
        "description": "The name that identifies the compute node group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "purchase_option": {
        "computed": true,
        "description": "Specifies how EC2 instances are purchased on your behalf. AWS PCS supports On-Demand, Spot and Capacity Block instances. For more information, see Instance purchasing options in the Amazon Elastic Compute Cloud User Guide. If you don't provide this option, it defaults to On-Demand.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_configuration": {
        "description": "Specifies the boundaries of the compute node group auto scaling.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_instance_count": {
              "description": "The upper bound of the number of instances allowed in the compute fleet.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_instance_count": {
              "description": "The lower bound of the number of instances allowed in the compute fleet.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "slurm_configuration": {
        "computed": true,
        "description": "Additional options related to the Slurm scheduler.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "slurm_custom_settings": {
              "computed": true,
              "description": "Additional Slurm-specific configuration that directly maps to Slurm settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "parameter_name": {
                    "computed": true,
                    "description": "AWS PCS supports configuration of the following Slurm parameters for compute node groups: Weight and RealMemory.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parameter_value": {
                    "computed": true,
                    "description": "The value for the configured Slurm setting.",
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
      "spot_options": {
        "computed": true,
        "description": "Additional configuration when you specify SPOT as the purchase option.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allocation_strategy": {
              "computed": true,
              "description": "The Amazon EC2 allocation strategy AWS PCS uses to provision EC2 instances. AWS PCS supports lowest price, capacity optimized, and price capacity optimized. If you don't provide this option, it defaults to price capacity optimized.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description": "The provisioning status of the compute node group. The provisioning status doesn't indicate the overall health of the compute node group.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "description": "The list of subnet IDs where instances are provisioned by the compute node group. The subnets must be in the same VPC as the cluster.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "AWS::PCS::ComputeNodeGroup resource creates an AWS PCS compute node group.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPcsComputeNodeGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcsComputeNodeGroup), &result)
	return &result
}
