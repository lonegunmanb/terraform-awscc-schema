package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSchedulerSchedule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the schedule.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the schedule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_date": {
        "computed": true,
        "description": "The date, in UTC, before which the schedule can invoke its target. Depending on the schedule's recurrence expression, invocations might stop on, or before, the EndDate you specify.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flexible_time_window": {
        "description": "Flexible time window allows configuration of a window within which a schedule can be invoked",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maximum_window_in_minutes": {
              "computed": true,
              "description": "The maximum time window during which a schedule can be invoked.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "mode": {
              "description": "Determines whether the schedule is executed within a flexible time window.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "group_name": {
        "computed": true,
        "description": "The name of the schedule group to associate with this schedule. If you omit this, the default schedule group is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN for a KMS Key that will be used to encrypt customer data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule_expression": {
        "description": "The scheduling expression.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_expression_timezone": {
        "computed": true,
        "description": "The timezone in which the scheduling expression is evaluated.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_date": {
        "computed": true,
        "description": "The date, in UTC, after which the schedule can begin invoking its target. Depending on the schedule's recurrence expression, invocations might occur on, or after, the StartDate you specify.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Specifies whether the schedule is enabled or disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target": {
        "description": "The schedule target.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "description": "The Amazon Resource Name (ARN) of the target.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "dead_letter_config": {
              "computed": true,
              "description": "A DeadLetterConfig object that contains information about a dead-letter queue configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The ARN of the SQS queue specified as the target for the dead-letter queue.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "ecs_parameters": {
              "computed": true,
              "description": "The custom parameters to be used when the target is an Amazon ECS task.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity_provider_strategy": {
                    "computed": true,
                    "description": "The capacity provider strategy to use for the task.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "base": {
                          "computed": true,
                          "description": "The base value designates how many tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. If no value is specified, the default value of 0 is used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "capacity_provider": {
                          "computed": true,
                          "description": "The short name of the capacity provider.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "weight": {
                          "computed": true,
                          "description": "The weight value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider. The weight value is taken into consideration after the base value, if defined, is satisfied.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "enable_ecs_managed_tags": {
                    "computed": true,
                    "description": "Specifies whether to enable Amazon ECS managed tags for the task. For more information, see Tagging Your Amazon ECS Resources in the Amazon Elastic Container Service Developer Guide.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_execute_command": {
                    "computed": true,
                    "description": "Whether or not to enable the execute command functionality for the containers in this task. If true, this enables execute command functionality on all containers in the task.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "group": {
                    "computed": true,
                    "description": "Specifies an ECS task group for the task. The maximum length is 255 characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "launch_type": {
                    "computed": true,
                    "description": "Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. The FARGATE value is supported only in the Regions where AWS Fargate with Amazon ECS is supported. For more information, see AWS Fargate on Amazon ECS in the Amazon Elastic Container Service Developer Guide.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "network_configuration": {
                    "computed": true,
                    "description": "This structure specifies the network configuration for an ECS task.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "awsvpc_configuration": {
                          "computed": true,
                          "description": "This structure specifies the VPC subnets and security groups for the task, and whether a public IP address is to be used. This structure is relevant only for ECS tasks that use the awsvpc network mode.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "assign_public_ip": {
                                "computed": true,
                                "description": "Specifies whether the task's elastic network interface receives a public IP address. You can specify ENABLED only when LaunchType in EcsParameters is set to FARGATE.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "security_groups": {
                                "computed": true,
                                "description": "Specifies the security groups associated with the task. These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnets": {
                                "computed": true,
                                "description": "Specifies the subnets associated with the task. These subnets must all be in the same VPC. You can specify as many as 16 subnets.",
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
                  "placement_constraints": {
                    "computed": true,
                    "description": "An array of placement constraint objects to use for the task. You can specify up to 10 constraints per task (including constraints in the task definition and those specified at runtime).",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "expression": {
                          "computed": true,
                          "description": "A cluster query language expression to apply to the constraint. You cannot specify an expression if the constraint type is distinctInstance. To learn more, see Cluster Query Language in the Amazon Elastic Container Service Developer Guide.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "The type of constraint. Use distinctInstance to ensure that each task in a particular group is running on a different container instance. Use memberOf to restrict the selection to a group of valid candidates.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "placement_strategy": {
                    "computed": true,
                    "description": "The placement strategy objects to use for the task. You can specify a maximum of five strategy rules per task.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "field": {
                          "computed": true,
                          "description": "The field to apply the placement strategy against. For the spread placement strategy, valid values are instanceId (or host, which has the same effect), or any platform or custom attribute that is applied to a container instance, such as attribute:ecs.availability-zone. For the binpack placement strategy, valid values are cpu and memory. For the random placement strategy, this field is not used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "The type of placement strategy. The random placement strategy randomly places tasks on available candidates. The spread placement strategy spreads placement across available candidates evenly based on the field parameter. The binpack strategy places tasks on available candidates that have the least available amount of the resource that is specified with the field parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory (but still enough to run the task).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "platform_version": {
                    "computed": true,
                    "description": "Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as 1.1.0.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "propagate_tags": {
                    "computed": true,
                    "description": "Specifies whether to propagate the tags from the task definition to the task. If no value is specified, the tags are not propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the TagResource API action.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "reference_id": {
                    "computed": true,
                    "description": "The reference ID to use for the task.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
                    "description": "The metadata that you apply to the task to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define. To learn more, see RunTask in the Amazon ECS API Reference.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      [
                        "list",
                        "string"
                      ]
                    ]
                  },
                  "task_count": {
                    "computed": true,
                    "description": "The number of tasks to create based on TaskDefinition. The default is 1.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "task_definition_arn": {
                    "computed": true,
                    "description": "The ARN of the task definition to use if the event target is an Amazon ECS task.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "event_bridge_parameters": {
              "computed": true,
              "description": "EventBridge PutEvent predefined target type.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "detail_type": {
                    "computed": true,
                    "description": "Free-form string, with a maximum of 128 characters, used to decide what fields to expect in the event detail.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source": {
                    "computed": true,
                    "description": "The source of the event.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "input": {
              "computed": true,
              "description": "The text, or well-formed JSON, passed to the target. If you are configuring a templated Lambda, AWS Step Functions, or Amazon EventBridge target, the input must be a well-formed JSON. For all other target types, a JSON is not required. If you do not specify anything for this field, EventBridge Scheduler delivers a default notification to the target.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kinesis_parameters": {
              "computed": true,
              "description": "The custom parameter you can use to control the shard to which EventBridge Scheduler sends the event.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "partition_key": {
                    "computed": true,
                    "description": "The custom parameter used as the Kinesis partition key. For more information, see Amazon Kinesis Streams Key Concepts in the Amazon Kinesis Streams Developer Guide.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "retry_policy": {
              "computed": true,
              "description": "A RetryPolicy object that includes information about the retry policy settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "maximum_event_age_in_seconds": {
                    "computed": true,
                    "description": "The maximum amount of time, in seconds, to continue to make retry attempts.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_retry_attempts": {
                    "computed": true,
                    "description": "The maximum number of retry attempts to make before the request fails. Retry attempts with exponential backoff continue until either the maximum number of attempts is made or until the duration of the MaximumEventAgeInSeconds is reached.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "role_arn": {
              "description": "The Amazon Resource Name (ARN) of the IAM role to be used for this target when the schedule is triggered.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sage_maker_pipeline_parameters": {
              "computed": true,
              "description": "These are custom parameters to use when the target is a SageMaker Model Building Pipeline that starts based on AWS EventBridge Scheduler schedules.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "pipeline_parameter_list": {
                    "computed": true,
                    "description": "List of Parameter names and values for SageMaker Model Building Pipeline execution.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "Name of parameter to start execution of a SageMaker Model Building Pipeline.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "Value of parameter to start execution of a SageMaker Model Building Pipeline.",
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
            "sqs_parameters": {
              "computed": true,
              "description": "Contains the message group ID to use when the target is a FIFO queue. If you specify an SQS FIFO queue as a target, the queue must have content-based deduplication enabled.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message_group_id": {
                    "computed": true,
                    "description": "The FIFO message group ID to use as the target.",
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
      }
    },
    "description": "Definition of AWS::Scheduler::Schedule Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSchedulerScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSchedulerSchedule), &result)
	return &result
}
