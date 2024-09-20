package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_name": {
        "computed": true,
        "description": "The name of the group to create. Do not include the path in this value.\n The group name must be unique within the account. Group names are not distinguished by case. For example, you cannot create groups named both \"ADMINS\" and \"admins\". If you don't specify a name, CFN generates a unique physical ID and uses that ID for the group name.\n  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.\n  If you specify a name, you must specify the ` + "`" + `` + "`" + `CAPABILITY_NAMED_IAM` + "`" + `` + "`" + ` value to acknowledge your template's capabilities. For more information, see [Acknowledging Resources in Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities).\n  Naming an IAM resource can cause an unrecoverable error if you reuse the same template in multiple Regions. To prevent this, we recommend using ` + "`" + `` + "`" + `Fn::Join` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `AWS::Region` + "`" + `` + "`" + ` to create a Region-specific name, as in the following example: ` + "`" + `` + "`" + `{\"Fn::Join\": [\"\", [{\"Ref\": \"AWS::Region\"}, {\"Ref\": \"MyResourceName\"}]]}` + "`" + `` + "`" + `.",
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
      "managed_policy_arns": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM policy you want to attach.\n For more information about ARNs, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "path": {
        "computed": true,
        "description": "The path to the group. For more information about paths, see [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.\n This parameter is optional. If it is not included, it defaults to a slash (/).\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (` + "`" + `` + "`" + `\\u0021` + "`" + `` + "`" + `) through the DEL character (` + "`" + `` + "`" + `\\u007F` + "`" + `` + "`" + `), including most punctuation characters, digits, and upper and lowercased letters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description": "Adds or updates an inline policy document that is embedded in the specified IAM group. To view AWS::IAM::Group snippets, see [Declaring an Group Resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-iam-group).\n  The name of each inline policy for a role, user, or group must be unique. If you don't choose unique names, updates to the IAM identity will fail. \n  For information about limits on the number of inline policies that you can embed in a group, see [Limitations on Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the *User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy_document": {
              "computed": true,
              "description": "The policy document.",
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
      }
    },
    "description": "Creates a new group.\n  For information about the number of groups you can create, see [Limitations on Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the *User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIamGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamGroup), &result)
	return &result
}
