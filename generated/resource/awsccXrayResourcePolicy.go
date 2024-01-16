package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccXrayResourcePolicy = `{
  "block": {
    "attributes": {
      "bypass_policy_lockout_check": {
        "computed": true,
        "description": "A flag to indicate whether to bypass the resource policy lockout safety check",
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
      "policy_document": {
        "description": "The resource policy document, which can be up to 5kb in size.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_name": {
        "description": "The name of the resource policy. Must be unique within a specific AWS account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "This schema provides construct and validation rules for AWS-XRay Resource Policy resource parameters.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccXrayResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccXrayResourcePolicy), &result)
	return &result
}
