package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpnConnectionRoute = `{
  "block": {
    "attributes": {
      "destination_cidr_block": {
        "computed": true,
        "description": "The CIDR block associated with the local subnet of the customer network.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpn_connection_id": {
        "computed": true,
        "description": "The ID of the VPN connection.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::VPNConnectionRoute",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpnConnectionRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpnConnectionRoute), &result)
	return &result
}
