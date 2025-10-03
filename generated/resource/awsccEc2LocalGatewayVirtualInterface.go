package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2LocalGatewayVirtualInterface = `{
  "block": {
    "attributes": {
      "configuration_state": {
        "computed": true,
        "description": "The current state of the local gateway virtual interface",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_address": {
        "description": "The local address.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_bgp_asn": {
        "computed": true,
        "description": "The Autonomous System Number(ASN) for the local Border Gateway Protocol (BGP)",
        "description_kind": "plain",
        "type": "number"
      },
      "local_gateway_id": {
        "computed": true,
        "description": "The ID of the local gateway",
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_virtual_interface_group_id": {
        "description": "The ID of the virtual interface group",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_gateway_virtual_interface_id": {
        "computed": true,
        "description": "The ID of the virtual interface",
        "description_kind": "plain",
        "type": "string"
      },
      "outpost_lag_id": {
        "description": "The Outpost LAG ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "The ID of the Amazon Web Services account that owns the local gateway virtual interface group",
        "description_kind": "plain",
        "type": "string"
      },
      "peer_address": {
        "description": "The peer address.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_bgp_asn": {
        "computed": true,
        "description": "The peer BGP ASN.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "peer_bgp_asn_extended": {
        "computed": true,
        "description": "The extended 32-bit ASN of the BGP peer for use with larger ASN values.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "vlan": {
        "description": "The ID of the VLAN.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for Local Gateway Virtual Interface which describes a virtual interface for AWS Outposts local gateways.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2LocalGatewayVirtualInterfaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2LocalGatewayVirtualInterface), &result)
	return &result
}
