package data

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
        "type": "string"
      },
      "encryption_key_arn": {
        "computed": true,
        "description": "ARN format",
        "description_kind": "plain",
        "type": "string"
      },
      "event_expiry_duration": {
        "computed": true,
        "description": "Duration in days until memory events expire",
        "description_kind": "plain",
        "type": "number"
      },
      "failure_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
                                      "type": "string"
                                    },
                                    "model_id": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
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
                                      "type": "string"
                                    },
                                    "model_id": {
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
                                      "type": "string"
                                    },
                                    "model_id": {
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
                                      "type": "string"
                                    },
                                    "model_id": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
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
                                      "type": "string"
                                    },
                                    "model_id": {
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
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "created_at": {
                    "computed": true,
                    "description": "Creation timestamp of the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Description of the Memory resource",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of the Memory resource",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "namespaces": {
                    "computed": true,
                    "description": "List of namespaces for memory strategy",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "status": {
                    "computed": true,
                    "description": "Status of the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "strategy_id": {
                    "computed": true,
                    "description": "Unique identifier for the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "updated_at": {
                    "computed": true,
                    "description": "Last update timestamp of the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
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
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Description of the Memory resource",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of the Memory resource",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "namespaces": {
                    "computed": true,
                    "description": "List of namespaces for memory strategy",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "status": {
                    "computed": true,
                    "description": "Status of the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "strategy_id": {
                    "computed": true,
                    "description": "Unique identifier for the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "updated_at": {
                    "computed": true,
                    "description": "Last update timestamp of the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
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
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Description of the Memory resource",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of the Memory resource",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "namespaces": {
                    "computed": true,
                    "description": "List of namespaces for memory strategy",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "status": {
                    "computed": true,
                    "description": "Status of the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "strategy_id": {
                    "computed": true,
                    "description": "Unique identifier for the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "updated_at": {
                    "computed": true,
                    "description": "Last update timestamp of the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
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
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Description of the Memory resource",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of the Memory resource",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "namespaces": {
                    "computed": true,
                    "description": "List of namespaces for memory strategy",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "status": {
                    "computed": true,
                    "description": "Status of the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "strategy_id": {
                    "computed": true,
                    "description": "Unique identifier for the memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of memory strategy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "updated_at": {
                    "computed": true,
                    "description": "Last update timestamp of the memory strategy",
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
      "name": {
        "computed": true,
        "description": "Name of the Memory resource",
        "description_kind": "plain",
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
    "description": "Data Source schema for AWS::BedrockAgentCore::Memory",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcoreMemorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreMemory), &result)
	return &result
}
