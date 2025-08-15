package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "public_ipv_4_pool": {
        "description": "The ID of the public IPv4 pool.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_table_id": {
        "description": "The ID of the route table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::IpPoolRouteTableAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2IpPoolRouteTableAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpPoolRouteTableAssociation), &result)
	return &result
}
