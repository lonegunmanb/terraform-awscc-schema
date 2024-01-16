package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "topic_rule_payload": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_alarm": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "alarm_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "state_reason": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "state_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "cloudwatch_logs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "log_group_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "cloudwatch_metric": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metric_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "metric_namespace": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "metric_timestamp": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "metric_unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "metric_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
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
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "dynamo_db": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "hash_key_field": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "hash_key_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "hash_key_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "payload_field": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "range_key_field": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "range_key_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "range_key_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "table_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "elasticsearch": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "endpoint": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "index": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "firehose": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "delivery_stream_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "separator": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
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
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "service_name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "signing_region": {
                                      "description_kind": "plain",
                                      "required": true,
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
                        "confirmation_url": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "headers": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "url": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "iot_analytics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "channel_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "iot_events": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "input_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "message_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "iot_site_wise": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "put_asset_property_value_entries": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "asset_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "entry_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_alias": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_values": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "quality": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "timestamp": {
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "offset_in_nanos": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "time_in_seconds": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "required": true
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "boolean_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "double_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "integer_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "string_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "required": true
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "required": true
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "kafka": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_properties": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "destination_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "headers": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "partition": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "topic": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "kinesis": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "partition_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "stream_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "lambda": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "function_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "location": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "device_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "latitude": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "longitude": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
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
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "tracker_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "open_search": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "endpoint": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "index": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
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
                                "optional": true,
                                "type": "string"
                              },
                              "correlation_data": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "message_expiry": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "payload_format_indicator": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "response_topic": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "user_properties": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "required": true,
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
                        "qos": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "topic": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "s3": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "canned_acl": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "sns": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "message_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "target_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "sqs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "queue_url": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "use_base_64": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "step_functions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "execution_name_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "state_machine_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "timestream": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "database_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "dimensions": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "required": true
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "table_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "timestamp": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "unit": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
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
                "nesting_mode": "list"
              },
              "required": true
            },
            "aws_iot_sql_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "state_reason": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "state_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "cloudwatch_logs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "log_group_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "cloudwatch_metric": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metric_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "metric_namespace": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "metric_timestamp": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "metric_unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "metric_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
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
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "dynamo_db": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "hash_key_field": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "hash_key_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "hash_key_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "payload_field": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "range_key_field": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "range_key_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "range_key_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "table_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "elasticsearch": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "endpoint": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "index": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "firehose": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "delivery_stream_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "separator": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
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
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "service_name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "signing_region": {
                                      "description_kind": "plain",
                                      "required": true,
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
                        "confirmation_url": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "headers": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "url": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "iot_analytics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "channel_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "iot_events": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "input_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "message_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "iot_site_wise": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "put_asset_property_value_entries": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "asset_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "entry_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_alias": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_values": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "quality": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "timestamp": {
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "offset_in_nanos": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "time_in_seconds": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "required": true
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "boolean_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "double_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "integer_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "string_value": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "required": true
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "required": true
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "kafka": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_properties": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "destination_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "headers": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "partition": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "topic": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "kinesis": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "partition_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "stream_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "lambda": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "function_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "location": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "device_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "latitude": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "longitude": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
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
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "tracker_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "open_search": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "endpoint": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "index": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
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
                                "optional": true,
                                "type": "string"
                              },
                              "correlation_data": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "message_expiry": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "payload_format_indicator": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "response_topic": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "user_properties": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "required": true,
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
                        "qos": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "topic": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "s3": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "canned_acl": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "sns": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "message_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "target_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "sqs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "queue_url": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "use_base_64": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "step_functions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "execution_name_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "state_machine_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "timestream": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "database_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "dimensions": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "required": true
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "table_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "timestamp": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "unit": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
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
            "rule_disabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sql": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Resource Type definition for AWS::IoT::TopicRule",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotTopicRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotTopicRule), &result)
	return &result
}
