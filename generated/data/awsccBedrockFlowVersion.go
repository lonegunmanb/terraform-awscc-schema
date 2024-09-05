package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockFlowVersion = `{
  "block": {
    "attributes": {
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
      "definition": {
        "computed": true,
        "description": "Flow definition",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connections": {
              "computed": true,
              "description": "List of connections",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "configuration": {
                    "computed": true,
                    "description": "Connection configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "conditional": {
                          "computed": true,
                          "description": "Conditional connection configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "condition": {
                                "computed": true,
                                "description": "Name of a condition in a flow",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "data": {
                          "computed": true,
                          "description": "Data connection configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "source_output": {
                                "computed": true,
                                "description": "Name of a node output in a flow",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "target_input": {
                                "computed": true,
                                "description": "Name of a node input in a flow",
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
                    "description": "Name of a connection in a flow",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source": {
                    "computed": true,
                    "description": "Name of a node in a flow",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target": {
                    "computed": true,
                    "description": "Name of a node in a flow",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Connection type",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "nodes": {
              "computed": true,
              "description": "List of nodes in a flow",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "configuration": {
                    "computed": true,
                    "description": "Node configuration in a flow",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "agent": {
                          "computed": true,
                          "description": "Agent flow node configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "agent_alias_arn": {
                                "computed": true,
                                "description": "Arn representation of the Agent Alias.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "collector": {
                          "computed": true,
                          "description": "Collector flow node configuration",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "condition": {
                          "computed": true,
                          "description": "Condition flow node configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "conditions": {
                                "computed": true,
                                "description": "List of conditions in a condition node",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "expression": {
                                      "computed": true,
                                      "description": "Expression for a condition in a flow",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "name": {
                                      "computed": true,
                                      "description": "Name of a condition in a flow",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "input": {
                          "computed": true,
                          "description": "Input flow node configuration",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "iterator": {
                          "computed": true,
                          "description": "Iterator flow node configuration",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "knowledge_base": {
                          "computed": true,
                          "description": "Knowledge base flow node configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "knowledge_base_id": {
                                "computed": true,
                                "description": "Identifier of the KnowledgeBase",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "model_id": {
                                "computed": true,
                                "description": "ARN or name of a Bedrock model.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "lambda_function": {
                          "computed": true,
                          "description": "Lambda function flow node configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "lambda_arn": {
                                "computed": true,
                                "description": "ARN of a Lambda.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "lex": {
                          "computed": true,
                          "description": "Lex flow node configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bot_alias_arn": {
                                "computed": true,
                                "description": "ARN of a Lex bot alias",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "locale_id": {
                                "computed": true,
                                "description": "Lex bot locale id",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "output": {
                          "computed": true,
                          "description": "Output flow node configuration",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "prompt": {
                          "computed": true,
                          "description": "Prompt flow node configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "source_configuration": {
                                "computed": true,
                                "description": "Prompt source configuration for prompt node",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "inline": {
                                      "computed": true,
                                      "description": "Inline prompt configuration for prompt node",
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
                                                      "top_k": {
                                                        "computed": true,
                                                        "description": "Sample from the k most likely next tokens",
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
                                            "description": "ARN or name of a Bedrock model.",
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
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "resource": {
                                      "computed": true,
                                      "description": "Resource prompt configuration for prompt node",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "prompt_arn": {
                                            "computed": true,
                                            "description": "ARN of a prompt resource possibly with a version",
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
                        "retrieval": {
                          "computed": true,
                          "description": "Retrieval flow node configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "service_configuration": {
                                "computed": true,
                                "description": "Retrieval service configuration for Retrieval node",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "s3": {
                                      "computed": true,
                                      "description": "s3 Retrieval configuration for Retrieval node",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "bucket_name": {
                                            "computed": true,
                                            "description": "bucket name of an s3 that will be used for Retrieval flow node configuration",
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
                        "storage": {
                          "computed": true,
                          "description": "Storage flow node configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "service_configuration": {
                                "computed": true,
                                "description": "storage service configuration for storage node",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "s3": {
                                      "computed": true,
                                      "description": "s3 storage configuration for storage node",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "bucket_name": {
                                            "computed": true,
                                            "description": "bucket name of an s3 that will be used for storage flow node configuration",
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
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "inputs": {
                    "computed": true,
                    "description": "List of node inputs in a flow",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "expression": {
                          "computed": true,
                          "description": "Expression for a node input in a flow",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description": "Name of a node input in a flow",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "Type of input/output for a node in a flow",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of a node in a flow",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "outputs": {
                    "computed": true,
                    "description": "List of node outputs in a flow",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "Name of a node output in a flow",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "Type of input/output for a node in a flow",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "type": {
                    "computed": true,
                    "description": "Flow node types",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "Description of the flow version",
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "ARN of a IAM role",
        "description_kind": "plain",
        "type": "string"
      },
      "flow_arn": {
        "computed": true,
        "description": "Arn representation of the Flow",
        "description_kind": "plain",
        "type": "string"
      },
      "flow_id": {
        "computed": true,
        "description": "Identifier for a Flow",
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
        "description": "Name for the flow",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Schema Type for Flow APIs",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "Numerical Version.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Bedrock::FlowVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockFlowVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockFlowVersion), &result)
	return &result
}
