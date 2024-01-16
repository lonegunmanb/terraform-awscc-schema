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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kinesis_stream_config": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "stream_arn": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sampling_rate": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::CloudFront::RealtimeLogConfig",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontRealtimeLogConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontRealtimeLogConfig), &result)
	return &result
}
