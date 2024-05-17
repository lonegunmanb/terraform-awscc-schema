package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBudgetsBudgetsAction = `{
  "block": {
    "attributes": {
      "action_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "action_threshold": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "action_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "approval_model": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "budget_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "definition": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "iam_action_definition": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "groups": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "policy_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "roles": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "users": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "scp_action_definition": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ssm_action_definition": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "region": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subtype": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "execution_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_tags": {
        "computed": true,
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
      },
      "subscribers": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Budgets::BudgetsAction",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBudgetsBudgetsActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBudgetsBudgetsAction), &result)
	return &result
}
