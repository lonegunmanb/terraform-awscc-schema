package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2RouteServer = `{
  "block": {
    "attributes": {
      "amazon_side_asn": {
        "computed": true,
        "description": "The Amazon-side ASN of the Route Server.",
        "description_kind": "plain",
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Route Server.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "persist_routes": {
        "computed": true,
        "description": "Whether to enable persistent routes",
        "description_kind": "plain",
        "type": "string"
      },
      "persist_routes_duration": {
        "computed": true,
        "description": "The duration of persistent routes in minutes",
        "description_kind": "plain",
        "type": "number"
      },
      "route_server_id": {
        "computed": true,
        "description": "The ID of the Route Server.",
        "description_kind": "plain",
        "type": "string"
      },
      "sns_notifications_enabled": {
        "computed": true,
        "description": "Whether to enable SNS notifications",
        "description_kind": "plain",
        "type": "bool"
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
      }
    },
    "description": "Data Source schema for AWS::EC2::RouteServer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2RouteServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2RouteServer), &result)
	return &result
}
