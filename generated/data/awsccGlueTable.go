package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueTable = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "computed": true,
        "description": "The ID of the Data Catalog in which to create the Table.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_name": {
        "computed": true,
        "description": "The name of the database where the table metadata resides. For Hive compatibility, this must be all lowercase.",
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
        "description": "The table name. For Hive compatibility, this is folded to lowercase when it is stored.",
        "description_kind": "plain",
        "type": "string"
      },
      "open_table_format_input": {
        "computed": true,
        "description": "Specifies an OpenTableFormatInput structure when creating an open format table.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "iceberg_input": {
              "computed": true,
              "description": "Specifies an IcebergInput structure that defines an Apache Iceberg metadata table.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "iceberg_table_input": {
                    "computed": true,
                    "description": "Specifies the Iceberg table configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "location": {
                          "computed": true,
                          "description": "The location of the Iceberg table.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "partition_spec": {
                          "computed": true,
                          "description": "The partition spec for the Iceberg table.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "fields": {
                                "computed": true,
                                "description": "The partition fields.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "field_id": {
                                      "computed": true,
                                      "description": "The partition field ID.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "name": {
                                      "computed": true,
                                      "description": "The partition field name.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "source_id": {
                                      "computed": true,
                                      "description": "The source field ID.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "transform": {
                                      "computed": true,
                                      "description": "The transform to apply.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "spec_id": {
                                "computed": true,
                                "description": "The spec ID.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "properties": {
                          "computed": true,
                          "description": "The properties for the Iceberg table.",
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "schema": {
                          "computed": true,
                          "description": "The schema for the Iceberg table.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "fields": {
                                "computed": true,
                                "description": "The fields in the schema.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "doc": {
                                      "computed": true,
                                      "description": "Documentation for the field.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "id": {
                                      "computed": true,
                                      "description": "The field ID.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "name": {
                                      "computed": true,
                                      "description": "The field name.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "required": {
                                      "computed": true,
                                      "description": "Whether the field is required.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "type": {
                                      "computed": true,
                                      "description": "The field type.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "identifier_field_ids": {
                                "computed": true,
                                "description": "The identifier field IDs.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "number"
                                ]
                              },
                              "schema_id": {
                                "computed": true,
                                "description": "The schema ID.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "type": {
                                "computed": true,
                                "description": "The type of the schema.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "write_order": {
                          "computed": true,
                          "description": "The write order for the Iceberg table.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "fields": {
                                "computed": true,
                                "description": "The sort fields.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "direction": {
                                      "computed": true,
                                      "description": "The sort direction.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "null_order": {
                                      "computed": true,
                                      "description": "The null ordering.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "source_id": {
                                      "computed": true,
                                      "description": "The source field ID.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "transform": {
                                      "computed": true,
                                      "description": "The transform to apply.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "order_id": {
                                "computed": true,
                                "description": "The order ID.",
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
                  "metadata_operation": {
                    "computed": true,
                    "description": "A required metadata operation. Can only be set to CREATE.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "version": {
                    "computed": true,
                    "description": "The table version of the Iceberg table. Defaults to 2.",
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
      "table_input": {
        "computed": true,
        "description": "A structure used to define a table.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "A description of the table.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The table name. For Hive compatibility, this is folded to lowercase when it is stored.",
              "description_kind": "plain",
              "type": "string"
            },
            "owner": {
              "computed": true,
              "description": "The table owner. Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.",
              "description_kind": "plain",
              "type": "string"
            },
            "parameters": {
              "computed": true,
              "description": "These key-value pairs define properties associated with the table.",
              "description_kind": "plain",
              "type": "string"
            },
            "partition_keys": {
              "computed": true,
              "description": "A list of columns by which the table is partitioned. Only primitive types are supported as partition keys.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comment": {
                    "computed": true,
                    "description": "A free-form text comment.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the Column.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "The data type of the Column.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "retention": {
              "computed": true,
              "description": "The retention time for this table.",
              "description_kind": "plain",
              "type": "number"
            },
            "storage_descriptor": {
              "computed": true,
              "description": "A storage descriptor containing information about the physical storage of this table.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_columns": {
                    "computed": true,
                    "description": "A list of reducer grouping columns, clustering columns, and bucketing columns in the table.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "columns": {
                    "computed": true,
                    "description": "A list of the Columns in the table.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comment": {
                          "computed": true,
                          "description": "A free-form text comment.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description": "The name of the Column.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "The data type of the Column.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "compressed": {
                    "computed": true,
                    "description": "True if the data in the table is compressed, or False if not.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "input_format": {
                    "computed": true,
                    "description": "The input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "location": {
                    "computed": true,
                    "description": "The physical location of the table. By default, this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "number_of_buckets": {
                    "computed": true,
                    "description": "Must be specified if the table contains any dimension columns.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "output_format": {
                    "computed": true,
                    "description": "The output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "parameters": {
                    "computed": true,
                    "description": "The user-supplied properties in key-value form.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "schema_reference": {
                    "computed": true,
                    "description": "An object that references a schema stored in the AWS Glue Schema Registry.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "schema_id": {
                          "computed": true,
                          "description": "A structure that contains schema identity fields. Either this or the SchemaVersionId has to be provided.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "registry_name": {
                                "computed": true,
                                "description": "The name of the schema registry that contains the schema.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "schema_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) of the schema. One of SchemaArn or SchemaName has to be provided.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "schema_name": {
                                "computed": true,
                                "description": "The name of the schema. One of SchemaArn or SchemaName has to be provided.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "schema_version_id": {
                          "computed": true,
                          "description": "The unique ID assigned to a version of the schema. Either this or the SchemaId has to be provided.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "schema_version_number": {
                          "computed": true,
                          "description": "The version number of the schema.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "serde_info": {
                    "computed": true,
                    "description": "The serialization/deserialization (SerDe) information.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "Name of the SerDe.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "parameters": {
                          "computed": true,
                          "description": "These key-value pairs define initialization parameters for the SerDe.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "serialization_library": {
                          "computed": true,
                          "description": "Usually the class that implements the SerDe. An example is org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "skewed_info": {
                    "computed": true,
                    "description": "The information about values that appear frequently in a column (skewed values).",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "skewed_column_names": {
                          "computed": true,
                          "description": "A list of names of columns that contain skewed values.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "skewed_column_value_location_maps": {
                          "computed": true,
                          "description": "A mapping of skewed values to the columns that contain them.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "skewed_column_values": {
                          "computed": true,
                          "description": "A list of names of columns that contain skewed values.",
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
                  "sort_columns": {
                    "computed": true,
                    "description": "A list specifying the sort order of each bucket in the table.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "column": {
                          "computed": true,
                          "description": "The name of the column.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sort_order": {
                          "computed": true,
                          "description": "Indicates that the column is sorted in ascending order (==1), or in descending order (==0).",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "stored_as_sub_directories": {
                    "computed": true,
                    "description": "True if the table data is stored in subdirectories, or False if not.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "table_type": {
              "computed": true,
              "description": "The type of this table. AWS Glue will create tables with the EXTERNAL_TABLE type. Other services, such as Athena, may create tables with additional table types.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_table": {
              "computed": true,
              "description": "A TableIdentifier structure that describes a target table for resource linking.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description": "The ID of the Data Catalog in which the table resides.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "The name of the catalog database that contains the target table.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the target table.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description": "The Region of the table.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "view_definition": {
              "computed": true,
              "description": "The view definition for the table.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "definer": {
                    "computed": true,
                    "description": "The definer of the view.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "is_protected": {
                    "computed": true,
                    "description": "Indicates whether the view is protected.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "representations": {
                    "computed": true,
                    "description": "The representations of the view.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dialect": {
                          "computed": true,
                          "description": "The dialect of the view representation.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "dialect_version": {
                          "computed": true,
                          "description": "The version of the dialect.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "validation_connection": {
                          "computed": true,
                          "description": "The connection used for validation.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "view_expanded_text": {
                          "computed": true,
                          "description": "The expanded text of the view.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "view_original_text": {
                          "computed": true,
                          "description": "The original text of the view.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "sub_objects": {
                    "computed": true,
                    "description": "The sub-objects of the view.",
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
            "view_expanded_text": {
              "computed": true,
              "description": "Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.",
              "description_kind": "plain",
              "type": "string"
            },
            "view_original_text": {
              "computed": true,
              "description": "Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Glue::Table",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueTable), &result)
	return &result
}
