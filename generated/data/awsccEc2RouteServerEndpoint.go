package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2RouteServerEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Route Server Endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "eni_address": {
        "computed": true,
        "description": "Elastic Network Interface IP address owned by the Route Server Endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "eni_id": {
        "computed": true,
        "description": "Elastic Network Interface ID owned by the Route Server Endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_server_endpoint_id": {
        "computed": true,
        "description": "The ID of the Route Server Endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "route_server_id": {
        "computed": true,
        "description": "Route Server ID",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description": "Subnet ID",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
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
          "nesting_mode": "list"
        }
      },
      "vpc_id": {
        "computed": true,
        "description": "VPC ID",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::RouteServerEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2RouteServerEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2RouteServerEndpoint), &result)
	return &result
}
