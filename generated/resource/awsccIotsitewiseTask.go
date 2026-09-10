package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseTask = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description of the task.",
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
      "status": {
        "computed": true,
        "description": "The current lifecycle status of the task.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "task_arn": {
        "computed": true,
        "description": "The ARN of the task.",
        "description_kind": "plain",
        "type": "string"
      },
      "task_configuration": {
        "description": "The task execution configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_task_configuration": {
              "description": "Configuration for running a custom container image on managed compute.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description": "The command to execute in the container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "ecr_uri": {
                    "description": "The Amazon ECR image URI for the task container.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "environment_variables": {
                    "computed": true,
                    "description": "A map of environment variable key-value pairs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "processing_type": {
                    "description": "The processing type for compute resources.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "processing_unit": {
                    "description": "The processing unit allocation that determines vCPU, memory, and GPU resources.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "task_execution_role": {
                    "description": "The ARN of the IAM role that grants the containerized workload permissions to access AWS resources.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timeout_seconds": {
                    "computed": true,
                    "description": "The timeout in seconds for task execution. Default: 3600 (1 hour).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "task_name": {
        "description": "The name of the task. Must be unique within the workspace.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workspace_name": {
        "description": "The name of the workspace.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::IoTSiteWise::Task. A task defines a reusable containerized compute workload that can be referenced by one or more pipeline compute nodes.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotsitewiseTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseTask), &result)
	return &result
}
