package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayMeteringPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "middlebox_attachment_ids": {
        "computed": true,
        "description": "Middle box attachment Ids",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "state": {
        "computed": true,
        "description": "State of the transit gateway metering policy",
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
      "transit_gateway_id": {
        "computed": true,
        "description": "The Id of transit gateway",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_metering_policy_id": {
        "computed": true,
        "description": "The Id of the transit gateway metering policy",
        "description_kind": "plain",
        "type": "string"
      },
      "update_effective_at": {
        "computed": true,
        "description": "The timestamp at which the latest action performed on the metering policy will become effective",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::TransitGatewayMeteringPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TransitGatewayMeteringPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayMeteringPolicy), &result)
	return &result
}
