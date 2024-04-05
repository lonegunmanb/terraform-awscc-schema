package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SubnetCidrBlock = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipv_6_cidr_block": {
        "computed": true,
        "description": "The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_6_ipam_pool_id": {
        "computed": true,
        "description": "The ID of an IPv6 Amazon VPC IP Address Manager (IPAM) pool from which to allocate, to get the subnet's CIDR",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_6_netmask_length": {
        "computed": true,
        "description": "The netmask length of the IPv6 CIDR to allocate to the subnet from an IPAM pool",
        "description_kind": "plain",
        "type": "number"
      },
      "subnet_cidr_block_id": {
        "computed": true,
        "description": "Information about the IPv6 association.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::SubnetCidrBlock",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2SubnetCidrBlockSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SubnetCidrBlock), &result)
	return &result
}
