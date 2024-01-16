package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NetworkInterface = `{
  "block": {
    "attributes": {
      "connection_tracking_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "tcp_established_timeout": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "udp_stream_timeout": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "udp_timeout": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A description for the network interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_primary_ipv_6": {
        "computed": true,
        "description": "If you have instances or ENIs that rely on the IPv6 address not changing, to avoid disrupting traffic to instances or ENIs, you can enable a primary IPv6 address. Enable this option to automatically assign an IPv6 associated with the ENI attached to your instance to be the primary IPv6 address. When you enable an IPv6 address to be a primary IPv6, you cannot disable it. Traffic will be routed to the primary IPv6 address until the instance is terminated or the ENI is detached. If you have multiple IPv6 addresses associated with an ENI and you enable a primary IPv6 address, the first IPv6 address associated with the ENI becomes the primary IPv6 address.",
        "description_kind": "plain",
        "type": "bool"
      },
      "group_set": {
        "computed": true,
        "description": "A list of security group IDs associated with this network interface.",
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
      "interface_type": {
        "computed": true,
        "description": "Indicates the type of network interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_4_prefix_count": {
        "computed": true,
        "description": "The number of IPv4 prefixes to assign to a network interface. When you specify a number of IPv4 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /28 prefixes. You can't specify a count of IPv4 prefixes if you've specified one of the following: specific IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.",
        "description_kind": "plain",
        "type": "number"
      },
      "ipv_4_prefixes": {
        "computed": true,
        "description": "Assigns a list of IPv4 prefixes to the network interface. If you want EC2 to automatically assign IPv4 prefixes, use the Ipv4PrefixCount property and do not specify this property. Presently, only /28 prefixes are supported. You can't specify IPv4 prefixes if you've specified one of the following: a count of IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ipv_4_prefix": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "ipv_6_address_count": {
        "computed": true,
        "description": "The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.",
        "description_kind": "plain",
        "type": "number"
      },
      "ipv_6_addresses": {
        "computed": true,
        "description": "One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ipv_6_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "ipv_6_prefix_count": {
        "computed": true,
        "description": "The number of IPv6 prefixes to assign to a network interface. When you specify a number of IPv6 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /80 prefixes. You can't specify a count of IPv6 prefixes if you've specified one of the following: specific IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.",
        "description_kind": "plain",
        "type": "number"
      },
      "ipv_6_prefixes": {
        "computed": true,
        "description": "Assigns a list of IPv6 prefixes to the network interface. If you want EC2 to automatically assign IPv6 prefixes, use the Ipv6PrefixCount property and do not specify this property. Presently, only /80 prefixes are supported. You can't specify IPv6 prefixes if you've specified one of the following: a count of IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ipv_6_prefix": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "primary_ipv_6_address": {
        "computed": true,
        "description": "The primary IPv6 address",
        "description_kind": "plain",
        "type": "string"
      },
      "primary_private_ip_address": {
        "computed": true,
        "description": "Returns the primary private IP address of the network interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description": "Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property. ",
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_addresses": {
        "computed": true,
        "description": "Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "primary": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "private_ip_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "secondary_private_ip_address_count": {
        "computed": true,
        "description": "The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses",
        "description_kind": "plain",
        "type": "number"
      },
      "secondary_private_ip_addresses": {
        "computed": true,
        "description": "Returns the secondary private IP addresses of the network interface.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "source_dest_check": {
        "computed": true,
        "description": "Indicates whether traffic to or from the instance is validated.",
        "description_kind": "plain",
        "type": "bool"
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet to associate with the network interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this network interface.",
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
    "description": "Data Source schema for AWS::EC2::NetworkInterface",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2NetworkInterfaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NetworkInterface), &result)
	return &result
}
