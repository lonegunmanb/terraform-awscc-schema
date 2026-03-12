package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53GlobalresolverFirewallRule = `{
  "block": {
    "attributes": {
      "action": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "block_override_dns_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "block_override_domain": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "block_override_ttl": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "block_response": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "confidence_threshold": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "dns_advanced_protection": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_view_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_domain_list_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firewall_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "q_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
    "description": "Resource schema for AWS::Route53GlobalResolver::FirewallRule",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53GlobalresolverFirewallRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53GlobalresolverFirewallRule), &result)
	return &result
}
