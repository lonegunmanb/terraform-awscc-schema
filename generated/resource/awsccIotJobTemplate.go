package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotJobTemplate = `{
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
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description": "The type of job action to take to initiate the job abort.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "failure_type": {
                    "computed": true,
                    "description": "The type of job execution failures that can initiate a job abort.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_number_of_executed_things": {
                    "computed": true,
                    "description": "The minimum number of things which must receive job execution notifications before the job can be aborted.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "threshold_percentage": {
                    "computed": true,
                    "description": "The minimum percentage of job execution failures that must occur to initiate the job abort.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description of the Job Template.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_package_versions": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "document": {
        "computed": true,
        "description": "The job document. Required if you don't specify a value for documentSource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "document_source": {
        "computed": true,
        "description": "An S3 link to the job document to use in the template. Required if you don't specify a value for document.",
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
      "job_arn": {
        "computed": true,
        "description": "Optional for copying a JobTemplate from a pre-existing Job configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_executions_retry_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "retry_criteria_list": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "failure_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "number_of_retries": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      "job_executions_rollout_config": {
        "computed": true,
        "description": "Allows you to create a staged rollout of a job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exponential_rollout_rate": {
              "computed": true,
              "description": "The rate of increase for a job rollout. This parameter allows you to define an exponential rate for a job rollout.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "base_rate_per_minute": {
                    "computed": true,
                    "description": "The minimum number of things that will be notified of a pending job, per minute at the start of job rollout. This parameter allows you to define the initial rate of rollout.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "increment_factor": {
                    "computed": true,
                    "description": "The exponential factor to increase the rate of rollout for a job.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "rate_increase_criteria": {
                    "computed": true,
                    "description": "The criteria to initiate the increase in rate of rollout for a job.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "number_of_notified_things": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "number_of_succeeded_things": {
                          "computed": true,
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
            "maximum_per_minute": {
              "computed": true,
              "description": "The maximum number of things that will be notified of a pending job, per minute. This parameter allows you to create a staged rollout.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "job_template_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maintenance_windows": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "duration_in_minutes": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "start_time": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "presigned_url_config": {
        "computed": true,
        "description": "Configuration for pre-signed S3 URLs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "expires_in_sec": {
              "computed": true,
              "description": "How number (in seconds) pre-signed URLs are valid.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of an IAM role that grants grants permission to download files from the S3 bucket where the job data/updates are stored. The role must also grant permission for IoT to download the files.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "Metadata that can be used to manage the JobTemplate.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "timeout_config": {
        "computed": true,
        "description": "Specifies the amount of time each device has to finish its execution of the job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "in_progress_timeout_in_minutes": {
              "computed": true,
              "description": "Specifies the amount of time, in minutes, this device has to finish execution of this job.",
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
    "description": "Resource Type definition for AWS::IoT::JobTemplate. Job templates enable you to preconfigure jobs so that you can deploy them to multiple sets of target devices.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotJobTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotJobTemplate), &result)
	return &result
}
