package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVpclatticeResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Retrieves information about the resource policy. The resource policy is an IAM policy created by AWS RAM on behalf of the resource owner when they share a resource.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccVpclatticeResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVpclatticeResourcePolicy), &result)
	return &result
}
