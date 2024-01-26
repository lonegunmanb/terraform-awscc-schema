package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Route = `{
  "block": {
    "attributes": {
      "carrier_gateway_id": {
        "computed": true,
        "description": "The ID of the carrier gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "cidr_block": {
        "computed": true,
        "description": "The primary identifier of the resource generated by the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the core network.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_cidr_block": {
        "computed": true,
        "description": "The IPv4 CIDR block used for the destination match.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_ipv_6_cidr_block": {
        "computed": true,
        "description": "The IPv6 CIDR block used for the destination match.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_prefix_list_id": {
        "computed": true,
        "description": "The ID of managed prefix list, it's a set of one or more CIDR blocks.",
        "description_kind": "plain",
        "type": "string"
      },
      "egress_only_internet_gateway_id": {
        "computed": true,
        "description": "The ID of the egress-only internet gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_id": {
        "computed": true,
        "description": "The ID of an internet gateway or virtual private gateway attached to your VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The ID of a NAT instance in your VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_id": {
        "computed": true,
        "description": "The ID of the local gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "nat_gateway_id": {
        "computed": true,
        "description": "The ID of a NAT gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description": "The ID of the network interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "route_table_id": {
        "computed": true,
        "description": "The ID of the route table. The routing table must be associated with the same VPC that the virtual private gateway is attached to.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_id": {
        "computed": true,
        "description": "The ID of a transit gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_endpoint_id": {
        "computed": true,
        "description": "The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_peering_connection_id": {
        "computed": true,
        "description": "The ID of a VPC peering connection.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::Route",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2RouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Route), &result)
	return &result
}