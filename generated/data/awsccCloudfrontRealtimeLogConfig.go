package data

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
        "computed": true,
        "description": "Contains information about the Amazon Kinesis data stream where you are sending real-time log data for this real-time log configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kinesis_stream_config": {
              "computed": true,
              "description": "Contains information about the Amazon Kinesis data stream where you are sending real-time log data in a real-time log configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "role_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of an IAMlong (IAM) role that CloudFront can use to send real-time log data to your Kinesis data stream.\n For more information the IAM role, see [Real-time log configuration IAM role](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-iam-role) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "stream_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the Kinesis data stream where you are sending real-time log data.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "stream_type": {
              "computed": true,
              "description": "The type of data stream where you are sending real-time log data. The only valid value is ` + "`" + `` + "`" + `Kinesis` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "fields": {
        "computed": true,
        "description": "A list of fields that are included in each real-time log record. In an API response, the fields are provided in the same order in which they are sent to the Amazon Kinesis data stream.\n For more information about fields, see [Real-time log configuration fields](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) in the *Amazon CloudFront Developer Guide*.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique name of this real-time log configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "sampling_rate": {
        "computed": true,
        "description": "The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. The sampling rate is an integer between 1 and 100, inclusive.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::CloudFront::RealtimeLogConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontRealtimeLogConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontRealtimeLogConfig), &result)
	return &result
}
