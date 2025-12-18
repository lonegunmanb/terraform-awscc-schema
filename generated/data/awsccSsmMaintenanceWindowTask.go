package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmMaintenanceWindowTask = `{
  "block": {
    "attributes": {
      "cutoff_behavior": {
        "computed": true,
        "description": "The specification for whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the task.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "logging_info": {
        "computed": true,
        "description": "Information about an Amazon S3 bucket to write Run Command task-level logs to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_bucket": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_prefix": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "max_concurrency": {
        "computed": true,
        "description": "The maximum number of targets this task can be run for, in parallel.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_errors": {
        "computed": true,
        "description": "The maximum number of errors allowed before this task stops being scheduled.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The task name.",
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description": "The priority of the task in the maintenance window. The lower the number, the higher the priority. Tasks that have the same priority are scheduled in parallel.",
        "description_kind": "plain",
        "type": "number"
      },
      "service_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM service role for AWS Systems Manager to assume when running a maintenance window task.",
        "description_kind": "plain",
        "type": "string"
      },
      "targets": {
        "computed": true,
        "description": "The targets (either instances or window target ids).",
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
      },
      "task_arn": {
        "computed": true,
        "description": "The resource that the task uses during execution.",
        "description_kind": "plain",
        "type": "string"
      },
      "task_invocation_parameters": {
        "computed": true,
        "description": "The parameters to pass to the task when it runs. Populate only the fields that match the task type. All other fields should be empty.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maintenance_window_automation_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "document_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "parameters": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "maintenance_window_lambda_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_context": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "payload": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "qualifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "maintenance_window_run_command_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_output_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cloudwatch_log_group_name": {
                          "computed": true,
                          "description": "The name of the CloudWatch log group where you want to send command output.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "cloudwatch_output_enabled": {
                          "computed": true,
                          "description": "Enables Systems Manager to send command output to CloudWatch Logs.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "comment": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "document_hash": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "document_hash_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "document_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "notification_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "notification_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "notification_events": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "notification_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "output_s3_bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_s3_key_prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "parameters": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "service_role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "timeout_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "maintenance_window_step_functions_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "input": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
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
      "task_parameters": {
        "computed": true,
        "description": "The parameters to pass to the task when it runs.",
        "description_kind": "plain",
        "type": "string"
      },
      "task_type": {
        "computed": true,
        "description": "The type of task.",
        "description_kind": "plain",
        "type": "string"
      },
      "window_id": {
        "computed": true,
        "description": "The ID of the maintenance window where the task is registered.",
        "description_kind": "plain",
        "type": "string"
      },
      "window_task_id": {
        "computed": true,
        "description": "Unique identifier of the maintenance window task.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSM::MaintenanceWindowTask",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmMaintenanceWindowTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmMaintenanceWindowTask), &result)
	return &result
}
