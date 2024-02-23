package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2RouteTable = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "route_table_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the route table.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The tag key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The tag value.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
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
    "description": "Specifies a route table for the specified VPC. After you create a route table, you can add routes and associate the table with a subnet.\n For more information, see [Route tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the *Amazon VPC User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2RouteTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2RouteTable), &result)
	return &result
}
