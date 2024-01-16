package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccControltowerEnabledControl = `{
  "block": {
    "attributes": {
      "control_identifier": {
        "description": "Arn of the control.",
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
      "target_identifier": {
        "description": "Arn for Organizational unit to which the control needs to be applied",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Enables a control on a specified target.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccControltowerEnabledControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccControltowerEnabledControl), &result)
	return &result
}
