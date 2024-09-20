package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayRouteTable = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the associated tag key-value pair",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the associated tag key-value pair",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "transit_gateway_id": {
        "description": "The ID of the transit gateway.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_route_table_id": {
        "computed": true,
        "description": "Transit Gateway Route Table primary identifier",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::TransitGatewayRouteTable",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayRouteTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayRouteTable), &result)
	return &result
}
