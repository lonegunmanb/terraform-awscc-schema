package resource

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
      "id": {
        "computed": true,
        "description": "Service profile Id. Returned after successful create.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "number"
            },
            "class_c_timeout": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "factory_preset_freqs_list": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "mac_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_duty_cycle": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_eirp": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ping_slot_dr": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ping_slot_freq": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ping_slot_period": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "reg_params_revision": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rf_region": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rx_data_rate_2": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rx_delay_1": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rx_dr_offset_1": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rx_freq_2": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "supports_32_bit_f_cnt": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "supports_class_b": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "supports_class_c": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "supports_join": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "name": {
        "computed": true,
        "description": "Name of service profile",
        "description_kind": "plain",
        "optional": true,
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
    "description": "Device Profile's resource schema demonstrating some basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotwirelessDeviceProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessDeviceProfile), &result)
	return &result
}
