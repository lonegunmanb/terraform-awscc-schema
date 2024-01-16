package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpnConnection = `{
  "block": {
    "attributes": {
      "customer_gateway_id": {
        "description": "The ID of the customer gateway at your end of the VPN connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "static_routes_only": {
        "computed": true,
        "description": "Indicates whether the VPN connection uses static routes only.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the VPN connection.",
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
      "transit_gateway_id": {
        "computed": true,
        "description": "The ID of the transit gateway associated with the VPN connection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The type of VPN connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpn_connection_id": {
        "computed": true,
        "description": "The provider-assigned unique ID for this managed resource",
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_gateway_id": {
        "computed": true,
        "description": "The ID of the virtual private gateway at the AWS side of the VPN connection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpn_tunnel_options_specifications": {
        "computed": true,
        "description": "The tunnel options for the VPN connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "pre_shared_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tunnel_inside_cidr": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::EC2::VPNConnection",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpnConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpnConnection), &result)
	return &result
}
