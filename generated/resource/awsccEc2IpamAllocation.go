package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpamAllocation = `{
  "block": {
    "attributes": {
      "cidr": {
        "computed": true,
        "description": "Represents an IPAM custom allocation of a single IPv4 or IPv6 CIDR",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
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
      "ipam_pool_allocation_id": {
        "computed": true,
        "description": "Id of the allocation.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_pool_id": {
        "description": "Id of the IPAM Pool.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "netmask_length": {
        "computed": true,
        "description": "The desired netmask length of the allocation. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource Schema of AWS::EC2::IPAMAllocation Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2IpamAllocationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpamAllocation), &result)
	return &result
}
