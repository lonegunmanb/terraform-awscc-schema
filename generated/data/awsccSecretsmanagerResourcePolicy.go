package data

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
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_policy": {
        "computed": true,
        "description": "A JSON-formatted string for an AWS resource-based policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_policy_id": {
        "computed": true,
        "description": "The Arn of the secret.",
        "description_kind": "plain",
        "type": "string"
      },
      "secret_id": {
        "computed": true,
        "description": "The ARN or name of the secret to attach the resource-based policy.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecretsManager::ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecretsmanagerResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecretsmanagerResourcePolicy), &result)
	return &result
}
