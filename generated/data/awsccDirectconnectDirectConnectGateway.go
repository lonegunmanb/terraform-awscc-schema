package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDirectconnectDirectConnectGateway = `{
  "block": {
    "attributes": {
      "amazon_side_asn": {
        "computed": true,
        "description": "The autonomous system number (ASN) for the Amazon side of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "direct_connect_gateway_arn": {
        "computed": true,
        "description": "The ARN of the Direct Connect gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "direct_connect_gateway_id": {
        "computed": true,
        "description": "The ID of the Direct Connect gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "direct_connect_gateway_name": {
        "computed": true,
        "description": "The name of the Direct Connect gateway.",
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
        "description": "The tags associated with the Direct Connect gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::DirectConnect::DirectConnectGateway",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDirectconnectDirectConnectGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDirectconnectDirectConnectGateway), &result)
	return &result
}
