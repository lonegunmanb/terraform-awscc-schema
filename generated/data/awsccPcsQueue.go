package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcsQueue = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The unique Amazon Resource Name (ARN) of the queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_id": {
        "computed": true,
        "description": "The ID of the cluster of the queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "compute_node_group_configurations": {
        "computed": true,
        "description": "The list of compute node group configurations associated with the queue. Queues assign jobs to associated compute node groups.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compute_node_group_id": {
              "computed": true,
              "description": "The compute node group ID for the compute node group configuration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "error_info": {
        "computed": true,
        "description": "The list of errors that occurred during queue provisioning.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code": {
              "computed": true,
              "description": "The short-form error code.",
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description": "The detailed error information.",
              "description_kind": "plain",
              "type": "string"
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
      "name": {
        "computed": true,
        "description": "The name that identifies the queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "queue_id": {
        "computed": true,
        "description": "The generated unique ID of the queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The provisioning status of the queue. The provisioning status doesn't indicate the overall health of the queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::PCS::Queue",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPcsQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcsQueue), &result)
	return &result
}
