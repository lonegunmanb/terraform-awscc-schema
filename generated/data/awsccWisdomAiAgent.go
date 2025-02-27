package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWisdomAiAgent = `{
  "block": {
    "attributes": {
      "ai_agent_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ai_agent_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assistant_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assistant_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "answer_recommendation_ai_agent_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "answer_generation_ai_guardrail_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "answer_generation_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "association_configurations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "association_configuration_data": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "knowledge_base_association_configuration_data": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "content_tag_filter": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "and_conditions": {
                                            "computed": true,
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
                                          "or_conditions": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "and_conditions": {
                                                  "computed": true,
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
                                                "tag_condition": {
                                                  "computed": true,
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
                                                    "nesting_mode": "single"
                                                  }
                                                }
                                              },
                                              "nesting_mode": "list"
                                            }
                                          },
                                          "tag_condition": {
                                            "computed": true,
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
                                              "nesting_mode": "single"
                                            }
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "max_results": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "override_knowledge_base_search_type": {
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
                        "association_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "association_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "intent_labeling_generation_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "locale": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "query_reformulation_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "manual_search_ai_agent_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "answer_generation_ai_guardrail_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "answer_generation_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "association_configurations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "association_configuration_data": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "knowledge_base_association_configuration_data": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "content_tag_filter": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "and_conditions": {
                                            "computed": true,
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
                                          "or_conditions": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "and_conditions": {
                                                  "computed": true,
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
                                                "tag_condition": {
                                                  "computed": true,
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
                                                    "nesting_mode": "single"
                                                  }
                                                }
                                              },
                                              "nesting_mode": "list"
                                            }
                                          },
                                          "tag_condition": {
                                            "computed": true,
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
                                              "nesting_mode": "single"
                                            }
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "max_results": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "override_knowledge_base_search_type": {
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
                        "association_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "association_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "locale": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "self_service_ai_agent_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "association_configurations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "association_configuration_data": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "knowledge_base_association_configuration_data": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "content_tag_filter": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "and_conditions": {
                                            "computed": true,
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
                                          "or_conditions": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "and_conditions": {
                                                  "computed": true,
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
                                                "tag_condition": {
                                                  "computed": true,
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
                                                    "nesting_mode": "single"
                                                  }
                                                }
                                              },
                                              "nesting_mode": "list"
                                            }
                                          },
                                          "tag_condition": {
                                            "computed": true,
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
                                              "nesting_mode": "single"
                                            }
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "max_results": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "override_knowledge_base_search_type": {
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
                        "association_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "association_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "self_service_ai_guardrail_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "self_service_answer_generation_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "self_service_pre_processing_ai_prompt_id": {
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
      "description": {
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
      "modified_time_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Wisdom::AIAgent",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWisdomAiAgentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWisdomAiAgent), &result)
	return &result
}
