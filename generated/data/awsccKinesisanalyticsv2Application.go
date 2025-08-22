package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKinesisanalyticsv2Application = `{
  "block": {
    "attributes": {
      "application_configuration": {
        "computed": true,
        "description": "Use this parameter to configure the application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_code_configuration": {
              "computed": true,
              "description": "The code location and type parameters for a Flink-based Kinesis Data Analytics application.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "code_content": {
                    "computed": true,
                    "description": "The location and type of the application code.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_content_location": {
                          "computed": true,
                          "description": "Information about the Amazon S3 bucket that contains the application code.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) for the S3 bucket containing the application code.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "file_key": {
                                "computed": true,
                                "description": "The file key for the object containing the application code.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "object_version": {
                                "computed": true,
                                "description": "The version of the object containing the application code.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "text_content": {
                          "computed": true,
                          "description": "The text-format code for a Flink-based Kinesis Data Analytics application.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "zip_file_content": {
                          "computed": true,
                          "description": "The zip-format code for a Flink-based Kinesis Data Analytics application.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "code_content_type": {
                    "computed": true,
                    "description": "Specifies whether the code content is in text or zip format.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "application_encryption_configuration": {
              "computed": true,
              "description": "Describes whether customer managed key is enabled and key details for customer data encryption",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key_id": {
                    "computed": true,
                    "description": "KMS KeyId. Can be either key uuid or full key arn or key alias arn or short key alias",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_type": {
                    "computed": true,
                    "description": "Specifies whether application data is encrypted using service key: AWS_OWNED_KEY or customer key: CUSTOMER_MANAGED_KEY",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "application_snapshot_configuration": {
              "computed": true,
              "description": "Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "snapshots_enabled": {
                    "computed": true,
                    "description": "Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "application_system_rollback_configuration": {
              "computed": true,
              "description": "Describes whether system initiated rollbacks are enabled for a Flink-based Kinesis Data Analytics application.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rollback_enabled": {
                    "computed": true,
                    "description": "Describes whether system initiated rollbacks are enabled for a Flink-based Kinesis Data Analytics application.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "environment_properties": {
              "computed": true,
              "description": "Describes execution properties for a Flink-based Kinesis Data Analytics application.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "property_groups": {
                    "computed": true,
                    "description": "Describes the execution property groups.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "property_group_id": {
                          "computed": true,
                          "description": "Describes the key of an application execution property key-value pair.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "property_map": {
                          "computed": true,
                          "description": "Describes the value of an application execution property key-value pair.",
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "flink_application_configuration": {
              "computed": true,
              "description": "The creation and update parameters for a Flink-based Kinesis Data Analytics application.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "checkpoint_configuration": {
                    "computed": true,
                    "description": "Describes an application's checkpointing configuration. Checkpointing is the process of persisting application state for fault tolerance. For more information, see Checkpoints for Fault Tolerance in the Apache Flink Documentation.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "checkpoint_interval": {
                          "computed": true,
                          "description": "Describes the interval in milliseconds between checkpoint operations.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "checkpointing_enabled": {
                          "computed": true,
                          "description": "Describes whether checkpointing is enabled for a Flink-based Kinesis Data Analytics application.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "configuration_type": {
                          "computed": true,
                          "description": "Describes whether the application uses Kinesis Data Analytics' default checkpointing behavior. You must set this property to ` + "`" + `CUSTOM` + "`" + ` in order to set the ` + "`" + `CheckpointingEnabled` + "`" + `, ` + "`" + `CheckpointInterval` + "`" + `, or ` + "`" + `MinPauseBetweenCheckpoints` + "`" + ` parameters.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "min_pause_between_checkpoints": {
                          "computed": true,
                          "description": "Describes the minimum time in milliseconds after a checkpoint operation completes that a new checkpoint operation can start. If a checkpoint operation takes longer than the CheckpointInterval, the application otherwise performs continual checkpoint operations. For more information, see Tuning Checkpointing in the Apache Flink Documentation.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "monitoring_configuration": {
                    "computed": true,
                    "description": "Describes configuration parameters for Amazon CloudWatch logging for an application.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "configuration_type": {
                          "computed": true,
                          "description": "Describes whether to use the default CloudWatch logging configuration for an application. You must set this property to CUSTOM in order to set the LogLevel or MetricsLevel parameters.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "log_level": {
                          "computed": true,
                          "description": "Describes the verbosity of the CloudWatch Logs for an application.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metrics_level": {
                          "computed": true,
                          "description": "Describes the granularity of the CloudWatch Logs for an application. The Parallelism level is not recommended for applications with a Parallelism over 64 due to excessive costs.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "parallelism_configuration": {
                    "computed": true,
                    "description": "Describes parameters for how an application executes multiple tasks simultaneously.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auto_scaling_enabled": {
                          "computed": true,
                          "description": "Describes whether the Kinesis Data Analytics service can increase the parallelism of the application in response to increased throughput.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "configuration_type": {
                          "computed": true,
                          "description": "Describes whether the application uses the default parallelism for the Kinesis Data Analytics service. You must set this property to ` + "`" + `CUSTOM` + "`" + ` in order to change your application's ` + "`" + `AutoScalingEnabled` + "`" + `, ` + "`" + `Parallelism` + "`" + `, or ` + "`" + `ParallelismPerKPU` + "`" + ` properties.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "parallelism": {
                          "computed": true,
                          "description": "Describes the initial number of parallel tasks that a Java-based Kinesis Data Analytics application can perform. The Kinesis Data Analytics service can increase this number automatically if ParallelismConfiguration:AutoScalingEnabled is set to true.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "parallelism_per_kpu": {
                          "computed": true,
                          "description": "Describes the number of parallel tasks that a Java-based Kinesis Data Analytics application can perform per Kinesis Processing Unit (KPU) used by the application. For more information about KPUs, see Amazon Kinesis Data Analytics Pricing.",
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
            "sql_application_configuration": {
              "computed": true,
              "description": "The creation and update parameters for a SQL-based Kinesis Data Analytics application.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "inputs": {
                    "computed": true,
                    "description": "The array of Input objects describing the input streams used by the application.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "input_parallelism": {
                          "computed": true,
                          "description": "Describes the number of in-application streams to create.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "count": {
                                "computed": true,
                                "description": "The number of in-application streams to create.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "input_processing_configuration": {
                          "computed": true,
                          "description": "The InputProcessingConfiguration for the input. An input processor transforms records as they are received from the stream, before the application's SQL code executes. Currently, the only input processing configuration available is InputLambdaProcessor.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "input_lambda_processor": {
                                "computed": true,
                                "description": "The InputLambdaProcessor that is used to preprocess the records in the stream before being processed by your application code.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "resource_arn": {
                                      "computed": true,
                                      "description": "The ARN of the Amazon Lambda function that operates on records in the stream.",
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
                        "input_schema": {
                          "computed": true,
                          "description": "Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "record_columns": {
                                "computed": true,
                                "description": "A list of ` + "`" + `RecordColumn` + "`" + ` objects.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "mapping": {
                                      "computed": true,
                                      "description": "A reference to the data element in the streaming input or the reference data source.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "name": {
                                      "computed": true,
                                      "description": "The name of the column that is created in the in-application input stream or reference table.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "sql_type": {
                                      "computed": true,
                                      "description": "The type of column created in the in-application input stream or reference table.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "record_encoding": {
                                "computed": true,
                                "description": "Specifies the encoding of the records in the streaming source. For example, UTF-8.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "record_format": {
                                "computed": true,
                                "description": "Specifies the format of the records on the streaming source.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "mapping_parameters": {
                                      "computed": true,
                                      "description": "When you configure application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "csv_mapping_parameters": {
                                            "computed": true,
                                            "description": "Provides additional mapping information when the record format uses delimiters (for example, CSV).",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "record_column_delimiter": {
                                                  "computed": true,
                                                  "description": "The column delimiter. For example, in a CSV format, a comma (\",\") is the typical column delimiter.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "record_row_delimiter": {
                                                  "computed": true,
                                                  "description": "The row delimiter. For example, in a CSV format, '\\n' is the typical row delimiter.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "json_mapping_parameters": {
                                            "computed": true,
                                            "description": "Provides additional mapping information when JSON is the record format on the streaming source.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "record_row_path": {
                                                  "computed": true,
                                                  "description": "The path to the top-level parent that contains the records.",
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
                                    "record_format_type": {
                                      "computed": true,
                                      "description": "The type of record format.",
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
                        "kinesis_firehose_input": {
                          "computed": true,
                          "description": "If the streaming source is an Amazon Kinesis Data Firehose delivery stream, identifies the delivery stream's ARN.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "resource_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) of the delivery stream.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "kinesis_streams_input": {
                          "computed": true,
                          "description": "If the streaming source is an Amazon Kinesis data stream, identifies the stream's Amazon Resource Name (ARN).",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "resource_arn": {
                                "computed": true,
                                "description": "The ARN of the input Kinesis data stream to read.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "name_prefix": {
                          "computed": true,
                          "description": "The name prefix to use when creating an in-application stream. Suppose that you specify a prefix ` + "`" + `\"MyInApplicationStream\"` + "`" + `. Kinesis Data Analytics then creates one or more (as per the InputParallelism count you specified) in-application streams with the names ` + "`" + `\"MyInApplicationStream_001\"` + "`" + `, ` + "`" + `\"MyInApplicationStream_002\"` + "`" + `, and so on.",
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
            "vpc_configurations": {
              "computed": true,
              "description": "The array of descriptions of VPC configurations available to the application.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description": "The array of SecurityGroup IDs used by the VPC configuration.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_ids": {
                    "computed": true,
                    "description": "The array of Subnet IDs used by the VPC configuration.",
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
            "zeppelin_application_configuration": {
              "computed": true,
              "description": "The configuration parameters for a Kinesis Data Analytics Studio notebook.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_configuration": {
                    "computed": true,
                    "description": "The Amazon Glue Data Catalog that you use in queries in a Kinesis Data Analytics Studio notebook.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "glue_data_catalog_configuration": {
                          "computed": true,
                          "description": "The configuration parameters for the default Amazon Glue database. You use this database for Apache Flink SQL queries and table API transforms that you write in a Kinesis Data Analytics Studio notebook.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) of the database.",
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
                  "custom_artifacts_configuration": {
                    "computed": true,
                    "description": "A list of CustomArtifactConfiguration objects.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "artifact_type": {
                          "computed": true,
                          "description": "Set this to either ` + "`" + `UDF` + "`" + ` or ` + "`" + `DEPENDENCY_JAR` + "`" + `. ` + "`" + `UDF` + "`" + ` stands for user-defined functions. This type of artifact must be in an S3 bucket. A ` + "`" + `DEPENDENCY_JAR` + "`" + ` can be in either Maven or an S3 bucket.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "maven_reference": {
                          "computed": true,
                          "description": "The parameters required to fully specify a Maven reference.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "artifact_id": {
                                "computed": true,
                                "description": "The artifact ID of the Maven reference.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "group_id": {
                                "computed": true,
                                "description": "The group ID of the Maven reference.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "version": {
                                "computed": true,
                                "description": "The version of the Maven reference.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "s3_content_location": {
                          "computed": true,
                          "description": "The location of the custom artifacts.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) for the S3 bucket containing the application code.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "file_key": {
                                "computed": true,
                                "description": "The file key for the object containing the application code.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "object_version": {
                                "computed": true,
                                "description": "The version of the object containing the application code.",
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
                  "deploy_as_application_configuration": {
                    "computed": true,
                    "description": "The information required to deploy a Kinesis Data Analytics Studio notebook as an application with durable state.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_content_location": {
                          "computed": true,
                          "description": "The description of an Amazon S3 object that contains the Amazon Data Analytics application, including the Amazon Resource Name (ARN) of the S3 bucket, the name of the Amazon S3 object that contains the data, and the version number of the Amazon S3 object that contains the data.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "base_path": {
                                "computed": true,
                                "description": "The base path for the S3 bucket.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) of the S3 bucket.",
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
                  "monitoring_configuration": {
                    "computed": true,
                    "description": "The monitoring configuration of a Kinesis Data Analytics Studio notebook.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "log_level": {
                          "computed": true,
                          "description": "The verbosity of the CloudWatch Logs for an application. You can set it to ` + "`" + `INFO` + "`" + `, ` + "`" + `WARN` + "`" + `, ` + "`" + `ERROR` + "`" + `, or ` + "`" + `DEBUG` + "`" + `.",
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
          "nesting_mode": "single"
        }
      },
      "application_description": {
        "computed": true,
        "description": "The description of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_maintenance_configuration": {
        "computed": true,
        "description": "Used to configure start of maintenance window.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_maintenance_window_start_time": {
              "computed": true,
              "description": "The start time for the maintenance window.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "application_mode": {
        "computed": true,
        "description": "To create a Kinesis Data Analytics Studio notebook, you must set the mode to ` + "`" + `INTERACTIVE` + "`" + `. However, for a Kinesis Data Analytics for Apache Flink application, the mode is optional.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_name": {
        "computed": true,
        "description": "The name of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "run_configuration": {
        "computed": true,
        "description": "Specifies run configuration (start parameters) of a Kinesis Data Analytics application. Evaluated on update for RUNNING applications an only.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_restore_configuration": {
              "computed": true,
              "description": "Describes the restore behavior of a restarting application.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "application_restore_type": {
                    "computed": true,
                    "description": "Specifies how the application should be restored.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "snapshot_name": {
                    "computed": true,
                    "description": "The identifier of an existing snapshot of application state to use to restart an application. The application uses this value if RESTORE_FROM_CUSTOM_SNAPSHOT is specified for the ApplicationRestoreType.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "flink_run_configuration": {
              "computed": true,
              "description": "Describes the starting parameters for a Flink-based Kinesis Data Analytics application.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allow_non_restored_state": {
                    "computed": true,
                    "description": "When restoring from a snapshot, specifies whether the runtime is allowed to skip a state that cannot be mapped to the new program. Defaults to false. If you update your application without specifying this parameter, AllowNonRestoredState will be set to false, even if it was previously set to true.",
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
      "runtime_environment": {
        "computed": true,
        "description": "The runtime environment for the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_execution_role": {
        "computed": true,
        "description": "Specifies the IAM role that the application uses to access external resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of one or more tags to assign to the application. A tag is a key-value pair that identifies an application. Note that the maximum number of application tags includes system tags. The maximum number of user-defined application tags is 50.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that's 0 to 256 characters in length.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::KinesisAnalyticsV2::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKinesisanalyticsv2ApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKinesisanalyticsv2Application), &result)
	return &result
}
