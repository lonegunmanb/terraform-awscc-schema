package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverFirewallRuleGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "Rfc3339TimeString",
        "description_kind": "plain",
        "type": "string"
      },
      "creator_request_id": {
        "computed": true,
        "description": "The id of the creator request.",
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_rules": {
        "computed": true,
        "description": "FirewallRules",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action": {
              "computed": true,
              "description": "Rule Action",
              "description_kind": "plain",
              "type": "string"
            },
            "block_override_dns_type": {
              "computed": true,
              "description": "BlockOverrideDnsType",
              "description_kind": "plain",
              "type": "string"
            },
            "block_override_domain": {
              "computed": true,
              "description": "BlockOverrideDomain",
              "description_kind": "plain",
              "type": "string"
            },
            "block_override_ttl": {
              "computed": true,
              "description": "BlockOverrideTtl",
              "description_kind": "plain",
              "type": "number"
            },
            "block_response": {
              "computed": true,
              "description": "BlockResponse",
              "description_kind": "plain",
              "type": "string"
            },
            "firewall_domain_list_id": {
              "computed": true,
              "description": "ResourceId",
              "description_kind": "plain",
              "type": "string"
            },
            "priority": {
              "computed": true,
              "description": "Rule Priority",
              "description_kind": "plain",
              "type": "number"
            },
            "qtype": {
              "computed": true,
              "description": "Qtype",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "modification_time": {
        "computed": true,
        "description": "Rfc3339TimeString",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "FirewallRuleGroupName",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "AccountId",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_count": {
        "computed": true,
        "description": "Count",
        "description_kind": "plain",
        "type": "number"
      },
      "share_status": {
        "computed": true,
        "description": "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description": "FirewallRuleGroupStatus",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Route53Resolver::FirewallRuleGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53ResolverFirewallRuleGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverFirewallRuleGroup), &result)
	return &result
}
