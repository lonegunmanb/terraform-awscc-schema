package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Route = `{
  "block": {
    "attributes": {
      "carrier_gateway_id": {
        "computed": true,
        "description": "The ID of the carrier gateway.\n You can only use this option when the VPC contains a subnet which is associated with a Wavelength Zone.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the core network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_cidr_block": {
        "computed": true,
        "description": "The IPv4 CIDR address block used for the destination match. Routing decisions are based on the most specific match. We modify the specified CIDR block to its canonical form; for example, if you specify ` + "`" + `` + "`" + `100.68.0.18/18` + "`" + `` + "`" + `, we modify it to ` + "`" + `` + "`" + `100.68.0.0/18` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_ipv_6_cidr_block": {
        "computed": true,
        "description": "The IPv6 CIDR block used for the destination match. Routing decisions are based on the most specific match.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_prefix_list_id": {
        "computed": true,
        "description": "The ID of a prefix list used for the destination match.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "egress_only_internet_gateway_id": {
        "computed": true,
        "description": "[IPv6 traffic only] The ID of an egress-only internet gateway.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_id": {
        "computed": true,
        "description": "The ID of an internet gateway or virtual private gateway attached to your VPC.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The ID of a NAT instance in your VPC. The operation fails if you specify an instance ID unless exactly one network interface is attached.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_gateway_id": {
        "computed": true,
        "description": "The ID of the local gateway.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_gateway_id": {
        "computed": true,
        "description": "[IPv4 traffic only] The ID of a NAT gateway.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description": "The ID of a network interface.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_id": {
        "description": "The ID of the route table for the route.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_id": {
        "computed": true,
        "description": "The ID of a transit gateway.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_endpoint_id": {
        "computed": true,
        "description": "The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_peering_connection_id": {
        "computed": true,
        "description": "The ID of a VPC peering connection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Specifies a route in a route table. For more information, see [Routes](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html#route-table-routes) in the *Amazon VPC User Guide*.\n You must specify either a destination CIDR block or prefix list ID. You must also specify exactly one of the resources as the target.\n If you create a route that references a transit gateway in the same template where you create the transit gateway, you must declare a dependency on the transit gateway attachment. The route table cannot use the transit gateway until it has successfully attached to the VPC. Add a [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) in the ` + "`" + `` + "`" + `AWS::EC2::Route` + "`" + `` + "`" + ` resource to explicitly declare a dependency on the ` + "`" + `` + "`" + `AWS::EC2::TransitGatewayAttachment` + "`" + `` + "`" + ` resource.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2RouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Route), &result)
	return &result
}
