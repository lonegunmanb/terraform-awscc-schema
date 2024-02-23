package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Subnet = `{
  "block": {
    "attributes": {
      "assign_ipv_6_address_on_creation": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_dns_64": {
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
      "ipv_4_ipam_pool_id": {
        "computed": true,
        "description": "The ID of an IPv4 IPAM pool you want to use for allocating this subnet's CIDR",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_4_netmask_length": {
        "computed": true,
        "description": "The netmask length of the IPv4 CIDR you want to allocate to this subnet from an Amazon VPC IP Address Manager (IPAM) pool",
        "description_kind": "plain",
        "type": "number"
      },
      "ipv_6_cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_6_cidr_blocks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "ipv_6_ipam_pool_id": {
        "computed": true,
        "description": "The ID of an IPv6 IPAM pool you want to use for allocating this subnet's CIDR",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_6_native": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ipv_6_netmask_length": {
        "computed": true,
        "description": "The netmask length of the IPv6 CIDR you want to allocate to this subnet from an Amazon VPC IP Address Manager (IPAM) pool",
        "description_kind": "plain",
        "type": "number"
      },
      "map_public_ip_on_launch": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "network_acl_association_id": {
        "computed": true,
        "description": "The ID of the network ACL that is associated with the subnet's VPC",
        "description_kind": "plain",
        "type": "string"
      },
      "outpost_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns_name_options_on_launch": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_resource_name_dns_a_record": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "enable_resource_name_dns_aaaa_record": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "hostname_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::Subnet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2SubnetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Subnet), &result)
	return &result
}
