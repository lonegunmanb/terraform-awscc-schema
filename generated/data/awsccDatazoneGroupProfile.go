package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneGroupProfile = `{
  "block": {
    "attributes": {
      "domain_id": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the group profile is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the group profile would be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_identifier": {
        "computed": true,
        "description": "The ID of the group.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_name": {
        "computed": true,
        "description": "The group-name of the Group Profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_profile_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone group profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the group profile.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataZone::GroupProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneGroupProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneGroupProfile), &result)
	return &result
}
