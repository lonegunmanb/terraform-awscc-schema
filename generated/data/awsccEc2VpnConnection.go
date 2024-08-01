package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpnConnection = `{
  "block": {
    "attributes": {
      "customer_gateway_id": {
        "computed": true,
        "description": "The ID of the customer gateway at your end of the VPN connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_acceleration": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "static_routes_only": {
        "computed": true,
        "description": "Indicates whether the VPN connection uses static routes only. Static routes must be used for devices that don't support BGP.\n If you are creating a VPN connection for a device that does not support Border Gateway Protocol (BGP), you must specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the VPN connection.",
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
      "transit_gateway_id": {
        "computed": true,
        "description": "The ID of the transit gateway associated with the VPN connection.\n You must specify either ` + "`" + `` + "`" + `TransitGatewayId` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `VpnGatewayId` + "`" + `` + "`" + `, but not both.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of VPN connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_connection_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_gateway_id": {
        "computed": true,
        "description": "The ID of the virtual private gateway at the AWS side of the VPN connection.\n You must specify either ` + "`" + `` + "`" + `TransitGatewayId` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `VpnGatewayId` + "`" + `` + "`" + `, but not both.",
        "description_kind": "plain",
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
              "description": "The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.\n Constraints: Allowed characters are alphanumeric characters, periods (.), and underscores (_). Must be between 8 and 64 characters in length and cannot start with zero (0).",
              "description_kind": "plain",
              "type": "string"
            },
            "tunnel_inside_cidr": {
              "computed": true,
              "description": "The range of inside IP addresses for the tunnel. Any specified CIDR blocks must be unique across all VPN connections that use the same virtual private gateway. \n Constraints: A size /30 CIDR block from the ` + "`" + `` + "`" + `169.254.0.0/16` + "`" + `` + "`" + ` range. The following CIDR blocks are reserved and cannot be used:\n  +   ` + "`" + `` + "`" + `169.254.0.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.1.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.2.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.3.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.4.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.5.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.169.252/30` + "`" + `` + "`" + `",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::EC2::VPNConnection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpnConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpnConnection), &result)
	return &result
}
