package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccControltowerEnabledControl = `{
  "block": {
    "attributes": {
      "control_identifier": {
        "computed": true,
        "description": "Arn of the control.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_identifier": {
        "computed": true,
        "description": "Arn for Organizational unit to which the control needs to be applied",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ControlTower::EnabledControl",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccControltowerEnabledControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccControltowerEnabledControl), &result)
	return &result
}
