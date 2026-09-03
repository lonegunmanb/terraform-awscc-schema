package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeMediaPipelineKinesisVideoStreamPool = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the Kinesis Video Stream Pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_timestamp": {
        "computed": true,
        "description": "The time at which the Kinesis Video Stream Pool was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "pool_id": {
        "computed": true,
        "description": "The unique identifier of the Kinesis Video Stream Pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "pool_name": {
        "description": "The name of the Kinesis Video Stream Pool.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pool_status": {
        "computed": true,
        "description": "The status of the Kinesis Video Stream Pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "stream_configuration": {
        "description": "The configuration settings for the Kinesis video stream.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_retention_in_hours": {
              "computed": true,
              "description": "The amount of time that data is retained, in hours.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "region": {
              "description": "The AWS Region of the video stream.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the Kinesis Video Stream Pool.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_timestamp": {
        "computed": true,
        "description": "The time at which the Kinesis Video Stream Pool was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for an Amazon Chime SDK Media Pipeline Kinesis Video Stream Pool",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccChimeMediaPipelineKinesisVideoStreamPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeMediaPipelineKinesisVideoStreamPool), &result)
	return &result
}
