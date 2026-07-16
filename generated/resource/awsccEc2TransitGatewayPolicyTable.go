package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "transit_gateway_id": {
        "description": "The Id of transit gateway",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_policy_table_id": {
        "computed": true,
        "description": "The Id of transit gateway policy table.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "AWS::EC2::TransitGatewayPolicyTable Resource Definition",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayPolicyTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayPolicyTable), &result)
	return &result
}
