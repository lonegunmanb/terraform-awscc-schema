package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccStepfunctionsStateMachineVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_machine_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state_machine_revision_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource schema for StateMachineVersion",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccStepfunctionsStateMachineVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccStepfunctionsStateMachineVersion), &result)
	return &result
}
