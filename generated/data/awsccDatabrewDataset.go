package data

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
                    "type": "string"
                  },
                  "header_row": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "excel": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "header_row": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "sheet_indexes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "sheet_names": {
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
            "json": {
              "computed": true,
              "description": "Json options",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "multi_line": {
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "input": {
        "computed": true,
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
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "Database name",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description": "Table name",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "temp_directory": {
                    "computed": true,
                    "description": "Input location",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bucket_owner": {
                          "computed": true,
                          "description": "Bucket owner",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
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
            "database_input_definition": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database_table_name": {
                    "computed": true,
                    "description": "Database table name",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "glue_connection_name": {
                    "computed": true,
                    "description": "Glue connection name",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "query_string": {
                    "computed": true,
                    "description": "Custom SQL to run against the provided AWS Glue connection. This SQL will be used as the input for DataBrew projects and jobs.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "temp_directory": {
                    "computed": true,
                    "description": "Input location",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bucket_owner": {
                          "computed": true,
                          "description": "Bucket owner",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
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
            "metadata": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source_arn": {
                    "computed": true,
                    "description": "Arn of the source of the dataset. For e.g.: AppFlow Flow ARN.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "s3_input_definition": {
              "computed": true,
              "description": "Input location",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "bucket_owner": {
                    "computed": true,
                    "description": "Bucket owner",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key": {
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
        "description": "Dataset name",
        "description_kind": "plain",
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
                    "computed": true,
                    "description": "Maximum number of files",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "order": {
                    "computed": true,
                    "description": "Order",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ordered_by": {
                    "computed": true,
                    "description": "Ordered by",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "last_modified_date_condition": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "expression": {
                    "computed": true,
                    "description": "Filtering expression for a parameter",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values_map": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value_reference": {
                          "computed": true,
                          "description": "Variable name",
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
            "parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dataset_parameter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "create_column": {
                          "computed": true,
                          "description": "Add the value of this parameter as a column in a dataset.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "datetime_options": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "format": {
                                "computed": true,
                                "description": "Date/time format of a date parameter",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "locale_code": {
                                "computed": true,
                                "description": "Locale code for a date parameter",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "timezone_offset": {
                                "computed": true,
                                "description": "Timezone offset",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "filter": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "expression": {
                                "computed": true,
                                "description": "Filtering expression for a parameter",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "values_map": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "value": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value_reference": {
                                      "computed": true,
                                      "description": "Variable name",
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
                        "name": {
                          "computed": true,
                          "description": "Parameter name",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "Parameter type",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "path_parameter_name": {
                    "computed": true,
                    "description": "Parameter name",
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
      "source": {
        "computed": true,
        "description": "Source type of the dataset",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
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
      }
    },
    "description": "Data Source schema for AWS::DataBrew::Dataset",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatabrewDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatabrewDataset), &result)
	return &result
}
