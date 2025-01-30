package data

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
        "description": "Name for a prompt resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "prompt_arn": {
        "computed": true,
        "description": "ARN of a prompt resource possibly with a version",
        "description_kind": "plain",
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
            "gen_ai_resource": {
              "computed": true,
              "description": "Target resource to invoke with Prompt",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "agent": {
                    "computed": true,
                    "description": "Target Agent to invoke with Prompt",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "agent_identifier": {
                          "computed": true,
                          "description": "Arn representation of the Agent Alias.",
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
                  "chat": {
                    "computed": true,
                    "description": "Configuration for chat prompt template",
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
                        "messages": {
                          "computed": true,
                          "description": "List of messages for chat prompt template",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content": {
                                "computed": true,
                                "description": "List of Content Blocks",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "text": {
                                      "computed": true,
                                      "description": "Configuration for chat prompt template",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "role": {
                                "computed": true,
                                "description": "Conversation roles for the chat prompt",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "system": {
                          "computed": true,
                          "description": "Configuration for chat prompt template",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "text": {
                                "computed": true,
                                "description": "Configuration for chat prompt template",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "tool_configuration": {
                          "computed": true,
                          "description": "Tool configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "tool_choice": {
                                "computed": true,
                                "description": "Tool choice",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "any": {
                                      "computed": true,
                                      "description": "Any Tool choice",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "auto": {
                                      "computed": true,
                                      "description": "Auto Tool choice",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "tool": {
                                      "computed": true,
                                      "description": "Specific Tool choice",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "Tool name",
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
                              "tools": {
                                "computed": true,
                                "description": "List of Tools",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "tool_spec": {
                                      "computed": true,
                                      "description": "Tool specification",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "description": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "input_schema": {
                                            "computed": true,
                                            "description": "Tool input schema",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "json": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "name": {
                                            "computed": true,
                                            "description": "Tool name",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
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
                  },
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
    "description": "Data Source schema for AWS::Bedrock::PromptVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockPromptVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockPromptVersion), &result)
	return &result
}
