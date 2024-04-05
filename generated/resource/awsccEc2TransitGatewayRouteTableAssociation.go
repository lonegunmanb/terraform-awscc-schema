package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayRouteTableAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_attachment_id": {
        "description": "The ID of transit gateway attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_route_table_id": {
        "description": "The ID of transit gateway route table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::TransitGatewayRouteTableAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayRouteTableAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayRouteTableAssociation), &result)
	return &result
}
