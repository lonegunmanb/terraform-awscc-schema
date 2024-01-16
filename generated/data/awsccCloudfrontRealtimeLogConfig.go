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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kinesis_stream_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "stream_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "stream_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "fields": {
        "computed": true,
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
        "description_kind": "plain",
        "type": "string"
      },
      "sampling_rate": {
        "computed": true,
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
