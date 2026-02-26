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
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_provider_name": {
        "computed": true,
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
              "description": "The maximum number of vCPUs that the capacity provider can provision across all compute instances.",
              "description_kind": "plain",
              "type": "number"
            },
            "scaling_mode": {
              "computed": true,
              "description": "The scaling mode that determines how the capacity provider responds to changes in demand.",
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
                    "description": "The predefined metric type to track for scaling decisions.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_value": {
                    "computed": true,
                    "description": "The target value for the metric that the scaling policy attempts to maintain through scaling actions.",
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
        "description": "The instance requirements for compute resources managed by the capacity provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowed_instance_types": {
              "computed": true,
              "description": "A list of EC2 instance types that the capacity provider is allowed to use. If not specified, all compatible instance types are allowed.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "architectures": {
              "computed": true,
              "description": "A list of supported CPU architectures for compute instances. Valid values include ` + "`" + `` + "`" + `x86_64` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `arm64` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "excluded_instance_types": {
              "computed": true,
              "description": "A list of EC2 instance types that the capacity provider should not use, even if they meet other requirements.",
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
        "description": "The ARN of the KMS key used to encrypt the capacity provider's resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "permissions_config": {
        "computed": true,
        "description": "The permissions configuration for the capacity provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_provider_operator_role_arn": {
              "computed": true,
              "description": "The ARN of the IAM role that the capacity provider uses to manage compute instances and other AWS resources.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "state": {
        "computed": true,
        "description": "The current state of the capacity provider. Indicates whether the provider is being created, is active and ready for use, has failed, or is being deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair that provides metadata for the capacity provider.",
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
        "description": "The VPC configuration for the capacity provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "A list of security group IDs that control network access for compute instances managed by the capacity provider.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "A list of subnet IDs where the capacity provider launches compute instances.",
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
