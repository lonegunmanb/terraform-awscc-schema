package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NatGateway = `{
  "block": {
    "attributes": {
      "allocation_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connectivity_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_drain_duration_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "nat_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_allocation_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "secondary_private_ip_address_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "secondary_private_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      }
    },
    "description": "Data Source schema for AWS::EC2::NatGateway",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2NatGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NatGateway), &result)
	return &result
}
