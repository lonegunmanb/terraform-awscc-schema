package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightAgent = `{
  "block": {
    "attributes": {
      "action_connectors": {
        "computed": true,
        "description": "A list of ActionConnector ARNs (max 10) attached to the agent.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "agent_id": {
        "computed": true,
        "description": "The unique identifier for the agent.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_lifecycle": {
        "computed": true,
        "description": "The lifecycle stage of the agent. PREVIEW or PUBLISHED.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_status": {
        "computed": true,
        "description": "The current status of the agent. One of ACTIVE, CREATING, UPDATING, or FAILED.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the agent.",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description": "The ID of the Amazon Web Services account where the agent is being created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time the agent was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "creator": {
        "computed": true,
        "description": "The ARN of the user who created the agent.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_prompt_input": {
        "computed": true,
        "description": "Custom prompt configuration. Specify either ExistingPrompt or NewPrompt.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "existing_prompt": {
              "computed": true,
              "description": "Reference to an existing custom prompt profile.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "model_profile_id": {
                    "computed": true,
                    "description": "The identifier of the model profile.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "qbs_aws_account_id": {
                    "computed": true,
                    "description": "The QBS AWS account identifier.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subscription_id": {
                    "computed": true,
                    "description": "The subscription identifier.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "new_prompt": {
              "computed": true,
              "description": "Parameters for creating a new custom prompt configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_instructions": {
                    "computed": true,
                    "description": "Custom instructions for the agent behavior.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "identity": {
                    "computed": true,
                    "description": "The identity or persona of the agent.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_style": {
                    "computed": true,
                    "description": "The output style for the agent responses.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "response_length": {
                    "computed": true,
                    "description": "The desired response length for the agent.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tone": {
                    "computed": true,
                    "description": "The tone used in agent responses.",
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
      "custom_prompt_interface": {
        "computed": true,
        "description": "Read-only view of the resolved custom prompt interface for the agent.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_instructions": {
              "computed": true,
              "description": "Custom instructions for the agent behavior.",
              "description_kind": "plain",
              "type": "string"
            },
            "identity": {
              "computed": true,
              "description": "The identity or persona of the agent.",
              "description_kind": "plain",
              "type": "string"
            },
            "model_profile_id": {
              "computed": true,
              "description": "The identifier of the model profile.",
              "description_kind": "plain",
              "type": "string"
            },
            "output_style": {
              "computed": true,
              "description": "The output style for the agent responses.",
              "description_kind": "plain",
              "type": "string"
            },
            "prompt_summary": {
              "computed": true,
              "description": "A summary of the resolved prompt.",
              "description_kind": "plain",
              "type": "string"
            },
            "qbs_aws_account_id": {
              "computed": true,
              "description": "The QBS AWS account identifier.",
              "description_kind": "plain",
              "type": "string"
            },
            "response_length": {
              "computed": true,
              "description": "The desired response length for the agent.",
              "description_kind": "plain",
              "type": "string"
            },
            "subscription_id": {
              "computed": true,
              "description": "The subscription identifier.",
              "description_kind": "plain",
              "type": "string"
            },
            "tone": {
              "computed": true,
              "description": "The tone used in agent responses.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the agent.",
        "description_kind": "plain",
        "type": "string"
      },
      "error_message": {
        "computed": true,
        "description": "The error message if the agent is in FAILED status.",
        "description_kind": "plain",
        "type": "string"
      },
      "icon_id": {
        "computed": true,
        "description": "The icon identifier for the agent.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The display name of the agent.",
        "description_kind": "plain",
        "type": "string"
      },
      "spaces": {
        "computed": true,
        "description": "A list of Space ARNs (max 10) attached to the agent.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "starter_prompts": {
        "computed": true,
        "description": "A list of up to 3 starter prompts displayed to users.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs to associate with the agent resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time the agent was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "welcome_message": {
        "computed": true,
        "description": "The welcome message displayed when a user opens the agent.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QuickSight::Agent",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightAgentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightAgent), &result)
	return &result
}
