package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBatchJobQueue = `{
  "block": {
    "attributes": {
      "compute_environment_order": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compute_environment": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "order": {
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
      "job_queue_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_queue_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_state_time_limit_actions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_time_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "reason": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "priority": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "scheduling_policy_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
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
    "description": "Resource Type definition for AWS::Batch::JobQueue",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBatchJobQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchJobQueue), &result)
	return &result
}
