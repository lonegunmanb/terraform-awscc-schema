package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineQueueEnvironment = `{
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "queue_environment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "queue_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Deadline::QueueEnvironment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDeadlineQueueEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineQueueEnvironment), &result)
	return &result
}
