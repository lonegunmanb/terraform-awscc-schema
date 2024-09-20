package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkfirewallFirewall = `{
  "block": {
    "attributes": {
      "delete_protection": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "firewall_arn": {
        "computed": true,
        "description": "A resource ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_policy_arn": {
        "description": "A resource ARN.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_policy_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "subnet_mappings": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_address_type": {
              "computed": true,
              "description": "A IPAddressType",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_id": {
              "description": "A SubnetId.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "required": true
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
          "nesting_mode": "set"
        },
        "optional": true
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::NetworkFirewall::Firewall",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkfirewallFirewallSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallFirewall), &result)
	return &result
}
