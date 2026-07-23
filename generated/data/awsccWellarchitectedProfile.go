package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedProfile = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The date and time the profile was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description": "The owner of the profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_arn": {
        "computed": true,
        "description": "The profile ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_description": {
        "computed": true,
        "description": "The profile description.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_name": {
        "computed": true,
        "description": "The name of the profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_questions": {
        "computed": true,
        "description": "The profile questions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "question_id": {
              "computed": true,
              "description": "The ID of the question.",
              "description_kind": "plain",
              "type": "string"
            },
            "selected_choice_ids": {
              "computed": true,
              "description": "The selected choices.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "profile_version": {
        "computed": true,
        "description": "The profile version.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time the profile was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::WellArchitected::Profile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWellarchitectedProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedProfile), &result)
	return &result
}
