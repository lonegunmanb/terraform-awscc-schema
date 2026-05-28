package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreHarness = `{
  "block": {
    "attributes": {
      "allowed_tools": {
        "computed": true,
        "description": "The tools that the agent is allowed to use.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the harness.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_configuration": {
        "computed": true,
        "description": "The inbound authorization configuration for authenticating incoming requests.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_jwt_authorizer": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_audience": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_clients": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_scopes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "custom_claims": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "authorizing_claim_match_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "claim_match_operator": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "claim_match_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "match_value_string": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "match_value_string_list": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "inbound_token_claim_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "inbound_token_claim_value_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "discovery_url": {
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
      "created_at": {
        "computed": true,
        "description": "The timestamp when the harness was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment": {
        "computed": true,
        "description": "The compute environment configuration for the harness, including underlying runtime information.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "agent_core_runtime_environment": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "agent_runtime_arn": {
                    "computed": true,
                    "description": "The ARN of the underlying AgentCore Runtime.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "agent_runtime_id": {
                    "computed": true,
                    "description": "The ID of the underlying AgentCore Runtime.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "agent_runtime_name": {
                    "computed": true,
                    "description": "The name of the underlying AgentCore Runtime.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "filesystem_configurations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "session_storage": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "mount_path": {
                                "computed": true,
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
                  },
                  "lifecycle_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "idle_runtime_session_timeout": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "max_lifetime": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "network_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "network_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "network_mode_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "security_groups": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnets": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "environment_artifact": {
        "computed": true,
        "description": "The environment artifact for the harness, such as a custom container image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_uri": {
                    "computed": true,
                    "description": "The ECR URI of the container.",
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
      "environment_variables": {
        "computed": true,
        "description": "Environment variables to set in the harness runtime environment.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role that the harness assumes when running.",
        "description_kind": "plain",
        "type": "string"
      },
      "harness_id": {
        "computed": true,
        "description": "The unique identifier of the harness.",
        "description_kind": "plain",
        "type": "string"
      },
      "harness_name": {
        "computed": true,
        "description": "The name of the harness.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_iterations": {
        "computed": true,
        "description": "The maximum number of iterations the agent loop can execute per invocation.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_tokens": {
        "computed": true,
        "description": "The maximum number of tokens the agent can generate per iteration.",
        "description_kind": "plain",
        "type": "number"
      },
      "memory": {
        "computed": true,
        "description": "The AgentCore Memory configuration for persisting conversation context.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "agent_core_memory_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "actor_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "messages_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "retrieval_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "relevance_score": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "strategy_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "top_k": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "map"
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
      "model": {
        "computed": true,
        "description": "The model configuration for the harness.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bedrock_model_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_tokens": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "model_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "temperature": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "top_p": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "gemini_model_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "api_key_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "max_tokens": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "model_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "temperature": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "top_k": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "top_p": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "open_ai_model_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "api_key_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "max_tokens": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "model_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "temperature": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "top_p": {
                    "computed": true,
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
      "skills": {
        "computed": true,
        "description": "The skills available to the agent.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "path": {
              "computed": true,
              "description": "The filesystem path to the skill definition.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "status": {
        "computed": true,
        "description": "The current status of the harness.",
        "description_kind": "plain",
        "type": "string"
      },
      "system_prompt": {
        "computed": true,
        "description": "The system prompt that defines the agent's behavior.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "text": {
              "computed": true,
              "description": "The text content of the system prompt block.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "Tags to apply to the harness resource.",
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
      "timeout_seconds": {
        "computed": true,
        "description": "The maximum duration in seconds for the agent loop execution per invocation.",
        "description_kind": "plain",
        "type": "number"
      },
      "tools": {
        "computed": true,
        "description": "The tools available to the agent.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "agent_core_browser": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "browser_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "agent_core_code_interpreter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "code_interpreter_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "agent_core_gateway": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "gateway_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "outbound_auth": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "aws_iam": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "none": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "oauth": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "custom_parameters": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    },
                                    "default_return_url": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "grant_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "provider_arn": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "scopes": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
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
                  "inline_function": {
                    "computed": true,
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
                          "description": "JSON Schema describing the tool's input parameters.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "remote_mcp": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "headers": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "url": {
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
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "truncation": {
        "computed": true,
        "description": "The truncation configuration for managing conversation context.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "sliding_window": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "messages_count": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "summarization": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "preserve_recent_messages": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "summarization_system_prompt": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "summary_ratio": {
                          "computed": true,
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
            "strategy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the harness was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::BedrockAgentCore::Harness",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcoreHarnessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreHarness), &result)
	return &result
}
