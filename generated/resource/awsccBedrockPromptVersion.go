package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockPromptVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "ARN of a prompt version resource",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Time Stamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_encryption_key_arn": {
        "computed": true,
        "description": "A KMS key ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "default_variant": {
        "computed": true,
        "description": "Name for a variant.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description for a prompt version resource.",
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
        "computed": true,
        "description": "Name for a prompt resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "prompt_arn": {
        "description": "ARN of a prompt resource possibly with a version",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "prompt_id": {
        "computed": true,
        "description": "Identifier for a Prompt",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of tag keys and values",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "updated_at": {
        "computed": true,
        "description": "Time Stamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "variants": {
        "computed": true,
        "description": "List of prompt variants",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "inference_configuration": {
              "computed": true,
              "description": "Model inference configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "text": {
                    "computed": true,
                    "description": "Prompt model inference configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_tokens": {
                          "computed": true,
                          "description": "Maximum length of output",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "stop_sequences": {
                          "computed": true,
                          "description": "List of stop sequences",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "temperature": {
                          "computed": true,
                          "description": "Controls randomness, higher values increase diversity",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "top_p": {
                          "computed": true,
                          "description": "Cumulative probability cutoff for token selection",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "model_id": {
              "computed": true,
              "description": "ARN or Id of a Bedrock Foundational Model or Inference Profile, or the ARN of a imported model, or a provisioned throughput ARN for custom models.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "Name for a variant.",
              "description_kind": "plain",
              "type": "string"
            },
            "template_configuration": {
              "computed": true,
              "description": "Prompt template configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "text": {
                    "computed": true,
                    "description": "Configuration for text prompt template",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "input_variables": {
                          "computed": true,
                          "description": "List of input variables",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "computed": true,
                                "description": "Name for an input variable",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "text": {
                          "computed": true,
                          "description": "Prompt content for String prompt template",
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
            "template_type": {
              "computed": true,
              "description": "Prompt template type",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "version": {
        "computed": true,
        "description": "Version.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::Bedrock::PromptVersion Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockPromptVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockPromptVersion), &result)
	return &result
}
