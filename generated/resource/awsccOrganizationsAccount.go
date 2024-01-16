package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOrganizationsAccount = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "If the account was created successfully, the unique identifier (ID) of the new account.",
        "description_kind": "plain",
        "type": "string"
      },
      "account_name": {
        "description": "The friendly name of the member account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the account.",
        "description_kind": "plain",
        "type": "string"
      },
      "email": {
        "description": "The email address of the owner to assign to the new member account.",
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
      "joined_method": {
        "computed": true,
        "description": "The method by which the account joined the organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "joined_timestamp": {
        "computed": true,
        "description": "The date the account became a part of the organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent_ids": {
        "computed": true,
        "description": "List of parent nodes for the member account. Currently only one parent at a time is supported. Default is root.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "role_name": {
        "computed": true,
        "description": "The name of an IAM role that AWS Organizations automatically preconfigures in the new member account. Default name is OrganizationAccountAccessRole if not specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the account in the organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags that you want to attach to the newly created account. For each tag in the list, you must specify both a tag key and a value.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key identifier, or name, of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The string value that's associated with the key of the tag. You can set the value of a tag to an empty string, but you can't set the value of a tag to null.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "You can use AWS::Organizations::Account to manage accounts in organization.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOrganizationsAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOrganizationsAccount), &result)
	return &result
}
