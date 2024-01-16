package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rpo_in_secs": {
              "computed": true,
              "description": "RPO in seconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "rto_in_secs": {
              "computed": true,
              "description": "RTO in seconds.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "map"
        }
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
        "type": "string"
      },
      "policy_name": {
        "computed": true,
        "description": "Name of Resiliency Policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tier": {
        "computed": true,
        "description": "Resiliency Policy Tier.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ResilienceHub::ResiliencyPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccResiliencehubResiliencyPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResiliencehubResiliencyPolicy), &result)
	return &result
}
