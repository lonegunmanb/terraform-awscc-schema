package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSqsQueueInlinePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description": "A policy document that contains permissions to add to the specified SQS queue",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "queue": {
        "description": "The URL of the SQS queue.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Schema for SQS QueueInlinePolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSqsQueueInlinePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSqsQueueInlinePolicy), &result)
	return &result
}
