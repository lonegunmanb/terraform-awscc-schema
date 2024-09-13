package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSnsSubscription = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Arn of the subscription",
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_policy": {
        "computed": true,
        "description": "The delivery policy JSON assigned to the subscription. Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "The subscription's endpoint. The endpoint value depends on the protocol that you specify. ",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_policy": {
        "computed": true,
        "description": "The filter policy JSON assigned to the subscription. Enables the subscriber to filter out unwanted messages.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_policy_scope": {
        "computed": true,
        "description": "This attribute lets you choose the filtering scope by using one of the following string value types: MessageAttributes (default) and MessageBody.",
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
      "protocol": {
        "description": "The subscription's protocol.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "raw_message_delivery": {
        "computed": true,
        "description": "When set to true, enables raw message delivery. Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "redrive_policy": {
        "computed": true,
        "description": "When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue. Messages that can't be delivered due to client errors are held in the dead-letter queue for further analysis or reprocessing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "For cross-region subscriptions, the region in which the topic resides.If no region is specified, AWS CloudFormation uses the region of the caller as the default.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replay_policy": {
        "computed": true,
        "description": "Specifies whether Amazon SNS resends the notification to the subscription when a message's attribute changes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_role_arn": {
        "computed": true,
        "description": "This property applies only to Amazon Data Firehose delivery stream subscriptions.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "topic_arn": {
        "description": "The ARN of the topic to subscribe to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SNS::Subscription",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSnsSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSnsSubscription), &result)
	return &result
}
