package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectFlowMediaStream = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the media stream, composed of the flow ARN followed by /mediaStream/ and the media stream name.",
        "description_kind": "plain",
        "type": "string"
      },
      "attributes": {
        "computed": true,
        "description": "Attributes that are related to the media stream.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fmtp": {
              "computed": true,
              "description": "A set of parameters that define the media stream.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "channel_order": {
                    "computed": true,
                    "description": "The format of the audio channel. Can only be specified for an audio media stream.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "colorimetry": {
                    "computed": true,
                    "description": "The format used for the representation of color.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "exact_framerate": {
                    "computed": true,
                    "description": "The frame rate for the video stream, in frames/second. For example: 60000/1001.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "par": {
                    "computed": true,
                    "description": "The pixel aspect ratio (PAR) of the video.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "range": {
                    "computed": true,
                    "description": "The encoding range of the video.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scan_mode": {
                    "computed": true,
                    "description": "The type of compression that was used to smooth the video's appearance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tcs": {
                    "computed": true,
                    "description": "The transfer characteristic system (TCS) that is used in the video.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "lang": {
              "computed": true,
              "description": "The audio language, in a format that is recognized by the receiver. Can only be specified for an audio media stream.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "clock_rate": {
        "computed": true,
        "description": "The sample rate (in Hz) for the stream. If the media stream type is video or ancillary data, set this value to 90000. If the media stream type is audio, set this value to either 48000 or 96000.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "computed": true,
        "description": "A description that can help you quickly identify what your media stream is used for.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flow_arn": {
        "description": "The Amazon Resource Name (ARN) of the flow that the media stream belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fmt": {
        "computed": true,
        "description": "The format type number (sometimes referred to as RTP payload type) of the media stream. MediaConnect assigns this value to the media stream.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "media_stream_id": {
        "description": "A unique identifier for the media stream.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "media_stream_name": {
        "description": "A name that helps you distinguish one media stream from another.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "media_stream_type": {
        "description": "The type of media stream.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The key-value pairs that can be used to tag and organize the media stream.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "video_format": {
        "computed": true,
        "description": "The resolution of the video. Required for a video media stream and rejected for other media stream types.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::MediaConnect::FlowMediaStream. A media stream represents a single track or stream of media containing video, audio, or ancillary data that is transported using the SMPTE 2110 JPEG XS or CDI protocol.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectFlowMediaStreamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectFlowMediaStream), &result)
	return &result
}
