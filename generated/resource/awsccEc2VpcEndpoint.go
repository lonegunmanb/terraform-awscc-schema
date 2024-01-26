package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcEndpoint = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_entries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "policy_document": {
        "computed": true,
        "description": "A policy to attach to the endpoint that controls access to the service.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_dns_enabled": {
        "computed": true,
        "description": "Indicate whether to associate a private hosted zone with the specified VPC.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "route_table_ids": {
        "computed": true,
        "description": "One or more route table IDs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "security_group_ids": {
        "computed": true,
        "description": "The ID of one or more security groups to associate with the endpoint network interface.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "service_name": {
        "description": "The service name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "The ID of one or more subnets in which to create an endpoint network interface.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "vpc_endpoint_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "description": "The ID of the VPC in which the endpoint will be used.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::VPCEndpoint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcEndpoint), &result)
	return &result
}