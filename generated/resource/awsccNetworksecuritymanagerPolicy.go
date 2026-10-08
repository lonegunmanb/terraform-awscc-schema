package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworksecuritymanagerPolicy = `{
  "block": {
    "attributes": {
      "associated_template_and_rule_list": {
        "computed": true,
        "description": "List of templates and rules associated with this policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rule_arn": {
              "computed": true,
              "description": "ARN of the associated rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "template_arn": {
              "computed": true,
              "description": "ARN of the associated template.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "firewall_type": {
        "description": "The type of firewall.",
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
      "policy_arn": {
        "computed": true,
        "description": "The ARN of the policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_configuration": {
        "description": "Configuration settings for policy behavior.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "remediation_enabled": {
              "computed": true,
              "description": "Controls automatic remediation of non-compliant resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "resources_clean_up": {
              "computed": true,
              "description": "Controls automatic cleanup of unused resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "waf_config": {
              "computed": true,
              "description": "WAF-specific policy settings. Populated only for WAF firewall type policies.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "conflict_resolution": {
                    "computed": true,
                    "description": "Conflict-resolution strategy applied to AWS WAF policies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "existing_customer_web_acl_resolution": {
                    "computed": true,
                    "description": "Controls how Network Security Manager handles remediation when a resource already has a customer-created WebACL.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "policy_description": {
        "computed": true,
        "description": "A description of the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description": "The unique identifier of the policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_name": {
        "description": "The name of the policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "priority": {
        "description": "The priority of the policy.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "The status of the policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the policy.",
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
      "updated_at": {
        "computed": true,
        "description": "An ISO 8601 timestamp indicating when the policy was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version number.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::NetworkSecurityManager::Policy. Creates and manages a Network Security Manager policy.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworksecuritymanagerPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworksecuritymanagerPolicy), &result)
	return &result
}
