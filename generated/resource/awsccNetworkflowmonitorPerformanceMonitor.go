package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkflowmonitorPerformanceMonitor = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the performance monitor.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time when the performance monitor was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description": "The date and time when the performance monitor was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_name": {
        "description": "The name of the performance monitor.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "monitor_status": {
        "computed": true,
        "description": "The status of the performance monitor.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_type": {
        "computed": true,
        "description": "The type of the performance monitor.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the scope for the performance monitor.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the performance monitor.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
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
      "targets": {
        "description": "The targets that are monitored by the performance monitor.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identifier": {
              "description": "The identifier of the target.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description": "The type of the target.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      }
    },
    "description": "Creates a performance monitor to monitor network performance between a set of targets for a workload.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkflowmonitorPerformanceMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkflowmonitorPerformanceMonitor), &result)
	return &result
}
