package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveMultiplex = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The unique arn of the multiplex.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zones": {
        "description": "A list of availability zones for the multiplex.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "destinations": {
        "computed": true,
        "description": "A list of the multiplex output destinations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "multiplex_media_connect_output_destination_settings": {
              "computed": true,
              "description": "Multiplex MediaConnect output destination settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "entitlement_arn": {
                    "computed": true,
                    "description": "The MediaConnect entitlement ARN available as a Flow source.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "multiplex_id": {
        "computed": true,
        "description": "The unique id of the multiplex.",
        "description_kind": "plain",
        "type": "string"
      },
      "multiplex_settings": {
        "description": "Configuration for a multiplex event.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maximum_video_buffer_delay_milliseconds": {
              "computed": true,
              "description": "Maximum video buffer delay in milliseconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "transport_stream_bitrate": {
              "description": "Transport stream bit rate.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "transport_stream_id": {
              "description": "Transport stream ID.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "transport_stream_reserved_bitrate": {
              "computed": true,
              "description": "Transport stream reserved bit rate.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "name": {
        "description": "Name of multiplex.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pipelines_running_count": {
        "computed": true,
        "description": "The number of currently healthy pipelines.",
        "description_kind": "plain",
        "type": "number"
      },
      "program_count": {
        "computed": true,
        "description": "The number of programs in the multiplex.",
        "description_kind": "plain",
        "type": "number"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of key-value pairs.",
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
      }
    },
    "description": "Resource schema for AWS::MediaLive::Multiplex",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMedialiveMultiplexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveMultiplex), &result)
	return &result
}
