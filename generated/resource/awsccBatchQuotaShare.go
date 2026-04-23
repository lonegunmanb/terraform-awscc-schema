package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBatchQuotaShare = `{
  "block": {
    "attributes": {
      "capacity_limits": {
        "description": "A list that specifies the quantity and type of compute capacity allocated to the quota share.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_unit": {
              "description": "The unit of compute capacity for the capacityLimit. For example, ` + "`" + `ml.m5.large` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "max_capacity": {
              "description": "The maximum capacity available for the quota share. This value represents the maximum quantity of a resource that can be allocated to jobs in the quota share without borrowing.",
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
        "description": "The AWS Batch job queue associated with the quota share. This can be the job queue name or ARN. A job queue must be in the ` + "`" + `VALID` + "`" + ` state before you can associate it with a quota share.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "preemption_configuration": {
        "description": "Specifies the preemption behavior for jobs in a quota share.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "in_share_preemption": {
              "description": "Specifies whether jobs within a quota share can be preempted by another, higher priority job in the same quota share.",
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
        "description": "The name of the quota share. It can be up to 128 characters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_sharing_configuration": {
        "description": "Specifies whether a quota share reserves, lends, or both lends and borrows idle compute capacity.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "borrow_limit": {
              "computed": true,
              "description": "The maximum percentage of additional capacity that the quota share can borrow from other shares. ` + "`" + `BorrowLimit` + "`" + ` can only be applied to quota shares with a strategy of ` + "`" + `LEND_AND_BORROW` + "`" + `. This value is expressed as a percentage of the quota share's configured CapacityLimits. The ` + "`" + `BorrowLimit` + "`" + ` is applied uniformly across all capacity units. For example, if the ` + "`" + `BorrowLimit` + "`" + ` is 200, the quota share can borrow up to 200% of its configured ` + "`" + `maxCapacity` + "`" + ` for each capacity unit. The default ` + "`" + `BorrowLimit` + "`" + ` is -1, which indicates unlimited borrowing.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "strategy": {
              "description": "The resource sharing strategy for the quota share. The ` + "`" + `RESERVE` + "`" + ` strategy allows a quota share to reserve idle capacity for itself. ` + "`" + `LEND` + "`" + ` configures the share to lend its idle capacity to another share in need of capacity. The ` + "`" + `LEND_AND_BORROW` + "`" + ` strategy configures the share to borrow idle capacity from an underutilized share, as well as lend to another share.",
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
        "description": "The state of the quota share. If the quota share is ` + "`" + `ENABLED` + "`" + `, it is able to accept jobs. If the quota share is ` + "`" + `DISABLED` + "`" + `, new jobs won't be accepted but jobs already submitted can finish. The default state is ` + "`" + `ENABLED` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags that you apply to the quota share to help you categorize and organize your resources. Each tag consists of a key and an optional value.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Creates an AWS Batch quota share. Each quota share operates as a virtual queue with a configured compute capacity, resource sharing strategy, and borrow limits.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBatchQuotaShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchQuotaShare), &result)
	return &result
}
