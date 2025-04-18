package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcCidrBlock = `{
  "block": {
    "attributes": {
      "amazon_provided_ipv_6_cidr_block": {
        "computed": true,
        "description": "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IPv6 addresses, or the size of the CIDR block.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cidr_block": {
        "computed": true,
        "description": "An IPv4 CIDR block to associate with the VPC.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_source": {
        "computed": true,
        "description": "The IP Source of an IPv6 VPC CIDR Block.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_4_ipam_pool_id": {
        "computed": true,
        "description": "The ID of the IPv4 IPAM pool to Associate a CIDR from to a VPC.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_4_netmask_length": {
        "computed": true,
        "description": "The netmask length of the IPv4 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ipv_6_address_attribute": {
        "computed": true,
        "description": "The value denoting whether an IPv6 VPC CIDR Block is public or private.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_6_cidr_block": {
        "computed": true,
        "description": "An IPv6 CIDR block from the IPv6 address pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_6_cidr_block_network_border_group": {
        "computed": true,
        "description": "The name of the location from which we advertise the IPV6 CIDR block.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_6_ipam_pool_id": {
        "computed": true,
        "description": "The ID of the IPv6 IPAM pool to Associate a CIDR from to a VPC.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_6_netmask_length": {
        "computed": true,
        "description": "The netmask length of the IPv6 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ipv_6_pool": {
        "computed": true,
        "description": "The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_cidr_block_id": {
        "computed": true,
        "description": "The Id of the VPC associated CIDR Block.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "description": "The ID of the VPC.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::VPCCidrBlock",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcCidrBlockSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcCidrBlock), &result)
	return &result
}
