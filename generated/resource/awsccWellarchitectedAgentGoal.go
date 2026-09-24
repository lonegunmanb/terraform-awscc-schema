package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedAgentGoal = `{
  "block": {
    "attributes": {
      "agent_goal_id": {
        "computed": true,
        "description": "The service-generated unique identifier of the Agent Goal.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The synthetic Amazon Resource Name (ARN) of the Agent Goal, composed of the parent profile ARN and the goal id.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the goal was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The identifier of the system or user that created this goal.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the Agent Goal.",
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
      "last_modified_at": {
        "computed": true,
        "description": "The timestamp when the goal was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_by": {
        "computed": true,
        "description": "The identifier of the system or user that last modified this goal.",
        "description_kind": "plain",
        "type": "string"
      },
      "pillars": {
        "description": "The list of Well-Architected pillars this goal targets.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "profile_arn": {
        "description": "The Amazon Resource Name (ARN) of the parent Agent Profile that owns this goal. Pass ` + "`" + `!Ref` + "`" + ` of the parent AWS::WellArchitected::AgentProfile to flow its ARN here.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "title": {
        "description": "The title of the Agent Goal.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::WellArchitected::AgentGoal. An Agent Goal expresses a Well-Architected objective (a title, optional description, and the pillars it targets) within an Agent Profile, used to focus Well-Architected Agent recommendations.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWellarchitectedAgentGoalSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedAgentGoal), &result)
	return &result
}
