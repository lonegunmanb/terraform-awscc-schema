package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockApplicationInferenceProfile = `{
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
        "description": "Description of the inference profile",
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
      "inference_profile_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "inference_profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "inference_profile_identifier": {
        "computed": true,
        "description": "Inference profile identifier. Supports both system-defined inference profile ids, and inference profile ARNs.",
        "description_kind": "plain",
        "type": "string"
      },
      "inference_profile_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_source": {
        "computed": true,
        "description": "Various ways to encode a list of models in a CreateInferenceProfile request",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "copy_from": {
              "computed": true,
              "description": "Source arns for a custom inference profile to copy its regional load balancing config from. This\ncan either be a foundation model or predefined inference profile ARN.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "models": {
        "computed": true,
        "description": "List of model configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "model_arn": {
              "computed": true,
              "description": "ARN for Foundation Models in Bedrock. These models can be used as base models for model customization jobs",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "status": {
        "computed": true,
        "description": "Status of the Inference Profile",
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
        "description": "Type of the Inference Profile",
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
    "description": "Definition of AWS::Bedrock::ApplicationInferenceProfile Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockApplicationInferenceProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockApplicationInferenceProfile), &result)
	return &result
}
