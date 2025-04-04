package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2RouteServerAssociation = `{
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
      "vpc_id": {
        "description": "VPC ID",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "VPC Route Server Association",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2RouteServerAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2RouteServerAssociation), &result)
	return &result
}
