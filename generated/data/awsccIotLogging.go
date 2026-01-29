package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotLogging = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
        "description_kind": "plain",
        "type": "string"
      },
      "default_log_level": {
        "computed": true,
        "description": "The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_configurations": {
        "computed": true,
        "description": "Configurations for event-based logging that specifies which event types to log and their logging settings. Overrides account-level logging for the specified event",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_type": {
              "computed": true,
              "description": "The type of event to log. These include event types like Connect, Publish, and Disconnect.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_destination": {
              "computed": true,
              "description": "CloudWatch Log Group for event-based logging. Specifies where log events should be sent. The log destination for event-based logging overrides default Log Group for the specified event type and applies to all resources associated with that event.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_level": {
              "computed": true,
              "description": "The logging level for the specified event type. Determines the verbosity of log messages generated for this event type.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of the role that allows IoT to write to Cloudwatch logs.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoT::Logging",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotLoggingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotLogging), &result)
	return &result
}
