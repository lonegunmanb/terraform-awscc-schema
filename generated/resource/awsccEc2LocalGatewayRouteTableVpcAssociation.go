package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2LocalGatewayRouteTableVpcAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_id": {
        "computed": true,
        "description": "The ID of the local gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_route_table_id": {
        "description": "The ID of the local gateway route table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_gateway_route_table_vpc_association_id": {
        "computed": true,
        "description": "The ID of the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the association.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "vpc_id": {
        "description": "The ID of the VPC.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Describes an association between a local gateway route table and a VPC.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2LocalGatewayRouteTableVpcAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2LocalGatewayRouteTableVpcAssociation), &result)
	return &result
}
