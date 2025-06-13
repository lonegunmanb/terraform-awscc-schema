package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the egress only internet gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
