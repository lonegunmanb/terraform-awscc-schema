package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMskReplicator = `{
  "block": {
    "attributes": {
      "current_version": {
        "computed": true,
        "description": "The current version of the MSK replicator.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A summary description of the replicator.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kafka_clusters": {
        "computed": true,
        "description": "Specifies a list of Kafka clusters which are targets of the replicator.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "amazon_msk_cluster": {
              "computed": true,
              "description": "Details of an Amazon MSK cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "msk_cluster_arn": {
                    "computed": true,
                    "description": "The ARN of an Amazon MSK cluster.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "apache_kafka_cluster": {
              "computed": true,
              "description": "Details of an Apache Kafka cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "apache_kafka_cluster_id": {
                    "computed": true,
                    "description": "The ID of the Apache Kafka cluster.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "bootstrap_broker_string": {
                    "computed": true,
                    "description": "The bootstrap broker string of the Apache Kafka cluster.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "client_authentication": {
              "computed": true,
              "description": "Details of the client authentication used by the Apache Kafka cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "sasl_scram": {
                    "computed": true,
                    "description": "Details for SASL/SCRAM client authentication.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "mechanism": {
                          "computed": true,
                          "description": "The SASL/SCRAM authentication mechanism.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "secret_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Secrets Manager secret.",
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
            "encryption_in_transit": {
              "computed": true,
              "description": "Details of encryption in transit to the Apache Kafka cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "encryption_type": {
                    "computed": true,
                    "description": "The type of encryption in transit to the Apache Kafka cluster.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "root_ca_certificate": {
                    "computed": true,
                    "description": "The root CA certificate.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vpc_config": {
              "computed": true,
              "description": "Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description": "The AWS security groups to associate with the elastic network interfaces in order to specify what the replicator has access to. If a security group is not specified, the default security group associated with the VPC is used.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnet_ids": {
                    "computed": true,
                    "description": "The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "set"
        }
      },
      "log_delivery": {
        "computed": true,
        "description": "Configuration for log delivery for the replicator.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "replicator_log_delivery": {
              "computed": true,
              "description": "The replicator logs configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_logs": {
                    "computed": true,
                    "description": "Details of the CloudWatch Logs destination for replicator logs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled": {
                          "computed": true,
                          "description": "Whether log delivery to CloudWatch Logs is enabled.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "log_group": {
                          "computed": true,
                          "description": "The CloudWatch log group that is the destination for log delivery.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "firehose": {
                    "computed": true,
                    "description": "Details of the Kinesis Data Firehose delivery stream that is the destination for replicator logs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delivery_stream": {
                          "computed": true,
                          "description": "The Firehose delivery stream that is the destination for log delivery.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "enabled": {
                          "computed": true,
                          "description": "Whether log delivery to Firehose is enabled.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "s3": {
                    "computed": true,
                    "description": "Details of the Amazon S3 destination for replicator logs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description": "The S3 bucket that is the destination for log delivery.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "enabled": {
                          "computed": true,
                          "description": "Whether log delivery to S3 is enabled.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "prefix": {
                          "computed": true,
                          "description": "The S3 prefix that is the destination for log delivery.",
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
      "replication_info_list": {
        "computed": true,
        "description": "A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "consumer_group_replication": {
              "computed": true,
              "description": "Configuration relating to consumer group replication.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "consumer_group_offset_sync_mode": {
                    "computed": true,
                    "description": "The consumer group offset synchronization mode.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "consumer_groups_to_exclude": {
                    "computed": true,
                    "description": "List of regular expression patterns indicating the consumer groups that should not be replicated.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "consumer_groups_to_replicate": {
                    "computed": true,
                    "description": "List of regular expression patterns indicating the consumer groups to copy.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "detect_and_copy_new_consumer_groups": {
                    "computed": true,
                    "description": "Whether to periodically check for new consumer groups.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "synchronise_consumer_group_offsets": {
                    "computed": true,
                    "description": "Whether to periodically write the translated offsets to __consumer_offsets topic in target cluster.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "source_kafka_cluster_arn": {
              "computed": true,
              "description": "Amazon Resource Name of the source Kafka cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "source_kafka_cluster_id": {
              "computed": true,
              "description": "The ID of the source Kafka cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_compression_type": {
              "computed": true,
              "description": "The type of compression to use writing records to target Kafka cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_kafka_cluster_arn": {
              "computed": true,
              "description": "Amazon Resource Name of the target Kafka cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_kafka_cluster_id": {
              "computed": true,
              "description": "The ID of the target Kafka cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "topic_replication": {
              "computed": true,
              "description": "Configuration relating to topic replication.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "copy_access_control_lists_for_topics": {
                    "computed": true,
                    "description": "Whether to periodically configure remote topic ACLs to match their corresponding upstream topics.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "copy_topic_configurations": {
                    "computed": true,
                    "description": "Whether to periodically configure remote topics to match their corresponding upstream topics.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "detect_and_copy_new_topics": {
                    "computed": true,
                    "description": "Whether to periodically check for new topics and partitions.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "starting_position": {
                    "computed": true,
                    "description": "Configuration for specifying the position in the topics to start replicating from.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description": "The type of replication starting position.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "topic_name_configuration": {
                    "computed": true,
                    "description": "Configuration for specifying replicated topic names should be the same as their corresponding upstream topics or prefixed with source cluster alias.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description": "The type of replicated topic name.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "topics_to_exclude": {
                    "computed": true,
                    "description": "List of regular expression patterns indicating the topics that should not be replicated.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "topics_to_replicate": {
                    "computed": true,
                    "description": "List of regular expression patterns indicating the topics to copy.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "set"
        }
      },
      "replicator_arn": {
        "computed": true,
        "description": "Amazon Resource Name for the created replicator.",
        "description_kind": "plain",
        "type": "string"
      },
      "replicator_name": {
        "computed": true,
        "description": "The name of the replicator.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_execution_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::MSK::Replicator",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMskReplicatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMskReplicator), &result)
	return &result
}
