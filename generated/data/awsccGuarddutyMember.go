package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGuarddutyMember = `{
  "block": {
    "attributes": {
      "detector_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disable_email_notification": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "email": {
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
      "member_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::GuardDuty::Member",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGuarddutyMemberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGuarddutyMember), &result)
	return &result
}
