package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcdhcpOptionsAssociation = `{
  "block": {
    "attributes": {
      "dhcp_options_id": {
        "description": "The ID of the DHCP options set, or default to associate no DHCP options with the VPC.",
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
      "vpc_id": {
        "description": "The ID of the VPC.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Associates a set of DHCP options with a VPC, or associates no DHCP options with the VPC.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcdhcpOptionsAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcdhcpOptionsAssociation), &result)
	return &result
}
