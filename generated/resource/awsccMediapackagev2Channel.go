package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackagev2Channel = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) associated with the resource.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "attached_multiview_channels": {
        "computed": true,
        "description": "\u003cp\u003eThe multiview channels, in the same channel group, that list this channel as an available source. This is a read-only field. You can't delete a channel while any multiview channel still lists it as a source. Use this field to find the multiview channels that you need to update first.\u003c/p\u003e",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "channel_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "channel_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "\u003cp\u003eThe date and time the channel was created.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "\u003cp\u003eEnter any descriptive text that helps you to identify the channel.\u003c/p\u003e",
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
      "ingest_endpoint_urls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "ingest_endpoints": {
        "computed": true,
        "description": "\u003cp\u003eThe list of ingest endpoints.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "\u003cp\u003eThe system-generated unique identifier for the IngestEndpoint.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
              "computed": true,
              "description": "\u003cp\u003eThe ingest domain URL where the source stream should be sent.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "input_switch_configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mqcs_input_switching": {
              "computed": true,
              "description": "\u003cp\u003eWhen true, AWS Elemental MediaPackage performs input switching based on the MQCS. Default is false. This setting is valid only when \u003ccode\u003eInputType\u003c/code\u003e is \u003ccode\u003eCMAF\u003c/code\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "preferred_input": {
              "computed": true,
              "description": "\u003cp\u003eFor CMAF inputs, indicates which input MediaPackage should prefer when both inputs have equal MQCS scores. Select \u003ccode\u003e1\u003c/code\u003e to prefer the first ingest endpoint, or \u003ccode\u003e2\u003c/code\u003e to prefer the second ingest endpoint. If you don't specify a preferred input, MediaPackage uses its default switching behavior when MQCS scores are equal.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "input_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description": "\u003cp\u003eThe date and time the channel was modified.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "multiview_configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe multiview configuration for a channel. A multiview channel composites video from several source channels into a single tiled output stream. Players receive one standard HLS or DASH stream instead of several separate streams. This setting is required when \u003ccode\u003eInputType\u003c/code\u003e is \u003ccode\u003eMULTIVIEW\u003c/code\u003e, and can't be set for any other input type.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "available_layouts": {
              "computed": true,
              "description": "\u003cp\u003eThe tile layouts that players can request from this multiview channel's origin endpoints. Only the layouts that you list here are available. Each layout must appear at most once.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "available_sources": {
              "computed": true,
              "description": "\u003cp\u003eThe channels that players can use as tiles in this multiview channel's output. Each source channel must be in the same channel group as the multiview channel, and must have an \u003ccode\u003eInputType\u003c/code\u003e of \u003ccode\u003eCMAF\u003c/code\u003e. Only the channels that you list here are available as tiles.\u003c/p\u003e",
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
      "output_header_configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "publish_mqcs": {
              "computed": true,
              "description": "\u003cp\u003eWhen true, AWS Elemental MediaPackage includes the MQCS in responses to the CDN. This setting is valid only when \u003ccode\u003eInputType\u003c/code\u003e is \u003ccode\u003eCMAF\u003c/code\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "output_locking_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
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
    "description": "\u003cp\u003eRepresents an entry point into AWS Elemental MediaPackage for an ABR video content stream sent from an upstream encoder such as AWS Elemental MediaLive. The channel continuously analyzes the content that it receives and prepares it to be distributed to consumers via one or more origin endpoints.\u003c/p\u003e",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediapackagev2ChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackagev2Channel), &result)
	return &result
}
