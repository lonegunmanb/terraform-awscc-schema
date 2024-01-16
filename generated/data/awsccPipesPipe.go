package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPipesPipe = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "current_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "desired_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enrichment": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enrichment_parameters": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "http_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "header_parameters": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "path_parameter_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "query_string_parameters": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "input_template": {
              "computed": true,
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
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "log_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs_log_destination": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "firehose_log_destination": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delivery_stream_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "include_execution_data": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "level": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_log_destination": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "bucket_owner": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_format": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_parameters": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "active_mq_broker_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "credentials": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "basic_auth": {
                          "computed": true,
                          "description": "Optional SecretManager ARN which stores the database credentials",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "queue_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "dynamo_db_stream_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "dead_letter_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "maximum_record_age_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "maximum_retry_attempts": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "on_partial_batch_item_failure": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "parallelization_factor": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "starting_position": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "filter_criteria": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "filters": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "pattern": {
                          "computed": true,
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
            "kinesis_stream_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "dead_letter_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "maximum_record_age_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "maximum_retry_attempts": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "on_partial_batch_item_failure": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "parallelization_factor": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "starting_position": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "starting_position_timestamp": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "managed_streaming_kafka_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "consumer_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "credentials": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_certificate_tls_auth": {
                          "computed": true,
                          "description": "Optional SecretManager ARN which stores the database credentials",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sasl_scram_512_auth": {
                          "computed": true,
                          "description": "Optional SecretManager ARN which stores the database credentials",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "starting_position": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "topic_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "rabbit_mq_broker_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "credentials": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "basic_auth": {
                          "computed": true,
                          "description": "Optional SecretManager ARN which stores the database credentials",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "queue_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "virtual_host": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "self_managed_kafka_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "additional_bootstrap_servers": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "consumer_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "credentials": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "basic_auth": {
                          "computed": true,
                          "description": "Optional SecretManager ARN which stores the database credentials",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "client_certificate_tls_auth": {
                          "computed": true,
                          "description": "Optional SecretManager ARN which stores the database credentials",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sasl_scram_256_auth": {
                          "computed": true,
                          "description": "Optional SecretManager ARN which stores the database credentials",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sasl_scram_512_auth": {
                          "computed": true,
                          "description": "Optional SecretManager ARN which stores the database credentials",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "server_root_ca_certificate": {
                    "computed": true,
                    "description": "Optional SecretManager ARN which stores the database credentials",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "starting_position": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "topic_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "vpc": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "security_group": {
                          "computed": true,
                          "description": "List of SecurityGroupId.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "subnets": {
                          "computed": true,
                          "description": "List of SubnetId.",
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
            "sqs_queue_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
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
      "state_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "target": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_parameters": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "batch_job_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "array_properties": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "size": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "container_overrides": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "command": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "environment": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
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
                            "nesting_mode": "list"
                          }
                        },
                        "instance_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "resource_requirements": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type": {
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
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "depends_on": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "job_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "job_definition": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "job_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "parameters": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "retry_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "attempts": {
                          "computed": true,
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
            "cloudwatch_logs_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_stream_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "timestamp": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ecs_task_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity_provider_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "base": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "capacity_provider": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "weight": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "enable_ecs_managed_tags": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "enable_execute_command": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "group": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "launch_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "awsvpc_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "assign_public_ip": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "security_groups": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnets": {
                                "computed": true,
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
                  "overrides": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "container_overrides": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "command": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "cpu": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "environment": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
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
                                  "nesting_mode": "list"
                                }
                              },
                              "environment_files": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "type": {
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
                                  "nesting_mode": "list"
                                }
                              },
                              "memory": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "memory_reservation": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "resource_requirements": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "type": {
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
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "cpu": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "ephemeral_storage": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "size_in_gi_b": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "execution_role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "inference_accelerator_overrides": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "device_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "device_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "memory": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "task_role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "placement_constraints": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "expression": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "placement_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "field": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "platform_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "propagate_tags": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "reference_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
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
                      "nesting_mode": "list"
                    }
                  },
                  "task_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "task_definition_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "event_bridge_event_bus_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "detail_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resources": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "source": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "time": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "http_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "header_parameters": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "path_parameter_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "query_string_parameters": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "input_template": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kinesis_stream_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "partition_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "lambda_function_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "invocation_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "redshift_data_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "computed": true,
                    "description": "Redshift Database",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "db_user": {
                    "computed": true,
                    "description": "Database user name",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secret_manager_arn": {
                    "computed": true,
                    "description": "Optional SecretManager ARN which stores the database credentials",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sqls": {
                    "computed": true,
                    "description": "A list of SQLs.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "statement_name": {
                    "computed": true,
                    "description": "A name for Redshift DataAPI statement which can be used as filter of ListStatement.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "with_event": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "sage_maker_pipeline_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "pipeline_parameter_list": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
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
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "sqs_queue_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message_deduplication_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "message_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "step_function_state_machine_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "invocation_type": {
                    "computed": true,
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
    "description": "Data Source schema for AWS::Pipes::Pipe",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPipesPipeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPipesPipe), &result)
	return &result
}
