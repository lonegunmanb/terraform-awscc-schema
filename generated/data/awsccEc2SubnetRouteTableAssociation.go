package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SubnetRouteTableAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_table_id": {
        "computed": true,
        "description": "The ID of the route table.\n The physical ID changes when the route table ID is changed.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_route_table_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::SubnetRouteTableAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2SubnetRouteTableAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SubnetRouteTableAssociation), &result)
	return &result
}
