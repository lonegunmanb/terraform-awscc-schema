package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneProjectProfile = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_unit_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_unit_identifier": {
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
      "identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project_profile_id": {
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
    "description": "Data Source schema for AWS::DataZone::ProjectProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneProjectProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneProjectProfile), &result)
	return &result
}
