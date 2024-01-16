package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2LocalGatewayRoute = `{
  "block": {
    "attributes": {
      "destination_cidr_block": {
        "computed": true,
        "description": "The CIDR block used for destination matches.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_gateway_route_table_id": {
        "computed": true,
        "description": "The ID of the local gateway route table.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_virtual_interface_group_id": {
        "computed": true,
        "description": "The ID of the virtual interface group.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description": "The ID of the network interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the route.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The route type.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::LocalGatewayRoute",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2LocalGatewayRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2LocalGatewayRoute), &result)
	return &result
}
