package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChatbotSlackChannelConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_name": {
        "description": "The name of the configuration",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "guardrail_policies": {
        "computed": true,
        "description": "The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "iam_role_arn": {
        "description": "The ARN of the IAM role that defines the permissions for AWS Chatbot",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_level": {
        "computed": true,
        "description": "Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "slack_channel_id": {
        "description": "The id of the Slack channel",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "slack_workspace_id": {
        "description": "The id of the Slack workspace",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sns_topic_arns": {
        "computed": true,
        "description": "ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "user_role_required": {
        "computed": true,
        "description": "Enables use of a user role requirement in your chat configuration",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Resource schema for AWS::Chatbot::SlackChannelConfiguration.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccChatbotSlackChannelConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChatbotSlackChannelConfiguration), &result)
	return &result
}
