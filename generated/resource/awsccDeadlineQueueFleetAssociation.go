package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineQueueFleetAssociation = `{
  "block": {
    "attributes": {
      "farm_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fleet_id": {
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
      "queue_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::Deadline::QueueFleetAssociation Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDeadlineQueueFleetAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineQueueFleetAssociation), &result)
	return &result
}
