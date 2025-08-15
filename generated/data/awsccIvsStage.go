package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvsStage = `{
  "block": {
    "attributes": {
      "active_session_id": {
        "computed": true,
        "description": "ID of the active session within the stage.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Stage ARN is automatically generated on creation and assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_participant_recording_configuration": {
        "computed": true,
        "description": "Configuration object for individual participant recording, to attach to the new stage.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "hls_configuration": {
              "computed": true,
              "description": "HLS configuration object for individual participant recording.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "participant_recording_hls_configuration": {
                    "computed": true,
                    "description": "An object representing a configuration of participant HLS recordings for individual participant recording.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "target_segment_duration_seconds": {
                          "computed": true,
                          "description": "Defines the target duration for recorded segments generated when recording a stage participant. Segments may have durations longer than the specified value when needed to ensure each segment begins with a keyframe. Default: 6.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "media_types": {
              "computed": true,
              "description": "Types of media to be recorded. Default: AUDIO_VIDEO.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "recording_reconnect_window_seconds": {
              "computed": true,
              "description": "If a stage publisher disconnects and then reconnects within the specified interval, the multiple recordings will be considered a single recording and merged together. The default value is 0, which disables merging.",
              "description_kind": "plain",
              "type": "number"
            },
            "storage_configuration_arn": {
              "computed": true,
              "description": "ARN of the StorageConfiguration resource to use for individual participant recording.",
              "description_kind": "plain",
              "type": "string"
            },
            "thumbnail_configuration": {
              "computed": true,
              "description": "A complex type that allows you to enable/disable the recording of thumbnails for individual participant recording and modify the interval at which thumbnails are generated for the live session.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "participant_thumbnail_configuration": {
                    "computed": true,
                    "description": "An object representing a configuration of thumbnails for recorded video from an individual participant.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "recording_mode": {
                          "computed": true,
                          "description": "Thumbnail recording mode. Default: DISABLED.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "storage": {
                          "computed": true,
                          "description": "Indicates the format in which thumbnails are recorded. SEQUENTIAL records all generated thumbnails in a serial manner, to the media/thumbnails/high directory. LATEST saves the latest thumbnail in media/latest_thumbnail/high/thumb.jpg and overwrites it at the interval specified by targetIntervalSeconds. You can enable both SEQUENTIAL and LATEST. Default: SEQUENTIAL.",
                          "description_kind": "plain",
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "target_interval_seconds": {
                          "computed": true,
                          "description": "The targeted thumbnail-generation interval in seconds. This is configurable only if recordingMode is INTERVAL. Default: 60.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Stage name",
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
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::IVS::Stage",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIvsStageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsStage), &result)
	return &result
}
