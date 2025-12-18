package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRamResourceShare = `{
  "block": {
    "attributes": {
      "allow_external_principals": {
        "computed": true,
        "description": "Specifies whether principals outside your organization in AWS Organizations can be associated with a resource share. A value of ` + "`" + `true` + "`" + ` lets you share with individual AWS accounts that are not in your organization. A value of ` + "`" + `false` + "`" + ` only has meaning if your account is a member of an AWS Organization. The default value is ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The date and time when the resource share was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "feature_set": {
        "computed": true,
        "description": "The feature set of the resource share.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The date and time when the resource share was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Specifies the name of the resource share.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owning_account_id": {
        "computed": true,
        "description": "The ID of the AWS account that owns the resource share.",
        "description_kind": "plain",
        "type": "string"
      },
      "permission_arns": {
        "computed": true,
        "description": "Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the AWS RAM permission to associate with the resource share. If you do not specify an ARN for the permission, AWS RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "principals": {
        "computed": true,
        "description": "Specifies the principals to associate with the resource share. The possible values are:\n\n- An AWS account ID\n\n- An Amazon Resource Name (ARN) of an organization in AWS Organizations\n\n- An ARN of an organizational unit (OU) in AWS Organizations\n\n- An ARN of an IAM role\n\n- An ARN of an IAM user",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "resource_arns": {
        "computed": true,
        "description": "Specifies a list of one or more ARNs of the resources to associate with the resource share.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "sources": {
        "computed": true,
        "description": "Specifies from which source accounts the service principal has access to the resources in this resource share.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "The current status of the resource share.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Specifies one or more tags to attach to the resource share itself. It doesn't attach the tags to the resources associated with the resource share.",
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
    "description": "Resource type definition for AWS::RAM::ResourceShare",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRamResourceShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRamResourceShare), &result)
	return &result
}
