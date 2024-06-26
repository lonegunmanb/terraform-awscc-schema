package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpnGateway = `{
  "block": {
    "attributes": {
      "amazon_side_asn": {
        "computed": true,
        "description": "The private Autonomous System Number (ASN) for the Amazon side of a BGP session.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the virtual private gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of VPN connection the virtual private gateway supports.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::VPNGateway",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpnGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpnGateway), &result)
	return &result
}
