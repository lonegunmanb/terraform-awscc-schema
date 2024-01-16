package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state_machine_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state_machine_revision_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::StepFunctions::StateMachineVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccStepfunctionsStateMachineVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccStepfunctionsStateMachineVersion), &result)
	return &result
}
