package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamUser = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that identifies the user. For more information about ARNs and how to use ARNs in policies, see IAM Identifiers in the IAM User Guide.",
        "description_kind": "plain",
        "type": "string"
      },
      "groups": {
        "computed": true,
        "description": "A list of group names to which you want to add the user.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "login_profile": {
        "computed": true,
        "description": "Creates a password for the specified IAM user. A password allows an IAM user to access AWS services through the AWS Management Console.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "password": {
              "description": "The user's password.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password_reset_required": {
              "computed": true,
              "description": "Specifies whether the user is required to set a new password on next sign-in.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "managed_policy_arns": {
        "computed": true,
        "description": "A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "path": {
        "computed": true,
        "description": "The path to the user. For more information about paths, see IAM identifiers in the IAM User Guide. The ARN of the policy used to set the permissions boundary for the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "permissions_boundary": {
        "computed": true,
        "description": "The ARN of the policy that is used to set the permissions boundary for the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description": "Adds or updates an inline policy document that is embedded in the specified IAM role.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy_document": {
              "description": "The policy document.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "policy_name": {
              "description": "The friendly name (not ARN) identifying the policy.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A list of tags that are associated with the user. For more information about tagging, see Tagging IAM resources in the IAM User Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "user_name": {
        "computed": true,
        "description": "The friendly name identifying the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::IAM::User",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIamUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamUser), &result)
	return &result
}
