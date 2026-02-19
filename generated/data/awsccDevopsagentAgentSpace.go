package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsagentAgentSpace = `{
  "block": {
    "attributes": {
      "agent_space_id": {
        "computed": true,
        "description": "The unique identifier of the AgentSpace",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AgentSpace.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the AgentSpace.",
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
        "description": "The name of the AgentSpace.",
        "description_kind": "plain",
        "type": "string"
      },
      "operator_app": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "iam": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "created_at": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "operator_app_role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "updated_at": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "idc": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "created_at": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "idc_application_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "idc_instance_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "operator_app_role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "updated_at": {
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
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the resource was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DevOpsAgent::AgentSpace",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDevopsagentAgentSpaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsagentAgentSpace), &result)
	return &result
}
