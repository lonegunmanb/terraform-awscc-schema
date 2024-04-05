package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvsEncoderConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Encoder configuration identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Encoder configuration name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "video": {
        "computed": true,
        "description": "Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30 fps",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bitrate": {
              "computed": true,
              "description": "Bitrate for generated output, in bps. Default: 2500000.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "framerate": {
              "computed": true,
              "description": "Video frame rate, in fps. Default: 30.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "height": {
              "computed": true,
              "description": "Video-resolution height. Note that the maximum value is determined by width times height, such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 720.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "width": {
              "computed": true,
              "description": "Video-resolution width. Note that the maximum value is determined by width times height, such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 1280.",
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
    "description": "Resource Type definition for AWS::IVS::EncoderConfiguration.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIvsEncoderConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsEncoderConfiguration), &result)
	return &result
}
