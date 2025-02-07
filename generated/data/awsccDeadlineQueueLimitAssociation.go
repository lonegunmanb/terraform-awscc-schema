package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineQueueLimitAssociation = `{
  "block": {
    "attributes": {
      "farm_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "limit_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "queue_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Deadline::QueueLimitAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDeadlineQueueLimitAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineQueueLimitAssociation), &result)
	return &result
}
