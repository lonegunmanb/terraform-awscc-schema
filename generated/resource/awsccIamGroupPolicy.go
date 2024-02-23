package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamGroupPolicy = `{
  "block": {
    "attributes": {
      "group_name": {
        "description": "The name of the group to associate the policy with.\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.",
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
      "policy_document": {
        "computed": true,
        "description": "The policy document.\n You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.\n The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:\n  +  Any printable ASCII character ranging from the space character (` + "`" + `` + "`" + `\\u0020` + "`" + `` + "`" + `) through the end of the ASCII character range\n  +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ` + "`" + `` + "`" + `\\u00FF` + "`" + `` + "`" + `)\n  +  The special characters tab (` + "`" + `` + "`" + `\\u0009` + "`" + `` + "`" + `), line feed (` + "`" + `` + "`" + `\\u000A` + "`" + `` + "`" + `), and carriage return (` + "`" + `` + "`" + `\\u000D` + "`" + `` + "`" + `)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_name": {
        "description": "The name of the policy document.\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Adds or updates an inline policy document that is embedded in the specified IAM group.\n A group can also have managed policies attached to it. To attach a managed policy to a group, use [AWS::IAM::Group](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html). To create a new managed policy, use [AWS::IAM::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html). For information about policies, see [Managed policies and inline policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *IAM User Guide*.\n For information about the maximum number of inline policies that you can embed in a group, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIamGroupPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamGroupPolicy), &result)
	return &result
}
