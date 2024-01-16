package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VolumeAttachment = `{
  "block": {
    "attributes": {
      "device": {
        "computed": true,
        "description": "The device name",
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
      "instance_id": {
        "description": "The ID of the instance to which the volume attaches",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "volume_id": {
        "description": "The ID of the Amazon EBS volume",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::VolumeAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VolumeAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VolumeAttachment), &result)
	return &result
}
