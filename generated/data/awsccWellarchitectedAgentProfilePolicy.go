package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedAgentProfilePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description": "The IAM resource-based policy document as a JSON string. Allow-only statements, no negation condition operators, maximum 51200 characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the parent Agent Profile the policy is attached to. Pass ` + "`" + `!GetAtt` + "`" + ` (or ` + "`" + `!Ref` + "`" + `) of the parent AWS::WellArchitected::AgentProfile to flow its ARN here. Also serves as the primary identifier of this policy attachment.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::WellArchitected::AgentProfilePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWellarchitectedAgentProfilePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedAgentProfilePolicy), &result)
	return &result
}
