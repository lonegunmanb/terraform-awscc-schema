package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTimestreamScheduledQuery = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name of the scheduled query that is generated upon creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description": "Using a ClientToken makes the call to CreateScheduledQuery idempotent, in other words, making the same request repeatedly will produce the same result. Making multiple identical CreateScheduledQuery requests has the same effect as making a single request. If CreateScheduledQuery is called without a ClientToken, the Query SDK generates a ClientToken on your behalf. After 8 hours, any request with the same ClientToken is treated as a new request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "error_report_configuration": {
        "description": "Configuration for error reporting. Error reports will be generated when a problem is encountered when writing the query results.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_configuration": {
              "description": "Details on S3 location for error reports that result from running a query.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "description": "Name of the S3 bucket under which error reports will be created.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "encryption_option": {
                    "computed": true,
                    "description": "Encryption at rest options for the error reports. If no encryption option is specified, Timestream will choose SSE_S3 as default.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "object_key_prefix": {
                    "computed": true,
                    "description": "Prefix for error report keys.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon KMS key used to encrypt the scheduled query resource, at-rest. If the Amazon KMS key is not specified, the scheduled query resource will be encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the name with alias/. If ErrorReportConfiguration uses SSE_KMS as encryption type, the same KmsKeyId is used to encrypt the error report at rest.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notification_configuration": {
        "description": "Notification configuration for the scheduled query. A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sns_configuration": {
              "description": "SNS configuration for notification upon scheduled query execution.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "topic_arn": {
                    "description": "SNS topic ARN that the scheduled query status notifications will be sent to.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "query_string": {
        "description": "The query string to run. Parameter names can be specified in the query string @ character followed by an identifier. The named Parameter @scheduled_runtime is reserved and can be used in the query to get the time at which the query is scheduled to run. The timestamp calculated according to the ScheduleConfiguration parameter, will be the value of @scheduled_runtime paramater for each query run. For example, consider an instance of a scheduled query executing on 2021-12-01 00:00:00. For this instance, the @scheduled_runtime parameter is initialized to the timestamp 2021-12-01 00:00:00 when invoking the query.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_configuration": {
        "description": "Configuration for when the scheduled query is executed.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "schedule_expression": {
              "description": "An expression that denotes when to trigger the scheduled query run. This can be a cron expression or a rate expression.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "scheduled_query_execution_role_arn": {
        "description": "The ARN for the IAM role that Timestream will assume when running the scheduled query.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scheduled_query_name": {
        "computed": true,
        "description": "The name of the scheduled query. Scheduled query names must be unique within each Region.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sq_error_report_configuration": {
        "computed": true,
        "description": "Configuration for error reporting. Error reports will be generated when a problem is encountered when writing the query results.",
        "description_kind": "plain",
        "type": "string"
      },
      "sq_kms_key_id": {
        "computed": true,
        "description": "The Amazon KMS key used to encrypt the scheduled query resource, at-rest. If the Amazon KMS key is not specified, the scheduled query resource will be encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the name with alias/. If ErrorReportConfiguration uses SSE_KMS as encryption type, the same KmsKeyId is used to encrypt the error report at rest.",
        "description_kind": "plain",
        "type": "string"
      },
      "sq_name": {
        "computed": true,
        "description": "The name of the scheduled query. Scheduled query names must be unique within each Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "sq_notification_configuration": {
        "computed": true,
        "description": "Notification configuration for the scheduled query. A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it.",
        "description_kind": "plain",
        "type": "string"
      },
      "sq_query_string": {
        "computed": true,
        "description": "The query string to run. Parameter names can be specified in the query string @ character followed by an identifier. The named Parameter @scheduled_runtime is reserved and can be used in the query to get the time at which the query is scheduled to run. The timestamp calculated according to the ScheduleConfiguration parameter, will be the value of @scheduled_runtime paramater for each query run. For example, consider an instance of a scheduled query executing on 2021-12-01 00:00:00. For this instance, the @scheduled_runtime parameter is initialized to the timestamp 2021-12-01 00:00:00 when invoking the query.",
        "description_kind": "plain",
        "type": "string"
      },
      "sq_schedule_configuration": {
        "computed": true,
        "description": "Configuration for when the scheduled query is executed.",
        "description_kind": "plain",
        "type": "string"
      },
      "sq_scheduled_query_execution_role_arn": {
        "computed": true,
        "description": "The ARN for the IAM role that Timestream will assume when running the scheduled query.",
        "description_kind": "plain",
        "type": "string"
      },
      "sq_target_configuration": {
        "computed": true,
        "description": "Configuration of target store where scheduled query results are written to.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs to label the scheduled query.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "target_configuration": {
        "computed": true,
        "description": "Configuration of target store where scheduled query results are written to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "timestream_configuration": {
              "computed": true,
              "description": "Configuration needed to write data into the Timestream database and table.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database_name": {
                    "computed": true,
                    "description": "Name of Timestream database to which the query result will be written.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dimension_mappings": {
                    "computed": true,
                    "description": "This is to allow mapping column(s) from the query result to the dimension in the destination table.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dimension_value_type": {
                          "computed": true,
                          "description": "Type for the dimension.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description": "Column name from query result.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "measure_name_column": {
                    "computed": true,
                    "description": "Name of the measure name column from the query result.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mixed_measure_mappings": {
                    "computed": true,
                    "description": "Specifies how to map measures to multi-measure records.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "measure_name": {
                          "computed": true,
                          "description": "Refers to the value of the measure name in a result row. This field is required if MeasureNameColumn is provided.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "measure_value_type": {
                          "computed": true,
                          "description": "Type of the value that is to be read from SourceColumn. If the mapping is for MULTI, use MeasureValueType.MULTI.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "multi_measure_attribute_mappings": {
                          "computed": true,
                          "description": "Required. Attribute mappings to be used for mapping query results to ingest data for multi-measure attributes.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "measure_value_type": {
                                "computed": true,
                                "description": "Value type of the measure value column to be read from the query result.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "source_column": {
                                "computed": true,
                                "description": "Source measure value column in the query result where the attribute value is to be read.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "target_multi_measure_attribute_name": {
                                "computed": true,
                                "description": "Custom name to be used for attribute name in derived table. If not provided, source column name would be used.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "source_column": {
                          "computed": true,
                          "description": "This field refers to the source column from which the measure value is to be read for result materialization.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "target_measure_name": {
                          "computed": true,
                          "description": "Target measure name to be used. If not provided, the target measure name by default would be MeasureName if provided, or SourceColumn otherwise.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "multi_measure_mappings": {
                    "computed": true,
                    "description": "Only one of MixedMeasureMappings or MultiMeasureMappings is to be provided. MultiMeasureMappings can be used to ingest data as multi measures in the derived table.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "multi_measure_attribute_mappings": {
                          "computed": true,
                          "description": "Required. Attribute mappings to be used for mapping query results to ingest data for multi-measure attributes.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "measure_value_type": {
                                "computed": true,
                                "description": "Value type of the measure value column to be read from the query result.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "source_column": {
                                "computed": true,
                                "description": "Source measure value column in the query result where the attribute value is to be read.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "target_multi_measure_attribute_name": {
                                "computed": true,
                                "description": "Custom name to be used for attribute name in derived table. If not provided, source column name would be used.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "target_multi_measure_name": {
                          "computed": true,
                          "description": "Name of the target multi-measure in the derived table. Required if MeasureNameColumn is not provided. If MeasureNameColumn is provided then the value from that column will be used as the multi-measure name.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "table_name": {
                    "computed": true,
                    "description": "Name of Timestream table that the query result will be written to. The table should be within the same database that is provided in Timestream configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "time_column": {
                    "computed": true,
                    "description": "Column from query result that should be used as the time column in destination table. Column type for this should be TIMESTAMP.",
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
    "description": "The AWS::Timestream::ScheduledQuery resource creates a Timestream Scheduled Query.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTimestreamScheduledQuerySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTimestreamScheduledQuery), &result)
	return &result
}
