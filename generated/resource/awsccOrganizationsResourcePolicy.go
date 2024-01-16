package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOrganizationsResourcePolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resource policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "content": {
        "description": "The policy document. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The unique identifier (ID) associated with this resource policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags that you want to attach to the resource policy",
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
    "description": "You can use AWS::Organizations::ResourcePolicy to delegate policy management for AWS Organizations to specified member accounts to perform policy actions that are by default available only to the management account.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOrganizationsResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOrganizationsResourcePolicy), &result)
	return &result
}
