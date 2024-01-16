package resource

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
              "optional": true,
              "type": "number"
            },
            "share_decay_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                    "optional": true,
                    "type": "string"
                  },
                  "weight_factor": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of Scheduling Policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource Type schema for AWS::Batch::SchedulingPolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBatchSchedulingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchSchedulingPolicy), &result)
	return &result
}
