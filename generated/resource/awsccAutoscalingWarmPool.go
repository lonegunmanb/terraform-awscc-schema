package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAutoscalingWarmPool = `{
  "block": {
    "attributes": {
      "auto_scaling_group_name": {
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
      "instance_reuse_policy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "reuse_on_scale_in": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "max_group_prepared_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "pool_state": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::AutoScaling::WarmPool.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAutoscalingWarmPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAutoscalingWarmPool), &result)
	return &result
}
