package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigConfigRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compliance": {
        "computed": true,
        "description": "Indicates whether an AWS resource or CC rule is compliant and provides the number of contributors that affect the compliance.",
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
        },
        "optional": true
      },
      "config_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "config_rule_name": {
        "computed": true,
        "description": "A name for the CC rule. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description that you provide for the CC rule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "evaluation_modes": {
        "computed": true,
        "description": "The modes the CC rule can be evaluated in. The valid values are distinct objects. By default, the value is Detective evaluation mode only.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mode": {
              "computed": true,
              "description": "The mode of an evaluation. The valid values are Detective or Proactive.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "input_parameters": {
        "computed": true,
        "description": "A string, in JSON format, that is passed to the CC rule Lambda function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maximum_execution_frequency": {
        "computed": true,
        "description": "The maximum frequency with which CC runs evaluations for a rule. You can specify a value for ` + "`" + `` + "`" + `MaximumExecutionFrequency` + "`" + `` + "`" + ` when:\n  +  You are using an AWS managed rule that is triggered at a periodic frequency.\n  +  Your custom rule is triggered when CC delivers the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html).\n  \n  By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ` + "`" + `` + "`" + `MaximumExecutionFrequency` + "`" + `` + "`" + ` parameter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description": "Defines which resources can trigger an evaluation for the rule. The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.\n  The scope can be empty.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compliance_resource_id": {
              "computed": true,
              "description": "The ID of the only AWS resource that you want to trigger an evaluation for the rule. If you specify a resource ID, you must specify one resource type for ` + "`" + `` + "`" + `ComplianceResourceTypes` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "compliance_resource_types": {
              "computed": true,
              "description": "The resource types of only those AWS resources that you want to trigger an evaluation for the rule. You can only specify one type if you also specify a resource ID for ` + "`" + `` + "`" + `ComplianceResourceId` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tag_key": {
              "computed": true,
              "description": "The tag key that is applied to only those AWS resources that you want to trigger an evaluation for the rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_value": {
              "computed": true,
              "description": "The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule. If you specify a value for ` + "`" + `` + "`" + `TagValue` + "`" + `` + "`" + `, you must also specify a value for ` + "`" + `` + "`" + `TagKey` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "source": {
        "description": "Provides the rule owner (` + "`" + `` + "`" + `` + "`" + `` + "`" + ` for managed rules, ` + "`" + `` + "`" + `CUSTOM_POLICY` + "`" + `` + "`" + ` for Custom Policy rules, and ` + "`" + `` + "`" + `CUSTOM_LAMBDA` + "`" + `` + "`" + ` for Custom Lambda rules), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_policy_details": {
              "computed": true,
              "description": "Provides the runtime system, policy definition, and whether debug logging is enabled. Required when owner is set to ` + "`" + `` + "`" + `CUSTOM_POLICY` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_debug_log_delivery": {
                    "computed": true,
                    "description": "The boolean expression for enabling debug logging for your CC Custom Policy rule. The default value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "policy_runtime": {
                    "computed": true,
                    "description": "The runtime system for your CC Custom Policy rule. Guard is a policy-as-code language that allows you to write policies that are enforced by CC Custom Policy rules. For more information about Guard, see the [Guard GitHub Repository](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-guard).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "policy_text": {
                    "computed": true,
                    "description": "The policy definition containing the logic for your CC Custom Policy rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "owner": {
              "description": "Indicates whether AWS or the customer owns and manages the CC rule.\n  CC Managed Rules are predefined rules owned by AWS. For more information, see [Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) in the *developer guide*.\n  CC Custom Rules are rules that you can develop either with Guard (` + "`" + `` + "`" + `CUSTOM_POLICY` + "`" + `` + "`" + `) or LAMlong (` + "`" + `` + "`" + `CUSTOM_LAMBDA` + "`" + `` + "`" + `). For more information, see [Custom Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html) in the *developer guide*.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_details": {
              "computed": true,
              "description": "Provides the source and the message types that cause CC to evaluate your AWS resources against a rule. It also provides the frequency with which you want CC to run evaluations for the rule if the trigger type is periodic.\n If the owner is set to ` + "`" + `` + "`" + `CUSTOM_POLICY` + "`" + `` + "`" + `, the only acceptable values for the CC rule trigger message type are ` + "`" + `` + "`" + `ConfigurationItemChangeNotification` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `OversizedConfigurationItemChangeNotification` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event_source": {
                    "description": "The source of the event, such as an AWS service, that triggers CC to evaluate your AWS resources.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "maximum_execution_frequency": {
                    "computed": true,
                    "description": "The frequency at which you want CC to run evaluations for a custom rule with a periodic trigger. If you specify a value for ` + "`" + `` + "`" + `MaximumExecutionFrequency` + "`" + `` + "`" + `, then ` + "`" + `` + "`" + `MessageType` + "`" + `` + "`" + ` must use the ` + "`" + `` + "`" + `ScheduledNotification` + "`" + `` + "`" + ` value.\n  By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ` + "`" + `` + "`" + `MaximumExecutionFrequency` + "`" + `` + "`" + ` parameter.\n Based on the valid value you choose, CC runs evaluations once for each valid value. For example, if you choose ` + "`" + `` + "`" + `Three_Hours` + "`" + `` + "`" + `, CC runs evaluations once every three hours. In this case, ` + "`" + `` + "`" + `Three_Hours` + "`" + `` + "`" + ` is the frequency of this rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "message_type": {
                    "description": "The type of notification that triggers CC to run an evaluation for a rule. You can specify the following notification types:\n  +   ` + "`" + `` + "`" + `ConfigurationItemChangeNotification` + "`" + `` + "`" + ` - Triggers an evaluation when CC delivers a configuration item as a result of a resource change.\n  +   ` + "`" + `` + "`" + `OversizedConfigurationItemChangeNotification` + "`" + `` + "`" + ` - Triggers an evaluation when CC delivers an oversized configuration item. CC may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.\n  +   ` + "`" + `` + "`" + `ScheduledNotification` + "`" + `` + "`" + ` - Triggers a periodic evaluation at the frequency specified for ` + "`" + `` + "`" + `MaximumExecutionFrequency` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `ConfigurationSnapshotDeliveryCompleted` + "`" + `` + "`" + ` - Triggers a periodic evaluation when CC delivers a configuration snapshot.\n  \n If you want your custom rule to be triggered by configuration changes, specify two SourceDetail objects, one for ` + "`" + `` + "`" + `ConfigurationItemChangeNotification` + "`" + `` + "`" + ` and one for ` + "`" + `` + "`" + `OversizedConfigurationItemChangeNotification` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "source_identifier": {
              "computed": true,
              "description": "For CC Managed rules, a predefined identifier from a list. For example, ` + "`" + `` + "`" + `IAM_PASSWORD_POLICY` + "`" + `` + "`" + ` is a managed rule. To reference a managed rule, see [List of Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html).\n For CC Custom Lambda rules, the identifier is the Amazon Resource Name (ARN) of the rule's LAMlong function, such as ` + "`" + `` + "`" + `arn:aws:lambda:us-east-2:123456789012:function:custom_rule_name` + "`" + `` + "`" + `.\n For CC Custom Policy rules, this field will be ignored.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "You must first create and start the CC configuration recorder in order to create CC managed rules with CFNlong. For more information, see [Managing the Configuration Recorder](https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html).\n Adds or updates an CC rule to evaluate if your AWS resources comply with your desired configurations. For information on how many CC rules you can have per account, see [Service Limits](https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html) in the *Developer Guide*.\n There are two types of rules: *Managed Rules* and *Custom Rules*. You can use the ` + "`" + `` + "`" + `ConfigRule` + "`" + `` + "`" + ` resource to create both CC Managed Rules and CC Custom Rules.\n CC Managed Rules are predefined, customizable rules created by CC. For a list of managed rules, see [List of Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html). If you are adding an CC managed rule, you must specify the rule's identifi",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConfigConfigRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigConfigRule), &result)
	return &result
}
