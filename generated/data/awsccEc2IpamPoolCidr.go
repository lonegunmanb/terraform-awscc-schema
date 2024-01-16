package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpamPoolCidr = `{
  "block": {
    "attributes": {
      "cidr": {
        "computed": true,
        "description": "Represents a single IPv4 or IPv6 CIDR",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_pool_cidr_id": {
        "computed": true,
        "description": "Id of the IPAM Pool Cidr.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_pool_id": {
        "computed": true,
        "description": "Id of the IPAM Pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "netmask_length": {
        "computed": true,
        "description": "The desired netmask length of the provision. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.",
        "description_kind": "plain",
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "Provisioned state of the cidr.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::IPAMPoolCidr",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2IpamPoolCidrSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpamPoolCidr), &result)
	return &result
}
