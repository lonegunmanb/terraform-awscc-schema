package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNimblestudioStreamingImage = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "\u003cp\u003eA human-readable description of the streaming image.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "ec_2_image_id": {
        "computed": true,
        "description": "\u003cp\u003eThe ID of an EC2 machine image with which to create this streaming image.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "\u003cp\u003eTODO\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key_arn": {
              "computed": true,
              "description": "\u003cp\u003eThe ARN for a KMS key that is used to encrypt studio data.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "key_type": {
              "computed": true,
              "description": "\u003cp/\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "eula_ids": {
        "computed": true,
        "description": "\u003cp\u003eThe list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.\u003c/p\u003e",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "\u003cp\u003eA friendly name for a streaming image resource.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description": "\u003cp\u003eThe owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "platform": {
        "computed": true,
        "description": "\u003cp\u003eThe platform of the streaming image, either WINDOWS or LINUX.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "streaming_image_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "studio_id": {
        "computed": true,
        "description": "\u003cp\u003eThe studioId. \u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::NimbleStudio::StreamingImage",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNimblestudioStreamingImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNimblestudioStreamingImage), &result)
	return &result
}
