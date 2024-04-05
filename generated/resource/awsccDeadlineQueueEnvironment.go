package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineQueueEnvironment = `{
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "queue_environment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "queue_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::Deadline::QueueEnvironment Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDeadlineQueueEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineQueueEnvironment), &result)
	return &result
}
