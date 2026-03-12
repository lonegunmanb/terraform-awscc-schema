package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53GlobalresolverFirewallRule = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "block_override_dns_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "block_override_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "block_override_ttl": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "block_response": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "confidence_threshold": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_advanced_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_view_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_domain_list_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_rule_id": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "q_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "query_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53GlobalResolver::FirewallRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53GlobalresolverFirewallRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53GlobalresolverFirewallRule), &result)
	return &result
}
