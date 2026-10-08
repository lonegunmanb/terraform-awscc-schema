package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2ClientVpnRoute = `{
  "block": {
    "attributes": {
      "client_vpn_endpoint_id": {
        "computed": true,
        "description": "The ID of the Client VPN endpoint to which to add the route.",
        "description_kind": "plain",
        "type": "string"
      },
      "client_vpn_route_id": {
        "computed": true,
        "description": "The CloudFormation-generated identifier for this Client VPN route, derived from the logical resource id and client request token.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A brief description of the route.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_cidr_block": {
        "computed": true,
        "description": "The IPv4 address range, in CIDR notation, of the route destination. For example: To add a route for Internet access, enter 0.0.0.0/0. To add a route for a peered VPC, enter the peered VPC's IPv4 CIDR range. To add a route for an on-premises network, enter the AWS Site-to-Site VPN connection's IPv4 CIDR range. To add a route for the local network, enter the client CIDR range.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_vpc_subnet_id": {
        "computed": true,
        "description": "The ID of the subnet through which you want to route traffic. The specified subnet must be an existing target network of the Client VPN endpoint. Alternatively, if you're adding a route for the local network, specify local.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::ClientVpnRoute",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2ClientVpnRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2ClientVpnRoute), &result)
	return &result
}
