package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMwaaEnvironment = `{
  "block": {
    "attributes": {
      "airflow_configuration_options": {
        "computed": true,
        "description": "Key/value pairs representing Airflow configuration variables.\n    Keys are prefixed by their section:\n\n    [core]\n    dags_folder={AIRFLOW_HOME}/dags\n\n    Would be represented as\n\n    \"core.dags_folder\": \"{AIRFLOW_HOME}/dags\"",
        "description_kind": "plain",
        "type": "string"
      },
      "airflow_version": {
        "computed": true,
        "description": "Version of airflow to deploy to the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "ARN for the MWAA environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "celery_executor_queue": {
        "computed": true,
        "description": "The celery executor queue associated with the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "dag_s3_path": {
        "computed": true,
        "description": "Represents an S3 prefix relative to the root of an S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_vpc_endpoint_service": {
        "computed": true,
        "description": "The database VPC endpoint service name.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_management": {
        "computed": true,
        "description": "Defines whether the VPC endpoints configured for the environment are created, and managed, by the customer or by Amazon MWAA.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_class": {
        "computed": true,
        "description": "Templated configuration for airflow processes and backing infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "IAM role to be used by tasks.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key": {
        "computed": true,
        "description": "The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for MWAA data encryption.\n\n    You can specify the CMK using any of the following:\n\n    Key ID. For example, key/1234abcd-12ab-34cd-56ef-1234567890ab.\n\n    Key alias. For example, alias/ExampleAlias.\n\n    Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.\n\n    Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.\n\n    AWS authenticates the CMK asynchronously. Therefore, if you specify an ID, alias, or ARN that is not valid, the action can appear to complete, but eventually fails.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_configuration": {
        "computed": true,
        "description": "Logging configuration for the environment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dag_processing_logs": {
              "computed": true,
              "description": "Logging configuration for a specific airflow component.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "log_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "scheduler_logs": {
              "computed": true,
              "description": "Logging configuration for a specific airflow component.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "log_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "task_logs": {
              "computed": true,
              "description": "Logging configuration for a specific airflow component.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "log_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "webserver_logs": {
              "computed": true,
              "description": "Logging configuration for a specific airflow component.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "log_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "worker_logs": {
              "computed": true,
              "description": "Logging configuration for a specific airflow component.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "log_level": {
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
      "max_webservers": {
        "computed": true,
        "description": "Maximum webserver compute units.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_workers": {
        "computed": true,
        "description": "Maximum worker compute units.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_webservers": {
        "computed": true,
        "description": "Minimum webserver compute units.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_workers": {
        "computed": true,
        "description": "Minimum worker compute units.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "Customer-defined identifier for the environment, unique per customer region.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_configuration": {
        "computed": true,
        "description": "Configures the network resources of the environment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "A list of security groups to use for the environment.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "A list of subnets to use for the environment. These must be private subnets, in the same VPC, in two different availability zones.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "plugins_s3_object_version": {
        "computed": true,
        "description": "Represents an version ID for an S3 object.",
        "description_kind": "plain",
        "type": "string"
      },
      "plugins_s3_path": {
        "computed": true,
        "description": "Represents an S3 prefix relative to the root of an S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "requirements_s3_object_version": {
        "computed": true,
        "description": "Represents an version ID for an S3 object.",
        "description_kind": "plain",
        "type": "string"
      },
      "requirements_s3_path": {
        "computed": true,
        "description": "Represents an S3 prefix relative to the root of an S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "schedulers": {
        "computed": true,
        "description": "Scheduler compute units.",
        "description_kind": "plain",
        "type": "number"
      },
      "source_bucket_arn": {
        "computed": true,
        "description": "ARN for the AWS S3 bucket to use as the source of DAGs and plugins for the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "startup_script_s3_object_version": {
        "computed": true,
        "description": "Represents an version ID for an S3 object.",
        "description_kind": "plain",
        "type": "string"
      },
      "startup_script_s3_path": {
        "computed": true,
        "description": "Represents an S3 prefix relative to the root of an S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of tags for the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "webserver_access_mode": {
        "computed": true,
        "description": "Choice for mode of webserver access including over public internet or via private VPC endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "webserver_url": {
        "computed": true,
        "description": "Url endpoint for the environment's Airflow UI.",
        "description_kind": "plain",
        "type": "string"
      },
      "webserver_vpc_endpoint_service": {
        "computed": true,
        "description": "The webserver VPC endpoint service name, applicable if private webserver access mode selected.",
        "description_kind": "plain",
        "type": "string"
      },
      "weekly_maintenance_window_start": {
        "computed": true,
        "description": "Start time for the weekly maintenance window.",
        "description_kind": "plain",
        "type": "string"
      },
      "worker_replacement_strategy": {
        "computed": true,
        "description": "The worker replacement strategy to use when updating the environment. Valid values: ` + "`" + `FORCED` + "`" + `, ` + "`" + `GRACEFUL` + "`" + `. FORCED means Apache Airflow workers will be stopped and replaced without waiting for tasks to complete before an update. GRACEFUL means Apache Airflow workers will be able to complete running tasks for up to 12 hours during an update before being stopped and replaced.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::MWAA::Environment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMwaaEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMwaaEnvironment), &result)
	return &result
}
