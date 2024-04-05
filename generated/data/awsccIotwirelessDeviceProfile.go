package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessDeviceProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Service profile Arn. Returned after successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "device_profile_id": {
        "computed": true,
        "description": "Service profile Id. Returned after successful create.",
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
        "description": "LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "class_b_timeout": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "class_c_timeout": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "factory_preset_freqs_list": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "number"
              ]
            },
            "mac_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "max_duty_cycle": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "max_eirp": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "ping_slot_dr": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "ping_slot_freq": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "ping_slot_period": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "reg_params_revision": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rf_region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rx_data_rate_2": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "rx_delay_1": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "rx_dr_offset_1": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "rx_freq_2": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "supports_32_bit_f_cnt": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "supports_class_b": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "supports_class_c": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "supports_join": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "Name of service profile",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the device profile.",
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
    "description": "Data Source schema for AWS::IoTWireless::DeviceProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotwirelessDeviceProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessDeviceProfile), &result)
	return &result
}
