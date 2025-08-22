package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventsRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the rule, such as arn:aws:events:us-east-2:123456789012:rule/example.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_bus_name": {
        "computed": true,
        "description": "The name or ARN of the event bus associated with the rule. If you omit this, the default event bus is used.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_pattern": {
        "computed": true,
        "description": "The event pattern of the rule. For more information, see Events and Event Patterns in the Amazon EventBridge User Guide.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the role that is used for target invocation.",
        "description_kind": "plain",
        "type": "string"
      },
      "schedule_expression": {
        "computed": true,
        "description": "The scheduling expression. For example, \"cron(0 20 * * ? *)\", \"rate(5 minutes)\". For more information, see Creating an Amazon EventBridge rule that runs on a schedule.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the event rule.",
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
      "targets": {
        "computed": true,
        "description": "Adds the specified targets to the specified rule, or updates the targets if they are already associated with the rule.\nTargets are the resources that are invoked when a rule is triggered.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "app_sync_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "graph_ql_operation": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "batch_parameters": {
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
            "ecs_parameters": {
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
                        "aws_vpc_configuration": {
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
                  "placement_strategies": {
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
                  "tag_list": {
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
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "input": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "input_path": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "input_transformer": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "input_paths_map": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
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
            "kinesis_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "partition_key_path": {
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
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "db_user": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secret_manager_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sql": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sqls": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "statement_name": {
                    "computed": true,
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
            "retry_policy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "maximum_event_age_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "maximum_retry_attempts": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "run_command_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "run_command_targets": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "values": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    }
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
            "sqs_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Events::Rule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEventsRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsRule), &result)
	return &result
}
