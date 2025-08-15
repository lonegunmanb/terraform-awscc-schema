package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpPoolRouteTableAssociation = `{
  "block": {
    "attributes": {
      "association_id": {
        "computed": true,
        "description": "The route table association ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "public_ipv_4_pool": {
        "computed": true,
        "description": "The ID of the public IPv4 pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "route_table_id": {
        "computed": true,
        "description": "The ID of the route table.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::IpPoolRouteTableAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2IpPoolRouteTableAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpPoolRouteTableAssociation), &result)
	return &result
}
