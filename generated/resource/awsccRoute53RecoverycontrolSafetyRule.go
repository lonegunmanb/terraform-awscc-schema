package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecoverycontrolSafetyRule = `{
  "block": {
    "attributes": {
      "assertion_rule": {
        "computed": true,
        "description": "An assertion rule enforces that, when a routing control state is changed, that the criteria set by the rule configuration is met. Otherwise, the change to the routing control is not accepted.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "asserted_controls": {
              "description": "The routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed. For example, you might include three routing controls, one for each of three AWS Regions.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "wait_period_ms": {
              "description": "An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent \"flapping\" of state. The wait period is 5000 ms by default, but you can choose a custom value.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "control_panel_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the control panel.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gating_rule": {
        "computed": true,
        "description": "A gating rule verifies that a set of gating controls evaluates as true, based on a rule configuration that you specify. If the gating rule evaluates to true, Amazon Route 53 Application Recovery Controller allows a set of routing control state changes to run and complete against the set of target controls.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "gating_controls": {
              "description": "The gating controls for the gating rule. That is, routing controls that are evaluated by the rule configuration that you specify.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "target_controls": {
              "description": "Routing controls that can only be set or unset if the specified RuleConfig evaluates to true for the specified GatingControls. For example, say you have three gating controls, one for each of three AWS Regions. Now you specify AtLeast 2 as your RuleConfig. With these settings, you can only change (set or unset) the routing controls that you have specified as TargetControls if that rule evaluates to true. \nIn other words, your ability to change the routing controls that you have specified as TargetControls is gated by the rule that you set for the routing controls in GatingControls.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "wait_period_ms": {
              "description": "An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent \"flapping\" of state. The wait period is 5000 ms by default, but you can choose a custom value.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name for the safety rule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_config": {
        "computed": true,
        "description": "The rule configuration for an assertion rule or gating rule. This is the criteria that you set for specific assertion controls (routing controls) or gating controls. This configuration specifies how many controls must be enabled after a transaction completes.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "inverted": {
              "description": "Logical negation of the rule. If the rule would usually evaluate true, it's evaluated as false, and vice versa.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "threshold": {
              "description": "The value of N, when you specify an ATLEAST rule type. That is, Threshold is the number of controls that must be set when you specify an ATLEAST type.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "type": {
              "description": "A rule can be one of the following: ATLEAST, AND, or OR.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "safety_rule_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the safety rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS Route53 Recovery Control basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53RecoverycontrolSafetyRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecoverycontrolSafetyRule), &result)
	return &result
}
