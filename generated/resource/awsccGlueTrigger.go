package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueTrigger = `{
  "block": {
    "attributes": {
      "actions": {
        "description": "The actions initiated by this trigger.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arguments": {
              "computed": true,
              "description": "The job arguments used when this trigger fires. For this job run, they replace the default arguments set in the job definition itself.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "crawler_name": {
              "computed": true,
              "description": "The name of the crawler to be used with this action.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "job_name": {
              "computed": true,
              "description": "The name of a job to be executed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "notification_property": {
              "computed": true,
              "description": "Specifies configuration properties of a job run notification.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "notify_delay_after": {
                    "computed": true,
                    "description": "After a job run starts, the number of minutes to wait before sending a job run delay notification",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "security_configuration": {
              "computed": true,
              "description": "The name of the SecurityConfiguration structure to be used with this action.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timeout": {
              "computed": true,
              "description": "The JobRun timeout in minutes. This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours). This overrides the timeout value set in the parent job.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "description": {
        "computed": true,
        "description": "A description of this trigger.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_batching_condition": {
        "computed": true,
        "description": "Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "batch_size": {
              "computed": true,
              "description": "Number of events that must be received from Amazon EventBridge before EventBridge event trigger fires.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "batch_window": {
              "computed": true,
              "description": "Window of time in seconds after which EventBridge event trigger fires. Window starts when first event is received.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the trigger.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "predicate": {
        "computed": true,
        "description": "The predicate of this trigger, which defines when it will fire.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "conditions": {
              "computed": true,
              "description": "A list of the conditions that determine when the trigger will fire.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "crawl_state": {
                    "computed": true,
                    "description": "The state of the crawler to which this condition applies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "crawler_name": {
                    "computed": true,
                    "description": "The name of the crawler to which this condition applies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "job_name": {
                    "computed": true,
                    "description": "The name of the job whose JobRuns this condition applies to, and on which this trigger waits.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "logical_operator": {
                    "computed": true,
                    "description": "A logical operator.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "state": {
                    "computed": true,
                    "description": "The condition state. Currently, the values supported are SUCCEEDED, STOPPED, TIMEOUT, and FAILED.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "logical": {
              "computed": true,
              "description": "An optional field if only one condition is listed. If multiple conditions are listed, then this field is required.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "schedule": {
        "computed": true,
        "description": "A cron expression used to specify the schedule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_on_creation": {
        "computed": true,
        "description": "Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is not supported for ON_DEMAND triggers.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "The tags to use with this trigger.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The type of trigger that this is.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workflow_name": {
        "computed": true,
        "description": "The name of the workflow associated with the trigger.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Glue::Trigger",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueTriggerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueTrigger), &result)
	return &result
}
