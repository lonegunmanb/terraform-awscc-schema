package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayRoute = `{
  "block": {
    "attributes": {
      "blackhole": {
        "computed": true,
        "description": "Indicates whether to drop traffic that matches this route.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "destination_cidr_block": {
        "description": "The CIDR range used for destination matches. Routing decisions are based on the most specific match.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_attachment_id": {
        "computed": true,
        "description": "The ID of transit gateway attachment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_gateway_route_table_id": {
        "description": "The ID of transit gateway route table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::TransitGatewayRoute",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayRoute), &result)
	return &result
}
