package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSupportappSlackWorkspaceConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "team_id": {
        "computed": true,
        "description": "The team ID in Slack, which uniquely identifies a workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_id": {
        "computed": true,
        "description": "An identifier used to update an existing Slack workspace configuration in AWS CloudFormation.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SupportApp::SlackWorkspaceConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSupportappSlackWorkspaceConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSupportappSlackWorkspaceConfiguration), &result)
	return &result
}
