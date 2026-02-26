package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccComputeoptimizerAutomationRule = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The AWS account ID that owns the automation rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_timestamp": {
        "computed": true,
        "description": "The timestamp when the automation rule was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "criteria": {
        "computed": true,
        "description": "Filter criteria that specify which recommended actions qualify for implementation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ebs_volume_size_in_gib": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ebs_volume_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "estimated_monthly_savings": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "look_back_period_in_days": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "region": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_arn": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_tag": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "restart_needed": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description": "The description of the automation rule.",
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
      "last_updated_timestamp": {
        "computed": true,
        "description": "The timestamp when the automation rule was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the automation rule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organization_configuration": {
        "computed": true,
        "description": "Organization configuration for organization rules, including rule apply order and account scope.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_ids": {
              "computed": true,
              "description": "List of account IDs where the organization rule applies",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "rule_apply_order": {
              "computed": true,
              "description": "When the rule should be applied relative to account rules",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "priority": {
        "computed": true,
        "description": "Rule priority within its group",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recommended_action_types": {
        "description": "The types of recommended actions this rule will implement.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "rule_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the automation rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_id": {
        "computed": true,
        "description": "The unique identifier of the automation rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_revision": {
        "computed": true,
        "description": "The revision number of the automation rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_type": {
        "description": "The type of automation rule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "description": "The schedule configuration for when the rule runs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "execution_window_in_minutes": {
              "computed": true,
              "description": "Execution window duration in minutes",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "schedule_expression": {
              "computed": true,
              "description": "Schedule expression (e.g., cron or rate expression)",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_expression_timezone": {
              "computed": true,
              "description": "IANA timezone identifier",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "status": {
        "description": "The status of the automation rule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the automation rule.",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Creates an AWS Compute Optimizer automation rule that automatically implements recommended actions based on your defined criteria and schedule. Automation rules are global resources that manage automated actions across all AWS Regions where Compute Optimizer Automation is available. Organization-level rules can only be created by the management account or delegated administrator.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccComputeoptimizerAutomationRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccComputeoptimizerAutomationRule), &result)
	return &result
}
