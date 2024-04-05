package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcdhcpOptionsAssociation = `{
  "block": {
    "attributes": {
      "dhcp_options_id": {
        "computed": true,
        "description": "The ID of the DHCP options set, or default to associate no DHCP options with the VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::VPCDHCPOptionsAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpcdhcpOptionsAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcdhcpOptionsAssociation), &result)
	return &result
}
