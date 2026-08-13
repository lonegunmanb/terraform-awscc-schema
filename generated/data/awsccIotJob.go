package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotJob = `{
  "block": {
    "attributes": {
      "abort_config": {
        "computed": true,
        "description": "The criteria that determine when and how a job abort takes place.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "criteria_list": {
              "computed": true,
              "description": "The list of criteria that determine when and how to abort the job.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description": "The type of job action to take to initiate the job abort.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "failure_type": {
                    "computed": true,
                    "description": "The type of job execution failures that can initiate a job abort.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "min_number_of_executed_things": {
                    "computed": true,
                    "description": "The minimum number of things which must receive job execution notifications before the job can be aborted.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "threshold_percentage": {
                    "computed": true,
                    "description": "The minimum percentage of job execution failures that must occur to initiate the job abort.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the job.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time when the job was created, in ISO 8601 date-time format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A short text description of the job.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_package_versions": {
        "computed": true,
        "description": "The package version Amazon Resource Names (ARNs) that are installed on the device when the job successfully completes.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "document": {
        "computed": true,
        "description": "The job document. Required if you don't specify a value for documentSource.",
        "description_kind": "plain",
        "type": "string"
      },
      "document_parameters": {
        "computed": true,
        "description": "Parameters of an Amazon Web Services managed template that you can specify to create the job document.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "document_source": {
        "computed": true,
        "description": "An S3 link, or S3 object URL, to the job document. The link is an Amazon S3 object URL and is required if you don't specify a value for document.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "job_executions_retry_config": {
        "computed": true,
        "description": "The configuration that determines how many retries are allowed for each failure type for a job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "criteria_list": {
              "computed": true,
              "description": "The list of criteria that determines how many retries are allowed for each failure type for a job.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "failure_type": {
                    "computed": true,
                    "description": "The type of job execution failures that can initiate a job retry.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "number_of_retries": {
                    "computed": true,
                    "description": "The number of retries allowed for a failure type for the job.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "job_executions_rollout_config": {
        "computed": true,
        "description": "Allows you to create a staged rollout of a job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exponential_rate": {
              "computed": true,
              "description": "Allows you to create an exponential rate of rollout for a job.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "base_rate_per_minute": {
                    "computed": true,
                    "description": "The minimum number of things that will be notified of a pending job, per minute at the start of job rollout.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "increment_factor": {
                    "computed": true,
                    "description": "The exponential factor to increase the rate of rollout for a job.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "rate_increase_criteria": {
                    "computed": true,
                    "description": "Allows you to define a criteria to initiate the increase in rate of rollout for a job.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "number_of_notified_things": {
                          "computed": true,
                          "description": "The threshold for number of notified things that will initiate the increase in rate of rollout.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "number_of_succeeded_things": {
                          "computed": true,
                          "description": "The threshold for number of succeeded things that will initiate the increase in rate of rollout.",
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
            "maximum_per_minute": {
              "computed": true,
              "description": "The maximum number of things that will be notified of a pending job, per minute. This parameter allows you to create a staged rollout.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "job_id": {
        "computed": true,
        "description": "A job identifier which must be unique for your AWS account. We recommend using a UUID. Alpha-numeric characters, '-' and '_' are valid for use here.",
        "description_kind": "plain",
        "type": "string"
      },
      "job_template_arn": {
        "computed": true,
        "description": "The ARN of the job template used to create the job.",
        "description_kind": "plain",
        "type": "string"
      },
      "presigned_url_config": {
        "computed": true,
        "description": "Configuration for pre-signed S3 URLs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "expires_in_sec": {
              "computed": true,
              "description": "How long (in seconds) pre-signed URLs are valid. Valid values are 60 - 3600, the default value is 3600 seconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of an IAM role that grants permission to download files from the S3 bucket where the job data/updates are stored.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "scheduling_config": {
        "computed": true,
        "description": "Specifies the date and time that a job will begin the rollout of the job document to all devices in the target group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_behavior": {
              "computed": true,
              "description": "Specifies the end behavior for all job executions after a job reaches the selected endTime.",
              "description_kind": "plain",
              "type": "string"
            },
            "end_time": {
              "computed": true,
              "description": "The time a job will stop rollout of the job document to all devices in the target group for a job.",
              "description_kind": "plain",
              "type": "string"
            },
            "maintenance_windows": {
              "computed": true,
              "description": "An optional configuration within the SchedulingConfig to setup a recurring maintenance window.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "duration_in_minutes": {
                    "computed": true,
                    "description": "Displays the duration of the next maintenance window.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "start_time": {
                    "computed": true,
                    "description": "Displays the start time of the next maintenance window.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "start_time": {
              "computed": true,
              "description": "The time a job will begin rollout of the job document to all devices in the target group for a job.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "Metadata which can be used to manage the job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "target_selection": {
        "computed": true,
        "description": "Specifies whether the job will continue to run (CONTINUOUS), or will be complete after all those things specified as targets have completed the job (SNAPSHOT).",
        "description_kind": "plain",
        "type": "string"
      },
      "targets": {
        "computed": true,
        "description": "A list of things and thing groups to which the job should be sent.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "timeout_config": {
        "computed": true,
        "description": "Specifies the amount of time each device has to finish its execution of the job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "in_progress_timeout_in_minutes": {
              "computed": true,
              "description": "Specifies the amount of time, in minutes, this device has to finish execution of this job. The timeout interval can be anywhere between 1 minute and 7 days (1 to 10080 minutes).",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::IoT::Job",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotJob), &result)
	return &result
}
