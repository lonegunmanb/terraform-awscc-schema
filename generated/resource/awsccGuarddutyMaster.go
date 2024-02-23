package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGuarddutyMaster = `{
  "block": {
    "attributes": {
      "detector_id": {
        "description": "Unique ID of the detector of the GuardDuty member account.",
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
      "invitation_id": {
        "computed": true,
        "description": "Value used to validate the master account to the member account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_id": {
        "description": "ID of the account used as the master account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "GuardDuty Master resource schema",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGuarddutyMasterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGuarddutyMaster), &result)
	return &result
}
