package resource

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
        "optional": true,
        "type": "string"
      },
      "fallback_model": {
        "description": "Model configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "model_arn": {
              "description": "Arn of underlying model which are added in the Prompt Router.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "models": {
        "description": "List of model configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "model_arn": {
              "description": "Arn of underlying model which are added in the Prompt Router.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "prompt_router_arn": {
        "computed": true,
        "description": "Arn of the Prompt Router.",
        "description_kind": "plain",
        "type": "string"
      },
      "prompt_router_name": {
        "description": "Name of the Prompt Router.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "routing_criteria": {
        "description": "Represents the criteria used for routing requests.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "response_quality_difference": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag Value",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
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
    "description": "Definition of AWS::Bedrock::IntelligentPromptRouter Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockIntelligentPromptRouterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockIntelligentPromptRouter), &result)
	return &result
}
