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
        "description_kind": "plain",
        "type": "string"
      },
      "route_table_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::SubnetRouteTableAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SubnetRouteTableAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SubnetRouteTableAssociation), &result)
	return &result
}
