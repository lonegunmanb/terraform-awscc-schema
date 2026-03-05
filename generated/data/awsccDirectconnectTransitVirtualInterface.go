package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDirectconnectTransitVirtualInterface = `{
  "block": {
    "attributes": {
      "allocate_transit_virtual_interface_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the role to allocate the TransitVifAllocation. Needs directconnect:AllocateTransitVirtualInterface permissions and tag permissions if applicable.",
        "description_kind": "plain",
        "type": "string"
      },
      "bgp_peers": {
        "computed": true,
        "description": "The BGP peers configured on this virtual interface..",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address_family": {
              "computed": true,
              "description": "The address family for the BGP peer.",
              "description_kind": "plain",
              "type": "string"
            },
            "amazon_address": {
              "computed": true,
              "description": "The IP address assigned to the Amazon interface.",
              "description_kind": "plain",
              "type": "string"
            },
            "asn": {
              "computed": true,
              "description": "The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "auth_key": {
              "computed": true,
              "description": "The authentication key for BGP configuration. This string has a minimum length of 6 characters and and a maximum length of 80 characters.",
              "description_kind": "plain",
              "type": "string"
            },
            "bgp_peer_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "customer_address": {
              "computed": true,
              "description": "The IP address assigned to the customer interface.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "connection_id": {
        "computed": true,
        "description": "The ID or ARN of the connection or LAG.",
        "description_kind": "plain",
        "type": "string"
      },
      "direct_connect_gateway_id": {
        "computed": true,
        "description": "The ID or ARN of the Direct Connect gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_site_link": {
        "computed": true,
        "description": "Indicates whether to enable or disable SiteLink.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mtu": {
        "computed": true,
        "description": "The maximum transmission unit (MTU), in bytes. The supported values are 1500 and 9001. The default value is 1500.",
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the private virtual interface.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "virtual_interface_arn": {
        "computed": true,
        "description": "The ARN of the virtual interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "virtual_interface_id": {
        "computed": true,
        "description": "The ID of the virtual interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "virtual_interface_name": {
        "computed": true,
        "description": "The name of the virtual interface assigned by the customer network. The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-).",
        "description_kind": "plain",
        "type": "string"
      },
      "vlan": {
        "computed": true,
        "description": "The ID of the VLAN.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::DirectConnect::TransitVirtualInterface",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDirectconnectTransitVirtualInterfaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDirectconnectTransitVirtualInterface), &result)
	return &result
}
