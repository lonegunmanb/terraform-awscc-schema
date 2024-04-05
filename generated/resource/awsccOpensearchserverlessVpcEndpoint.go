package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserverlessVpcEndpoint = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the VPC Endpoint",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description": "The ID of one or more security groups to associate with the endpoint network interface",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_ids": {
        "description": "The ID of one or more subnets in which to create an endpoint network interface",
        "description_kind": "plain",
        "required": true,
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
        "description": "The ID of the VPC in which the endpoint will be used.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Amazon OpenSearchServerless vpc endpoint resource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOpensearchserverlessVpcEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserverlessVpcEndpoint), &result)
	return &result
}
