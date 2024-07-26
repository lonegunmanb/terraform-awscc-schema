package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlobalacceleratorCrossAccountAttachment = `{
  "block": {
    "attributes": {
      "attachment_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The Friendly identifier of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "principals": {
        "computed": true,
        "description": "Principals to share the resources with.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "resources": {
        "computed": true,
        "description": "Resources shared using the attachment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "endpoint_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Key of the tag. Value can be 1 to 127 characters.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Value for the tag. Value can be 1 to 255 characters.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::GlobalAccelerator::CrossAccountAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlobalacceleratorCrossAccountAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlobalacceleratorCrossAccountAttachment), &result)
	return &result
}
