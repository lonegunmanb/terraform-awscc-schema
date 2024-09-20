package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediatailorChannel = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe ARN of the channel.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "audiences": {
        "computed": true,
        "description": "\u003cp\u003eThe list of audiences defined in channel.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "channel_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filler_slate": {
        "computed": true,
        "description": "\u003cp\u003eSlate VOD source configuration.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_location_name": {
              "computed": true,
              "description": "\u003cp\u003eThe name of the source location where the slate VOD source is stored.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vod_source_name": {
              "computed": true,
              "description": "\u003cp\u003eThe slate VOD source name. The VOD source must already exist in a source location before it can be used for slate.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe log configuration for the channel.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_types": {
              "computed": true,
              "description": "\u003cp\u003eThe log types.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "outputs": {
        "description": "\u003cp\u003eThe channel's output properties.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dash_playlist_settings": {
              "computed": true,
              "description": "\u003cp\u003eDash manifest configuration parameters.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "manifest_window_seconds": {
                    "computed": true,
                    "description": "\u003cp\u003eThe total duration (in seconds) of each manifest. Minimum value: \u003ccode\u003e30\u003c/code\u003e seconds. Maximum value: \u003ccode\u003e3600\u003c/code\u003e seconds.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_buffer_time_seconds": {
                    "computed": true,
                    "description": "\u003cp\u003eMinimum amount of content (measured in seconds) that a player must keep available in the buffer. Minimum value: \u003ccode\u003e2\u003c/code\u003e seconds. Maximum value: \u003ccode\u003e60\u003c/code\u003e seconds.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_update_period_seconds": {
                    "computed": true,
                    "description": "\u003cp\u003eMinimum amount of time (in seconds) that the player should wait before requesting updates to the manifest. Minimum value: \u003ccode\u003e2\u003c/code\u003e seconds. Maximum value: \u003ccode\u003e60\u003c/code\u003e seconds.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "suggested_presentation_delay_seconds": {
                    "computed": true,
                    "description": "\u003cp\u003eAmount of time (in seconds) that the player should be from the live point at the end of the manifest. Minimum value: \u003ccode\u003e2\u003c/code\u003e seconds. Maximum value: \u003ccode\u003e60\u003c/code\u003e seconds.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "hls_playlist_settings": {
              "computed": true,
              "description": "\u003cp\u003eHLS playlist configuration parameters.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ad_markup_type": {
                    "computed": true,
                    "description": "\u003cp\u003eDetermines the type of SCTE 35 tags to use in ad markup. Specify \u003ccode\u003eDATERANGE\u003c/code\u003e to use \u003ccode\u003eDATERANGE\u003c/code\u003e tags (for live or VOD content). Specify \u003ccode\u003eSCTE35_ENHANCED\u003c/code\u003e to use \u003ccode\u003eEXT-X-CUE-OUT\u003c/code\u003e and \u003ccode\u003eEXT-X-CUE-IN\u003c/code\u003e tags (for VOD content only).\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "manifest_window_seconds": {
                    "computed": true,
                    "description": "\u003cp\u003eThe total duration (in seconds) of each manifest. Minimum value: \u003ccode\u003e30\u003c/code\u003e seconds. Maximum value: \u003ccode\u003e3600\u003c/code\u003e seconds.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "manifest_name": {
              "description": "\u003cp\u003eThe name of the manifest for the channel. The name appears in the \u003ccode\u003ePlaybackUrl\u003c/code\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_group": {
              "description": "\u003cp\u003eA string used to match which \u003ccode\u003eHttpPackageConfiguration\u003c/code\u003e is used for each \u003ccode\u003eVodSource\u003c/code\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "playback_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to assign to the channel.",
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
          "nesting_mode": "set"
        },
        "optional": true
      },
      "tier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_shift_configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe configuration for time-shifted viewing.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_time_delay_seconds": {
              "computed": true,
              "description": "\u003cp\u003eThe maximum time delay for time-shifted viewing. The minimum allowed maximum time delay is 0 seconds, and the maximum allowed maximum time delay is 21600 seconds (6 hours).\u003c/p\u003e",
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
    "description": "Definition of AWS::MediaTailor::Channel Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediatailorChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediatailorChannel), &result)
	return &result
}
