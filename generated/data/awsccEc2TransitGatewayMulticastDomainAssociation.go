package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayMulticastDomainAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
      "state": {
        "computed": true,
        "description": "The state of the subnet association.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description": "The IDs of the subnets to associate with the transit gateway multicast domain.",
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
    "description": "Data Source schema for AWS::EC2::TransitGatewayMulticastDomainAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TransitGatewayMulticastDomainAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayMulticastDomainAssociation), &result)
	return &result
}
