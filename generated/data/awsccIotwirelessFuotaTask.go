package data

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
        "type": "string"
      },
      "associate_wireless_device": {
        "computed": true,
        "description": "Wireless device to associate. Only for update request.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "FUOTA task description",
        "description_kind": "plain",
        "type": "string"
      },
      "disassociate_multicast_group": {
        "computed": true,
        "description": "Multicast group to disassociate. Only for update request.",
        "description_kind": "plain",
        "type": "string"
      },
      "disassociate_wireless_device": {
        "computed": true,
        "description": "Wireless device to disassociate. Only for update request.",
        "description_kind": "plain",
        "type": "string"
      },
      "firmware_update_image": {
        "computed": true,
        "description": "FUOTA task firmware update image binary S3 link",
        "description_kind": "plain",
        "type": "string"
      },
      "firmware_update_role": {
        "computed": true,
        "description": "FUOTA task firmware IAM role for reading S3",
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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lo_ra_wan": {
        "computed": true,
        "description": "FUOTA task LoRaWAN",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rf_region": {
              "computed": true,
              "description": "FUOTA task LoRaWAN RF region",
              "description_kind": "plain",
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
        }
      },
      "name": {
        "computed": true,
        "description": "Name of FUOTA task",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::IoTWireless::FuotaTask",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotwirelessFuotaTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessFuotaTask), &result)
	return &result
}
