package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSqsQueueInlinePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "A policy document that contains permissions to add to the specified SQS queue",
        "description_kind": "plain",
        "type": "string"
      },
      "queue": {
        "computed": true,
        "description": "The URL of the SQS queue.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SQS::QueueInlinePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSqsQueueInlinePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSqsQueueInlinePolicy), &result)
	return &result
}
