package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEmrStudioSessionMapping = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_name": {
        "computed": true,
        "description": "The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_type": {
        "computed": true,
        "description": "Specifies whether the identity to map to the Studio is a user or a group.",
        "description_kind": "plain",
        "type": "string"
      },
      "session_policy_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.",
        "description_kind": "plain",
        "type": "string"
      },
      "studio_id": {
        "computed": true,
        "description": "The ID of the Amazon EMR Studio to which the user or group will be mapped.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EMR::StudioSessionMapping",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEmrStudioSessionMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEmrStudioSessionMapping), &result)
	return &result
}
