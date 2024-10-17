package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueCrawler = `{
  "block": {
    "attributes": {
      "classifiers": {
        "computed": true,
        "description": "A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "configuration": {
        "computed": true,
        "description": "Crawler configuration information. This versioned JSON string allows users to specify aspects of a crawler's behavior.",
        "description_kind": "plain",
        "type": "string"
      },
      "crawler_security_configuration": {
        "computed": true,
        "description": "The name of the SecurityConfiguration structure to be used by this crawler.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_name": {
        "computed": true,
        "description": "The name of the database in which the crawler's output is stored.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the crawler.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lake_formation_configuration": {
        "computed": true,
        "description": "Specifies AWS Lake Formation configuration settings for the crawler",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_id": {
              "computed": true,
              "description": "Required for cross account crawls. For same account crawls as the target data, this can be left as null.",
              "description_kind": "plain",
              "type": "string"
            },
            "use_lake_formation_credentials": {
              "computed": true,
              "description": "Specifies whether to use AWS Lake Formation credentials for the crawler instead of the IAM role credentials.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "The name of the crawler.",
        "description_kind": "plain",
        "type": "string"
      },
      "recrawl_policy": {
        "computed": true,
        "description": "When crawling an Amazon S3 data source after the first crawl is complete, specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run. For more information, see Incremental Crawls in AWS Glue in the developer guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "recrawl_behavior": {
              "computed": true,
              "description": "Specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run. A value of CRAWL_EVERYTHING specifies crawling the entire dataset again. A value of CRAWL_NEW_FOLDERS_ONLY specifies crawling only folders that were added since the last crawler run. A value of CRAWL_EVENT_MODE specifies crawling only the changes identified by Amazon S3 events.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "role": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.",
        "description_kind": "plain",
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description": "A scheduling object using a cron statement to schedule an event.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "schedule_expression": {
              "computed": true,
              "description": "A cron expression used to specify the schedule. For more information, see Time-Based Schedules for Jobs and Crawlers. For example, to run something every day at 12:15 UTC, specify cron(15 12 * * ? *).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "schema_change_policy": {
        "computed": true,
        "description": "The policy that specifies update and delete behaviors for the crawler. The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The SchemaChangePolicy does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the SchemaChangePolicy on a crawler. The SchemaChangePolicy consists of two components, UpdateBehavior and DeleteBehavior.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "delete_behavior": {
              "computed": true,
              "description": "The deletion behavior when the crawler finds a deleted object. A value of LOG specifies that if a table or partition is found to no longer exist, do not delete it, only log that it was found to no longer exist. A value of DELETE_FROM_DATABASE specifies that if a table or partition is found to have been removed, delete it from the database. A value of DEPRECATE_IN_DATABASE specifies that if a table has been found to no longer exist, to add a property to the table that says 'DEPRECATED' and includes a timestamp with the time of deprecation.",
              "description_kind": "plain",
              "type": "string"
            },
            "update_behavior": {
              "computed": true,
              "description": "The update behavior when the crawler finds a changed schema. A value of LOG specifies that if a table or a partition already exists, and a change is detected, do not update it, only log that a change was detected. Add new tables and new partitions (including on existing tables). A value of UPDATE_IN_DATABASE specifies that if a table or partition already exists, and a change is detected, update it. Add new tables and partitions.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "table_prefix": {
        "computed": true,
        "description": "The prefix added to the names of tables that are created.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to use with this crawler.",
        "description_kind": "plain",
        "type": "string"
      },
      "targets": {
        "computed": true,
        "description": "Specifies data stores to crawl.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "catalog_targets": {
              "computed": true,
              "description": "Specifies AWS Glue Data Catalog targets.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connection_name": {
                    "computed": true,
                    "description": "The name of the connection for an Amazon S3-backed Data Catalog table to be a target of the crawl when using a Catalog connection type paired with a NETWORK Connection type.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "The name of the database to be synchronized.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "dlq_event_queue_arn": {
                    "computed": true,
                    "description": "A valid Amazon dead-letter SQS ARN. For example, arn:aws:sqs:region:account:deadLetterQueue.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "event_queue_arn": {
                    "computed": true,
                    "description": "A valid Amazon SQS ARN. For example, arn:aws:sqs:region:account:sqs.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tables": {
                    "computed": true,
                    "description": "A list of the tables to be synchronized.",
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
            "delta_targets": {
              "computed": true,
              "description": "Specifies an array of Delta data store targets.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connection_name": {
                    "computed": true,
                    "description": "The name of the connection to use to connect to the Delta table target.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "create_native_delta_table": {
                    "computed": true,
                    "description": "Specifies whether the crawler will create native tables, to allow integration with query engines that support querying of the Delta transaction log directly.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "delta_tables": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "write_manifest": {
                    "computed": true,
                    "description": "Specifies whether to write the manifest files to the Delta table path.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "dynamo_db_targets": {
              "computed": true,
              "description": "Specifies Amazon DynamoDB targets.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "path": {
                    "computed": true,
                    "description": "The name of the DynamoDB table to crawl.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "iceberg_targets": {
              "computed": true,
              "description": "Specifies Apache Iceberg data store targets.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connection_name": {
                    "computed": true,
                    "description": "The name of the connection to use to connect to the Iceberg target.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "exclusions": {
                    "computed": true,
                    "description": "A list of global patterns used to exclude from the crawl.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "maximum_traversal_depth": {
                    "computed": true,
                    "description": "The maximum depth of Amazon S3 paths that the crawler can traverse to discover the Iceberg metadata folder in your Amazon S3 path. Used to limit the crawler run time.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "paths": {
                    "computed": true,
                    "description": "One or more Amazon S3 paths that contains Iceberg metadata folders as s3://bucket/prefix .",
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
            "jdbc_targets": {
              "computed": true,
              "description": "Specifies JDBC targets.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connection_name": {
                    "computed": true,
                    "description": "The name of the connection to use to connect to the JDBC target.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enable_additional_metadata": {
                    "computed": true,
                    "description": "Specify a value of RAWTYPES or COMMENTS to enable additional metadata in table responses. RAWTYPES provides the native-level datatype. COMMENTS provides comments associated with a column or table in the database.\n\nIf you do not need additional metadata, keep the field empty.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exclusions": {
                    "computed": true,
                    "description": "A list of glob patterns used to exclude from the crawl. For more information, see Catalog Tables with a Crawler.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "path": {
                    "computed": true,
                    "description": "The path of the JDBC target.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "mongo_db_targets": {
              "computed": true,
              "description": "A list of Mongo DB targets.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connection_name": {
                    "computed": true,
                    "description": "The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description": "The path of the Amazon DocumentDB or MongoDB target (database/collection).",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "s3_targets": {
              "computed": true,
              "description": "Specifies Amazon Simple Storage Service (Amazon S3) targets.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connection_name": {
                    "computed": true,
                    "description": "The name of a connection which allows a job or crawler to access data in Amazon S3 within an Amazon Virtual Private Cloud environment (Amazon VPC).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "dlq_event_queue_arn": {
                    "computed": true,
                    "description": "A valid Amazon dead-letter SQS ARN. For example, arn:aws:sqs:region:account:deadLetterQueue.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "event_queue_arn": {
                    "computed": true,
                    "description": "A valid Amazon SQS ARN. For example, arn:aws:sqs:region:account:sqs.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "exclusions": {
                    "computed": true,
                    "description": "A list of glob patterns used to exclude from the crawl.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "path": {
                    "computed": true,
                    "description": "The path to the Amazon S3 target.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sample_size": {
                    "computed": true,
                    "description": "Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Glue::Crawler",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueCrawlerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueCrawler), &result)
	return &result
}
