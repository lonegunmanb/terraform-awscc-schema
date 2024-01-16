package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEmrStudioSessionMapping = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_name": {
        "description": "The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_type": {
        "description": "Specifies whether the identity to map to the Studio is a user or a group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "session_policy_arn": {
        "description": "The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "studio_id": {
        "description": "The ID of the Amazon EMR Studio to which the user or group will be mapped.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "An example resource schema demonstrating some basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEmrStudioSessionMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEmrStudioSessionMapping), &result)
	return &result
}
