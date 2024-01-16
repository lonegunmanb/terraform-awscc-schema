package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2GatewayRouteTableAssociation = `{
  "block": {
    "attributes": {
      "association_id": {
        "computed": true,
        "description": "The route table association ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_id": {
        "description": "The ID of the gateway.",
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
      "route_table_id": {
        "description": "The ID of the route table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Associates a gateway with a route table. The gateway and route table must be in the same VPC. This association causes the incoming traffic to the gateway to be routed according to the routes in the route table.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2GatewayRouteTableAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2GatewayRouteTableAssociation), &result)
	return &result
}
