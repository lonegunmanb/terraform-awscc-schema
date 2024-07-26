package data

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
      "channel_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "channel_name": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
      "input_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description": "\u003cp\u003eThe date and time the channel was modified.\u003c/p\u003e",
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
    "description": "Data Source schema for AWS::MediaPackageV2::Channel",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediapackagev2ChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackagev2Channel), &result)
	return &result
}
