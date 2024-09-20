package resource

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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "action_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "approval_model": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "budget_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "definition": {
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
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "policy_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "roles": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "users": {
                    "computed": true,
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
            },
            "scp_action_definition": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_ids": {
                    "computed": true,
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
            },
            "ssm_action_definition": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "region": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subtype": {
                    "computed": true,
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
      "execution_role_arn": {
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
      "notification_type": {
        "description_kind": "plain",
        "required": true,
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
      },
      "subscribers": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      }
    },
    "description": "An example resource schema demonstrating some basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBudgetsBudgetsActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBudgetsBudgetsAction), &result)
	return &result
}
