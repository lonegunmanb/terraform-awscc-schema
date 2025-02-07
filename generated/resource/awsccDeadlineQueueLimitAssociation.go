package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineQueueLimitAssociation = `{
  "block": {
    "attributes": {
      "farm_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "limit_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "queue_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::Deadline::QueueLimitAssociation Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDeadlineQueueLimitAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineQueueLimitAssociation), &result)
	return &result
}
