package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBatchQuotaShare = `{
  "block": {
    "attributes": {
      "capacity_limits": {
        "computed": true,
        "description": "The capacity limits for the quota share.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_unit": {
              "computed": true,
              "description": "The unit of compute capacity for the capacityLimit.",
              "description_kind": "plain",
              "type": "string"
            },
            "max_capacity": {
              "computed": true,
              "description": "The maximum capacity available for the quota share. This value represents the maximum amount of resources that can be allocated to jobs in the quota share without borrowing",
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
      },
      "job_queue": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) or name of the job queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "preemption_configuration": {
        "computed": true,
        "description": "The preemption configuration for the quota share.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "in_share_preemption": {
              "computed": true,
              "description": "Whether preemption is enabled within the quota share.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "quota_share_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the quota share.",
        "description_kind": "plain",
        "type": "string"
      },
      "quota_share_name": {
        "computed": true,
        "description": "The name of the quota share.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_sharing_configuration": {
        "computed": true,
        "description": "The resource sharing configuration for the quota share.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "borrow_limit": {
              "computed": true,
              "description": "The maximum amount of compute capacity that can be borrowed. Use -1 for unlimited borrowing.",
              "description_kind": "plain",
              "type": "number"
            },
            "strategy": {
              "computed": true,
              "description": "The resource sharing strategy.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "state": {
        "computed": true,
        "description": "The state of the quota share.",
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
    "description": "Data Source schema for AWS::Batch::QuotaShare",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBatchQuotaShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchQuotaShare), &result)
	return &result
}
