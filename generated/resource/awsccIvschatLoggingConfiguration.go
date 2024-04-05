package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvschatLoggingConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "LoggingConfiguration ARN is automatically generated on creation and assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_configuration": {
        "description": "Destination configuration for IVS Chat logging.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs": {
              "computed": true,
              "description": "CloudWatch destination configuration for IVS Chat logging.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_group_name": {
                    "description": "Name of the Amazon CloudWatch Logs log group where chat activity will be logged.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "firehose": {
              "computed": true,
              "description": "Kinesis Firehose destination configuration for IVS Chat logging.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delivery_stream_name": {
                    "description": "Name of the Amazon Kinesis Firehose delivery stream where chat activity will be logged.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "s3": {
              "computed": true,
              "description": "S3 destination configuration for IVS Chat logging.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "description": "Name of the Amazon S3 bucket where chat activity will be logged.",
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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_configuration_id": {
        "computed": true,
        "description": "The system-generated ID of the logging configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the logging configuration. The value does not need to be unique.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the logging configuration. When the state is ACTIVE, the configuration is ready to log chat content.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource type definition for AWS::IVSChat::LoggingConfiguration.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIvschatLoggingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvschatLoggingConfiguration), &result)
	return &result
}
