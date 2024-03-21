package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDynamodbTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "attribute_definitions": {
        "computed": true,
        "description": "A list of attributes that describe the key schema for the table and indexes.\n This property is required to create a DDB table.\n Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). Replacement if you edit an existing AttributeDefinition.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attribute_name": {
              "computed": true,
              "description": "A name for the attribute.",
              "description_kind": "plain",
              "type": "string"
            },
            "attribute_type": {
              "computed": true,
              "description": "The data type for the attribute, where:\n  +   ` + "`" + `` + "`" + `S` + "`" + `` + "`" + ` - the attribute is of type String\n  +   ` + "`" + `` + "`" + `N` + "`" + `` + "`" + ` - the attribute is of type Number\n  +   ` + "`" + `` + "`" + `B` + "`" + `` + "`" + ` - the attribute is of type Binary",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "billing_mode": {
        "computed": true,
        "description": "Specify how you are charged for read and write throughput and how you manage capacity.\n Valid values include:\n  +   ` + "`" + `` + "`" + `PROVISIONED` + "`" + `` + "`" + ` - We recommend using ` + "`" + `` + "`" + `PROVISIONED` + "`" + `` + "`" + ` for predictable workloads. ` + "`" + `` + "`" + `PROVISIONED` + "`" + `` + "`" + ` sets the billing mode to [Provisioned Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).\n  +   ` + "`" + `` + "`" + `PAY_PER_REQUEST` + "`" + `` + "`" + ` - We recommend using ` + "`" + `` + "`" + `PAY_PER_REQUEST` + "`" + `` + "`" + ` for unpredictable workloads. ` + "`" + `` + "`" + `PAY_PER_REQUEST` + "`" + `` + "`" + ` sets the billing mode to [On-Demand Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).\n  \n If not specified, the default is ` + "`" + `` + "`" + `PROVISIONED` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "contributor_insights_specification": {
        "computed": true,
        "description": "The settings used to enable or disable CloudWatch Contributor Insights for the specified table.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Indicates whether CloudWatch Contributor Insights are to be enabled (true) or disabled (false).",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description": "Determines if a table is protected from deletion. When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Developer Guide*.",
        "description_kind": "plain",
        "type": "bool"
      },
      "global_secondary_indexes": {
        "computed": true,
        "description": "Global secondary indexes to be created on the table. You can create up to 20 global secondary indexes.\n  If you update a table to include a new global secondary index, CFNlong initiates the index creation and then proceeds with the stack update. CFNlong doesn't wait for the index to complete creation because the backfilling phase can take a long time, depending on the size of the table. You can't use the index or update the table until the index's status is ` + "`" + `` + "`" + `ACTIVE` + "`" + `` + "`" + `. You can track its status by using the DynamoDB [DescribeTable](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html) command.\n If you add or delete an index during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new index, you must manually delete the index. \n Updates are not supported. The following are exceptions:\n  +  If you update either the contributor insights specification or the provisioned throughput values of global secondary indexes, you can update the table without interruption.\n  +  You can delete or add one global secondary index without interruption. If you do both in the same update (for example, by changing the index's logical ID), the update fails.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "contributor_insights_specification": {
              "computed": true,
              "description": "The settings used to enable or disable CloudWatch Contributor Insights for the specified global secondary index.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "Indicates whether CloudWatch Contributor Insights are to be enabled (true) or disabled (false).",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "index_name": {
              "computed": true,
              "description": "The name of the global secondary index. The name must be unique among all other indexes on this table.",
              "description_kind": "plain",
              "type": "string"
            },
            "key_schema": {
              "computed": true,
              "description": "The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:\n  +   ` + "`" + `` + "`" + `HASH` + "`" + `` + "`" + ` - partition key\n  +   ` + "`" + `` + "`" + `RANGE` + "`" + `` + "`" + ` - sort key\n  \n  The partition key of an item is also known as its *hash attribute*. The term \"hash attribute\" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.\n The sort key of an item is also known as its *range attribute*. The term \"range attribute\" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_name": {
                    "computed": true,
                    "description": "The name of a key attribute.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_type": {
                    "computed": true,
                    "description": "The role that this key attribute will assume:\n  +   ` + "`" + `` + "`" + `HASH` + "`" + `` + "`" + ` - partition key\n  +   ` + "`" + `` + "`" + `RANGE` + "`" + `` + "`" + ` - sort key\n  \n  The partition key of an item is also known as its *hash attribute*. The term \"hash attribute\" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.\n The sort key of an item is also known as its *range attribute*. The term \"range attribute\" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "projection": {
              "computed": true,
              "description": "Represents attributes that are copied (projected) from the table into the global secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "non_key_attributes": {
                    "computed": true,
                    "description": "Represents the non-key attribute names which will be projected into the index.\n For local secondary indexes, the total count of ` + "`" + `` + "`" + `NonKeyAttributes` + "`" + `` + "`" + ` summed across all of the local secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "projection_type": {
                    "computed": true,
                    "description": "The set of attributes that are projected into the index:\n  +   ` + "`" + `` + "`" + `KEYS_ONLY` + "`" + `` + "`" + ` - Only the index and primary keys are projected into the index.\n  +   ` + "`" + `` + "`" + `INCLUDE` + "`" + `` + "`" + ` - In addition to the attributes described in ` + "`" + `` + "`" + `KEYS_ONLY` + "`" + `` + "`" + `, the secondary index will include other non-key attributes that you specify.\n  +   ` + "`" + `` + "`" + `ALL` + "`" + `` + "`" + ` - All of the table attributes are projected into the index.\n  \n When using the DynamoDB console, ` + "`" + `` + "`" + `ALL` + "`" + `` + "`" + ` is selected by default.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "provisioned_throughput": {
              "computed": true,
              "description": "Represents the provisioned throughput settings for the specified global secondary index.\n For current minimum and maximum provisioned throughput values, see [Service, Account, and Table Quotas](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the *Amazon DynamoDB Developer Guide*.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "read_capacity_units": {
                    "computed": true,
                    "description": "The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ` + "`" + `` + "`" + `ThrottlingException` + "`" + `` + "`" + `. For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html) in the *Amazon DynamoDB Developer Guide*.\n If read/write capacity mode is ` + "`" + `` + "`" + `PAY_PER_REQUEST` + "`" + `` + "`" + ` the value is set to 0.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "write_capacity_units": {
                    "computed": true,
                    "description": "The maximum number of writes consumed per second before DynamoDB returns a ` + "`" + `` + "`" + `ThrottlingException` + "`" + `` + "`" + `. For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html) in the *Amazon DynamoDB Developer Guide*.\n If read/write capacity mode is ` + "`" + `` + "`" + `PAY_PER_REQUEST` + "`" + `` + "`" + ` the value is set to 0.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "import_source_specification": {
        "computed": true,
        "description": "Specifies the properties of data being imported from the S3 bucket source to the table.\n  If you specify the ` + "`" + `` + "`" + `ImportSourceSpecification` + "`" + `` + "`" + ` property, and also specify either the ` + "`" + `` + "`" + `StreamSpecification` + "`" + `` + "`" + `, the ` + "`" + `` + "`" + `TableClass` + "`" + `` + "`" + ` property, or the ` + "`" + `` + "`" + `DeletionProtectionEnabled` + "`" + `` + "`" + ` property, the IAM entity creating/updating stack must have ` + "`" + `` + "`" + `UpdateTable` + "`" + `` + "`" + ` permission.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "input_compression_type": {
              "computed": true,
              "description": "Type of compression to be used on the input coming from the imported table.",
              "description_kind": "plain",
              "type": "string"
            },
            "input_format": {
              "computed": true,
              "description": "The format of the source data. Valid values for ` + "`" + `` + "`" + `ImportFormat` + "`" + `` + "`" + ` are ` + "`" + `` + "`" + `CSV` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `DYNAMODB_JSON` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ION` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            },
            "input_format_options": {
              "computed": true,
              "description": "Additional properties that specify how the input is formatted,",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "csv": {
                    "computed": true,
                    "description": "The options for imported source files in CSV format. The values are Delimiter and HeaderList.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delimiter": {
                          "computed": true,
                          "description": "The delimiter used for separating items in the CSV file being imported.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "header_list": {
                          "computed": true,
                          "description": "List of the headers used to specify a common header for all source CSV files being imported. If this field is specified then the first line of each CSV file is treated as data instead of the header. If this field is not specified the the first line of each CSV file is treated as the header.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "s3_bucket_source": {
              "computed": true,
              "description": "The S3 bucket that provides the source for the import.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_bucket": {
                    "computed": true,
                    "description": "The S3 bucket that is being imported from.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_bucket_owner": {
                    "computed": true,
                    "description": "The account number of the S3 bucket that is being imported from. If the bucket is owned by the requester this is optional.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_key_prefix": {
                    "computed": true,
                    "description": "The key prefix shared by all S3 Objects that are being imported.",
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
      "key_schema": {
        "computed": true,
        "description": "Specifies the attributes that make up the primary key for the table. The attributes in the ` + "`" + `` + "`" + `KeySchema` + "`" + `` + "`" + ` property must also be defined in the ` + "`" + `` + "`" + `AttributeDefinitions` + "`" + `` + "`" + ` property.",
        "description_kind": "plain",
        "type": "string"
      },
      "kinesis_stream_specification": {
        "computed": true,
        "description": "The Kinesis Data Streams configuration for the specified table.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "approximate_creation_date_time_precision": {
              "computed": true,
              "description": "The precision for the time and date that the stream was created.",
              "description_kind": "plain",
              "type": "string"
            },
            "stream_arn": {
              "computed": true,
              "description": "The ARN for a specific Kinesis data stream.\n Length Constraints: Minimum length of 37. Maximum length of 1024.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "local_secondary_indexes": {
        "computed": true,
        "description": "Local secondary indexes to be created on the table. You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "index_name": {
              "computed": true,
              "description": "The name of the local secondary index. The name must be unique among all other indexes on this table.",
              "description_kind": "plain",
              "type": "string"
            },
            "key_schema": {
              "computed": true,
              "description": "The complete key schema for the local secondary index, consisting of one or more pairs of attribute names and key types:\n  +   ` + "`" + `` + "`" + `HASH` + "`" + `` + "`" + ` - partition key\n  +   ` + "`" + `` + "`" + `RANGE` + "`" + `` + "`" + ` - sort key\n  \n  The partition key of an item is also known as its *hash attribute*. The term \"hash attribute\" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.\n The sort key of an item is also known as its *range attribute*. The term \"range attribute\" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_name": {
                    "computed": true,
                    "description": "The name of a key attribute.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_type": {
                    "computed": true,
                    "description": "The role that this key attribute will assume:\n  +   ` + "`" + `` + "`" + `HASH` + "`" + `` + "`" + ` - partition key\n  +   ` + "`" + `` + "`" + `RANGE` + "`" + `` + "`" + ` - sort key\n  \n  The partition key of an item is also known as its *hash attribute*. The term \"hash attribute\" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.\n The sort key of an item is also known as its *range attribute*. The term \"range attribute\" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "projection": {
              "computed": true,
              "description": "Represents attributes that are copied (projected) from the table into the local secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "non_key_attributes": {
                    "computed": true,
                    "description": "Represents the non-key attribute names which will be projected into the index.\n For local secondary indexes, the total count of ` + "`" + `` + "`" + `NonKeyAttributes` + "`" + `` + "`" + ` summed across all of the local secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "projection_type": {
                    "computed": true,
                    "description": "The set of attributes that are projected into the index:\n  +   ` + "`" + `` + "`" + `KEYS_ONLY` + "`" + `` + "`" + ` - Only the index and primary keys are projected into the index.\n  +   ` + "`" + `` + "`" + `INCLUDE` + "`" + `` + "`" + ` - In addition to the attributes described in ` + "`" + `` + "`" + `KEYS_ONLY` + "`" + `` + "`" + `, the secondary index will include other non-key attributes that you specify.\n  +   ` + "`" + `` + "`" + `ALL` + "`" + `` + "`" + ` - All of the table attributes are projected into the index.\n  \n When using the DynamoDB console, ` + "`" + `` + "`" + `ALL` + "`" + `` + "`" + ` is selected by default.",
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
      "point_in_time_recovery_specification": {
        "computed": true,
        "description": "The settings used to enable point in time recovery.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "point_in_time_recovery_enabled": {
              "computed": true,
              "description": "Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "provisioned_throughput": {
        "computed": true,
        "description": "Throughput for the specified table, which consists of values for ` + "`" + `` + "`" + `ReadCapacityUnits` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `WriteCapacityUnits` + "`" + `` + "`" + `. For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html). \n If you set ` + "`" + `` + "`" + `BillingMode` + "`" + `` + "`" + ` as ` + "`" + `` + "`" + `PROVISIONED` + "`" + `` + "`" + `, you must specify this property. If you set ` + "`" + `` + "`" + `BillingMode` + "`" + `` + "`" + ` as ` + "`" + `` + "`" + `PAY_PER_REQUEST` + "`" + `` + "`" + `, you cannot specify this property.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "read_capacity_units": {
              "computed": true,
              "description": "The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ` + "`" + `` + "`" + `ThrottlingException` + "`" + `` + "`" + `. For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html) in the *Amazon DynamoDB Developer Guide*.\n If read/write capacity mode is ` + "`" + `` + "`" + `PAY_PER_REQUEST` + "`" + `` + "`" + ` the value is set to 0.",
              "description_kind": "plain",
              "type": "number"
            },
            "write_capacity_units": {
              "computed": true,
              "description": "The maximum number of writes consumed per second before DynamoDB returns a ` + "`" + `` + "`" + `ThrottlingException` + "`" + `` + "`" + `. For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html) in the *Amazon DynamoDB Developer Guide*.\n If read/write capacity mode is ` + "`" + `` + "`" + `PAY_PER_REQUEST` + "`" + `` + "`" + ` the value is set to 0.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_policy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy_document": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "sse_specification": {
        "computed": true,
        "description": "Specifies the settings to enable server-side encryption.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_master_key_id": {
              "computed": true,
              "description": "The KMS key that should be used for the KMS encryption. To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB key ` + "`" + `` + "`" + `alias/aws/dynamodb` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            },
            "sse_enabled": {
              "computed": true,
              "description": "Indicates whether server-side encryption is done using an AWS managed key or an AWS owned key. If enabled (true), server-side encryption type is set to ` + "`" + `` + "`" + `KMS` + "`" + `` + "`" + ` and an AWS managed key is used (KMS charges apply). If disabled (false) or not specified, server-side encryption is set to AWS owned key.",
              "description_kind": "plain",
              "type": "bool"
            },
            "sse_type": {
              "computed": true,
              "description": "Server-side encryption type. The only supported value is:\n  +   ` + "`" + `` + "`" + `KMS` + "`" + `` + "`" + ` - Server-side encryption that uses KMSlong. The key is stored in your account and is managed by KMS (KMS charges apply).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "stream_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stream_specification": {
        "computed": true,
        "description": "The settings for the DDB table stream, which capture changes to items stored in the table.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_policy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "policy_document": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "stream_view_type": {
              "computed": true,
              "description": "When an item in the table is modified, ` + "`" + `` + "`" + `StreamViewType` + "`" + `` + "`" + ` determines what information is written to the stream for this table. Valid values for ` + "`" + `` + "`" + `StreamViewType` + "`" + `` + "`" + ` are:\n  +   ` + "`" + `` + "`" + `KEYS_ONLY` + "`" + `` + "`" + ` - Only the key attributes of the modified item are written to the stream.\n  +   ` + "`" + `` + "`" + `NEW_IMAGE` + "`" + `` + "`" + ` - The entire item, as it appears after it was modified, is written to the stream.\n  +   ` + "`" + `` + "`" + `OLD_IMAGE` + "`" + `` + "`" + ` - The entire item, as it appeared before it was modified, is written to the stream.\n  +   ` + "`" + `` + "`" + `NEW_AND_OLD_IMAGES` + "`" + `` + "`" + ` - Both the new and the old item images of the item are written to the stream.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "table_class": {
        "computed": true,
        "description": "The table class of the new table. Valid values are ` + "`" + `` + "`" + `STANDARD` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `STANDARD_INFREQUENT_ACCESS` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_name": {
        "computed": true,
        "description": "A name for the table. If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.\n For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag. Tag keys are case sensitive. Each DynamoDB table can only have up to one tag with the same key. If you try to add an existing tag (same key), the existing tag value will be updated to the new value.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag. Tag values are case-sensitive and can be null.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "time_to_live_specification": {
        "computed": true,
        "description": "Specifies the Time to Live (TTL) settings for the table.\n  For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attribute_name": {
              "computed": true,
              "description": "The name of the TTL attribute used to store the expiration time for items in the table.\n   + The ` + "`" + `` + "`" + `AttributeName` + "`" + `` + "`" + ` property is required when enabling the TTL, or when TTL is already enabled.\n  +  To update this property, you must first disable TTL and then enable TTL with the new attribute name.",
              "description_kind": "plain",
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description": "Indicates whether TTL is to be enabled (true) or disabled (false) on the table.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::DynamoDB::Table",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDynamodbTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDynamodbTable), &result)
	return &result
}
