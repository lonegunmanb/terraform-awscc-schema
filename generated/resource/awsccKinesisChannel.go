package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKinesisChannel = `{
  "block": {
    "attributes": {
      "channel_arn": {
        "computed": true,
        "description": "The ARN of the channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_creation_timestamp": {
        "computed": true,
        "description": "Timestamp of when the channel was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_id": {
        "computed": true,
        "description": "The service-generated unique identifier for the channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_name": {
        "description": "The name of the channel. The name's uniqueness is scoped per AWS account and region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "channel_status": {
        "computed": true,
        "description": "The status of the channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "Server-side encryption configuration for data at rest.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_type": {
              "computed": true,
              "description": "The encryption type. KMS is the only supported value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_id": {
              "computed": true,
              "description": "The customer-managed AWS KMS key. Accepts a key GUID, key ARN, alias ARN, or alias name prefixed by 'alias/'. The Kinesis Data Streams managed alias 'aws/kinesis' is not accepted - the key must be customer-owned so it can also be used by readers of the destination.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
      "logging_configuration": {
        "computed": true,
        "description": "Configuration for delivering channel operational logs. Defaults to CloudWatch Logs disabled.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs": {
              "computed": true,
              "description": "CloudWatch Logs configuration block. When provided, controls whether and where the channel writes operational logs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "Whether CloudWatch Logs delivery is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group_name": {
                    "computed": true,
                    "description": "The CloudWatch log group name. When Enabled is true and LogGroupName is omitted, the service uses the default '/aws/kinesis/\u003cchannelName\u003e/\u003cchannelId\u003e'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_stream_name": {
                    "computed": true,
                    "description": "The CloudWatch log stream name. Defaults to the literal string 'DestinationDelivery' when omitted.",
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
        "description": "Configuration for delivery to a vanilla S3 bucket destination. Exactly one of S3DestinationConfiguration and S3TablesDestinationConfiguration must be specified.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_freshness_in_seconds": {
              "computed": true,
              "description": "The maximum time in seconds the channel buffers records before delivery if the minimum target file size is not reached.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dead_letter_queue_s3_configuration": {
              "computed": true,
              "description": "Optional dead-letter queue (DLQ) configuration for records that cannot be delivered to the destination. When omitted, the service auto-fills using the storage BucketARN with an error prefix.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "The ARN of the S3 bucket for storing failed records.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "error_output_prefix": {
                    "computed": true,
                    "description": "Optional S3 key prefix under which error records are organized. When omitted, the service uses the default 'kinesis-channel/errors/\u003cchannelName\u003e/\u003cchannelId\u003e/'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expected_bucket_owner": {
                    "computed": true,
                    "description": "The AWS account ID of the expected owner of the dead-letter queue S3 bucket. Used to verify bucket ownership before delivery.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "storage_configuration": {
              "computed": true,
              "description": "S3 storage configuration including the destination bucket, output key template, storage class, and compression type.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "The ARN of the S3 bucket for record delivery. Different channels can deliver to the same bucket. Buckets can be cross-account but must be in the same region as the channel.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "compression_type": {
                    "computed": true,
                    "description": "The compression algorithm applied to delivered objects.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expected_bucket_owner": {
                    "computed": true,
                    "description": "The AWS account ID of the expected owner of the destination S3 bucket. Used to verify bucket ownership before delivery.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_key_template": {
                    "computed": true,
                    "description": "Optional template for the S3 object key path. Supports placeholders in the form !{name}: !{channel-name}, !{channel-id}, !{stream-name}, !{yyyy}, !{yy}, !{MM}, !{dd}, !{HH}, !{mm}, and !{extension} (a literal file extension can be supplied as !{extension:.json.gz}). When omitted, the service uses the default 'kinesis-channel/!{channel-name}/!{channel-id}/!{yyyy}/!{MM}/!{dd}/!{HH}/!{channel-name}-!{channel-id}-!{yyyy}-!{MM}-!{dd}-!{HH}-!{mm}!{extension}'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_class": {
                    "computed": true,
                    "description": "The S3 storage class for delivered objects.",
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
      "s3_tables_destination_configuration": {
        "computed": true,
        "description": "Configuration for delivery to S3 Tables destinations. Exactly one of S3DestinationConfiguration and S3TablesDestinationConfiguration must be specified.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_freshness_in_seconds": {
              "computed": true,
              "description": "The maximum time in seconds the channel buffers records before delivery if the minimum target file size is not reached.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dead_letter_queue_s3_configuration": {
              "computed": true,
              "description": "The dead-letter queue (DLQ) configuration for records that cannot be delivered to the S3 Tables destination. Required for S3 Tables: there is no safe fallback because S3 Tables metadata writes are critical-path.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "The ARN of the S3 bucket for storing failed records.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "error_output_prefix": {
                    "computed": true,
                    "description": "Optional S3 key prefix under which error records are organized. When omitted, the service uses the default 'kinesis-channel/errors/\u003cchannelName\u003e/\u003cchannelId\u003e/'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expected_bucket_owner": {
                    "computed": true,
                    "description": "The AWS account ID of the expected owner of the dead-letter queue S3 bucket. Used to verify bucket ownership before delivery.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "s3_tables_configuration_list": {
              "computed": true,
              "description": "The list of S3 Tables destinations. v1 supports a single element; the list shape allows future extensibility to fan out to multiple tables.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "compression_type": {
                    "computed": true,
                    "description": "The compression algorithm applied to objects delivered to the S3 Tables destination.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "namespace": {
                    "computed": true,
                    "description": "The name of the S3 Tables namespace that contains the destination table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partition_spec": {
                    "computed": true,
                    "description": "The partition specification used by the destination Iceberg table.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "partition_fields": {
                          "computed": true,
                          "description": "List of partition fields that define how records are partitioned when written to the destination table.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "source_name": {
                                "computed": true,
                                "description": "The name of the source column on which the transform is applied.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "transform": {
                                "computed": true,
                                "description": "The partitioning transform applied to the SourceName column.",
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
                  "table_bucket_arn": {
                    "computed": true,
                    "description": "The ARN of the S3 Tables table bucket for record delivery. Buckets can be cross-account but must be in the same region as the channel.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description": "The name of the destination S3 Tables table. The table is created for the customer if it does not yet exist.",
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
      "service_execution_role_arn": {
        "description": "The ARN of the IAM role that the channel assumes to read from the source stream, deliver records to the destination, and (when enabled) write CloudWatch Logs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stream_configuration_list": {
        "description": "List of stream configurations associated with the channel. v1 supports a single element; the list shape allows future extensibility to fan in from multiple streams.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "record_configuration": {
              "description": "The configuration that describes how records on the source stream are encoded.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "gsr_schema_arn": {
                    "computed": true,
                    "description": "The ARN of the AWS Glue Schema Registry (GSR) schema. Required for the S3 Tables destination, where it is used to create the S3 Table and to validate that the record format matches the table schema. Also used when RecordFormatType is GSR_JSON to interpret records read from the source stream. Vanilla S3 delivery writes records as S3 objects and does not need a schema. The schema must be in the same account and region as the channel.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "record_format_type": {
                    "description": "The format used to interpret records read from the source stream.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "stream_arn": {
              "description": "The Amazon resource name (ARN) of the Kinesis data stream that the channel reads from.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) to associate with the Kinesis channel.",
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
              "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "Resource Type definition for AWS::Kinesis::Channel",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKinesisChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKinesisChannel), &result)
	return &result
}
