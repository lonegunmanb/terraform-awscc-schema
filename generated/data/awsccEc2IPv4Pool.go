package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IPv4Pool = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the public IPv4 pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the public IPv4 pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_border_group": {
        "computed": true,
        "description": "The name of the location from which the address pool is advertised.",
        "description_kind": "plain",
        "type": "string"
      },
      "pool_id": {
        "computed": true,
        "description": "The ID of the public IPv4 pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the public IPv4 pool.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "total_address_count": {
        "computed": true,
        "description": "The total number of addresses in the pool.",
        "description_kind": "plain",
        "type": "number"
      },
      "total_available_address_count": {
        "computed": true,
        "description": "The total number of available addresses in the pool.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::EC2::IPv4Pool",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2IPv4PoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IPv4Pool), &result)
	return &result
}
