package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSupportappSlackWorkspaceConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "team_id": {
        "description": "The team ID in Slack, which uniquely identifies a workspace.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version_id": {
        "computed": true,
        "description": "An identifier used to update an existing Slack workspace configuration in AWS CloudFormation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "An AWS Support App resource that creates, updates, lists, and deletes Slack workspace configurations.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSupportappSlackWorkspaceConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSupportappSlackWorkspaceConfiguration), &result)
	return &result
}
