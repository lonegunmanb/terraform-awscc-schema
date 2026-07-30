package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayPolicyTableEntry = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_rule": {
        "description": "The policy rule associated with the entry.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_cidr_block": {
              "computed": true,
              "description": "The destination CIDR block for the transit gateway policy rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_port_range": {
              "computed": true,
              "description": "The destination port range for the transit gateway policy rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "computed": true,
              "description": "The protocol for the transit gateway policy rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_cidr_block": {
              "computed": true,
              "description": "The source CIDR block for the transit gateway policy rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_port_range": {
              "computed": true,
              "description": "The source port range for the transit gateway policy rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "policy_rule_number": {
        "description": "The rule number for the policy table entry.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the policy table entry.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_route_table_id": {
        "description": "The ID of the target route table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_policy_table_id": {
        "description": "The ID of the transit gateway policy table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "AWS::EC2::TransitGatewayPolicyTableEntry Resource Definition",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayPolicyTableEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayPolicyTableEntry), &result)
	return &result
}
