package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvsChannel = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Channel ARN is automatically generated on creation and assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorized": {
        "computed": true,
        "description": "Whether the channel is authorized.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "container_format": {
        "computed": true,
        "description": "Indicates which content-packaging format is used (MPEG-TS or fMP4). If multitrackInputConfiguration is specified and enabled is true, then containerFormat is required and must be set to FRAGMENTED_MP4. Otherwise, containerFormat may be set to TS or FRAGMENTED_MP4. Default: TS.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ingest_endpoint": {
        "computed": true,
        "description": "Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.",
        "description_kind": "plain",
        "type": "string"
      },
      "insecure_ingest": {
        "computed": true,
        "description": "Whether the channel allows insecure ingest.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "latency_mode": {
        "computed": true,
        "description": "Channel latency mode.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multitrack_input_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Indicates whether multitrack input is enabled. Can be set to true only if channel type is STANDARD. Setting enabled to true with any other channel type will cause an exception. If true, then policy, maximumResolution, and containerFormat are required, and containerFormat must be set to FRAGMENTED_MP4. Default: false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "maximum_resolution": {
              "computed": true,
              "description": "Maximum resolution for multitrack input. Required if enabled is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy": {
              "computed": true,
              "description": "Indicates whether multitrack input is allowed or required. Required if enabled is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "name": {
        "computed": true,
        "description": "Channel",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "playback_url": {
        "computed": true,
        "description": "Channel Playback URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "preset": {
        "computed": true,
        "description": "Optional transcode preset for the channel. This is selectable only for ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default preset is HIGHER_BANDWIDTH_DELIVERY. For other channel types (BASIC and STANDARD), preset is the empty string (\"\").",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recording_configuration_arn": {
        "computed": true,
        "description": "Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: \"\" (recording is disabled).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the asset model.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "type": {
        "computed": true,
        "description": "Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::IVS::Channel",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIvsChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsChannel), &result)
	return &result
}
