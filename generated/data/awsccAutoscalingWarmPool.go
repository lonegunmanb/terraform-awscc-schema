package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAutoscalingWarmPool = `{
  "block": {
    "attributes": {
      "auto_scaling_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_reuse_policy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "reuse_on_scale_in": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "max_group_prepared_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "min_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "pool_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AutoScaling::WarmPool",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAutoscalingWarmPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAutoscalingWarmPool), &result)
	return &result
}
