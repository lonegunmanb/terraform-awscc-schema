package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightTopic = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "config_options": {
        "computed": true,
        "description": "Model for configuration of a Topic",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "q_business_insights_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "data_sets": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "calculated_fields": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aggregation": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "allowed_aggregations": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "calculated_field_description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "calculated_field_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "calculated_field_synonyms": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cell_value_synonyms": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cell_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "synonyms": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "column_data_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "comparative_order": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "specifed_order": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "treat_undefined_specified_values": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "use_ordering": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "default_formatting": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "display_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "display_format_options": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "blank_cell_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "currency_symbol": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "decimal_separator": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "fraction_digits": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "grouping_separator": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "negative_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "prefix": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "suffix": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "suffix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "unit_scaler": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "use_blank_cell_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "use_grouping": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "disable_indexing": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "expression": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "is_included_in_topic": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "never_aggregate_in_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "non_additive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "not_allowed_aggregations": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "semantic_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "falsey_cell_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "falsey_cell_value_synonyms": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "sub_type_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "truthy_cell_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "truthy_cell_value_synonyms": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "type_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type_parameters": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "time_granularity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "columns": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aggregation": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "allowed_aggregations": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cell_value_synonyms": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cell_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "synonyms": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "column_data_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "column_description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "column_friendly_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "column_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "column_synonyms": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "comparative_order": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "specifed_order": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "treat_undefined_specified_values": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "use_ordering": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "default_formatting": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "display_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "display_format_options": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "blank_cell_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "currency_symbol": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "decimal_separator": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "fraction_digits": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "grouping_separator": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "negative_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "prefix": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "suffix": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "suffix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "unit_scaler": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "use_blank_cell_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "use_grouping": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "disable_indexing": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "is_included_in_topic": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "never_aggregate_in_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "non_additive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "not_allowed_aggregations": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "semantic_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "falsey_cell_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "falsey_cell_value_synonyms": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "sub_type_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "truthy_cell_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "truthy_cell_value_synonyms": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "type_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type_parameters": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "time_granularity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "data_aggregation": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dataset_row_date_granularity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "default_date_column_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "dataset_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "dataset_description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "dataset_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "filters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "category_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "category_filter_function": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "category_filter_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "constant": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "collective_constant": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "value_list": {
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
                              },
                              "constant_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "singular_constant": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "inverse": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "date_range_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "constant": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "constant_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "range_constant": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "maximum": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "minimum": {
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
                        "inclusive": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "filter_class": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "filter_description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "filter_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "filter_synonyms": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "filter_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "numeric_equality_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "aggregation": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "constant": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "constant_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "singular_constant": {
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
                  "numeric_range_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "aggregation": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "constant": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "constant_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "range_constant": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "maximum": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "minimum": {
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
                        "inclusive": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "operand_field_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "relative_date_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "constant": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "constant_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "singular_constant": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "relative_date_filter_function": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "time_granularity": {
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
            "named_entities": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "definition": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "aggregation": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "aggregation_function_parameters": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "property_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "property_role": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "property_usage": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "entity_description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "entity_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "entity_synonyms": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "semantic_entity_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "sub_type_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type_parameters": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
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
          "nesting_mode": "list"
        }
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "folder_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "topic_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_experience_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QuickSight::Topic",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightTopic), &result)
	return &result
}
