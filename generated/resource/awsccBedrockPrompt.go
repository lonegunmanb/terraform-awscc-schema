package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockPrompt = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "ARN of a prompt resource possibly with a version",
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
        "optional": true,
        "type": "string"
      },
      "default_variant": {
        "computed": true,
        "description": "Name for a variant.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Name for a prompt resource.",
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
        "description": "Name for a prompt resource.",
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
            "additional_model_request_fields": {
              "computed": true,
              "description": "Contains model-specific configurations",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
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
                          "optional": true,
                          "type": "number"
                        },
                        "stop_sequences": {
                          "computed": true,
                          "description": "List of stop sequences",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "temperature": {
                          "computed": true,
                          "description": "Controls randomness, higher values increase diversity",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "top_p": {
                          "computed": true,
                          "description": "Cumulative probability cutoff for token selection",
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
            "model_id": {
              "computed": true,
              "description": "ARN or Id of a Bedrock Foundational Model or Inference Profile, or the ARN of a imported model, or a provisioned throughput ARN for custom models.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "Name for a variant.",
              "description_kind": "plain",
              "optional": true,
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
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
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
                                    "cache_point": {
                                      "computed": true,
                                      "description": "CachePointBlock",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "type": {
                                            "computed": true,
                                            "description": "CachePoint types for CachePointBlock",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "text": {
                                      "computed": true,
                                      "description": "Configuration for chat prompt template",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "optional": true
                              },
                              "role": {
                                "computed": true,
                                "description": "Conversation roles for the chat prompt",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "system": {
                          "computed": true,
                          "description": "Configuration for chat prompt template",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "cache_point": {
                                "computed": true,
                                "description": "CachePointBlock",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "type": {
                                      "computed": true,
                                      "description": "CachePoint types for CachePointBlock",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "text": {
                                "computed": true,
                                "description": "Configuration for chat prompt template",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
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
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "auto": {
                                      "computed": true,
                                      "description": "Auto Tool choice",
                                      "description_kind": "plain",
                                      "optional": true,
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
                              "tools": {
                                "computed": true,
                                "description": "List of Tools",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "cache_point": {
                                      "computed": true,
                                      "description": "CachePointBlock",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "type": {
                                            "computed": true,
                                            "description": "CachePoint types for CachePointBlock",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "tool_spec": {
                                      "computed": true,
                                      "description": "Tool specification",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "description": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
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
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "name": {
                                            "computed": true,
                                            "description": "Tool name",
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
                  },
                  "text": {
                    "computed": true,
                    "description": "Configuration for text prompt template",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cache_point": {
                          "computed": true,
                          "description": "CachePointBlock",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type": {
                                "computed": true,
                                "description": "CachePoint types for CachePointBlock",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
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
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "text": {
                          "computed": true,
                          "description": "Prompt content for String prompt template",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "text_s3_location": {
                          "computed": true,
                          "description": "The identifier for the S3 resource.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket": {
                                "computed": true,
                                "description": "A bucket in S3",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key": {
                                "computed": true,
                                "description": "A object key in S3",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "version": {
                                "computed": true,
                                "description": "The version of the the S3 object to use",
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "template_type": {
              "computed": true,
              "description": "Prompt template type",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "version": {
        "computed": true,
        "description": "Draft Version.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::Bedrock::Prompt Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockPromptSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockPrompt), &result)
	return &result
}
