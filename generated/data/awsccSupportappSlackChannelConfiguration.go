package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSupportappSlackChannelConfiguration = `{
  "block": {
    "attributes": {
      "channel_id": {
        "computed": true,
        "description": "The channel ID in Slack, which identifies a channel within a workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_name": {
        "computed": true,
        "description": "The channel name in Slack.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notify_on_add_correspondence_to_case": {
        "computed": true,
        "description": "Whether to notify when a correspondence is added to a case.",
        "description_kind": "plain",
        "type": "bool"
      },
      "notify_on_case_severity": {
        "computed": true,
        "description": "The severity level of a support case that a customer wants to get notified for.",
        "description_kind": "plain",
        "type": "string"
      },
      "notify_on_create_or_reopen_case": {
        "computed": true,
        "description": "Whether to notify when a case is created or reopened.",
        "description_kind": "plain",
        "type": "bool"
      },
      "notify_on_resolve_case": {
        "computed": true,
        "description": "Whether to notify when a case is resolved.",
        "description_kind": "plain",
        "type": "bool"
      },
      "team_id": {
        "computed": true,
        "description": "The team ID in Slack, which uniquely identifies a workspace.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SupportApp::SlackChannelConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSupportappSlackChannelConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSupportappSlackChannelConfiguration), &result)
	return &result
}
