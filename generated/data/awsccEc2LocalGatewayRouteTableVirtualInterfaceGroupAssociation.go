package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2LocalGatewayRouteTableVirtualInterfaceGroupAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_gateway_id": {
        "computed": true,
        "description": "The ID of the local gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_route_table_arn": {
        "computed": true,
        "description": "The ARN of the local gateway route table.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_route_table_id": {
        "computed": true,
        "description": "The ID of the local gateway route table.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_route_table_virtual_interface_group_association_id": {
        "computed": true,
        "description": "The ID of the local gateway route table virtual interface group association.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_virtual_interface_group_id": {
        "computed": true,
        "description": "The ID of the local gateway route table virtual interface group.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "The owner of the local gateway route table virtual interface group association.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the local gateway route table virtual interface group association.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the local gateway route table virtual interface group association.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2LocalGatewayRouteTableVirtualInterfaceGroupAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2LocalGatewayRouteTableVirtualInterfaceGroupAssociation), &result)
	return &result
}
