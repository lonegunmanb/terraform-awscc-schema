package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamRole = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the role.",
        "description_kind": "plain",
        "type": "string"
      },
      "assume_role_policy_document": {
        "computed": true,
        "description": "The trust policy that is associated with this role.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the role that you provide.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_policy_arns": {
        "computed": true,
        "description": "A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role. ",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "max_session_duration": {
        "computed": true,
        "description": "The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours. ",
        "description_kind": "plain",
        "type": "number"
      },
      "path": {
        "computed": true,
        "description": "The path to the role.",
        "description_kind": "plain",
        "type": "string"
      },
      "permissions_boundary": {
        "computed": true,
        "description": "The ARN of the policy used to set the permissions boundary for the role.",
        "description_kind": "plain",
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description": "Adds or updates an inline policy document that is embedded in the specified IAM role. ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy_document": {
              "computed": true,
              "description": "The policy document.",
              "description_kind": "plain",
              "type": "string"
            },
            "policy_name": {
              "computed": true,
              "description": "The friendly name (not ARN) identifying the policy.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "role_id": {
        "computed": true,
        "description": "The stable and unique string identifying the role.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_name": {
        "computed": true,
        "description": "A name for the IAM role, up to 64 characters in length.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags that are attached to the role.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::IAM::Role",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIamRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamRole), &result)
	return &result
}
