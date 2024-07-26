package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NetworkInterfaceAttachment = `{
  "block": {
    "attributes": {
      "attachment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delete_on_termination": {
        "computed": true,
        "description": "Whether to delete the network interface when the instance terminates. By default, this value is set to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "bool"
      },
      "device_index": {
        "computed": true,
        "description": "The network interface's position in the attachment order. For example, the first attached network interface has a ` + "`" + `` + "`" + `DeviceIndex` + "`" + `` + "`" + ` of 0.",
        "description_kind": "plain",
        "type": "string"
      },
      "ena_srd_specification": {
        "computed": true,
        "description": "Configures ENA Express for the network interface that this action attaches to the instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ena_srd_enabled": {
              "computed": true,
              "description": "Indicates whether ENA Express is enabled for the network interface.",
              "description_kind": "plain",
              "type": "bool"
            },
            "ena_srd_udp_specification": {
              "computed": true,
              "description": "Configures ENA Express for UDP network traffic.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ena_srd_udp_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The ID of the instance to which you will attach the ENI.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description": "The ID of the ENI that you want to attach.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::NetworkInterfaceAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2NetworkInterfaceAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NetworkInterfaceAttachment), &result)
	return &result
}
