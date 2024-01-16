package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectGateway = `{
  "block": {
    "attributes": {
      "egress_cidr_blocks": {
        "description": "The range of IP addresses that contribute content or initiate output requests for flows communicating with this gateway. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "gateway_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_state": {
        "computed": true,
        "description": "The current status of the gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the gateway. This name can not be modified after the gateway is created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "networks": {
        "description": "The list of networks in the gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr_block": {
              "description": "A unique IP address range to use for this network. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description": "The name of the network. This name is used to reference the network and must be unique among networks in this gateway.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      }
    },
    "description": "Resource schema for AWS::MediaConnect::Gateway",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectGateway), &result)
	return &result
}
