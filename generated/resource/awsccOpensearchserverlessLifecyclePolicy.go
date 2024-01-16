package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserverlessLifecyclePolicy = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the policy",
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
      "name": {
        "description": "The name of the policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "description": "The JSON policy document that is the content for the policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "The type of lifecycle policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Amazon OpenSearchServerless lifecycle policy resource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOpensearchserverlessLifecyclePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserverlessLifecyclePolicy), &result)
	return &result
}
