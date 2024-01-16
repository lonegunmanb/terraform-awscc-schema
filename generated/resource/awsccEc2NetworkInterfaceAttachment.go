package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NetworkInterfaceAttachment = `{
  "block": {
    "attributes": {
      "attachment_id": {
        "computed": true,
        "description": "The ID of the network interface attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_on_termination": {
        "computed": true,
        "description": "Whether to delete the network interface when the instance terminates. By default, this value is set to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "device_index": {
        "description": "The network interface's position in the attachment order. For example, the first attached network interface has a DeviceIndex of 0.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "The ID of the instance to which you will attach the ENI.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_interface_id": {
        "description": "The ID of the ENI that you want to attach.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::NetworkInterfaceAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2NetworkInterfaceAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NetworkInterfaceAttachment), &result)
	return &result
}
