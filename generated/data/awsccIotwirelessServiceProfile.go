package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessServiceProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Service profile Arn. Returned after successful create.",
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
        "description": "LoRaWAN supports all LoRa specific attributes for service profile for CreateServiceProfile operation",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "add_gw_metadata": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "channel_mask": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "dev_status_req_freq": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "dl_bucket_size": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "dl_rate": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "dl_rate_policy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "dr_max": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "dr_min": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "hr_allowed": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "min_gw_diversity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "nwk_geo_loc": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "pr_allowed": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "ra_allowed": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "report_dev_status_battery": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "report_dev_status_margin": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "target_per": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "ul_bucket_size": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "ul_rate": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "ul_rate_policy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
        "description": "A list of key-value pairs that contain metadata for the service profile.",
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
    "description": "Data Source schema for AWS::IoTWireless::ServiceProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotwirelessServiceProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessServiceProfile), &result)
	return &result
}
