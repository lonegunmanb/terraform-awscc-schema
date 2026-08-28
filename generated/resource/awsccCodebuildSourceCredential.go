package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodebuildSourceCredential = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the SourceCredential resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "auth_type": {
        "description": "The type of authentication used by the credentials. Valid options are OAUTH, BASIC_AUTH, PERSONAL_ACCESS_TOKEN, CODECONNECTIONS, or SECRETS_MANAGER. ",
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
      "server_type": {
        "description": "The type of source provider. The valid options are GITHUB, GITHUB_ENTERPRISE, GITLAB, GITLAB_SELF_MANAGED, or BITBUCKET.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "token": {
        "description": "For GitHub or GitHub Enterprise, this is the personal access token. For Bitbucket, this is either the access token or the app password. For the authType CODECONNECTIONS, this is the connectionArn. For the authType SECRETS_MANAGER, this is the secretArn.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "username": {
        "computed": true,
        "description": " The Bitbucket username when the authType is BASIC_AUTH. This parameter is not valid for other types of source providers or connections.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CodeBuild::SourceCredential",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCodebuildSourceCredentialSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodebuildSourceCredential), &result)
	return &result
}
