package data

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
        "type": "string"
      },
      "evaluator_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the evaluator.",
        "description_kind": "plain",
        "type": "string"
      },
      "evaluator_config": {
        "computed": true,
        "description": "The configuration for the evaluator, including LLM-as-a-Judge, code-based, or derived settings.",
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
                          "type": "string"
                        },
                        "lambda_timeout_in_seconds": {
                          "computed": true,
                          "description": "The timeout in seconds for the Lambda function invocation.",
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
            "derived": {
              "computed": true,
              "description": "The configuration for an evaluator derived from an existing base evaluator (a built-in or third-party evaluator), run on your own model. The base evaluator supplies the prompt and scoring logic.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "base_evaluator_id": {
                    "computed": true,
                    "description": "The identifier of the base evaluator whose logic to run (a Builtin.* or ThirdParty.* evaluator).",
                    "description_kind": "plain",
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
                                      "type": "number"
                                    },
                                    "stop_sequences": {
                                      "computed": true,
                                      "description": "The sequences that make the model stop generating further tokens.",
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "temperature": {
                                      "computed": true,
                                      "description": "The temperature value that controls randomness in the model's responses.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "top_p": {
                                      "computed": true,
                                      "description": "The top-p sampling parameter that controls the diversity of the model's responses.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "model_id": {
                                "computed": true,
                                "description": "The identifier of the Amazon Bedrock model to use for evaluation.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "responses_evaluator_model_config": {
                          "computed": true,
                          "description": "The configuration for using OpenResponses-compatible models in evaluator assessments.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max_output_tokens": {
                                "computed": true,
                                "description": "The maximum number of output tokens to generate, including visible output and reasoning tokens.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "model_id": {
                                "computed": true,
                                "description": "The identifier of the model to use for evaluation.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "reasoning": {
                                "computed": true,
                                "description": "The reasoning configuration for reasoning models.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "effort": {
                                      "computed": true,
                                      "description": "The level of reasoning effort the model applies.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "temperature": {
                                "computed": true,
                                "description": "The sampling temperature between 0 and 2.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "top_p": {
                                "computed": true,
                                "description": "The nucleus sampling probability mass between 0 and 1.",
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
                  }
                },
                "nesting_mode": "single"
              }
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
                                      "type": "number"
                                    },
                                    "stop_sequences": {
                                      "computed": true,
                                      "description": "The sequences that make the model stop generating further tokens.",
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "temperature": {
                                      "computed": true,
                                      "description": "The temperature value that controls randomness in the model's responses.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "top_p": {
                                      "computed": true,
                                      "description": "The top-p sampling parameter that controls the diversity of the model's responses.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "model_id": {
                                "computed": true,
                                "description": "The identifier of the Amazon Bedrock model to use for evaluation.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "responses_evaluator_model_config": {
                          "computed": true,
                          "description": "The configuration for using OpenResponses-compatible models in evaluator assessments.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max_output_tokens": {
                                "computed": true,
                                "description": "The maximum number of output tokens to generate, including visible output and reasoning tokens.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "model_id": {
                                "computed": true,
                                "description": "The identifier of the model to use for evaluation.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "reasoning": {
                                "computed": true,
                                "description": "The reasoning configuration for reasoning models.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "effort": {
                                      "computed": true,
                                      "description": "The level of reasoning effort the model applies.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "temperature": {
                                "computed": true,
                                "description": "The sampling temperature between 0 and 2.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "top_p": {
                                "computed": true,
                                "description": "The nucleus sampling probability mass between 0 and 1.",
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
                                "type": "string"
                              },
                              "label": {
                                "computed": true,
                                "description": "The label of this categorical rating option.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "multi_dimensional": {
                          "computed": true,
                          "description": "The dimensions an LLM-as-a-Judge evaluator scores in a single judge call. Each dimension has its own instruction and rating scale.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "dimension_instruction": {
                                "computed": true,
                                "description": "The instruction to the evaluator model describing what to judge for this dimension.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "dimension_name": {
                                "computed": true,
                                "description": "The name of the dimension. Must be unique within the evaluator.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "dimension_rating_scale": {
                                "computed": true,
                                "description": "The rating scale for one dimension of an LLM-as-a-Judge evaluator.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "categorical": {
                                      "computed": true,
                                      "description": "The categorical rating scale for this dimension.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "definition": {
                                            "computed": true,
                                            "description": "The description that explains what this categorical rating represents.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "label": {
                                            "computed": true,
                                            "description": "The label of this categorical rating option.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      }
                                    },
                                    "numerical": {
                                      "computed": true,
                                      "description": "The numerical rating scale for this dimension.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "definition": {
                                            "computed": true,
                                            "description": "The description that explains what this numerical rating represents.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "label": {
                                            "computed": true,
                                            "description": "The label that describes this numerical rating option.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description": "The numerical value for this rating scale option.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      }
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              }
                            },
                            "nesting_mode": "list"
                          }
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
                                "type": "string"
                              },
                              "label": {
                                "computed": true,
                                "description": "The label that describes this numerical rating option.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "The numerical value for this rating scale option.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "evaluator_id": {
        "computed": true,
        "description": "The unique identifier of the evaluator.",
        "description_kind": "plain",
        "type": "string"
      },
      "evaluator_name": {
        "computed": true,
        "description": "The name of the evaluator. Must be unique within your account.",
        "description_kind": "plain",
        "type": "string"
      },
      "evaluator_type": {
        "computed": true,
        "description": "The kind of evaluator resource. Builtin and ThirdParty are AWS-managed global evaluators that you reference but don't create; Custom, CustomCode, and CustomDerived are evaluators that you create.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the KMS key used to encrypt evaluator data.",
        "description_kind": "plain",
        "type": "string"
      },
      "level": {
        "computed": true,
        "description": "The evaluation level that determines the scope of evaluation. Required for all evaluator configuration types. For Derived evaluators it must match the level of the base evaluator identified by BaseEvaluatorId.",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_name": {
        "computed": true,
        "description": "Where the evaluator's logic comes from: AWS, a third-party library (DeepEval or AutoEval), or you. Derived evaluators inherit the provider of their base evaluator.",
        "description_kind": "plain",
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
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the evaluator was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::BedrockAgentCore::Evaluator",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcoreEvaluatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreEvaluator), &result)
	return &result
}
