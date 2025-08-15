package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmParameter = `{
  "block": {
    "attributes": {
      "allowed_pattern": {
        "computed": true,
        "description": "A regular expression used to validate the parameter value. For example, for ` + "`" + `` + "`" + `String` + "`" + `` + "`" + ` types with values restricted to numbers, you can specify the following: ` + "`" + `` + "`" + `AllowedPattern=^\\d+$` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "data_type": {
        "computed": true,
        "description": "The data type of the parameter, such as ` + "`" + `` + "`" + `text` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `aws:ec2:image` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `text` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Information about the parameter.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the parameter.\n  The reported maximum length of 2048 characters for a parameter name includes 1037 characters that are reserved for internal use by SYS. The maximum length for a parameter name that you specify is 1011 characters.\n This count of 1011 characters includes the characters in the ARN that precede the name you specify. This ARN length will vary depending on your partition and Region. For example, the following 45 characters count toward the 1011 character maximum for a parameter created in the US East (Ohio) Region: ` + "`" + `` + "`" + `arn:aws:ssm:us-east-2:111122223333:parameter/` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description": "Information about the policies assigned to a parameter.\n [Assigning parameter policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html) in the *User Guide*.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs). Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a SYS parameter to identify the type of resource to which it applies, the environment, or the type of configuration data referenced by the parameter.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tier": {
        "computed": true,
        "description": "The parameter tier.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of parameter.\n  Parameters of type ` + "`" + `` + "`" + `SecureString` + "`" + `` + "`" + ` are not supported by CFNlong.",
        "description_kind": "plain",
        "type": "string"
      },
      "value": {
        "computed": true,
        "description": "The parameter value.\n  If type is ` + "`" + `` + "`" + `StringList` + "`" + `` + "`" + `, the system returns a comma-separated string with no spaces between commas in the ` + "`" + `` + "`" + `Value` + "`" + `` + "`" + ` field.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSM::Parameter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmParameterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmParameter), &result)
	return &result
}
