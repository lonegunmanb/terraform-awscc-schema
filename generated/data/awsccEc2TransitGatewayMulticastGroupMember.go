package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayMulticastGroupMember = `{
  "block": {
    "attributes": {
      "group_ip_address": {
        "computed": true,
        "description": "The IP address assigned to the transit gateway multicast group.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_member": {
        "computed": true,
        "description": "Indicates that the resource is a transit gateway multicast group member.",
        "description_kind": "plain",
        "type": "bool"
      },
      "group_source": {
        "computed": true,
        "description": "Indicates that the resource is a transit gateway multicast group member.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "member_type": {
        "computed": true,
        "description": "The member type (for example, static).",
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description": "The ID of the transit gateway attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "The ID of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The type of resource, for example a VPC attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_attachment_id": {
        "computed": true,
        "description": "The ID of the transit gateway attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_multicast_domain_id": {
        "computed": true,
        "description": "The ID of the transit gateway multicast domain.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::TransitGatewayMulticastGroupMember",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TransitGatewayMulticastGroupMemberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayMulticastGroupMember), &result)
	return &result
}
