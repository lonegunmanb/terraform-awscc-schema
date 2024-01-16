package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesConfigurationSetEventDestination = `{
  "block": {
    "attributes": {
      "configuration_set_name": {
        "description": "The name of the configuration set that contains the event destination.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_destination": {
        "description": "The event destination object.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_destination": {
              "computed": true,
              "description": "An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dimension_configurations": {
                    "computed": true,
                    "description": "A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "default_dimension_value": {
                          "description": "The default value of the dimension that is published to Amazon CloudWatch if you do not provide the value of the dimension when you send an email.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "dimension_name": {
                          "description": "The name of an Amazon CloudWatch dimension associated with an email sending metric.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "dimension_value_source": {
                          "description": "The place where Amazon SES finds the value of a dimension to publish to Amazon CloudWatch. To use the message tags that you specify using an X-SES-MESSAGE-TAGS header or a parameter to the SendEmail/SendRawEmail API, specify messageTag. To use your own email headers, specify emailHeader. To put a custom tag on any link included in your email, specify linkTag.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "enabled": {
              "computed": true,
              "description": "Sets whether Amazon SES publishes events to this destination when you send an email with the associated configuration set. Set to true to enable publishing to this destination; set to false to prevent publishing to this destination. The default value is false.   ",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kinesis_firehose_destination": {
              "computed": true,
              "description": "An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delivery_stream_arn": {
                    "description": "The ARN of the Amazon Kinesis Firehose stream that email sending events should be published to.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "iam_role_arn": {
                    "description": "The ARN of the IAM role under which Amazon SES publishes email sending events to the Amazon Kinesis Firehose stream.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "matching_event_types": {
              "description": "The type of email sending events, send, reject, bounce, complaint, delivery, open, click, renderingFailure, deliveryDelay, and subscription.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "The name of the event destination set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sns_destination": {
              "computed": true,
              "description": "An object that contains SNS topic ARN associated event destination.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "topic_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SES::ConfigurationSetEventDestination",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesConfigurationSetEventDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesConfigurationSetEventDestination), &result)
	return &result
}
