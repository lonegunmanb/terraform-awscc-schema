package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneProjectMembership = `{
  "block": {
    "attributes": {
      "designation": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
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
      "member": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "group_identifier": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "user_identifier": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "project_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataZone::ProjectMembership",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneProjectMembershipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneProjectMembership), &result)
	return &result
}
