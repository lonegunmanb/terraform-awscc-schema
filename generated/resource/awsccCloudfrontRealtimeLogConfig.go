package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontRealtimeLogConfig = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "end_points": {
        "description": "Contains information about the Amazon Kinesis data stream where you are sending real-time log data for this real-time log configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kinesis_stream_config": {
              "description": "Contains information about the Amazon Kinesis data stream where you are sending real-time log data.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "role_arn": {
                    "description": "The Amazon Resource Name (ARN) of an IAMlong (IAM) role that CloudFront can use to send real-time log data to your Kinesis data stream.\n For more information the IAM role, see [Real-time log configuration IAM role](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-iam-role) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "stream_arn": {
                    "description": "The Amazon Resource Name (ARN) of the Kinesis data stream where you are sending real-time log data.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "stream_type": {
              "description": "The type of data stream where you are sending real-time log data. The only valid value is ` + "`" + `` + "`" + `Kinesis` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "fields": {
        "description": "A list of fields that are included in each real-time log record. In an API response, the fields are provided in the same order in which they are sent to the Amazon Kinesis data stream.\n For more information about fields, see [Real-time log configuration fields](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) in the *Amazon CloudFront Developer Guide*.",
        "description_kind": "plain",
        "required": true,
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
      "name": {
        "description": "The unique name of this real-time log configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sampling_rate": {
        "description": "The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. The sampling rate is an integer between 1 and 100, inclusive.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description": "A real-time log configuration.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontRealtimeLogConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontRealtimeLogConfig), &result)
	return &result
}
