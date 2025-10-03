package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneProjectMembership = `{
  "block": {
    "attributes": {
      "designation": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_identifier": {
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
      "member": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "group_identifier": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_identifier": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "member_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "member_identifier_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::DataZone::ProjectMembership Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneProjectMembershipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneProjectMembership), &result)
	return &result
}
