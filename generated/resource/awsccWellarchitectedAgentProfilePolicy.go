package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedAgentProfilePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "description": "The IAM resource-based policy document as a JSON string. Allow-only statements, no negation condition operators, maximum 51200 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_arn": {
        "description": "The Amazon Resource Name (ARN) of the parent Agent Profile the policy is attached to. Pass ` + "`" + `!GetAtt` + "`" + ` (or ` + "`" + `!Ref` + "`" + `) of the parent AWS::WellArchitected::AgentProfile to flow its ARN here. Also serves as the primary identifier of this policy attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::WellArchitected::AgentProfilePolicy. Attaches an IAM resource-based policy to an Agent Profile controlling cross-account access to the profile and its child resources (recommendations, goals, contexts). One policy per profile: the parent profile's ARN is the primary identifier. Standard CloudFormation policy-attachment pattern (see AWS::S3::BucketPolicy).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWellarchitectedAgentProfilePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedAgentProfilePolicy), &result)
	return &result
}
