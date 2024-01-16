package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2EgressOnlyInternetGateway = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC for which to create the egress-only internet gateway.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::EgressOnlyInternetGateway",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2EgressOnlyInternetGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2EgressOnlyInternetGateway), &result)
	return &result
}
