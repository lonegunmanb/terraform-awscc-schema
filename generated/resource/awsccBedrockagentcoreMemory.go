package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreMemory = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the Memory resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_key_arn": {
        "computed": true,
        "description": "ARN format",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_expiry_duration": {
        "description": "Duration in days until memory events expire",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "failure_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "memory_arn": {
        "computed": true,
        "description": "ARN of the Memory resource",
        "description_kind": "plain",
        "type": "string"
      },
      "memory_execution_role_arn": {
        "computed": true,
        "description": "ARN format",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "memory_id": {
        "computed": true,
        "description": "Unique identifier for the Memory resource",
        "description_kind": "plain",
        "type": "string"
      },
      "memory_strategies": {
        "computed": true,
        "description": "List of memory strategies attached to this memory",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_memory_strategy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "episodic_override": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "consolidation": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "append_to_prompt": {
                                      "computed": true,
                                      "description": "Text prompt for model instructions",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "model_id": {
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
                              "extraction": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "append_to_prompt": {
                                      "computed": true,
                                      "description": "Text prompt for model instructions",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "model_id": {
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
                              "reflection": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "append_to_prompt": {
                                      "computed": true,
                                      "description": "Text prompt for model instructions",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "model_id": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "namespaces": {
                                      "computed": true,
                                      "description": "List of namespaces for memory strategy",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
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
                        "self_managed_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "historical_context_window_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "invocation_configuration": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "payload_delivery_bucket_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "topic_arn": {
                                      "computed": true,
                                      "description": "ARN format",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "trigger_conditions": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "message_based_trigger": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "message_count": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "time_based_trigger": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "idle_session_timeout": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "token_based_trigger": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "token_count": {
                                            "computed": true,
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
                                  "nesting_mode": "list"
                                },
                                "optional": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "semantic_override": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "consolidation": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "append_to_prompt": {
                                      "computed": true,
                                      "description": "Text prompt for model instructions",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "model_id": {
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
                              "extraction": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "append_to_prompt": {
                                      "computed": true,
                                      "description": "Text prompt for model instructions",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "model_id": {
                                      "computed": true,
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
                        "summary_override": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "consolidation": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "append_to_prompt": {
                                      "computed": true,
                                      "description": "Text prompt for model instructions",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "model_id": {
                                      "computed": true,
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
                        "user_preference_override": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "consolidation": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "append_to_prompt": {
                                      "computed": true,
                                      "description": "Text prompt for model instructions",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "model_id": {
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
                              "extraction": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "append_to_prompt": {
                                      "computed": true,
                                      "description": "Text prompt for model instructions",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "model_id": {
                                      "computed": true,
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
                  "created_at": {
                    "computed": true,
                    "description": "Creation timestamp of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Description of the Memory resource",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of the Memory resource",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "namespaces": {
                    "computed": true,
                    "description": "List of namespaces for memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "status": {
                    "computed": true,
                    "description": "Status of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "strategy_id": {
                    "computed": true,
                    "description": "Unique identifier for the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "updated_at": {
                    "computed": true,
                    "description": "Last update timestamp of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "episodic_memory_strategy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "created_at": {
                    "computed": true,
                    "description": "Creation timestamp of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Description of the Memory resource",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of the Memory resource",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "namespaces": {
                    "computed": true,
                    "description": "List of namespaces for memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "reflection_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "namespaces": {
                          "computed": true,
                          "description": "List of namespaces for memory strategy",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "status": {
                    "computed": true,
                    "description": "Status of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "strategy_id": {
                    "computed": true,
                    "description": "Unique identifier for the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "updated_at": {
                    "computed": true,
                    "description": "Last update timestamp of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "semantic_memory_strategy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "created_at": {
                    "computed": true,
                    "description": "Creation timestamp of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Description of the Memory resource",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of the Memory resource",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "namespaces": {
                    "computed": true,
                    "description": "List of namespaces for memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "status": {
                    "computed": true,
                    "description": "Status of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "strategy_id": {
                    "computed": true,
                    "description": "Unique identifier for the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "updated_at": {
                    "computed": true,
                    "description": "Last update timestamp of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "summary_memory_strategy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "created_at": {
                    "computed": true,
                    "description": "Creation timestamp of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Description of the Memory resource",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of the Memory resource",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "namespaces": {
                    "computed": true,
                    "description": "List of namespaces for memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "status": {
                    "computed": true,
                    "description": "Status of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "strategy_id": {
                    "computed": true,
                    "description": "Unique identifier for the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "updated_at": {
                    "computed": true,
                    "description": "Last update timestamp of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "user_preference_memory_strategy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "created_at": {
                    "computed": true,
                    "description": "Creation timestamp of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Description of the Memory resource",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of the Memory resource",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "namespaces": {
                    "computed": true,
                    "description": "List of namespaces for memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "status": {
                    "computed": true,
                    "description": "Status of the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "strategy_id": {
                    "computed": true,
                    "description": "Unique identifier for the memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of memory strategy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "updated_at": {
                    "computed": true,
                    "description": "Last update timestamp of the memory strategy",
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
      },
      "name": {
        "description": "Name of the Memory resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Status of the Memory resource",
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
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::BedrockAgentCore::Memory",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreMemorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreMemory), &result)
	return &result
}
