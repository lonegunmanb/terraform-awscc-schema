package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description": "The IAM policy document defining access permissions for the guardrail and guardrail profile resources",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "description": "The ARN of the Bedrock Guardrail or Guardrail Profile resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::Bedrock::ResourcePolicy Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockResourcePolicy), &result)
	return &result
}
