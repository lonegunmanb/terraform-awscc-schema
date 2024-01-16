package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2CustomerGateway = `{
  "block": {
    "attributes": {
      "bgp_asn": {
        "computed": true,
        "description": "For devices that support BGP, the customer gateway's BGP ASN.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "customer_gateway_id": {
        "computed": true,
        "description": "CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "device_name": {
        "computed": true,
        "description": "A name for the customer gateway device.",
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
      "ip_address": {
        "description": "The internet-routable IP address for the customer gateway's outside interface. The address must be static.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags for the customer gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "type": {
        "description": "The type of VPN connection that this customer gateway supports.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::CustomerGateway",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2CustomerGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2CustomerGateway), &result)
	return &result
}
