package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The Friendly identifier of the attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principals": {
        "computed": true,
        "description": "Principals to share the resources with.",
        "description_kind": "plain",
        "optional": true,
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
            "endpoint_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "region": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "Key of the tag. Value can be 1 to 127 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "Value for the tag. Value can be 1 to 255 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::GlobalAccelerator::CrossAccountAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlobalacceleratorCrossAccountAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlobalacceleratorCrossAccountAttachment), &result)
	return &result
}
