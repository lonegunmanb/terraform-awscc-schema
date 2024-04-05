package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessFuotaTask = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "FUOTA task arn. Returned after successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "associate_multicast_group": {
        "computed": true,
        "description": "Multicast group to associate. Only for update request.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "FUOTA task description",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disassociate_multicast_group": {
        "computed": true,
        "description": "Multicast group to disassociate. Only for update request.",
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
      "firmware_update_image": {
        "description": "FUOTA task firmware update image binary S3 link",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firmware_update_role": {
        "description": "FUOTA task firmware IAM role for reading S3",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fuota_task_id": {
        "computed": true,
        "description": "FUOTA task id. Returned after successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "fuota_task_status": {
        "computed": true,
        "description": "FUOTA task status. Returned after successful read.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "lo_ra_wan": {
        "description": "FUOTA task LoRaWAN",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rf_region": {
              "description": "FUOTA task LoRaWAN RF region",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "start_time": {
              "computed": true,
              "description": "FUOTA task LoRaWAN start time",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "name": {
        "computed": true,
        "description": "Name of FUOTA task",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the FUOTA task.",
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
    "description": "Create and manage FUOTA tasks.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotwirelessFuotaTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessFuotaTask), &result)
	return &result
}
