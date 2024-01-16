package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectSecurityProfile = `{
  "block": {
    "attributes": {
      "allowed_access_control_tags": {
        "computed": true,
        "description": "The list of tags that a security profile uses to restrict access to resources in Amazon Connect.",
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
          "nesting_mode": "set"
        }
      },
      "description": {
        "computed": true,
        "description": "The description of the security profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description": "Permissions assigned to the security profile.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "security_profile_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the security profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_profile_name": {
        "computed": true,
        "description": "The name of the security profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "tag_restricted_resources": {
        "computed": true,
        "description": "The list of resources that a security profile applies tag restrictions to in Amazon Connect.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The tags used to organize, track, or control access for this resource.",
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Connect::SecurityProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectSecurityProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectSecurityProfile), &result)
	return &result
}
