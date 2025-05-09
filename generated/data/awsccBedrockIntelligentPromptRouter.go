package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockIntelligentPromptRouter = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "Time Stamp",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the Prompt Router.",
        "description_kind": "plain",
        "type": "string"
      },
      "fallback_model": {
        "computed": true,
        "description": "Model configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "model_arn": {
              "computed": true,
              "description": "Arn of underlying model which are added in the Prompt Router.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "models": {
        "computed": true,
        "description": "List of model configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "model_arn": {
              "computed": true,
              "description": "Arn of underlying model which are added in the Prompt Router.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "prompt_router_arn": {
        "computed": true,
        "description": "Arn of the Prompt Router.",
        "description_kind": "plain",
        "type": "string"
      },
      "prompt_router_name": {
        "computed": true,
        "description": "Name of the Prompt Router.",
        "description_kind": "plain",
        "type": "string"
      },
      "routing_criteria": {
        "computed": true,
        "description": "Represents the criteria used for routing requests.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "response_quality_difference": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "Status of a PromptRouter",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "List of Tags",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag Key",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag Value",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "Type of a Prompt Router",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "Time Stamp",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Bedrock::IntelligentPromptRouter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockIntelligentPromptRouterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockIntelligentPromptRouter), &result)
	return &result
}
