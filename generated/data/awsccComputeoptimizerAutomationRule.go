package data

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
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "ebs_volume_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "estimated_monthly_savings": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "look_back_period_in_days": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "region": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "resource_arn": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "resource_tag": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "restart_needed": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "The description of the automation rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_timestamp": {
        "computed": true,
        "description": "The timestamp when the automation rule was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the automation rule.",
        "description_kind": "plain",
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
              "type": [
                "list",
                "string"
              ]
            },
            "rule_apply_order": {
              "computed": true,
              "description": "When the rule should be applied relative to account rules",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "priority": {
        "computed": true,
        "description": "Rule priority within its group",
        "description_kind": "plain",
        "type": "string"
      },
      "recommended_action_types": {
        "computed": true,
        "description": "The types of recommended actions this rule will implement.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "The type of automation rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description": "The schedule configuration for when the rule runs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "execution_window_in_minutes": {
              "computed": true,
              "description": "Execution window duration in minutes",
              "description_kind": "plain",
              "type": "number"
            },
            "schedule_expression": {
              "computed": true,
              "description": "Schedule expression (e.g., cron or rate expression)",
              "description_kind": "plain",
              "type": "string"
            },
            "schedule_expression_timezone": {
              "computed": true,
              "description": "IANA timezone identifier",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "The status of the automation rule.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::ComputeOptimizer::AutomationRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccComputeoptimizerAutomationRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccComputeoptimizerAutomationRule), &result)
	return &result
}
