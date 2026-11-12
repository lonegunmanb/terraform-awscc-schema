package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchAlert = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The AWS account ID that owns the alert.",
        "description_kind": "plain",
        "type": "string"
      },
      "alert_id": {
        "computed": true,
        "description": "The ID of the alert, assigned by CloudWatch when the alert is created. It doesn't change when the alert is updated, and it is the resource ID in the alert's ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the alert.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time the alert was created, in ISO 8601 format (UTC).",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the alert.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The display name of the alert. Names don't need to be unique within a space; the alert is identified by AlertId.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_rules": {
        "computed": true,
        "description": "The notification rules that determine when and where notifications are sent.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "target": {
              "computed": true,
              "description": "Where the notification is sent.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The ARN of the target. For sns, the ARN of the Amazon SNS topic. For slack and pagerduty, the ARN of a CloudWatch Omni integration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "metadata": {
                    "computed": true,
                    "description": "Settings specific to the target, as key-value pairs (up to 20). The keys depend on the integration type. Not used for sns targets.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "type": {
                    "computed": true,
                    "description": "The target type. Valid values: sns, slack, pagerduty.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "trigger": {
              "computed": true,
              "description": "The alert states that trigger the notification rule.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "state_values": {
                    "computed": true,
                    "description": "The alert states that trigger this rule. A notification is sent when the alert enters any listed state. If omitted or empty, every state triggers the rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "notifications_enabled": {
        "computed": true,
        "description": "Whether the alert sends notifications. When false, the alert is still evaluated and records its state, but sends nothing. Defaults to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "profile_id": {
        "description": "The ID of the access profile the alert uses to run its query and send notifications.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule": {
        "description": "The rule that defines how the alert is evaluated.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "telemetry_rule": {
              "description": "A rule that queries telemetry, evaluates the result against a condition, and sets how missing data is handled.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition": {
                    "description": "The thresholds that set the alert's state. Specify at least one of WarningThreshold and CriticalThreshold.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparator": {
                          "computed": true,
                          "description": "The comparison operator. Valid values: GT, GTE, LT, LTE (greater than, greater than or equal to, less than, less than or equal to).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "critical_threshold": {
                          "computed": true,
                          "description": "The value at which the alert enters the CRITICAL state.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "threshold_field": {
                          "computed": true,
                          "description": "The result field to compare against the threshold. Required when ThresholdMode is FIELD_VALUE.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "threshold_mode": {
                          "computed": true,
                          "description": "How the threshold is applied. COUNT_OF_RESULTS compares the number of rows the query returns. FIELD_VALUE compares the value of ThresholdField in each result row; each row is tracked separately.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "warning_threshold": {
                          "computed": true,
                          "description": "The value at which the alert enters the WARNING state.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "evaluation": {
                    "description": "How often the alert is evaluated, and how long a change must last before the state changes.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "interval_seconds": {
                          "description": "The time between evaluations, in seconds. Valid values: 30, 60, 120, 300, 600, 900, 1800, 3600.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "pending_duration_seconds": {
                          "computed": true,
                          "description": "How long a threshold breach must last before the alert changes state, in seconds. Must be a multiple of IntervalSeconds. 0 changes state immediately.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "recovery_duration_seconds": {
                          "computed": true,
                          "description": "How long the result must stay within its thresholds before the alert returns to OK, in seconds. Must be a multiple of IntervalSeconds. 0 returns to OK immediately.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "no_data": {
                    "computed": true,
                    "description": "How the alert behaves when an evaluation returns no data. If omitted, the alert reports NODATA.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "treat_as": {
                          "computed": true,
                          "description": "The state to report when an evaluation returns no data. Valid values: OK, WARNING, CRITICAL, NODATA.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "query": {
                    "description": "The query the alert runs, and the language it is written in.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "expression": {
                          "description": "The query expression. For PromQL, the series value is returned in a field named value.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "language": {
                          "description": "The query language. Valid values: SQL, PROMQL.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "space_id": {
        "description": "The ID of the space to create the alert in. Changing this value replaces the alert.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
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
        "description": "The time the alert was last updated, in ISO 8601 format (UTC).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudWatch::Alert. An alert runs a SQL or PromQL query in a CloudWatch Omni space on a fixed interval, compares the result against warning and critical thresholds, and sends notifications to Amazon SNS topics or integrations when its state changes.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchAlertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchAlert), &result)
	return &result
}
