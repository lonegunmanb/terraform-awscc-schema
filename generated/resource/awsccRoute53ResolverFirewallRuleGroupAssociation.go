package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverFirewallRuleGroupAssociation = `{
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
      "firewall_rule_group_association_id": {
        "computed": true,
        "description": "Id",
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_rule_group_id": {
        "description": "FirewallRuleGroupId",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "managed_owner_name": {
        "computed": true,
        "description": "ServicePrincipal",
        "description_kind": "plain",
        "type": "string"
      },
      "modification_time": {
        "computed": true,
        "description": "Rfc3339TimeString",
        "description_kind": "plain",
        "type": "string"
      },
      "mutation_protection": {
        "computed": true,
        "description": "MutationProtectionStatus",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "FirewallRuleGroupAssociationName",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description": "Priority",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description": "FirewallDomainListAssociationStatus",
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
        "description": "VpcId",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Route53Resolver::FirewallRuleGroupAssociation.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53ResolverFirewallRuleGroupAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverFirewallRuleGroupAssociation), &result)
	return &result
}
