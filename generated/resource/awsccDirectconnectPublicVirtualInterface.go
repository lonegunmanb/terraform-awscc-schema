package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDirectconnectPublicVirtualInterface = `{
  "block": {
    "attributes": {
      "allocate_public_virtual_interface_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the role to allocate the public virtual interface. Needs directconnect:AllocatePublicVirtualInterface permissions and tag permissions if applicable.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bgp_peers": {
        "description": "The BGP peers configured on this virtual interface.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address_family": {
              "description": "The address family for the BGP peer.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "amazon_address": {
              "computed": true,
              "description": "The IP address assigned to the Amazon interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "asn": {
              "description": "The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "auth_key": {
              "computed": true,
              "description": "The authentication key for BGP configuration. This string has a minimum length of 6 characters and and a maximum length of 80 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bgp_peer_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "customer_address": {
              "computed": true,
              "description": "The IP address assigned to the customer interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "connection_id": {
        "description": "The ID or ARN of the connection or LAG.",
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
      "route_filter_prefixes": {
        "computed": true,
        "description": "The routes to be advertised to the AWS network in this region.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the public virtual interface.",
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
        "description": "The name of the virtual interface assigned by the customer network. The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vlan": {
        "description": "The ID of the VLAN.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::DirectConnect::PublicVirtualInterface",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDirectconnectPublicVirtualInterfaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDirectconnectPublicVirtualInterface), &result)
	return &result
}
