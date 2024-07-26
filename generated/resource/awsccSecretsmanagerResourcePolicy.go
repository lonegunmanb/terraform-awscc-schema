package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecretsmanagerResourcePolicy = `{
  "block": {
    "attributes": {
      "block_public_policy": {
        "computed": true,
        "description": "Specifies whether to block resource-based policies that allow broad access to the secret.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_policy": {
        "description": "A JSON-formatted string for an AWS resource-based policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_policy_id": {
        "computed": true,
        "description": "The Arn of the secret.",
        "description_kind": "plain",
        "type": "string"
      },
      "secret_id": {
        "description": "The ARN or name of the secret to attach the resource-based policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SecretsManager::ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecretsmanagerResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecretsmanagerResourcePolicy), &result)
	return &result
}
