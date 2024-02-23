package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGuarddutyMaster = `{
  "block": {
    "attributes": {
      "detector_id": {
        "computed": true,
        "description": "Unique ID of the detector of the GuardDuty member account.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "invitation_id": {
        "computed": true,
        "description": "Value used to validate the master account to the member account.",
        "description_kind": "plain",
        "type": "string"
      },
      "master_id": {
        "computed": true,
        "description": "ID of the account used as the master account.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::GuardDuty::Master",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGuarddutyMasterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGuarddutyMaster), &result)
	return &result
}
