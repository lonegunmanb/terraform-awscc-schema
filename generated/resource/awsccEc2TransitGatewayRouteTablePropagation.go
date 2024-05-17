package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayRouteTablePropagation = `{
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
    "description": "AWS::EC2::TransitGatewayRouteTablePropagation Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayRouteTablePropagationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayRouteTablePropagation), &result)
	return &result
}
