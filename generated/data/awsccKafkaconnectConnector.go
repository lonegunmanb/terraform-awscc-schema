package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKafkaconnectConnector = `{
  "block": {
    "attributes": {
      "capacity": {
        "computed": true,
        "description": "Information about the capacity allocated to the connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_scaling": {
              "computed": true,
              "description": "Details about auto scaling of a connector.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_autoscaling_task_count": {
                    "computed": true,
                    "description": "The maximum number of tasks allocated to the connector during autoscaling operations.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "max_worker_count": {
                    "computed": true,
                    "description": "The maximum number of workers for a connector.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "mcu_count": {
                    "computed": true,
                    "description": "Specifies how many MSK Connect Units (MCU) as the minimum scaling unit.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min_worker_count": {
                    "computed": true,
                    "description": "The minimum number of workers for a connector.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "scale_in_policy": {
                    "computed": true,
                    "description": "Information about the scale in policy of the connector.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cpu_utilization_percentage": {
                          "computed": true,
                          "description": "Specifies the CPU utilization percentage threshold at which connector scale in should trigger.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "scale_out_policy": {
                    "computed": true,
                    "description": "Information about the scale out policy of the connector.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cpu_utilization_percentage": {
                          "computed": true,
                          "description": "Specifies the CPU utilization percentage threshold at which connector scale out should trigger.",
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
            "provisioned_capacity": {
              "computed": true,
              "description": "Details about a fixed capacity allocated to a connector.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "mcu_count": {
                    "computed": true,
                    "description": "Specifies how many MSK Connect Units (MCU) are allocated to the connector.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "worker_count": {
                    "computed": true,
                    "description": "Number of workers for a connector.",
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
      "connector_arn": {
        "computed": true,
        "description": "Amazon Resource Name for the created Connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_configuration": {
        "computed": true,
        "description": "The configuration for the connector.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "connector_description": {
        "computed": true,
        "description": "A summary description of the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_name": {
        "computed": true,
        "description": "The name of the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kafka_cluster": {
        "computed": true,
        "description": "Details of how to connect to the Kafka cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "apache_kafka_cluster": {
              "computed": true,
              "description": "Details of how to connect to an Apache Kafka cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bootstrap_servers": {
                    "computed": true,
                    "description": "The bootstrap servers string of the Apache Kafka cluster.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "vpc": {
                    "computed": true,
                    "description": "Information about a VPC used with the connector.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "security_groups": {
                          "computed": true,
                          "description": "The AWS security groups to associate with the elastic network interfaces in order to specify what the connector has access to.",
                          "description_kind": "plain",
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "subnets": {
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
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "kafka_cluster_client_authentication": {
        "computed": true,
        "description": "Details of the client authentication used by the Kafka cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authentication_type": {
              "computed": true,
              "description": "The type of client authentication used to connect to the Kafka cluster. Value NONE means that no client authentication is used.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "kafka_cluster_encryption_in_transit": {
        "computed": true,
        "description": "Details of encryption in transit to the Kafka cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_type": {
              "computed": true,
              "description": "The type of encryption in transit to the Kafka cluster.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "kafka_connect_version": {
        "computed": true,
        "description": "The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_delivery": {
        "computed": true,
        "description": "Details of what logs are delivered and where they are delivered.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "worker_log_delivery": {
              "computed": true,
              "description": "Specifies where worker logs are delivered.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_logs": {
                    "computed": true,
                    "description": "Details about delivering logs to Amazon CloudWatch Logs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled": {
                          "computed": true,
                          "description": "Specifies whether the logs get sent to the specified CloudWatch Logs destination.",
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
                    "description": "Details about delivering logs to Amazon Kinesis Data Firehose.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delivery_stream": {
                          "computed": true,
                          "description": "The Kinesis Data Firehose delivery stream that is the destination for log delivery.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "enabled": {
                          "computed": true,
                          "description": "Specifies whether the logs get sent to the specified Kinesis Data Firehose delivery stream.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "s3": {
                    "computed": true,
                    "description": "Details about delivering logs to Amazon S3.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description": "The name of the S3 bucket that is the destination for log delivery.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "enabled": {
                          "computed": true,
                          "description": "Specifies whether the logs get sent to the specified Amazon S3 destination.",
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
      "network_type": {
        "computed": true,
        "description": "The network type of the Connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "plugins": {
        "computed": true,
        "description": "List of plugins to use with the connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_plugin": {
              "computed": true,
              "description": "Details about a custom plugin.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_plugin_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the custom plugin to use.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "revision": {
                    "computed": true,
                    "description": "The revision of the custom plugin to use.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "set"
        }
      },
      "service_execution_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.",
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
      },
      "worker_configuration": {
        "computed": true,
        "description": "Specifies the worker configuration to use with the connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "revision": {
              "computed": true,
              "description": "The revision of the worker configuration to use.",
              "description_kind": "plain",
              "type": "number"
            },
            "worker_configuration_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the worker configuration to use.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::KafkaConnect::Connector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKafkaconnectConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKafkaconnectConnector), &result)
	return &result
}
