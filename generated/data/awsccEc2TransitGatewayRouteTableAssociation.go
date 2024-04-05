package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayRouteTableAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_attachment_id": {
        "computed": true,
        "description": "The ID of transit gateway attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_route_table_id": {
        "computed": true,
        "description": "The ID of transit gateway route table.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::TransitGatewayRouteTableAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TransitGatewayRouteTableAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayRouteTableAssociation), &result)
	return &result
}
