package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description": "The date and time when the performance monitor was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_name": {
        "computed": true,
        "description": "The name of the performance monitor.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "targets": {
        "computed": true,
        "description": "The targets that are monitored by the performance monitor.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identifier": {
              "computed": true,
              "description": "The identifier of the target.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of the target.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::NetworkFlowMonitor::PerformanceMonitor",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkflowmonitorPerformanceMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkflowmonitorPerformanceMonitor), &result)
	return &result
}
