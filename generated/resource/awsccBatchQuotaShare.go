package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBatchQuotaShare = `{
  "block": {
    "attributes": {
      "capacity_limits": {
        "description": "The capacity limits for the quota share.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_unit": {
              "description": "The unit of compute capacity for the capacityLimit.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "max_capacity": {
              "description": "The maximum capacity available for the quota share. This value represents the maximum amount of resources that can be allocated to jobs in the quota share without borrowing",
              "description_kind": "plain",
              "required": true,
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
      },
      "job_queue": {
        "description": "The Amazon Resource Name (ARN) or name of the job queue.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "preemption_configuration": {
        "description": "The preemption configuration for the quota share.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "in_share_preemption": {
              "description": "Whether preemption is enabled within the quota share.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "quota_share_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the quota share.",
        "description_kind": "plain",
        "type": "string"
      },
      "quota_share_name": {
        "description": "The name of the quota share.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_sharing_configuration": {
        "description": "The resource sharing configuration for the quota share.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "borrow_limit": {
              "computed": true,
              "description": "The maximum amount of compute capacity that can be borrowed. Use -1 for unlimited borrowing.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "strategy": {
              "description": "The resource sharing strategy.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "state": {
        "computed": true,
        "description": "The state of the quota share.",
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
    "description": "Resource Type definition for AWS::Batch::QuotaShare",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBatchQuotaShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchQuotaShare), &result)
	return &result
}
