package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserverlessVpcEndpoint = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the VPC Endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description": "The ID of one or more security groups to associate with the endpoint network interface",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_ids": {
        "computed": true,
        "description": "The ID of one or more subnets in which to create an endpoint network interface",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_endpoint_id": {
        "computed": true,
        "description": "The identifier of the VPC Endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC in which the endpoint will be used.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::OpenSearchServerless::VpcEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchserverlessVpcEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserverlessVpcEndpoint), &result)
	return &result
}
