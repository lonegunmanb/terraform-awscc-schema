package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotTopicRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_name": {
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
      "topic_rule_payload": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_alarm": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "alarm_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "state_reason": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "state_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "cloudwatch_logs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "log_group_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "cloudwatch_metric": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metric_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_namespace": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_timestamp": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_unit": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "dynamo_d_bv_2": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "put_item": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "table_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "dynamo_db": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "hash_key_field": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "hash_key_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "hash_key_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "payload_field": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "range_key_field": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "range_key_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "range_key_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "table_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "elasticsearch": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "endpoint": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "index": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "firehose": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "delivery_stream_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "separator": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "http": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auth": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "sigv_4": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "role_arn": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "service_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "signing_region": {
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
                        "confirmation_url": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "headers": {
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
                        "url": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "iot_analytics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "channel_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "iot_events": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "input_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "message_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "iot_site_wise": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "put_asset_property_value_entries": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "asset_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "entry_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "property_alias": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "property_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "property_values": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "quality": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "timestamp": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "offset_in_nanos": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "time_in_seconds": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "value": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "boolean_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "double_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "integer_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "string_value": {
                                            "computed": true,
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
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "kafka": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_properties": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "destination_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "headers": {
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
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "partition": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "topic": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "kinesis": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "partition_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "stream_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "lambda": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "function_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "location": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "device_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "latitude": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "longitude": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "timestamp": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "unit": {
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
                            "nesting_mode": "single"
                          }
                        },
                        "tracker_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "open_search": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "endpoint": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "index": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "republish": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "headers": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "correlation_data": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "message_expiry": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "payload_format_indicator": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "response_topic": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "user_properties": {
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
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "qos": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "topic": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "s3": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "canned_acl": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "sns": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "message_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "target_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "sqs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "queue_url": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "use_base_64": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "step_functions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "execution_name_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "state_machine_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "timestream": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "database_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "dimensions": {
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
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "table_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "timestamp": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "unit": {
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
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "aws_iot_sql_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "error_action": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_alarm": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "alarm_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "state_reason": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "state_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "cloudwatch_logs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "log_group_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "cloudwatch_metric": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metric_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_namespace": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_timestamp": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_unit": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "dynamo_d_bv_2": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "put_item": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "table_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "dynamo_db": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "hash_key_field": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "hash_key_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "hash_key_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "payload_field": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "range_key_field": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "range_key_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "range_key_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "table_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "elasticsearch": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "endpoint": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "index": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "firehose": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "delivery_stream_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "separator": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "http": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auth": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "sigv_4": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "role_arn": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "service_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "signing_region": {
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
                        "confirmation_url": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "headers": {
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
                        "url": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "iot_analytics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "channel_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "iot_events": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "input_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "message_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "iot_site_wise": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "put_asset_property_value_entries": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "asset_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "entry_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "property_alias": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "property_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "property_values": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "quality": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "timestamp": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "offset_in_nanos": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "time_in_seconds": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "value": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "boolean_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "double_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "integer_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "string_value": {
                                            "computed": true,
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
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "kafka": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_properties": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "destination_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "headers": {
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
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "partition": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "topic": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "kinesis": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "partition_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "stream_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "lambda": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "function_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "location": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "device_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "latitude": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "longitude": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "timestamp": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "unit": {
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
                            "nesting_mode": "single"
                          }
                        },
                        "tracker_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "open_search": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "endpoint": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "index": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "republish": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "headers": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "correlation_data": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "message_expiry": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "payload_format_indicator": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "response_topic": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "user_properties": {
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
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "qos": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "topic": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "s3": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "canned_acl": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "sns": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "message_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "target_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "sqs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "queue_url": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "use_base_64": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "step_functions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "execution_name_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "state_machine_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "timestream": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "database_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "dimensions": {
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
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "table_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "timestamp": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "unit": {
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
            "rule_disabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "sql": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::IoT::TopicRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotTopicRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotTopicRule), &result)
	return &result
}
