package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3TablesTableBucket = `{
  "block": {
    "attributes": {
      "encryption_configuration": {
        "computed": true,
        "description": "Specifies encryption settings for the table bucket",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "ARN of the KMS key to use for encryption",
              "description_kind": "plain",
              "type": "string"
            },
            "sse_algorithm": {
              "computed": true,
              "description": "Server-side encryption algorithm",
              "description_kind": "plain",
              "type": "string"
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
      "logging_configuration": {
        "computed": true,
        "description": "Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a table bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "target_bucket": {
              "computed": true,
              "description": "The name of the general purpose bucket that logs are delivered to. The bucket must be in the same AWS Region and owned by the same AWS account as the table bucket. The bucket can't use SSE-KMS or DSSE-KMS default encryption, can't have S3 Object Lock enabled, and can't have Requester Pays enabled. You can deliver logs from multiple table buckets to the same target bucket. In this case, choose a different TargetPrefix for each table bucket so that you can tell their log objects apart.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_object_key_format": {
              "computed": true,
              "description": "The key format of the log objects. Only PartitionedPrefix is supported.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "partitioned_prefix": {
                    "computed": true,
                    "description": "Amazon S3 keys for log objects are partitioned in the following format: [TargetPrefix][SourceAccountId]/[SourceRegion]/[SourceTableBucket]/[YYYY]/[MM]/[DD]/[YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString]",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "partition_date_source": {
                          "computed": true,
                          "description": "The date to use for the log object key partitions. Valid values are EventTime, the time that the requests in the log were made, and DeliveryTime, the time that the log was delivered.",
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
            "target_prefix": {
              "computed": true,
              "description": "The prefix for the keys of the log objects that are delivered to the target bucket.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "metrics_configuration": {
        "computed": true,
        "description": "Settings governing the Metric configuration for the table bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "computed": true,
              "description": "Indicates whether Metrics are enabled.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "replication_configuration": {
        "computed": true,
        "description": "Specifies replication configuration for the table bucket",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "role": {
              "computed": true,
              "description": "The ARN of the IAM role to use for replication",
              "description_kind": "plain",
              "type": "string"
            },
            "rules": {
              "computed": true,
              "description": "List of replication rules",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destinations": {
                    "computed": true,
                    "description": "List of replication destinations",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "destination_table_bucket_arn": {
                          "computed": true,
                          "description": "The ARN of the destination table bucket",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "storage_class_configuration": {
        "computed": true,
        "description": "Specifies storage class settings for the table bucket",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "storage_class": {
              "computed": true,
              "description": "The storage class for the table bucket",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "table_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified table bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_bucket_name": {
        "computed": true,
        "description": "A name for the table bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "User tags (key-value pairs) to associate with the table bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag key must be between 1 to 128 characters in length. Tag key cannot start with 'aws:' and can only contain alphanumeric characters, spaces, _, ., /, =, +, -, and @.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag value must be between 0 to 256 characters in length. Tag value can only contain alphanumeric characters, spaces, _, ., /, =, +, -, and @.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "unreferenced_file_removal": {
        "computed": true,
        "description": "Settings governing the Unreferenced File Removal maintenance action. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "noncurrent_days": {
              "computed": true,
              "description": "S3 permanently deletes noncurrent objects after the number of days specified by the NoncurrentDays property.",
              "description_kind": "plain",
              "type": "number"
            },
            "status": {
              "computed": true,
              "description": "Indicates whether the Unreferenced File Removal maintenance action is enabled.",
              "description_kind": "plain",
              "type": "string"
            },
            "unreferenced_days": {
              "computed": true,
              "description": "For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::S3Tables::TableBucket",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3TablesTableBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3TablesTableBucket), &result)
	return &result
}
