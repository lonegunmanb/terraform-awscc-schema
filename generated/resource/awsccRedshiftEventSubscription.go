package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftEventSubscription = `{
  "block": {
    "attributes": {
      "cust_subscription_id": {
        "computed": true,
        "description": "The name of the Amazon Redshift event notification subscription.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_aws_id": {
        "computed": true,
        "description": "The AWS account associated with the Amazon Redshift event notification subscription.",
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description": "A boolean value; set to true to activate the subscription, and set to false to create the subscription but not activate it.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "event_categories": {
        "computed": true,
        "description": "Specifies the Amazon Redshift event categories to be published by the event notification subscription.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "event_categories_list": {
        "computed": true,
        "description": "The list of Amazon Redshift event categories specified in the event notification subscription.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "severity": {
        "computed": true,
        "description": "Specifies the Amazon Redshift event severity to be published by the event notification subscription.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sns_topic_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_ids": {
        "computed": true,
        "description": "A list of one or more identifiers of Amazon Redshift source objects.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "source_ids_list": {
        "computed": true,
        "description": "A list of the sources that publish events to the Amazon Redshift event notification subscription.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "source_type": {
        "computed": true,
        "description": "The type of source that will be generating the events.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Amazon Redshift event notification subscription.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_creation_time": {
        "computed": true,
        "description": "The date and time the Amazon Redshift event notification subscription was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_name": {
        "description": "The name of the Amazon Redshift event notification subscription",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "The ` + "`" + `AWS::Redshift::EventSubscription` + "`" + ` resource creates an Amazon Redshift Event Subscription.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftEventSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftEventSubscription), &result)
	return &result
}
