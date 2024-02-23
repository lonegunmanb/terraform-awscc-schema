package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackagev2ChannelGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) associated with the resource.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "\u003cp\u003eThe date and time the channel group was created.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "\u003cp\u003eEnter any descriptive text that helps you to identify the channel group.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "egress_domain": {
        "computed": true,
        "description": "\u003cp\u003eThe output domain where the source stream should be sent. Integrate the domain with a downstream CDN (such as Amazon CloudFront) or playback device.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description": "\u003cp\u003eThe date and time the channel group was modified.\u003c/p\u003e",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::MediaPackageV2::ChannelGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediapackagev2ChannelGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackagev2ChannelGroup), &result)
	return &result
}
