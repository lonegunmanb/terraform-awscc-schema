package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightDataSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "column_groups": {
        "computed": true,
        "description": "\u003cp\u003eGroupings of columns that work together in certain QuickSight features. Currently, only geospatial hierarchy is supported.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "geo_spatial_column_group": {
              "computed": true,
              "description": "\u003cp\u003eGeospatial column group that denotes a hierarchy.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "columns": {
                    "computed": true,
                    "description": "\u003cp\u003eColumns in this hierarchy.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "country_code": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "\u003cp\u003eA display name for the hierarchy.\u003c/p\u003e",
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
      "column_level_permission_rules": {
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
            "principals": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "consumed_spice_capacity_in_bytes": {
        "computed": true,
        "description": "\u003cp\u003eThe amount of SPICE capacity used by this dataset. This is 0 if the dataset isn't\n            imported into SPICE.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "number"
      },
      "created_time": {
        "computed": true,
        "description": "\u003cp\u003eThe time that this dataset was created.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "data_set_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_set_usage_configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe dataset usage configuration for the dataset.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "disable_use_as_direct_query_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disable_use_as_imported_source": {
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
      "field_folders": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "columns": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "map"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "import_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ingestion_wait_policy": {
        "computed": true,
        "description": "\u003cp\u003eWait policy to use when creating/updating dataset. Default is to wait for SPICE ingestion to finish with timeout of 36 hours.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ingestion_wait_time_in_hours": {
              "computed": true,
              "description": "\u003cp\u003eThe maximum time (in hours) to wait for Ingestion to complete. Default timeout is 36 hours.\n Applicable only when DataSetImportMode mode is set to SPICE and WaitForSpiceIngestion is set to true.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "wait_for_spice_ingestion": {
              "computed": true,
              "description": "\u003cp\u003eWait for SPICE ingestion to finish to mark dataset creation/update successful. Default (true).\n  Applicable only when DataSetImportMode mode is set to SPICE.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "last_updated_time": {
        "computed": true,
        "description": "\u003cp\u003eThe last time that this dataset was updated.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "logical_table_map": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alias": {
              "computed": true,
              "description": "\u003cp\u003eA display name for the logical table.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_transforms": {
              "computed": true,
              "description": "\u003cp\u003eTransform operations that act on this logical table.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cast_column_type_operation": {
                    "computed": true,
                    "description": "\u003cp\u003eA transform operation that casts a column to a different type.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "column_name": {
                          "computed": true,
                          "description": "\u003cp\u003eColumn name.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "format": {
                          "computed": true,
                          "description": "\u003cp\u003eWhen casting a column from string to datetime type, you can supply a string in a\n            format supported by Amazon QuickSight to denote the source data format.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "new_column_type": {
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
                  "create_columns_operation": {
                    "computed": true,
                    "description": "\u003cp\u003eA transform operation that creates calculated columns. Columns created in one such\n            operation form a lexical closure.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "columns": {
                          "computed": true,
                          "description": "\u003cp\u003eCalculated columns to create.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "column_id": {
                                "computed": true,
                                "description": "\u003cp\u003eA unique ID to identify a calculated column. During a dataset update, if the column ID\n            of a calculated column matches that of an existing calculated column, Amazon QuickSight\n            preserves the existing calculated column.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "column_name": {
                                "computed": true,
                                "description": "\u003cp\u003eColumn name.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "expression": {
                                "computed": true,
                                "description": "\u003cp\u003eAn expression that defines the calculated column.\u003c/p\u003e",
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
                  "filter_operation": {
                    "computed": true,
                    "description": "\u003cp\u003eA transform operation that filters rows based on a condition.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "condition_expression": {
                          "computed": true,
                          "description": "\u003cp\u003eAn expression that must evaluate to a Boolean value. Rows for which the expression\n            evaluates to true are kept in the dataset.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "project_operation": {
                    "computed": true,
                    "description": "\u003cp\u003eA transform operation that projects columns. Operations that come after a projection\n            can only refer to projected columns.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "projected_columns": {
                          "computed": true,
                          "description": "\u003cp\u003eProjected columns.\u003c/p\u003e",
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
                  "rename_column_operation": {
                    "computed": true,
                    "description": "\u003cp\u003eA transform operation that renames a column.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "column_name": {
                          "computed": true,
                          "description": "\u003cp\u003eThe name of the column to be renamed.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "new_column_name": {
                          "computed": true,
                          "description": "\u003cp\u003eThe new name for the column.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "tag_column_operation": {
                    "computed": true,
                    "description": "\u003cp\u003eA transform operation that tags a column with additional information.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "column_name": {
                          "computed": true,
                          "description": "\u003cp\u003eThe column that this operation acts on.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "tags": {
                          "computed": true,
                          "description": "\u003cp\u003eThe dataset column tag, currently only used for geospatial type tagging. .\u003c/p\u003e\n        \u003cnote\u003e\n            \u003cp\u003eThis is not tags for the AWS tagging feature. .\u003c/p\u003e\n        \u003c/note\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "column_description": {
                                "computed": true,
                                "description": "\u003cp\u003eMetadata that contains a description for a column.\u003c/p\u003e",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "text": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe text of a description for a column.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "column_geographic_role": {
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
                "nesting_mode": "list"
              },
              "optional": true
            },
            "source": {
              "computed": true,
              "description": "\u003cp\u003eInformation about the source of a logical table. This is a variant type structure. For\n            this structure to be valid, only one of the attributes can be non-null.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_set_arn": {
                    "computed": true,
                    "description": "\u003cp\u003eThe Amazon Resource Name (ARN) for the dataset.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "join_instruction": {
                    "computed": true,
                    "description": "\u003cp\u003eJoin instruction.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "left_join_key_properties": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "unique_key": {
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
                        "left_operand": {
                          "computed": true,
                          "description": "\u003cp\u003eLeft operand.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "on_clause": {
                          "computed": true,
                          "description": "\u003cp\u003eOn Clause.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "right_join_key_properties": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "unique_key": {
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
                        "right_operand": {
                          "computed": true,
                          "description": "\u003cp\u003eRight operand.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
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
                  "physical_table_id": {
                    "computed": true,
                    "description": "\u003cp\u003ePhysical table ID.\u003c/p\u003e",
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
          "nesting_mode": "map"
        },
        "optional": true
      },
      "name": {
        "computed": true,
        "description": "\u003cp\u003eThe display name for the dataset.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_columns": {
        "computed": true,
        "description": "\u003cp\u003eThe list of columns after all transforms. These columns are available in templates,\n            analyses, and dashboards.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "\u003cp\u003eA description for a column.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "\u003cp\u003eA display name for the dataset.\u003c/p\u003e",
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
      "permissions": {
        "computed": true,
        "description": "\u003cp\u003eA list of resource permissions on the dataset.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description": "\u003cp\u003eThe IAM action to grant or revoke permissions on.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
              "computed": true,
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:\u003c/p\u003e\n        \u003cul\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) \u003c/p\u003e\n            \u003c/li\u003e\n         \u003c/ul\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "physical_table_map": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_sql": {
              "computed": true,
              "description": "\u003cp\u003eA physical table type built from the results of the custom SQL query.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "columns": {
                    "computed": true,
                    "description": "\u003cp\u003eThe column schema from the SQL query result set.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "\u003cp\u003eThe name of this column in the underlying data source.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
                  "data_source_arn": {
                    "computed": true,
                    "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "\u003cp\u003eA display name for the SQL query result.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sql_query": {
                    "computed": true,
                    "description": "\u003cp\u003eThe SQL query.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "relational_table": {
              "computed": true,
              "description": "\u003cp\u003eA physical table type for relational data sources.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog": {
                    "computed": true,
                    "description": "\u003cp\u003eThe catalog associated with a table.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_source_arn": {
                    "computed": true,
                    "description": "\u003cp\u003eThe Amazon Resource Name (ARN) for the data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "input_columns": {
                    "computed": true,
                    "description": "\u003cp\u003eThe column schema of the table.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "\u003cp\u003eThe name of this column in the underlying data source.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
                  "name": {
                    "computed": true,
                    "description": "\u003cp\u003eThe name of the relational table.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "schema": {
                    "computed": true,
                    "description": "\u003cp\u003eThe schema name. This name applies to certain relational database engines.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "s3_source": {
              "computed": true,
              "description": "\u003cp\u003eA physical table type for as S3 data source.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_source_arn": {
                    "computed": true,
                    "description": "\u003cp\u003eThe amazon Resource Name (ARN) for the data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "input_columns": {
                    "computed": true,
                    "description": "\u003cp\u003eA physical table type for as S3 data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "\u003cp\u003eThe name of this column in the underlying data source.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
                  "upload_settings": {
                    "computed": true,
                    "description": "\u003cp\u003eInformation about the format for a source file or files.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "contains_header": {
                          "computed": true,
                          "description": "\u003cp\u003eWhether the file has a header row, or the files each have a header row.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "delimiter": {
                          "computed": true,
                          "description": "\u003cp\u003eThe delimiter between values in the file.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "format": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "start_from_row": {
                          "computed": true,
                          "description": "\u003cp\u003eA row number to start reading data from.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "text_qualifier": {
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
          "nesting_mode": "map"
        },
        "optional": true
      },
      "row_level_permission_data_set": {
        "computed": true,
        "description": "\u003cp\u003eThe row-level security configuration for the dataset.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the permission dataset.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "format_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "computed": true,
              "description": "\u003cp\u003eThe namespace associated with the row-level permissions dataset.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "permission_policy": {
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
      "tags": {
        "computed": true,
        "description": "\u003cp\u003eContains a map of the key-value pairs for the resource tag or tags assigned to the dataset.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
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
    "description": "Definition of the AWS::QuickSight::DataSet Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightDataSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightDataSet), &result)
	return &result
}
