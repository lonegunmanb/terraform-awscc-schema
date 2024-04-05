package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2EipAssociation = `{
  "block": {
    "attributes": {
      "allocation_id": {
        "computed": true,
        "description": "The allocation ID. This is required for EC2-VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "eip": {
        "computed": true,
        "description": "The Elastic IP address to associate with the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "eip_association_id": {
        "computed": true,
        "description": "Composite ID of non-empty properties, to determine the identification.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The ID of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description": "The ID of the network interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description": "The primary or secondary private IP address to associate with the Elastic IP address.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::EIPAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2EipAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2EipAssociation), &result)
	return &result
}
