package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsCapacityProvider = `{
  "block": {
    "attributes": {
      "auto_scaling_group_provider": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_scaling_group_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "managed_draining": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "managed_scaling": {
              "computed": true,
              "description": "The managed scaling settings for the Auto Scaling group capacity provider.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_warmup_period": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "maximum_scaling_step_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "minimum_scaling_step_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "managed_termination_protection": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
      "name": {
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
    "description": "Data Source schema for AWS::ECS::CapacityProvider",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcsCapacityProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsCapacityProvider), &result)
	return &result
}
