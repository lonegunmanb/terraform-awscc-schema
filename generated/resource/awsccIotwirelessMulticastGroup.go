package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessMulticastGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Multicast group arn. Returned after successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "associate_wireless_device": {
        "computed": true,
        "description": "Wireless device to associate. Only for update request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Multicast group description",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disassociate_wireless_device": {
        "computed": true,
        "description": "Wireless device to disassociate. Only for update request.",
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
      "lo_ra_wan": {
        "description": "Multicast group LoRaWAN",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dl_class": {
              "description": "Multicast group LoRaWAN DL Class",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "number_of_devices_in_group": {
              "computed": true,
              "description": "Multicast group number of devices in group. Returned after successful read.",
              "description_kind": "plain",
              "type": "number"
            },
            "number_of_devices_requested": {
              "computed": true,
              "description": "Multicast group number of devices requested. Returned after successful read.",
              "description_kind": "plain",
              "type": "number"
            },
            "rf_region": {
              "description": "Multicast group LoRaWAN RF region",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "multicast_group_id": {
        "computed": true,
        "description": "Multicast group id. Returned after successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of Multicast group",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Multicast group status. Returned after successful read.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the Multicast group.",
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Create and manage Multicast groups.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotwirelessMulticastGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessMulticastGroup), &result)
	return &result
}
