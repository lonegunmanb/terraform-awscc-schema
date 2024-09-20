package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChatbotMicrosoftTeamsChannelConfiguration = `{
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
      "tags": {
        "computed": true,
        "description": "The tags to add to the configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "team_id": {
        "description": "The id of the Microsoft Teams team",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "teams_channel_id": {
        "description": "The id of the Microsoft Teams channel",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "teams_tenant_id": {
        "description": "The id of the Microsoft Teams tenant",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_role_required": {
        "computed": true,
        "description": "Enables use of a user role requirement in your chat configuration",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Resource schema for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccChatbotMicrosoftTeamsChannelConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChatbotMicrosoftTeamsChannelConfiguration), &result)
	return &result
}
