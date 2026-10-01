package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventsv2Subscriber = `{
  "block": {
    "attributes": {
      "batch_configuration": {
        "computed": true,
        "description": "Configuration for batching events into a single delivery to the target.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_batch_size": {
              "computed": true,
              "description": "The maximum number of events in a single batch delivered to the target. The maximum depends on the target: 500 for Kinesis Data Streams and Amazon Data Firehose, 100 for Lambda, Step Functions, and AWS::EventsV2::EventBus targets, 10 for Amazon SQS, Amazon SNS, and AWS::Events::EventBus targets, and 1 for API Gateway, API destinations, and universal service integration targets. The service rejects a value above the target's maximum. Fewer events may be delivered when the batch window elapses. When omitted, the default is 10 for Lambda and Step Functions targets and the target's maximum for other targets. The resolved value applied by the service is returned on read.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_batch_window_in_seconds": {
              "computed": true,
              "description": "The maximum time in seconds to wait for a batch to fill before delivering it, 0-300. The default is 0 (no wait). The resolved value applied by the service is returned on read.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "bus_name": {
        "computed": true,
        "description": "The name of the event bus this subscriber belongs to. This property is read-only.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "Creation timestamp (ISO-8601), read-only.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the subscriber. Control characters and Unicode line separators are not allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_bus_arn": {
        "description": "The ARN of the event bus this subscriber belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filter_configuration": {
        "computed": true,
        "description": "Configuration for filtering which events are delivered to the target. An event must match every filter to be delivered.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filters": {
              "computed": true,
              "description": "The list of filters, 1-50 entries. An event must match every filter to be delivered.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "pattern": {
                    "computed": true,
                    "description": "The event pattern, as a JSON string.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scope": {
                    "computed": true,
                    "description": "Which part of the event the pattern is evaluated against: DATA (the event payload), METADATA (event metadata), or SYSTEM_METADATA (service-generated metadata).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "language": {
              "computed": true,
              "description": "The filter language. The default is EVENT_BRIDGE_PATTERN.",
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
      "invoke_configuration": {
        "description": "Configuration for how the subscriber invokes its target, including the target ARN, the IAM role used to invoke it, and, optionally, the target-specific parameters object that matches the target type.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_bus_v2_parameters": {
              "computed": true,
              "description": "Parameters for forwarding events to another EventBridge event bus, used when TargetArn is an event bus ARN of the form arn:{partition}:events:{region}:{account}:event-busv2/{name}/{id}.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "deduplication_configuration": {
                    "computed": true,
                    "description": "Deduplication settings applied to the forwarded events on the downstream event bus.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "deduplication_type": {
                          "computed": true,
                          "description": "How duplicate events are detected: CONTENT_BASED deduplicates by a hash of the event content. To deduplicate by a caller-supplied token instead, omit DeduplicationConfiguration and set SystemMetadata.DeduplicationId.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "metadata": {
                    "computed": true,
                    "description": "Metadata forwarded with each event, as key-value string pairs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "system_metadata": {
                    "computed": true,
                    "description": "System metadata attached to each forwarded event, controlling FIFO ordering and deduplication on the downstream event bus.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "deduplication_id": {
                          "computed": true,
                          "description": "The deduplication ID for FIFO deduplication on the downstream event bus. Accepts a literal value or a JSONata expression.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "event_group_id": {
                          "computed": true,
                          "description": "The event group ID for FIFO ordering on the downstream event bus. Accepts a literal value or a JSONata expression.",
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
            "http_parameters": {
              "computed": true,
              "description": "Parameters for invoking an HTTP endpoint target, such as an Amazon API Gateway endpoint or an EventBridge API destination.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "header_parameters": {
                    "computed": true,
                    "description": "HTTP headers to add to the request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "invocation_timeout_seconds": {
                    "computed": true,
                    "description": "The timeout in seconds for each invocation of the target, written as a string. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path_parameter_values": {
                    "computed": true,
                    "description": "Values for the path parameters (wildcards) in the target URL, in order.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "query_string_parameters": {
                    "computed": true,
                    "description": "Query string parameters to add to the request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "kinesis_parameters": {
              "computed": true,
              "description": "Parameters for writing events to an Amazon Kinesis Data Streams target.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "explicit_hash_key": {
                    "computed": true,
                    "description": "An explicit hash key that overrides the partition key's shard assignment. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partition_key": {
                    "computed": true,
                    "description": "The partition key that determines which shard each record is written to. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "lambda_parameters": {
              "computed": true,
              "description": "Parameters for invoking an AWS Lambda function target.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "durable_execution_name": {
                    "computed": true,
                    "description": "A unique name for a durable function execution. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invocation_timeout_seconds": {
                    "computed": true,
                    "description": "The timeout in seconds for each invocation of the target, written as a string. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invocation_type": {
                    "computed": true,
                    "description": "How the function is invoked: EVENT (asynchronous) or REQUEST_RESPONSE (synchronous).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "qualifier": {
                    "computed": true,
                    "description": "The version or alias of the Lambda function to invoke. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tenant_id": {
                    "computed": true,
                    "description": "The tenant identifier for multi-tenant Lambda functions. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "role_arn": {
              "description": "The ARN of the IAM role the service assumes to invoke the target. The role must belong to the same account as the subscriber.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sns_parameters": {
              "computed": true,
              "description": "Parameters for publishing events to an Amazon SNS topic target.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message_attributes": {
                    "computed": true,
                    "description": "Custom message attributes to attach to each message; Amazon SNS subscription filter policies can match on them.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "binary_value": {
                          "computed": true,
                          "description": "The attribute value for the Binary data type, Base64-encoded.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "data_type": {
                          "computed": true,
                          "description": "The attribute data type. For Amazon SQS targets, specify String, Number, or Binary, optionally with a custom label suffix such as Number.float. For Amazon SNS targets, specify String, String.Array, Number, or Binary.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "string_value": {
                          "computed": true,
                          "description": "The attribute value for the String and Number data types (and String.Array for Amazon SNS targets).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "map"
                    },
                    "optional": true
                  },
                  "message_deduplication_id": {
                    "computed": true,
                    "description": "The message deduplication ID to use when the target is a FIFO topic. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "message_group_id": {
                    "computed": true,
                    "description": "The message group ID to use when the target is a FIFO topic. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "message_structure": {
                    "computed": true,
                    "description": "Set to json to send a different message per delivery protocol. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject": {
                    "computed": true,
                    "description": "The subject line to use for email-protocol subscriptions. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "sqs_parameters": {
              "computed": true,
              "description": "Parameters for sending events to an Amazon SQS queue target.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delay_seconds": {
                    "computed": true,
                    "description": "The delay in seconds for the message, written as a string. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "message_attributes": {
                    "computed": true,
                    "description": "Custom message attributes to attach to each message.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "binary_value": {
                          "computed": true,
                          "description": "The attribute value for the Binary data type, Base64-encoded.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "data_type": {
                          "computed": true,
                          "description": "The attribute data type. For Amazon SQS targets, specify String, Number, or Binary, optionally with a custom label suffix such as Number.float. For Amazon SNS targets, specify String, String.Array, Number, or Binary.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "string_value": {
                          "computed": true,
                          "description": "The attribute value for the String and Number data types (and String.Array for Amazon SNS targets).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "map"
                    },
                    "optional": true
                  },
                  "message_deduplication_id": {
                    "computed": true,
                    "description": "The message deduplication ID to use when the target is a FIFO queue. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "message_group_id": {
                    "computed": true,
                    "description": "The message group ID to use when the target is a FIFO queue. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "message_system_attributes": {
                    "computed": true,
                    "description": "Message system attributes to attach to each message, such as AWSTraceHeader.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "binary_value": {
                          "computed": true,
                          "description": "The attribute value for the Binary data type, Base64-encoded.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "data_type": {
                          "computed": true,
                          "description": "The attribute data type. For Amazon SQS targets, specify String, Number, or Binary, optionally with a custom label suffix such as Number.float. For Amazon SNS targets, specify String, String.Array, Number, or Binary.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "string_value": {
                          "computed": true,
                          "description": "The attribute value for the String and Number data types (and String.Array for Amazon SNS targets).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "map"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "step_functions_parameters": {
              "computed": true,
              "description": "Parameters for starting an AWS Step Functions state machine execution target.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "invocation_timeout_seconds": {
                    "computed": true,
                    "description": "The timeout in seconds for each invocation of the target, written as a string. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invocation_type": {
                    "computed": true,
                    "description": "How the execution is started: EVENT (StartExecution, asynchronous) or REQUEST_RESPONSE (StartSyncExecution, synchronous).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "A name for the execution. Must be unique for the account, Region, and state machine. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "trace_header": {
                    "computed": true,
                    "description": "The AWS X-Ray trace header for distributed tracing. Accepts a literal value or a JSONata expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "target_arn": {
              "description": "The Amazon Resource Name (ARN) of the target that the subscriber invokes. For universal service integration targets, use the form arn:{partition}:events:::aws-sdk:{service}:{apiAction}.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "universal_target_parameters": {
              "computed": true,
              "description": "Parameters for invoking an AWS service API as a universal service integration target, used when TargetArn has the form arn:{partition}:events:::aws-sdk:{service}:{apiAction}.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "input": {
                    "computed": true,
                    "description": "JSON string or JSONata expression that produces the API request. Supports {% ... %} JSONata expressions for dynamic values from the event.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invocation_timeout_seconds": {
                    "computed": true,
                    "description": "Timeout in seconds for each invocation of the target (1-30, default 30). Must be a literal integer written as a string; JSONata expressions are not supported for this field.",
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
        "required": true
      },
      "last_modified_time": {
        "computed": true,
        "description": "Last-modification timestamp (ISO-8601), read-only.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_configuration": {
        "computed": true,
        "description": "Delivery logging configuration for the subscriber.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "include_payload": {
              "computed": true,
              "description": "Whether the event payload is included in emitted log records: FULL includes it in every emitted record, and ON_ERROR_ONLY includes it only in error records. The default is ON_ERROR_ONLY.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "level": {
              "computed": true,
              "description": "The minimum log level: OFF (no logging), ERROR, or INFO. Records below this level are not emitted. The default is OFF.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "name": {
        "description": "The name of the subscriber. The first character must be alphanumeric; the remaining characters may also include '.', '-', and '_'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "on_failure_configuration": {
        "computed": true,
        "description": "The destination for events that could not be delivered to the target.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "The ARN of the destination that receives events that could not be delivered. An Amazon SQS queue is the supported destination.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "point_in_time_configuration": {
        "computed": true,
        "description": "The point in time to start delivering events from. Used when StartingPosition is POINT_IN_TIME.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_point": {
              "computed": true,
              "description": "An optional time to stop delivering events at, in seconds since the Unix epoch.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "point_type": {
              "computed": true,
              "description": "Where to start: HORIZON starts from the earliest available event; TIMESTAMP starts from the StartingPoint timestamp.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "starting_point": {
              "computed": true,
              "description": "The time to start delivering events from, in seconds since the Unix epoch. Required when PointType is TIMESTAMP.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "resume_position": {
        "computed": true,
        "description": "Resume-time control, never returned by the service. Applied only when an update transitions State from STOPPED to RUNNING: LAST_PROCESSED (default) resumes from the last processed event, LATEST skips to the newest. Ignored on create and on any update that does not perform that transition.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retry_policy": {
        "computed": true,
        "description": "The retry policy for failed deliveries to the target.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_event_age_in_seconds": {
              "computed": true,
              "description": "The maximum age of an event in seconds, 60-86400 (24 hours). When an event reaches this age, retries stop; if OnFailureConfiguration is set, the event is delivered to that destination, otherwise it is dropped. The default is 300.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_retry_attempts": {
              "computed": true,
              "description": "The maximum number of retry attempts, 0-185. When the attempts are exhausted, retries stop; if OnFailureConfiguration is set, the event is delivered to that destination, otherwise it is dropped. The default is 5.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "retry_strategy": {
              "computed": true,
              "description": "Which errors are retried. ALL retries all errors. The default is ALL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "starting_position": {
        "computed": true,
        "description": "Where the subscriber starts reading events: LATEST starts from the newest events; POINT_IN_TIME starts from the point specified in PointInTimeConfiguration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The run state of the subscriber. Events are delivered only while the state is RUNNING. Setting the state to STOPPED pauses delivery. When an update sets a stopped subscriber back to RUNNING, ResumePosition controls where delivery resumes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscriber_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the subscriber.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the subscriber.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key. For each resource, each tag key must be unique and each key can have only one value; keys are case sensitive. A key cannot begin or end with a whitespace character; whitespace inside the key is allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value. May be empty. A value cannot begin or end with a whitespace character; whitespace inside the value is allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "transformer": {
        "computed": true,
        "description": "Configuration for transforming events before delivery to the target.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "jsonata_configuration": {
              "computed": true,
              "description": "The JSONata expression configuration. Required when Type is JSONATA.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "expression": {
                    "computed": true,
                    "description": "The JSONata expression that transforms the event, enclosed in {% %} delimiters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "computed": true,
              "description": "The transform type: RAW delivers the event payload only; WITH_METADATA delivers the event with its metadata envelope; JSONATA delivers the output of the JSONata expression in JsonataConfiguration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "type": {
        "computed": true,
        "description": "The delivery ordering mode of the subscriber. FIFO delivers events in order within an event group; UNORDERED delivers without an ordering guarantee.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::EventsV2::Subscriber, an Amazon EventBridge subscription that delivers events from an event bus to a target. Canonical identity is the combination of bus, name, and target. Replacement uses delete_then_create.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventsv2SubscriberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsv2Subscriber), &result)
	return &result
}
