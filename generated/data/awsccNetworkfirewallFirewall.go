package data

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
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled_analysis_types": {
        "computed": true,
        "description": "The types of analysis to enable for the firewall. Can be TLS_SNI, HTTP_HOST, or both.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_policy_arn": {
        "computed": true,
        "description": "A resource ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_policy_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "subnet_mappings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_address_type": {
              "computed": true,
              "description": "A IPAddressType",
              "description_kind": "plain",
              "type": "string"
            },
            "subnet_id": {
              "computed": true,
              "description": "A SubnetId.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
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
          "nesting_mode": "set"
        }
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkFirewall::Firewall",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkfirewallFirewallSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallFirewall), &result)
	return &result
}
