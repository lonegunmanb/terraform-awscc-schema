package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsClusterCapacityProviderAssociations = `{
  "block": {
    "attributes": {
      "capacity_providers": {
        "description": "List of capacity providers to associate with the cluster",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "cluster": {
        "description": "The name of the cluster",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_capacity_provider_strategy": {
        "description": "List of capacity providers to associate with the cluster",
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
              "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
              "description_kind": "plain",
              "required": true,
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
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Associate a set of ECS Capacity Providers with a specified ECS Cluster",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcsClusterCapacityProviderAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsClusterCapacityProviderAssociations), &result)
	return &result
}
