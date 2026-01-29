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
            "case_summarization_ai_agent_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "case_summarization_ai_guardrail_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "case_summarization_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
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
            "email_generative_answer_ai_agent_configuration": {
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
                  "email_generative_answer_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "email_query_reformulation_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
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
            "email_overview_ai_agent_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "email_overview_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
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
            "email_response_ai_agent_configuration": {
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
                  "email_query_reformulation_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "email_response_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
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
            "note_taking_ai_agent_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "locale": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "note_taking_ai_guardrail_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "note_taking_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "orchestration_ai_agent_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connect_instance_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "locale": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "orchestration_ai_guardrail_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "orchestration_ai_prompt_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tool_configurations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "annotations": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "description": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "input_schema": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "instruction": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "examples": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "instruction": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "output_filters": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "json_path": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "output_configuration": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "output_variable_name_override": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "session_data_namespace": {
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
                        "output_schema": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "override_input_values": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "json_path": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "constant": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "type": {
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
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "title": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "tool_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "tool_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "tool_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "user_interaction_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "is_user_confirmation_required": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
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
