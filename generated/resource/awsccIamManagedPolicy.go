package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamManagedPolicy = `{
  "block": {
    "attributes": {
      "attachment_count": {
        "computed": true,
        "description": "The number of entities (users, groups, and roles) that the policy is attached to.",
        "description_kind": "plain",
        "type": "number"
      },
      "create_date": {
        "computed": true,
        "description": "The date and time, in ISO 8601 date-time format, when the policy was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_version_id": {
        "computed": true,
        "description": "The identifier for the version of the policy that is set as the default version.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A friendly description of the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "groups": {
        "computed": true,
        "description": "The name (friendly name, not ARN) of the group to attach the policy to.",
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
      "is_attachable": {
        "computed": true,
        "description": "Specifies whether the policy can be attached to an IAM user, group, or role.",
        "description_kind": "plain",
        "type": "bool"
      },
      "managed_policy_name": {
        "computed": true,
        "description": "The friendly name of the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path": {
        "computed": true,
        "description": "The path for the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "permissions_boundary_usage_count": {
        "computed": true,
        "description": "The number of entities (users and roles) for which the policy is used to set the permissions boundary.",
        "description_kind": "plain",
        "type": "number"
      },
      "policy_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the managed policy",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description": "The JSON policy document that you want to use as the content for the new policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description": "The stable and unique string identifying the policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "roles": {
        "computed": true,
        "description": "The name (friendly name, not ARN) of the role to attach the policy to.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "update_date": {
        "computed": true,
        "description": "The date and time, in ISO 8601 date-time format, when the policy was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "users": {
        "computed": true,
        "description": "The name (friendly name, not ARN) of the IAM user to attach the policy to.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Resource Type definition for AWS::IAM::ManagedPolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIamManagedPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamManagedPolicy), &result)
	return &result
}
