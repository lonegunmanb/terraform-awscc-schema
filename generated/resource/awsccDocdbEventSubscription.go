package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDocdbEventSubscription = `{
  "block": {
    "attributes": {
      "enabled": {
        "computed": true,
        "description": "A Boolean value; set to true to activate the subscription, set to false to create the subscription but not active it.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "event_categories": {
        "computed": true,
        "description": "A list of event categories for a SourceType that you want to subscribe to.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "sns_topic_arn": {
        "description": "The Amazon Resource Name (ARN) of the SNS topic created for event notification. Amazon SNS creates the ARN when you create a topic and subscribe to it.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_ids": {
        "computed": true,
        "description": "The list of identifiers of the event sources for which events are returned",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "source_type": {
        "computed": true,
        "description": "The type of source that is generating the events.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_name": {
        "computed": true,
        "description": "The name of the subscription.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::DocDB::EventSubscription",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDocdbEventSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDocdbEventSubscription), &result)
	return &result
}
