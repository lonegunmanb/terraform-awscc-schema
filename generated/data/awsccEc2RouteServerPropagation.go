package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2RouteServerPropagation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_server_id": {
        "computed": true,
        "description": "Route Server ID",
        "description_kind": "plain",
        "type": "string"
      },
      "route_table_id": {
        "computed": true,
        "description": "Route Table ID",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::RouteServerPropagation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2RouteServerPropagationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2RouteServerPropagation), &result)
	return &result
}
