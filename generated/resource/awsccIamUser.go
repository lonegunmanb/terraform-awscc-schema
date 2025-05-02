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
        "description": "Creates a password for the specified IAM user. A password allows an IAM user to access AWS services through the console.\n You can use the CLI, the AWS API, or the *Users* page in the IAM console to create a password for any IAM user. Use [ChangePassword](https://docs.aws.amazon.com/IAM/latest/APIReference/API_ChangePassword.html) to update your own existing password in the *My Security Credentials* page in the console.\n For more information about managing passwords, see [Managing passwords](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html) in the *User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "password": {
              "computed": true,
              "description": "The user's password.",
              "description_kind": "plain",
              "optional": true,
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
        "description": "A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the user.\n For more information about ARNs, see [Amazon Resource Names (ARNs) and Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "path": {
        "computed": true,
        "description": "The path for the user name. For more information about paths, see [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.\n This parameter is optional. If it is not included, it defaults to a slash (/).\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (` + "`" + `` + "`" + `\\u0021` + "`" + `` + "`" + `) through the DEL character (` + "`" + `` + "`" + `\\u007F` + "`" + `` + "`" + `), including most punctuation characters, digits, and upper and lowercased letters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "permissions_boundary": {
        "computed": true,
        "description": "The ARN of the managed policy that is used to set the permissions boundary for the user.\n A permissions boundary policy defines the maximum permissions that identity-based policies can grant to an entity, but does not grant permissions. Permissions boundaries do not define the maximum permissions that a resource-based policy can grant to an entity. To learn more, see [Permissions boundaries for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide*.\n For more information about policy types, see [Policy types](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policy-types) in the *IAM User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description": "Adds or updates an inline policy document that is embedded in the specified IAM user. To view AWS::IAM::User snippets, see [Declaring an User Resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-iam-user).\n  The name of each policy for a role, user, or group must be unique. If you don't choose unique names, updates to the IAM identity will fail. \n  For information about limits on the number of inline policies that you can embed in a user, see [Limitations on Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the *User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy_document": {
              "computed": true,
              "description": "The entire contents of the policy that defines permissions. For more information, see [Overview of JSON policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy_name": {
              "computed": true,
              "description": "The friendly name (not ARN) identifying the policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A list of tags that you want to attach to the new user. Each tag consists of a key name and an associated value. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide*.\n  If any one of the tags is invalid or if you exceed the allowed maximum number of tags, then the entire request fails and the resource is not created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name that can be used to look up or retrieve the associated value. For example, ` + "`" + `` + "`" + `Department` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Cost Center` + "`" + `` + "`" + ` are common choices.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value associated with this tag. For example, tags with a key name of ` + "`" + `` + "`" + `Department` + "`" + `` + "`" + ` could have values such as ` + "`" + `` + "`" + `Human Resources` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Accounting` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `Support` + "`" + `` + "`" + `. Tags with a key name of ` + "`" + `` + "`" + `Cost Center` + "`" + `` + "`" + ` might have values that consist of the number associated with the different cost centers in your company. Typically, many resources have tags with the same key name but with different values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "user_name": {
        "computed": true,
        "description": "The name of the user to create. Do not include the path in this value.\n This parameter allows (per its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-. The user name must be unique within the account. User names are not distinguished by case. For example, you cannot create users named both \"John\" and \"john\".\n If you don't specify a name, CFN generates a unique physical ID and uses that ID for the user name.\n If you specify a name, you must specify the ` + "`" + `` + "`" + `CAPABILITY_NAMED_IAM` + "`" + `` + "`" + ` value to acknowledge your template's capabilities. For more information, see [Acknowledging Resources in Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities).\n  Naming an IAM resource can cause an unrecoverable error if you reuse the same template in multiple Regions. To prevent this, we recommend using ` + "`" + `` + "`" + `Fn::Join` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `AWS::Region` + "`" + `` + "`" + ` to create a Region-specific name, as in the following example: ` + "`" + `` + "`" + `{\"Fn::Join\": [\"\", [{\"Ref\": \"AWS::Region\"}, {\"Ref\": \"MyResourceName\"}]]}` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Creates a new IAM user for your AWS-account.\n  For information about quotas for the number of IAM users you can create, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIamUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamUser), &result)
	return &result
}
