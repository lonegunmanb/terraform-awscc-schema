package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2EgressOnlyInternetGateway = `{
  "block": {
    "attributes": {
      "egress_only_internet_gateway_id": {
        "computed": true,
        "description": "Service Generated ID of the EgressOnlyInternetGateway",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "description": "The ID of the VPC for which to create the egress-only internet gateway.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::EgressOnlyInternetGateway",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2EgressOnlyInternetGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2EgressOnlyInternetGateway), &result)
	return &result
}
