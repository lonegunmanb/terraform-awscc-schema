package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2DhcpOptions = `{
  "block": {
    "attributes": {
      "dhcp_options_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "This value is used to complete unqualified DNS hostnames.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name_servers": {
        "computed": true,
        "description": "The IPv4 addresses of up to four domain name servers, or AmazonProvidedDNS.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipv_6_address_preferred_lease_time": {
        "computed": true,
        "description": "The preferred Lease Time for ipV6 address in seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "netbios_name_servers": {
        "computed": true,
        "description": "The IPv4 addresses of up to four NetBIOS name servers.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "netbios_node_type": {
        "computed": true,
        "description": "The NetBIOS node type (1, 2, 4, or 8).",
        "description_kind": "plain",
        "type": "number"
      },
      "ntp_servers": {
        "computed": true,
        "description": "The IPv4 addresses of up to four Network Time Protocol (NTP) servers.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the DHCP options set.",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::EC2::DHCPOptions",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2DhcpOptionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2DhcpOptions), &result)
	return &result
}
