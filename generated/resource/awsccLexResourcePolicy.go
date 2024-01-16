package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLexResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "The Physical ID of the resource policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "description": "A resource policy to add to the resource. The policy is a JSON structure following the IAM syntax that contains one or more statements that define the policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "description": "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description": "The current revision of the resource policy. Use the revision ID to make sure that you are updating the most current version of a resource policy when you add a policy statement to a resource, delete a resource, or update a resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "A resource policy with specified policy statements that attaches to a Lex bot or bot alias.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLexResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLexResourcePolicy), &result)
	return &result
}
