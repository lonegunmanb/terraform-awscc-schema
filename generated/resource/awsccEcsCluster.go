package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon ECS cluster, such as arn:aws:ecs:us-east-2:123456789012:cluster/MyECSCluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_providers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "cluster_name": {
        "computed": true,
        "description": "A user-generated string that you use to identify your cluster. If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
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
      "configuration": {
        "computed": true,
        "description": "The configurations to be set at cluster level.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "execute_command_configuration": {
              "computed": true,
              "description": "The configuration for ExecuteCommand.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_configuration": {
                    "computed": true,
                    "description": "The session logging configuration for ExecuteCommand.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cloudwatch_encryption_enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "cloudwatch_log_group_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_bucket_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_encryption_enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "s3_key_prefix": {
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
                  "logging": {
                    "computed": true,
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
      "default_capacity_provider_strategy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "capacity_provider": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weight": {
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_connect_defaults": {
        "computed": true,
        "description": "Service Connect Configuration default for all services or tasks within this cluster",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "namespace": {
              "computed": true,
              "description": "Service Connect Namespace Name or ARN default for all services or tasks within this cluster",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
    "description": "Create an Elastic Container Service (ECS) cluster.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcsClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsCluster), &result)
	return &result
}
