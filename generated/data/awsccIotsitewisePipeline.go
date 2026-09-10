package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewisePipeline = `{
  "block": {
    "attributes": {
      "computations": {
        "computed": true,
        "description": "The list of compute nodes that form the pipeline DAG.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compute_node_name": {
              "computed": true,
              "description": "The unique name for this compute node within the pipeline.",
              "description_kind": "plain",
              "type": "string"
            },
            "depends_on": {
              "computed": true,
              "description": "A list of compute node names that must complete successfully before this node can start.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "environment_variables": {
              "computed": true,
              "description": "A map of environment variable key-value pairs.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "task_name": {
              "computed": true,
              "description": "The name of the task to execute for this compute node.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_variables": {
        "computed": true,
        "description": "Environment variables shared across all compute nodes in the pipeline.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pipeline_arn": {
        "computed": true,
        "description": "The ARN of the pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "pipeline_name": {
        "computed": true,
        "description": "The name of the pipeline. Must be unique within the workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current lifecycle status of the pipeline.",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "workspace_name": {
        "computed": true,
        "description": "The name of the workspace.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoTSiteWise::Pipeline",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotsitewisePipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewisePipeline), &result)
	return &result
}
