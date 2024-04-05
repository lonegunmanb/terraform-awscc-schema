package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SubnetRouteTableAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "route_table_id": {
        "description": "The ID of the route table.\n The physical ID changes when the route table ID is changed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_id": {
        "description": "The ID of the subnet.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_route_table_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Associates a subnet with a route table. The subnet and route table must be in the same VPC. This association causes traffic originating from the subnet to be routed according to the routes in the route table. A route table can be associated with multiple subnets. To create a route table, see [AWS::EC2::RouteTable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routetable.html).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SubnetRouteTableAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SubnetRouteTableAssociation), &result)
	return &result
}
