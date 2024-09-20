package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectSecurityProfile = `{
  "block": {
    "attributes": {
      "allowed_access_control_hierarchy_group_id": {
        "computed": true,
        "description": "The identifier of the hierarchy group that a security profile uses to restrict access to resources in Amazon Connect.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "applications": {
        "computed": true,
        "description": "A list of third-party applications that the security profile will give access to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_permissions": {
              "computed": true,
              "description": "The permissions that the agent is granted on the application",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "namespace": {
              "computed": true,
              "description": "Namespace of the application that you want to give access to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description": "The description of the security profile.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hierarchy_restricted_resources": {
        "computed": true,
        "description": "The list of resources that a security profile applies hierarchy restrictions to in Amazon Connect.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_region": {
        "computed": true,
        "description": "The AWS Region where this resource was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The timestamp when this resource was last modified.",
        "description_kind": "plain",
        "type": "number"
      },
      "permissions": {
        "computed": true,
        "description": "Permissions assigned to the security profile.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The name of the security profile.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tag_restricted_resources": {
        "computed": true,
        "description": "The list of resources that a security profile applies tag restrictions to in Amazon Connect.",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Connect::SecurityProfile",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectSecurityProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectSecurityProfile), &result)
	return &result
}
