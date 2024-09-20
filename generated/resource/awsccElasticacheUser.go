package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticacheUser = `{
  "block": {
    "attributes": {
      "access_string": {
        "computed": true,
        "description": "Access permissions string used for this user account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the user account.",
        "description_kind": "plain",
        "type": "string"
      },
      "authentication_mode": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "passwords": {
              "computed": true,
              "description": "Passwords used for this user account. You can create up to two passwords for each user.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "type": {
              "computed": true,
              "description": "Authentication Type",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "engine": {
        "description": "Must be redis.",
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
      "no_password_required": {
        "computed": true,
        "description": "Indicates a password is not required for this user account.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "passwords": {
        "computed": true,
        "description": "Passwords used for this user account. You can create up to two passwords for each user.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "Indicates the user status. Can be \"active\", \"modifying\" or \"deleting\".",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this user.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "user_id": {
        "description": "The ID of the user.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_name": {
        "description": "The username of the user.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ElastiCache::User",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccElasticacheUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticacheUser), &result)
	return &result
}
