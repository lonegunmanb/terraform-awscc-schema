package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvsRecordingConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_configuration": {
        "description": "Recording Destination Configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3": {
              "computed": true,
              "description": "Recording S3 Destination Configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
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
      "name": {
        "computed": true,
        "description": "Recording Configuration Name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recording_reconnect_window_seconds": {
        "computed": true,
        "description": "Recording Reconnect Window Seconds. (0 means disabled)",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "rendition_configuration": {
        "computed": true,
        "description": "Rendition Configuration describes which renditions should be recorded for a stream.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rendition_selection": {
              "computed": true,
              "description": "Resolution Selection indicates which set of renditions are recorded for a stream.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "renditions": {
              "computed": true,
              "description": "Renditions indicates which renditions are recorded for a stream.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "state": {
        "computed": true,
        "description": "Recording Configuration State.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the asset model.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "thumbnail_configuration": {
        "computed": true,
        "description": "Recording Thumbnail Configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "recording_mode": {
              "computed": true,
              "description": "Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resolution": {
              "computed": true,
              "description": "Resolution indicates the desired resolution of recorded thumbnails.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage": {
              "computed": true,
              "description": "Storage indicates the format in which thumbnails are recorded.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "target_interval_seconds": {
              "computed": true,
              "description": "Target Interval Seconds defines the interval at which thumbnails are recorded. This field is required if RecordingMode is INTERVAL.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::IVS::RecordingConfiguration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIvsRecordingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsRecordingConfiguration), &result)
	return &result
}
