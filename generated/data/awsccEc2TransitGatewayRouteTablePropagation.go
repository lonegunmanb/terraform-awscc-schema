package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayRouteTablePropagation = `{
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
    "description": "Data Source schema for AWS::EC2::TransitGatewayRouteTablePropagation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TransitGatewayRouteTablePropagationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayRouteTablePropagation), &result)
	return &result
}
