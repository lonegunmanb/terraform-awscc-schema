package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessWirelessGateway = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Arn for Wireless Gateway. Returned upon successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of Wireless Gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_uplink_received_at": {
        "computed": true,
        "description": "The date and time when the most recent uplink was received.",
        "description_kind": "plain",
        "type": "string"
      },
      "lo_ra_wan": {
        "computed": true,
        "description": "The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "gateway_eui": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rf_region": {
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
        "description": "Name of Wireless Gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the gateway.",
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
      },
      "thing_arn": {
        "computed": true,
        "description": "Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "thing_name": {
        "computed": true,
        "description": "Thing Name. If there is a Thing created, this can be returned with a Get call.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoTWireless::WirelessGateway",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotwirelessWirelessGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessWirelessGateway), &result)
	return &result
}
