package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAmplifyWebhook = `{
  "block": {
    "attributes": {
      "app_id": {
        "computed": true,
        "description": "The unique ID for an Amplify app.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the webhook.",
        "description_kind": "plain",
        "type": "string"
      },
      "branch_name": {
        "description": "The name for a branch that is part of an Amplify app.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description for a webhook.",
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
      "tags": {
        "computed": true,
        "description": "Tags for the webhook.",
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
      "webhook_id": {
        "computed": true,
        "description": "The unique ID for a webhook.",
        "description_kind": "plain",
        "type": "string"
      },
      "webhook_url": {
        "computed": true,
        "description": "The URL of the webhook.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Amplify::Webhook. Creates a webhook on an Amplify app.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAmplifyWebhookSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAmplifyWebhook), &result)
	return &result
}
