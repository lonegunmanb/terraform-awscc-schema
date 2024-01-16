package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaEventSourceMapping = `{
  "block": {
    "attributes": {
      "amazon_managed_kafka_event_source_config": {
        "computed": true,
        "description": "Specific configuration settings for an MSK event source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "consumer_group_id": {
              "computed": true,
              "description": "The identifier for the Kafka Consumer Group to join.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "batch_size": {
        "computed": true,
        "description": "The maximum number of items to retrieve in a single batch.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "bisect_batch_on_function_error": {
        "computed": true,
        "description": "(Streams) If the function returns an error, split the batch in two and retry.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "destination_config": {
        "computed": true,
        "description": "(Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "on_failure": {
              "computed": true,
              "description": "The destination configuration for failed invocations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the destination resource.",
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
      "document_db_event_source_config": {
        "computed": true,
        "description": "Document db event source config.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "collection_name": {
              "computed": true,
              "description": "The collection name to connect to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "database_name": {
              "computed": true,
              "description": "The database name to connect to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "full_document": {
              "computed": true,
              "description": "Include full document in change stream response. The default option will only send the changes made to documents to Lambda. If you want the complete document sent to Lambda, set this to UpdateLookup.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "enabled": {
        "computed": true,
        "description": "Disables the event source mapping to pause polling and invocation.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "event_source_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the event source.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_criteria": {
        "computed": true,
        "description": "The filter criteria to control event filtering.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filters": {
              "computed": true,
              "description": "List of filters of this FilterCriteria",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "pattern": {
                    "computed": true,
                    "description": "The filter pattern that defines which events should be passed for invocations.",
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
      "function_name": {
        "description": "The name of the Lambda function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_response_types": {
        "computed": true,
        "description": "(Streams) A list of response types supported by the function.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Event Source Mapping Identifier UUID.",
        "description_kind": "plain",
        "type": "string"
      },
      "maximum_batching_window_in_seconds": {
        "computed": true,
        "description": "(Streams) The maximum amount of time to gather records before invoking the function, in seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maximum_record_age_in_seconds": {
        "computed": true,
        "description": "(Streams) The maximum age of a record that Lambda sends to a function for processing.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maximum_retry_attempts": {
        "computed": true,
        "description": "(Streams) The maximum number of times to retry when the function returns an error.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "parallelization_factor": {
        "computed": true,
        "description": "(Streams) The number of batches to process from each shard concurrently.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "queues": {
        "computed": true,
        "description": "(ActiveMQ) A list of ActiveMQ queues.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "scaling_config": {
        "computed": true,
        "description": "The scaling configuration for the event source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maximum_concurrency": {
              "computed": true,
              "description": "The maximum number of concurrent functions that the event source can invoke.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "self_managed_event_source": {
        "computed": true,
        "description": "Self-managed event source endpoints.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoints": {
              "computed": true,
              "description": "The endpoints for a self-managed event source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kafka_bootstrap_servers": {
                    "computed": true,
                    "description": "A list of Kafka server endpoints.",
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
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "self_managed_kafka_event_source_config": {
        "computed": true,
        "description": "Specific configuration settings for a Self-Managed Apache Kafka event source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "consumer_group_id": {
              "computed": true,
              "description": "The identifier for the Kafka Consumer Group to join.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "source_access_configurations": {
        "computed": true,
        "description": "A list of SourceAccessConfiguration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description": "The type of source access configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uri": {
              "computed": true,
              "description": "The URI for the source access configuration resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "starting_position": {
        "computed": true,
        "description": "The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB Streams sources.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "starting_position_timestamp": {
        "computed": true,
        "description": "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "topics": {
        "computed": true,
        "description": "(Kafka) A list of Kafka topics.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tumbling_window_in_seconds": {
        "computed": true,
        "description": "(Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::EventSourceMapping",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaEventSourceMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaEventSourceMapping), &result)
	return &result
}
