package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneUserProfile = `{
  "block": {
    "attributes": {
      "details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "iam": {
              "computed": true,
              "description": "The details of the IAM User Profile.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The ARN of the IAM User Profile.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "sso": {
              "computed": true,
              "description": "The details of the SSO User Profile.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "first_name": {
                    "computed": true,
                    "description": "The First Name of the IAM User Profile.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "last_name": {
                    "computed": true,
                    "description": "The Last Name of the IAM User Profile.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "The username of the SSO User Profile.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain_id": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the user profile is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the user profile would be created.",
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
        "description": "The status of the user profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of the user profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_identifier": {
        "computed": true,
        "description": "The ID of the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_profile_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone user profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_type": {
        "computed": true,
        "description": "The type of the user.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataZone::UserProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneUserProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneUserProfile), &result)
	return &result
}
