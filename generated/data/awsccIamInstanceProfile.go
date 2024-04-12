package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamInstanceProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_profile_name": {
        "computed": true,
        "description": "The name of the instance profile to create.\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
        "description_kind": "plain",
        "type": "string"
      },
      "path": {
        "computed": true,
        "description": "The path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.\n This parameter is optional. If it is not included, it defaults to a slash (/).\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (` + "`" + `` + "`" + `\\u0021` + "`" + `` + "`" + `) through the DEL character (` + "`" + `` + "`" + `\\u007F` + "`" + `` + "`" + `), including most punctuation characters, digits, and upper and lowercased letters.",
        "description_kind": "plain",
        "type": "string"
      },
      "roles": {
        "computed": true,
        "description": "The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::IAM::InstanceProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIamInstanceProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamInstanceProfile), &result)
	return &result
}
