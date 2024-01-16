package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsClusterCapacityProviderAssociations = `{
  "block": {
    "attributes": {
      "capacity_providers": {
        "computed": true,
        "description": "List of capacity providers to associate with the cluster",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cluster": {
        "computed": true,
        "description": "The name of the cluster",
        "description_kind": "plain",
        "type": "string"
      },
      "default_capacity_provider_strategy": {
        "computed": true,
        "description": "List of capacity providers to associate with the cluster",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "capacity_provider": {
              "computed": true,
              "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
              "description_kind": "plain",
              "type": "string"
            },
            "weight": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECS::ClusterCapacityProviderAssociations",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcsClusterCapacityProviderAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsClusterCapacityProviderAssociations), &result)
	return &result
}
