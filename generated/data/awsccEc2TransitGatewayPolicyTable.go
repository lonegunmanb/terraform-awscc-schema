package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayPolicyTable = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "Creation time of the transit gateway policy table",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the transit gateway policy table",
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
      "transit_gateway_policy_table_id": {
        "computed": true,
        "description": "The Id of transit gateway policy table.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::TransitGatewayPolicyTable",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TransitGatewayPolicyTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayPolicyTable), &result)
	return &result
}
