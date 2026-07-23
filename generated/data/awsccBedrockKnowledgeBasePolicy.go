package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockKnowledgeBasePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "knowledge_base_id": {
        "computed": true,
        "description": "The unique identifier of the knowledge base",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "The IAM policy document defining access permissions for the knowledge base",
        "description_kind": "plain",
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description": "The revision identifier for the policy, used for optimistic concurrency control",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Bedrock::KnowledgeBasePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockKnowledgeBasePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockKnowledgeBasePolicy), &result)
	return &result
}
