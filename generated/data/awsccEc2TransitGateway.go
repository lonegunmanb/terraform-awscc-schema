package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGateway = `{
  "block": {
    "attributes": {
      "amazon_side_asn": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "association_default_route_table_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_accept_shared_attachments": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_route_table_association": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_route_table_propagation": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_support": {
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
      "multicast_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "propagation_default_route_table_id": {
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
      },
      "transit_gateway_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_cidr_blocks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "transit_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_ecmp_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::TransitGateway",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TransitGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGateway), &result)
	return &result
}
