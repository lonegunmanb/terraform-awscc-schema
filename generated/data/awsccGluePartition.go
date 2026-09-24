package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGluePartition = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "computed": true,
        "description": "The name of the catalog database in which to create the partition.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_name": {
        "computed": true,
        "description": "The AWS account ID of the catalog in which the partion is to be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identifier_partition_input_values": {
        "computed": true,
        "description": "A hashed string equivalent to the partition values list for use as a component of the resource's identifier in CFN",
        "description_kind": "plain",
        "type": "string"
      },
      "partition_input": {
        "computed": true,
        "description": "The structure used to create and update a partition.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "parameters": {
              "computed": true,
              "description": "Key-value pairs defining partition parameters.",
              "description_kind": "plain",
              "type": "string"
            },
            "storage_descriptor": {
              "computed": true,
              "description": "Provides information about the physical location where the partition is stored.",
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
                    "description": "The number of buckets. You must specify this property if the partition contains any dimension columns.",
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
                          "description": "A list of values that appear so frequently as to be considered skewed.",
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
                          "description": "Indicates that the column is sorted in ascending order (== 1), or in descending order (==0).",
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
            "values": {
              "computed": true,
              "description": "The values of the partition. Although this parameter is not required by the SDK, you must specify this parameter for a valid input. The values for the keys for the new partition must be passed as an array of String objects that must be ordered in the same order as the partition keys appearing in the Amazon S3 prefix. Otherwise AWS Glue will add the values to the wrong keys.",
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
      "table_name": {
        "computed": true,
        "description": "The name of the metadata table in which the partition is to be created.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Glue::Partition",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGluePartitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGluePartition), &result)
	return &result
}
