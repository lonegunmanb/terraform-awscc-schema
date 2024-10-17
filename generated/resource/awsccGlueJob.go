package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueJob = `{
  "block": {
    "attributes": {
      "allocated_capacity": {
        "computed": true,
        "description": "The number of capacity units that are allocated to this job.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "command": {
        "description": "The code that executes a job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of the job command",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "python_version": {
              "computed": true,
              "description": "The Python version being used to execute a Python shell job.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "runtime": {
              "computed": true,
              "description": "Runtime is used to specify the versions of Ray, Python and additional libraries available in your environment",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "script_location": {
              "computed": true,
              "description": "Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that executes a job",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "connections": {
        "computed": true,
        "description": "Specifies the connections used by a job",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connections": {
              "computed": true,
              "description": "A list of connections used by the job.",
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
      },
      "default_arguments": {
        "computed": true,
        "description": "The default arguments for this job, specified as name-value pairs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the job.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_class": {
        "computed": true,
        "description": "Indicates whether the job is run with a standard or flexible execution class.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_property": {
        "computed": true,
        "description": "The maximum number of concurrent runs that are allowed for this job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_concurrent_runs": {
              "computed": true,
              "description": "The maximum number of concurrent runs allowed for the job.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "glue_version": {
        "computed": true,
        "description": "Glue version determines the versions of Apache Spark and Python that AWS Glue supports.",
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
      "job_mode": {
        "computed": true,
        "description": "Property description not available.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_run_queuing_enabled": {
        "computed": true,
        "description": "Property description not available.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "log_uri": {
        "computed": true,
        "description": "This field is reserved for future use.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maintenance_window": {
        "computed": true,
        "description": "Property description not available.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_capacity": {
        "computed": true,
        "description": "The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_retries": {
        "computed": true,
        "description": "The maximum number of times to retry this job after a JobRun fails",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "The name you assign to the job definition",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "non_overridable_arguments": {
        "computed": true,
        "description": "Non-overridable arguments for this job, specified as name-value pairs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notification_property": {
        "computed": true,
        "description": "Specifies configuration properties of a notification.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "notify_delay_after": {
              "computed": true,
              "description": "It is the number of minutes to wait before sending a job run delay notification after a job run starts",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "number_of_workers": {
        "computed": true,
        "description": "The number of workers of a defined workerType that are allocated when a job runs.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "role": {
        "description": "The name or Amazon Resource Name (ARN) of the IAM role associated with this job.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_configuration": {
        "computed": true,
        "description": "The name of the SecurityConfiguration structure to be used with this job.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to use with this job.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timeout": {
        "computed": true,
        "description": "The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "worker_type": {
        "computed": true,
        "description": "TThe type of predefined worker that is allocated when a job runs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Glue::Job",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueJob), &result)
	return &result
}
