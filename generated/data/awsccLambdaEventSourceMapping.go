package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaEventSourceMapping = `{
  "block": {
    "attributes": {
      "amazon_managed_kafka_event_source_config": {
        "computed": true,
        "description": "Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "consumer_group_id": {
              "computed": true,
              "description": "The identifier for the Kafka consumer group to join. The consumer group ID must be unique among all your Kafka event sources. After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value. For more information, see [Customizable consumer group ID](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-consumer-group-id).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "batch_size": {
        "computed": true,
        "description": "The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function. Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).\n  +   *Amazon Kinesis* – Default 100. Max 10,000.\n  +   *Amazon DynamoDB Streams* – Default 100. Max 10,000.\n  +   *Amazon Simple Queue Service* – Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.\n  +   *Amazon Managed Streaming for Apache Kafka* – Default 100. Max 10,000.\n  +   *Self-managed Apache Kafka* – Default 100. Max 10,000.\n  +   *Amazon MQ (ActiveMQ and RabbitMQ)* – Default 100. Max 10,000.\n  +   *DocumentDB* – Default 100. Max 10,000.",
        "description_kind": "plain",
        "type": "number"
      },
      "bisect_batch_on_function_error": {
        "computed": true,
        "description": "(Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry. The default value is false.",
        "description_kind": "plain",
        "type": "bool"
      },
      "destination_config": {
        "computed": true,
        "description": "(Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka event sources only) A configuration object that specifies the destination of an event after Lambda processes it.",
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
                    "description": "The Amazon Resource Name (ARN) of the destination resource.\n To retain records of [asynchronous invocations](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations), you can configure an Amazon SNS topic, Amazon SQS queue, Lambda function, or Amazon EventBridge event bus as the destination.\n To retain records of failed invocations from [Kinesis and DynamoDB event sources](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#event-source-mapping-destinations), you can configure an Amazon SNS topic or Amazon SQS queue as the destination.\n To retain records of failed invocations from [self-managed Kafka](https://docs.aws.amazon.com/lambda/latest/dg/with-kafka.html#services-smaa-onfailure-destination) or [Amazon MSK](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-onfailure-destination), you can configure an Amazon SNS topic, Amazon SQS queue, or Amazon S3 bucket as the destination.",
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
      "document_db_event_source_config": {
        "computed": true,
        "description": "Specific configuration settings for a DocumentDB event source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "collection_name": {
              "computed": true,
              "description": "The name of the collection to consume within the database. If you do not specify a collection, Lambda consumes all collections.",
              "description_kind": "plain",
              "type": "string"
            },
            "database_name": {
              "computed": true,
              "description": "The name of the database to consume within the DocumentDB cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "full_document": {
              "computed": true,
              "description": "Determines what DocumentDB sends to your event stream during document update operations. If set to UpdateLookup, DocumentDB sends a delta describing the changes, along with a copy of the entire document. Otherwise, DocumentDB sends only a partial document that contains the changes.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "enabled": {
        "computed": true,
        "description": "When true, the event source mapping is active. When false, Lambda pauses polling and invocation.\n Default: True",
        "description_kind": "plain",
        "type": "bool"
      },
      "event_source_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the event source.\n  +   *Amazon Kinesis* – The ARN of the data stream or a stream consumer.\n  +   *Amazon DynamoDB Streams* – The ARN of the stream.\n  +   *Amazon Simple Queue Service* – The ARN of the queue.\n  +   *Amazon Managed Streaming for Apache Kafka* – The ARN of the cluster or the ARN of the VPC connection (for [cross-account event source mappings](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#msk-multi-vpc)).\n  +   *Amazon MQ* – The ARN of the broker.\n  +   *Amazon DocumentDB* – The ARN of the DocumentDB change stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_source_mapping_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "filter_criteria": {
        "computed": true,
        "description": "An object that defines the filter criteria that determine whether Lambda should process an event. For more information, see [Lambda event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filters": {
              "computed": true,
              "description": "A list of filters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "pattern": {
                    "computed": true,
                    "description": "A filter pattern. For more information on the syntax of a filter pattern, see [Filter rule syntax](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-syntax).",
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
      "function_name": {
        "computed": true,
        "description": "The name or ARN of the Lambda function.\n  **Name formats**\n +   *Function name* – ` + "`" + `` + "`" + `MyFunction` + "`" + `` + "`" + `.\n  +   *Function ARN* – ` + "`" + `` + "`" + `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` + "`" + `` + "`" + `.\n  +   *Version or Alias ARN* – ` + "`" + `` + "`" + `arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD` + "`" + `` + "`" + `.\n  +   *Partial ARN* – ` + "`" + `` + "`" + `123456789012:function:MyFunction` + "`" + `` + "`" + `.\n  \n The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_response_types": {
        "computed": true,
        "description": "(Streams and SQS) A list of current response type enums applied to the event source mapping.\n Valid Values: ` + "`" + `` + "`" + `ReportBatchItemFailures` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maximum_batching_window_in_seconds": {
        "computed": true,
        "description": "The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.\n  *Default (, , event sources)*: 0\n  *Default (, Kafka, , event sources)*: 500 ms\n  *Related setting:* For SQS event sources, when you set ` + "`" + `` + "`" + `BatchSize` + "`" + `` + "`" + ` to a value greater than 10, you must set ` + "`" + `` + "`" + `MaximumBatchingWindowInSeconds` + "`" + `` + "`" + ` to at least 1.",
        "description_kind": "plain",
        "type": "number"
      },
      "maximum_record_age_in_seconds": {
        "computed": true,
        "description": "(Kinesis and DynamoDB Streams only) Discard records older than the specified age. The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records.\n  The minimum valid value for maximum record age is 60s. Although values less than 60 and greater than -1 fall within the parameter's absolute range, they are not allowed",
        "description_kind": "plain",
        "type": "number"
      },
      "maximum_retry_attempts": {
        "computed": true,
        "description": "(Kinesis and DynamoDB Streams only) Discard records after the specified number of retries. The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source.",
        "description_kind": "plain",
        "type": "number"
      },
      "parallelization_factor": {
        "computed": true,
        "description": "(Kinesis and DynamoDB Streams only) The number of batches to process concurrently from each shard. The default value is 1.",
        "description_kind": "plain",
        "type": "number"
      },
      "queues": {
        "computed": true,
        "description": "(Amazon MQ) The name of the Amazon MQ broker destination queue to consume.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "scaling_config": {
        "computed": true,
        "description": "(Amazon SQS only) The scaling configuration for the event source. For more information, see [Configuring maximum concurrency for Amazon SQS event sources](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maximum_concurrency": {
              "computed": true,
              "description": "Limits the number of concurrent instances that the SQS event source can invoke.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "self_managed_event_source": {
        "computed": true,
        "description": "The self-managed Apache Kafka cluster for your event source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoints": {
              "computed": true,
              "description": "The list of bootstrap servers for your Kafka brokers in the following format: ` + "`" + `` + "`" + `\"KafkaBootstrapServers\": [\"abc.xyz.com:xxxx\",\"abc2.xyz.com:xxxx\"]` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kafka_bootstrap_servers": {
                    "computed": true,
                    "description": "The list of bootstrap servers for your Kafka brokers in the following format: ` + "`" + `` + "`" + `\"KafkaBootstrapServers\": [\"abc.xyz.com:xxxx\",\"abc2.xyz.com:xxxx\"]` + "`" + `` + "`" + `.",
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
      "self_managed_kafka_event_source_config": {
        "computed": true,
        "description": "Specific configuration settings for a self-managed Apache Kafka event source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "consumer_group_id": {
              "computed": true,
              "description": "The identifier for the Kafka consumer group to join. The consumer group ID must be unique among all your Kafka event sources. After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value. For more information, see [Customizable consumer group ID](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-consumer-group-id).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "source_access_configurations": {
        "computed": true,
        "description": "An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description": "The type of authentication protocol, VPC components, or virtual host for your event source. For example: ` + "`" + `` + "`" + `\"Type\":\"SASL_SCRAM_512_AUTH\"` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `BASIC_AUTH` + "`" + `` + "`" + ` – (Amazon MQ) The ASMlong secret that stores your broker credentials.\n  +   ` + "`" + `` + "`" + `BASIC_AUTH` + "`" + `` + "`" + ` – (Self-managed Apache Kafka) The Secrets Manager ARN of your secret key used for SASL/PLAIN authentication of your Apache Kafka brokers.\n  +   ` + "`" + `` + "`" + `VPC_SUBNET` + "`" + `` + "`" + ` – (Self-managed Apache Kafka) The subnets associated with your VPC. Lambda connects to these subnets to fetch data from your self-managed Apache Kafka cluster.\n  +   ` + "`" + `` + "`" + `VPC_SECURITY_GROUP` + "`" + `` + "`" + ` – (Self-managed Apache Kafka) The VPC security group used to manage access to your self-managed Apache Kafka brokers.\n  +   ` + "`" + `` + "`" + `SASL_SCRAM_256_AUTH` + "`" + `` + "`" + ` – (Self-managed Apache Kafka) The Secrets Manager ARN of your secret key used for SASL SCRAM-256 authentication of your self-managed Apache Kafka brokers.\n  +   ` + "`" + `` + "`" + `SASL_SCRAM_512_AUTH` + "`" + `` + "`" + ` – (Amazon MSK, Self-managed Apache Kafka) The Secrets Manager ARN of your secret key used for SASL SCRAM-512 authentication of your self-managed Apache Kafka brokers.\n  +   ` + "`" + `` + "`" + `VIRTUAL_HOST` + "`" + `` + "`" + ` –- (RabbitMQ) The name of the virtual host in your RabbitMQ broker. Lambda uses this RabbitMQ host as the event source. This property cannot be specified in an UpdateEventSourceMapping API call.\n  +   ` + "`" + `` + "`" + `CLIENT_CERTIFICATE_TLS_AUTH` + "`" + `` + "`" + ` – (Amazon MSK, self-managed Apache Kafka) The Secrets Manager ARN of your secret key containing the certificate chain (X.509 PEM), private key (PKCS#8 PEM), and private key password (optional) used for mutual TLS authentication of your MSK/Apache Kafka brokers.\n  +   ` + "`" + `` + "`" + `SERVER_ROOT_CA_CERTIFICATE` + "`" + `` + "`" + ` – (Self-managed Apache Kafka) The Secrets Manager ARN of your secret key containing the root CA certificate (X.509 PEM) used for TLS encryption of your Apache Kafka brokers.",
              "description_kind": "plain",
              "type": "string"
            },
            "uri": {
              "computed": true,
              "description": "The value for your chosen configuration in ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + `. For example: ` + "`" + `` + "`" + `\"URI\": \"arn:aws:secretsmanager:us-east-1:01234567890:secret:MyBrokerSecretName\"` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "starting_position": {
        "computed": true,
        "description": "The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB.\n  +   *LATEST* - Read only new records.\n  +   *TRIM_HORIZON* - Process all available records.\n  +   *AT_TIMESTAMP* - Specify a time from which to start reading records.",
        "description_kind": "plain",
        "type": "string"
      },
      "starting_position_timestamp": {
        "computed": true,
        "description": "With ` + "`" + `` + "`" + `StartingPosition` + "`" + `` + "`" + ` set to ` + "`" + `` + "`" + `AT_TIMESTAMP` + "`" + `` + "`" + `, the time from which to start reading, in Unix time seconds. ` + "`" + `` + "`" + `StartingPositionTimestamp` + "`" + `` + "`" + ` cannot be in the future.",
        "description_kind": "plain",
        "type": "number"
      },
      "topics": {
        "computed": true,
        "description": "The name of the Kafka topic.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tumbling_window_in_seconds": {
        "computed": true,
        "description": "(Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources. A value of 0 seconds indicates no tumbling window.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::Lambda::EventSourceMapping",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaEventSourceMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaEventSourceMapping), &result)
	return &result
}
