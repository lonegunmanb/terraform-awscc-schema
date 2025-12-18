package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKafkaconnectConnector = `{
  "block": {
    "attributes": {
      "capacity": {
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
                  "max_worker_count": {
                    "computed": true,
                    "description": "The maximum number of workers for a connector.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "mcu_count": {
                    "computed": true,
                    "description": "Specifies how many MSK Connect Units (MCU) as the minimum scaling unit.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_worker_count": {
                    "computed": true,
                    "description": "The minimum number of workers for a connector.",
                    "description_kind": "plain",
                    "optional": true,
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
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
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
                          "optional": true,
                          "type": "number"
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
                    "optional": true,
                    "type": "number"
                  },
                  "worker_count": {
                    "computed": true,
                    "description": "Number of workers for a connector.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      "connector_arn": {
        "computed": true,
        "description": "Amazon Resource Name for the created Connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_configuration": {
        "description": "The configuration for the connector.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "connector_description": {
        "computed": true,
        "description": "A summary description of the connector.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connector_name": {
        "description": "The name of the connector.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kafka_cluster": {
        "description": "Details of how to connect to the Kafka cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "apache_kafka_cluster": {
              "description": "Details of how to connect to an Apache Kafka cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bootstrap_servers": {
                    "description": "The bootstrap servers string of the Apache Kafka cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "vpc": {
                    "description": "Information about a VPC used with the connector.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "security_groups": {
                          "description": "The AWS security groups to associate with the elastic network interfaces in order to specify what the connector has access to.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "subnets": {
                          "description": "The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "kafka_cluster_client_authentication": {
        "description": "Details of the client authentication used by the Kafka cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authentication_type": {
              "description": "The type of client authentication used to connect to the Kafka cluster. Value NONE means that no client authentication is used.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "kafka_cluster_encryption_in_transit": {
        "description": "Details of encryption in transit to the Kafka cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_type": {
              "description": "The type of encryption in transit to the Kafka cluster.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "kafka_connect_version": {
        "description": "The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.",
        "description_kind": "plain",
        "required": true,
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
                          "optional": true,
                          "type": "bool"
                        },
                        "log_group": {
                          "computed": true,
                          "description": "The CloudWatch log group that is the destination for log delivery.",
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
                    "description": "Details about delivering logs to Amazon Kinesis Data Firehose.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delivery_stream": {
                          "computed": true,
                          "description": "The Kinesis Data Firehose delivery stream that is the destination for log delivery.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "enabled": {
                          "computed": true,
                          "description": "Specifies whether the logs get sent to the specified Kinesis Data Firehose delivery stream.",
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
                    "description": "Details about delivering logs to Amazon S3.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description": "The name of the S3 bucket that is the destination for log delivery.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "enabled": {
                          "computed": true,
                          "description": "Specifies whether the logs get sent to the specified Amazon S3 destination.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "prefix": {
                          "computed": true,
                          "description": "The S3 prefix that is the destination for log delivery.",
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
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "network_type": {
        "computed": true,
        "description": "The network type of the Connector.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "plugins": {
        "description": "List of plugins to use with the connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_plugin": {
              "description": "Details about a custom plugin.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_plugin_arn": {
                    "description": "The Amazon Resource Name (ARN) of the custom plugin to use.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "revision": {
                    "description": "The revision of the custom plugin to use.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "service_execution_role_arn": {
        "description": "The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
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
              "optional": true,
              "type": "number"
            },
            "worker_configuration_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the worker configuration to use.",
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
    "description": "Resource Type definition for AWS::KafkaConnect::Connector",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKafkaconnectConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKafkaconnectConnector), &result)
	return &result
}
