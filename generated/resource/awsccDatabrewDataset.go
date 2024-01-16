package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatabrewDataset = `{
  "block": {
    "attributes": {
      "format": {
        "computed": true,
        "description": "Dataset format",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "format_options": {
        "computed": true,
        "description": "Format options for dataset",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "csv": {
              "computed": true,
              "description": "Csv options",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delimiter": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "header_row": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "excel": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "header_row": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "sheet_indexes": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "sheet_names": {
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
            "json": {
              "computed": true,
              "description": "Json options",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "multi_line": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "input": {
        "description": "Input",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_catalog_input_definition": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description": "Catalog id",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "Database name",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description": "Table name",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "temp_directory": {
                    "computed": true,
                    "description": "Input location",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "key": {
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
            "database_input_definition": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database_table_name": {
                    "computed": true,
                    "description": "Database table name",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "glue_connection_name": {
                    "description": "Glue connection name",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "query_string": {
                    "computed": true,
                    "description": "Custom SQL to run against the provided AWS Glue connection. This SQL will be used as the input for DataBrew projects and jobs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "temp_directory": {
                    "computed": true,
                    "description": "Input location",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "key": {
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
            "metadata": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source_arn": {
                    "computed": true,
                    "description": "Arn of the source of the dataset. For e.g.: AppFlow Flow ARN.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "s3_input_definition": {
              "computed": true,
              "description": "Input location",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
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
        "required": true
      },
      "name": {
        "description": "Dataset name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "path_options": {
        "computed": true,
        "description": "PathOptions",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "files_limit": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_files": {
                    "description": "Maximum number of files",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "order": {
                    "computed": true,
                    "description": "Order",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ordered_by": {
                    "computed": true,
                    "description": "Ordered by",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "last_modified_date_condition": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "expression": {
                    "description": "Filtering expression for a parameter",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values_map": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value_reference": {
                          "description": "Variable name",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dataset_parameter": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "create_column": {
                          "computed": true,
                          "description": "Add the value of this parameter as a column in a dataset.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "datetime_options": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "format": {
                                "description": "Date/time format of a date parameter",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "locale_code": {
                                "computed": true,
                                "description": "Locale code for a date parameter",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "timezone_offset": {
                                "computed": true,
                                "description": "Timezone offset",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "filter": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "expression": {
                                "description": "Filtering expression for a parameter",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "values_map": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "value": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value_reference": {
                                      "description": "Variable name",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "name": {
                          "description": "Parameter name",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "Parameter type",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "path_parameter_name": {
                    "description": "Parameter name",
                    "description_kind": "plain",
                    "required": true,
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::DataBrew::Dataset.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatabrewDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatabrewDataset), &result)
	return &result
}
