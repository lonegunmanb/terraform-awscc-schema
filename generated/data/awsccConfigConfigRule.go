package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigConfigRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "ARN generated for the AWS Config rule ",
        "description_kind": "plain",
        "type": "string"
      },
      "compliance": {
        "computed": true,
        "description": "Compliance details of the Config rule",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description": "Compliance type determined by the Config rule",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "config_rule_id": {
        "computed": true,
        "description": "ID of the config rule",
        "description_kind": "plain",
        "type": "string"
      },
      "config_rule_name": {
        "computed": true,
        "description": "Name for the AWS Config rule",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description provided for the AWS Config rule",
        "description_kind": "plain",
        "type": "string"
      },
      "evaluation_modes": {
        "computed": true,
        "description": "List of EvaluationModeConfiguration objects",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mode": {
              "computed": true,
              "description": "Mode of evaluation of AWS Config rule",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "input_parameters": {
        "computed": true,
        "description": "JSON string passed the Lambda function",
        "description_kind": "plain",
        "type": "string"
      },
      "maximum_execution_frequency": {
        "computed": true,
        "description": "Maximum frequency at which the rule has to be evaluated",
        "description_kind": "plain",
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description": "Scope to constrain which resources can trigger the AWS Config rule",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compliance_resource_id": {
              "computed": true,
              "description": "ID of the only one resource which we want to trigger the rule",
              "description_kind": "plain",
              "type": "string"
            },
            "compliance_resource_types": {
              "computed": true,
              "description": "Resource types of resources which we want to trigger the rule",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "tag_key": {
              "computed": true,
              "description": "Tag key applied only to resources which we want to trigger the rule",
              "description_kind": "plain",
              "type": "string"
            },
            "tag_value": {
              "computed": true,
              "description": "Tag value applied only to resources which we want to trigger the rule",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "source": {
        "computed": true,
        "description": "Source of events for the AWS Config rule",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_policy_details": {
              "computed": true,
              "description": "Custom policy details when rule is custom owned",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_debug_log_delivery": {
                    "computed": true,
                    "description": "Logging toggle for custom policy rule",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "policy_runtime": {
                    "computed": true,
                    "description": "Runtime system for custom policy rule",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "policy_text": {
                    "computed": true,
                    "description": "Policy definition containing logic for custom policy rule",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "owner": {
              "computed": true,
              "description": "Owner of the config rule",
              "description_kind": "plain",
              "type": "string"
            },
            "source_details": {
              "computed": true,
              "description": "List of message types that can trigger the rule",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event_source": {
                    "computed": true,
                    "description": "Source of event that can trigger the rule",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "maximum_execution_frequency": {
                    "computed": true,
                    "description": "Frequency at which the rule has to be evaluated",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "message_type": {
                    "computed": true,
                    "description": "Notification type that can trigger the rule",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "source_identifier": {
              "computed": true,
              "description": "Identifier for the source of events",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Config::ConfigRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConfigConfigRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigConfigRule), &result)
	return &result
}
