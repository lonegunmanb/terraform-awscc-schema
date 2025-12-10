package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaCapacityProvider = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the capacity provider. This is a read-only property that is automatically generated when the capacity provider is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_provider_name": {
        "computed": true,
        "description": "The name of the capacity provider. The name must be unique within your AWS account and region. If you don't specify a name, CloudFormation generates one.",
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_provider_scaling_config": {
        "computed": true,
        "description": "The scaling configuration for the capacity provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_v_cpu_count": {
              "computed": true,
              "description": "The maximum number of EC2 instances that the capacity provider can scale up to.",
              "description_kind": "plain",
              "type": "number"
            },
            "scaling_mode": {
              "computed": true,
              "description": "The scaling mode for the capacity provider.",
              "description_kind": "plain",
              "type": "string"
            },
            "scaling_policies": {
              "computed": true,
              "description": "A list of target tracking scaling policies for the capacity provider.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "predefined_metric_type": {
                    "computed": true,
                    "description": "The predefined metric for target tracking.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_value": {
                    "computed": true,
                    "description": "The target value for the metric as a percentage (for example, 70.0 for 70%).",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_requirements": {
        "computed": true,
        "description": "Specifications for the types of EC2 instances that the capacity provider can use.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowed_instance_types": {
              "computed": true,
              "description": "A list of instance types that the capacity provider can use. Supports wildcards (for example, m5.*).",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "architectures": {
              "computed": true,
              "description": "The instruction set architecture for EC2 instances. Specify either x86_64 or arm64.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "excluded_instance_types": {
              "computed": true,
              "description": "A list of instance types that the capacity provider should not use. Takes precedence over AllowedInstanceTypes.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the AWS Key Management Service (KMS) key used by the capacity provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "permissions_config": {
        "computed": true,
        "description": "IAM permissions configuration for the capacity provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_provider_operator_role_arn": {
              "computed": true,
              "description": "The ARN of the IAM role that Lambda assumes to manage the capacity provider.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "state": {
        "computed": true,
        "description": "The current state of the capacity provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the capacity provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "vpc_config": {
        "computed": true,
        "description": "VPC configuration for the capacity provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "A list of security group IDs to associate with EC2 instances.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "A list of subnet IDs where the capacity provider can launch EC2 instances.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Lambda::CapacityProvider",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaCapacityProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaCapacityProvider), &result)
	return &result
}
