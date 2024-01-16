package resource

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
        "optional": true,
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_dns_64": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_4_ipam_pool_id": {
        "computed": true,
        "description": "The ID of an IPv4 IPAM pool you want to use for allocating this subnet's CIDR",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_4_netmask_length": {
        "computed": true,
        "description": "The netmask length of the IPv4 CIDR you want to allocate to this subnet from an Amazon VPC IP Address Manager (IPAM) pool",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ipv_6_cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_6_cidr_blocks": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ipv_6_ipam_pool_id": {
        "computed": true,
        "description": "The ID of an IPv6 IPAM pool you want to use for allocating this subnet's CIDR",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_6_native": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ipv_6_netmask_length": {
        "computed": true,
        "description": "The netmask length of the IPv6 CIDR you want to allocate to this subnet from an Amazon VPC IP Address Manager (IPAM) pool",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "map_public_ip_on_launch": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "network_acl_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "outpost_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "bool"
            },
            "enable_resource_name_dns_aaaa_record": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hostname_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::Subnet",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SubnetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Subnet), &result)
	return &result
}
