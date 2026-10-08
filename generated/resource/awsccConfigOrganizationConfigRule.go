package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigOrganizationConfigRule = `{
  "block": {
    "attributes": {
      "excluded_accounts": {
        "computed": true,
        "description": "A comma-separated list of accounts that you want to exclude from an organization AWS Config rule.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_config_rule_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an organization AWS Config rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_config_rule_name": {
        "description": "The name that you assign to an organization AWS Config rule. Required.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organization_custom_policy_rule_metadata": {
        "computed": true,
        "description": "This object specifies metadata for your organization's AWS Config Custom Policy rule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "debug_log_delivery_accounts": {
              "computed": true,
              "description": "A list of accounts that you can enable debug logging for your organization AWS Config Custom Policy rule. ",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "description": {
              "computed": true,
              "description": "The description that you provide for your organization AWS Config rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "input_parameters": {
              "computed": true,
              "description": "A string, in JSON format, that is passed to your organization AWS Config Custom Policy rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organization_config_rule_trigger_types": {
              "computed": true,
              "description": "The type of notification that initiates AWS Config to run an evaluation for a rule.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "policy_text": {
              "computed": true,
              "description": "The policy definition containing the logic for your organization AWS Config Custom Policy rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_id_scope": {
              "computed": true,
              "description": "The ID of the AWS resource that was evaluated.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_types_scope": {
              "computed": true,
              "description": "The type of the AWS resource that was evaluated.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "runtime": {
              "computed": true,
              "description": "The runtime system for your organization AWS Config Custom Policy rules.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_key_scope": {
              "computed": true,
              "description": "One part of a key-value pair that make up a tag. A key is a general label that acts like a category for more specific tag values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_value_scope": {
              "computed": true,
              "description": "The optional part of a key-value pair that make up a tag. A value acts as a descriptor within a tag category (key).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "organization_custom_rule_metadata": {
        "computed": true,
        "description": "This object specifies organization custom rule metadata such as resource type, resource ID of AWS resource, Lambda function ARN, and organization trigger types that trigger AWS Config to evaluate your AWS resources against a rule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "The description that you provide for your organization AWS Config rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "input_parameters": {
              "computed": true,
              "description": "A string, in JSON format, that is passed to your organization AWS Config rule Lambda function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lambda_function_arn": {
              "computed": true,
              "description": "The lambda function ARN.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "maximum_execution_frequency": {
              "computed": true,
              "description": "The maximum frequency with which AWS Config runs evaluations for a rule.Allowed values: One_Hour | Three_Hours | Six_Hours | Twelve_Hours | TwentyFour_Hours.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organization_config_rule_trigger_types": {
              "computed": true,
              "description": "The type of notification that triggers AWS Config to run an evaluation for a rule. You can specify the following notification types:",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resource_id_scope": {
              "computed": true,
              "description": "The ID of the AWS resource that was evaluated.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_types_scope": {
              "computed": true,
              "description": "The type of the AWS resource that was evaluated.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tag_key_scope": {
              "computed": true,
              "description": "One part of a key-value pair that make up a tag. A key is a general label that acts like a category for more specific tag values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_value_scope": {
              "computed": true,
              "description": "The optional part of a key-value pair that make up a tag. A value acts as a descriptor within a tag category (key).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "organization_managed_rule_metadata": {
        "computed": true,
        "description": "This object specifies organization managed rule metadata such as resource type and ID of AWS resource along with the rule identifier.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "The description that you provide for your organization AWS Config rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "input_parameters": {
              "computed": true,
              "description": "A string, in JSON format, that is passed to your organization AWS Config rule Lambda function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "maximum_execution_frequency": {
              "computed": true,
              "description": "The maximum frequency with which AWS Config runs evaluations for a rule. Valid Values: One_Hour | Three_Hours | Six_Hours | Twelve_Hours | TwentyFour_Hours.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_id_scope": {
              "computed": true,
              "description": "The ID of the AWS resource that was evaluated.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_types_scope": {
              "computed": true,
              "description": "The type of the AWS resource that was evaluated.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "rule_identifier": {
              "computed": true,
              "description": "Required. For organization config managed rules, a predefined identifier from a list. For example, IAM_PASSWORD_POLICY is a managed rule. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_key_scope": {
              "computed": true,
              "description": "One part of a key-value pair that make up a tag. A key is a general label that acts like a category for more specific tag values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_value_scope": {
              "computed": true,
              "description": "The optional part of a key-value pair that make up a tag. A value acts as a descriptor within a tag category (key).\n\n",
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
    "description": "Resource Type definition for AWS::Config::OrganizationConfigRule",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConfigOrganizationConfigRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigOrganizationConfigRule), &result)
	return &result
}
