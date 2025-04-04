package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2RouteServerPropagation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "route_server_id": {
        "description": "Route Server ID",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_table_id": {
        "description": "Route Table ID",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "VPC Route Server Propagation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2RouteServerPropagationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2RouteServerPropagation), &result)
	return &result
}
