package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMskChannel = `{
  "block": {
    "attributes": {
      "channel_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that uniquely identifies the channel",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_name": {
        "description": "Name of the channel",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the cluster",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "Encryption configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "The ARN of the KMS key for encryption",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "iceberg_destination_configuration": {
        "computed": true,
        "description": "Iceberg destination configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "append_only": {
              "computed": true,
              "description": "Append only mode",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "catalog": {
              "computed": true,
              "description": "Catalog configuration of the destination",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_arn": {
                    "computed": true,
                    "description": "The ARN of the catalog",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "warehouse_location": {
                    "computed": true,
                    "description": "The warehouse location",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "compression_type": {
              "computed": true,
              "description": "Compression codec for Iceberg table data files. Defaults to ZSTD.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_freshness_in_seconds": {
              "computed": true,
              "description": "Data freshness in seconds",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dead_letter_queue_s3": {
              "computed": true,
              "description": "Dead letter queue S3 configuration of the destination",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "The ARN of the S3 bucket",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "error_output_prefix": {
                    "computed": true,
                    "description": "The error output prefix",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expected_bucket_owner": {
                    "computed": true,
                    "description": "Optional 12-digit AWS account ID expected to own the dead-letter S3 bucket",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "destination_table_list": {
              "computed": true,
              "description": "List of destination tables",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_database_name": {
                    "computed": true,
                    "description": "The destination database name",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "destination_table_name": {
                    "computed": true,
                    "description": "The destination table name",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partition_spec": {
                    "computed": true,
                    "description": "Partition specification",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "partition_strategy": {
                          "computed": true,
                          "description": "Partition strategy for MSK channel",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "source_list": {
                          "computed": true,
                          "description": "Source list",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "source_name": {
                                "computed": true,
                                "description": "Source name",
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
            "schema_evolution": {
              "computed": true,
              "description": "Schema evolution configuration of the destination",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_schema_evolution": {
                    "computed": true,
                    "description": "Whether schema evolution is enabled",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "service_execution_role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of an IAM role used by MSK to access the table",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "table_creation": {
              "computed": true,
              "description": "Table creation configuration of the destination",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_table_creation": {
                    "computed": true,
                    "description": "Whether table creation is enabled",
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
      "logging_info": {
        "computed": true,
        "description": "Log configuration details for Channel",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs": {
              "computed": true,
              "description": "CloudWatch Logs log destination details",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "Whether CloudWatch Logs logging is enabled",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group": {
                    "computed": true,
                    "description": "The CloudWatch log group for log delivery",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "firehose": {
              "computed": true,
              "description": "Firehose log destination details",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delivery_stream": {
                    "computed": true,
                    "description": "The Firehose delivery stream for log delivery",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description": "Whether Firehose logging is enabled",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "s3": {
              "computed": true,
              "description": "S3 log destination details",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description": "The name of the S3 bucket for log delivery",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description": "Whether S3 logging is enabled",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "prefix": {
                    "computed": true,
                    "description": "The S3 prefix for log delivery",
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
      "s3_destination_configuration": {
        "computed": true,
        "description": "S3 destination configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_freshness_in_seconds": {
              "computed": true,
              "description": "Data freshness in seconds",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dead_letter_queue_s3": {
              "computed": true,
              "description": "Dead letter queue S3 configuration of the destination",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "The ARN of the S3 bucket",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "error_output_prefix": {
                    "computed": true,
                    "description": "The error output prefix",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expected_bucket_owner": {
                    "computed": true,
                    "description": "Optional 12-digit AWS account ID expected to own the dead-letter S3 bucket",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "service_execution_role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of an IAM role used by MSK to access S3",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage": {
              "computed": true,
              "description": "S3 storage configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "ARN of the S3 bucket",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "compression_type": {
                    "computed": true,
                    "description": "S3 compression type",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expected_bucket_owner": {
                    "computed": true,
                    "description": "Optional 12-digit AWS account ID expected to own the S3 bucket",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_key_template": {
                    "computed": true,
                    "description": "Template for S3 key for output objects, used for partitioning",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_prefix": {
                    "computed": true,
                    "description": "Optional prefix for output objects",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_class": {
                    "computed": true,
                    "description": "S3 storage class",
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
      "state_info": {
        "computed": true,
        "description": "Includes information about the channel state",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code": {
              "computed": true,
              "description": "Code for channel state",
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description": "Message for channel state",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "Status of a channel resource",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags attached to the channel",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "topic_configuration_list": {
        "description": "Topic configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "record_converter": {
              "description": "Record converter configuration for a topic",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "value_converter": {
                    "description": "Value converter for topic data",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "record_schema": {
              "computed": true,
              "description": "Record schema configuration for a topic",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "gsr_arn": {
                    "computed": true,
                    "description": "ARN of Glue Schema Registry resource used for table schema",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "topic_arn": {
              "description": "The Amazon Resource Name (ARN) that uniquely identifies the topic",
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
    "description": "Resource Type definition for AWS::MSK::Channel",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMskChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMskChannel), &result)
	return &result
}
