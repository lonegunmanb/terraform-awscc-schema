package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBatchSchedulingPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "ARN of the Scheduling Policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "fairshare_policy": {
        "computed": true,
        "description": "Fair Share Policy for the Job Queue.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compute_reservation": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "share_decay_seconds": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "share_distribution": {
              "computed": true,
              "description": "List of Share Attributes",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "share_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "weight_factor": {
                    "computed": true,
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
      "name": {
        "computed": true,
        "description": "Name of Scheduling Policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Batch::SchedulingPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBatchSchedulingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchSchedulingPolicy), &result)
	return &result
}
