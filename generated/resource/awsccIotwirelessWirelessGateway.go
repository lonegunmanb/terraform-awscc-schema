package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_uplink_received_at": {
        "computed": true,
        "description": "The date and time when the most recent uplink was received.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lo_ra_wan": {
        "description": "The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "gateway_eui": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rf_region": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "name": {
        "computed": true,
        "description": "Name of Wireless Gateway.",
        "description_kind": "plain",
        "optional": true,
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
      },
      "thing_arn": {
        "computed": true,
        "description": "Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "thing_name": {
        "computed": true,
        "description": "Thing Name. If there is a Thing created, this can be returned with a Get call.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "wireless_gateway_id": {
        "computed": true,
        "description": "Id for Wireless Gateway. Returned upon successful create.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Create and manage wireless gateways, including LoRa gateways.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotwirelessWirelessGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessWirelessGateway), &result)
	return &result
}
