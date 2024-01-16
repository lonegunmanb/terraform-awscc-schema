package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBatchJobQueue = `{
  "block": {
    "attributes": {
      "compute_environment_order": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compute_environment": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "order": {
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
      },
      "job_queue_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_queue_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "scheduling_policy_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
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
    "description": "Data Source schema for AWS::Batch::JobQueue",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBatchJobQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchJobQueue), &result)
	return &result
}
