package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResiliencehubResiliencyPolicy = `{
  "block": {
    "attributes": {
      "data_location_constraint": {
        "computed": true,
        "description": "Data Location Constraint of the Policy.",
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
      "policy": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rpo_in_secs": {
              "computed": true,
              "description": "RPO in seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rto_in_secs": {
              "computed": true,
              "description": "RTO in seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "map"
        },
        "required": true
      },
      "policy_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the Resiliency Policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_description": {
        "computed": true,
        "description": "Description of Resiliency Policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_name": {
        "description": "Name of Resiliency Policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tier": {
        "description": "Resiliency Policy Tier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type Definition for Resiliency Policy.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccResiliencehubResiliencyPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResiliencehubResiliencyPolicy), &result)
	return &result
}
