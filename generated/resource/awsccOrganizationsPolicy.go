package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOrganizationsPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "ARN of the Policy",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_managed": {
        "computed": true,
        "description": "A boolean value that indicates whether the specified policy is an AWS managed policy. If true, then you can attach the policy to roots, OUs, or accounts, but you cannot edit it.",
        "description_kind": "plain",
        "type": "bool"
      },
      "content": {
        "description": "The Policy text content. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Human readable description of the policy",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the Policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description": "Id of the Policy",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags that you want to attach to the newly created policy. For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key identifier, or name, of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The string value that's associated with the key of the tag. You can set the value of a tag to an empty string, but you can't set the value of a tag to null.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_ids": {
        "computed": true,
        "description": "List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "type": {
        "description": "The type of policy to create. You can specify one of the following values: AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, TAG_POLICY, CHATBOT_POLICY, RESOURCE_CONTROL_POLICY",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Policies in AWS Organizations enable you to manage different features of the AWS accounts in your organization.  You can use policies when all features are enabled in your organization.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOrganizationsPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOrganizationsPolicy), &result)
	return &result
}
