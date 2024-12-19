package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCleanroomsConfiguredTable = `{
  "block": {
    "attributes": {
      "allowed_columns": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "analysis_method": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "analysis_rules": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "v1": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "aggregation": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "additional_analyses": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "aggregate_columns": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "column_names": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "function": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "optional": true
                              },
                              "allowed_join_operators": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "dimension_columns": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "join_columns": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "join_required": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "output_constraints": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "column_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "minimum": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "optional": true
                              },
                              "scalar_functions": {
                                "computed": true,
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
                        "custom": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "additional_analyses": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "allowed_analyses": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allowed_analysis_providers": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "differential_privacy": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "columns": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
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
                              "disallowed_output_columns": {
                                "computed": true,
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
                        "list": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "additional_analyses": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "allowed_join_operators": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "join_columns": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "list_columns": {
                                "computed": true,
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configured_table_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "table_reference": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "athena": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_location": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "work_group": {
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
            "glue": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_name": {
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
            "snowflake": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "schema_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secret_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_schema": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "v1": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "column_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "column_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
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
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Represents a table that can be associated with collaborations",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCleanroomsConfiguredTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCleanroomsConfiguredTable), &result)
	return &result
}
