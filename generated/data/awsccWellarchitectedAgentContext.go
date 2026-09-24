package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedAgentContext = `{
  "block": {
    "attributes": {
      "agent_context_id": {
        "computed": true,
        "description": "The service-generated unique identifier of the Agent Context.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_type": {
        "computed": true,
        "description": "Type of the application described by this context. Mirrors the value stored in ` + "`" + `Content.applicationType` + "`" + ` and is surfaced as a typed read-only attribute by the service for discoverability.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The synthetic Amazon Resource Name (ARN) of the Agent Context, composed of the parent profile ARN and the context id.",
        "description_kind": "plain",
        "type": "string"
      },
      "content": {
        "computed": true,
        "description": "The free-form content of the Agent Context, supplied as an arbitrary JSON object.",
        "description_kind": "plain",
        "type": "string"
      },
      "context_type": {
        "computed": true,
        "description": "The type of the Agent Context.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the context was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The identifier of the system or user that created this context.",
        "description_kind": "plain",
        "type": "string"
      },
      "criticality": {
        "computed": true,
        "description": "Business criticality of the application described by this context. Mirrors the value stored in ` + "`" + `Content.criticality` + "`" + ` and is surfaced as a typed read-only attribute by the service for discoverability.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_at": {
        "computed": true,
        "description": "The timestamp when the context was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_by": {
        "computed": true,
        "description": "The identifier of the system or user that last modified this context.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the parent Agent Profile that owns this context. Pass ` + "`" + `!Ref` + "`" + ` of the parent AWS::WellArchitected::AgentProfile to flow its ARN here.",
        "description_kind": "plain",
        "type": "string"
      },
      "title": {
        "computed": true,
        "description": "The title of the Agent Context.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::WellArchitected::AgentContext",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWellarchitectedAgentContextSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedAgentContext), &result)
	return &result
}
