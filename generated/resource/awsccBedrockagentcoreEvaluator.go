package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreEvaluator = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the evaluator was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the evaluator.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "evaluator_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the evaluator.",
        "description_kind": "plain",
        "type": "string"
      },
      "evaluator_config": {
        "description": "The configuration for the evaluator.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code_based": {
              "computed": true,
              "description": "The configuration for code-based evaluation using a Lambda function.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda_config": {
                    "computed": true,
                    "description": "The Lambda function configuration for code-based evaluation.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "lambda_arn": {
                          "computed": true,
                          "description": "The ARN of the Lambda function used for evaluation.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lambda_timeout_in_seconds": {
                          "computed": true,
                          "description": "The timeout in seconds for the Lambda function invocation.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "llm_as_a_judge": {
              "computed": true,
              "description": "The configuration for LLM-as-a-Judge evaluation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instructions": {
                    "computed": true,
                    "description": "The evaluation instructions that guide the language model in assessing agent performance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "model_config": {
                    "computed": true,
                    "description": "The model configuration that specifies which foundation model to use for evaluation.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bedrock_evaluator_model_config": {
                          "computed": true,
                          "description": "The configuration for using Amazon Bedrock models in evaluator assessments.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "additional_model_request_fields": {
                                "computed": true,
                                "description": "Additional model-specific request fields.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "inference_config": {
                                "computed": true,
                                "description": "The inference configuration parameters that control model behavior during evaluation.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max_tokens": {
                                      "computed": true,
                                      "description": "The maximum number of tokens to generate in the model response.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "temperature": {
                                      "computed": true,
                                      "description": "The temperature value that controls randomness in the model's responses.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "top_p": {
                                      "computed": true,
                                      "description": "The top-p sampling parameter that controls the diversity of the model's responses.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "model_id": {
                                "computed": true,
                                "description": "The identifier of the Amazon Bedrock model to use for evaluation.",
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
                    "optional": true
                  },
                  "rating_scale": {
                    "computed": true,
                    "description": "The rating scale that defines how evaluators should score agent performance.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "categorical": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "definition": {
                                "computed": true,
                                "description": "The description that explains what this categorical rating represents.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "label": {
                                "computed": true,
                                "description": "The label of this categorical rating option.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "numerical": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "definition": {
                                "computed": true,
                                "description": "The description that explains what this numerical rating represents.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "label": {
                                "computed": true,
                                "description": "The label that describes this numerical rating option.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "The numerical value for this rating scale option.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
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
      "evaluator_id": {
        "computed": true,
        "description": "The unique identifier of the evaluator.",
        "description_kind": "plain",
        "type": "string"
      },
      "evaluator_name": {
        "description": "The name of the evaluator. Must be unique within your account.",
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
      "level": {
        "description": "The evaluation level that determines the scope of evaluation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the evaluator.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to assign to the evaluator.",
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
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the evaluator was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::BedrockAgentCore::Evaluator - Creates a custom evaluator for agent quality assessment using LLM-as-a-Judge configurations.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreEvaluatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreEvaluator), &result)
	return &result
}
