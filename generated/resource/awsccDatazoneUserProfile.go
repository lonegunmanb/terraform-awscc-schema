package resource

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
        "description": "The identifier of the Amazon DataZone domain in which the user profile would be created.",
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
      "status": {
        "computed": true,
        "description": "The status of the user profile.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of the user profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_identifier": {
        "description": "The ID of the user.",
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      }
    },
    "description": "A user profile represents Amazon DataZone users. Amazon DataZone supports both IAM roles and SSO identities to interact with the Amazon DataZone Management Console and the data portal for different purposes. Domain administrators use IAM roles to perform the initial administrative domain-related work in the Amazon DataZone Management Console, including creating new Amazon DataZone domains, configuring metadata form types, and implementing policies. Data workers use their SSO corporate identities via Identity Center to log into the Amazon DataZone Data Portal and access projects where they have memberships.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneUserProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneUserProfile), &result)
	return &result
}
