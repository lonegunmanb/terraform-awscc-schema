package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkflowmonitorMonitor = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the monitor.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time when the monitor was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_resources": {
        "computed": true,
        "description": "The local resources to monitor.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identifier": {
              "computed": true,
              "description": "The identifier of the local resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of the local resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "modified_at": {
        "computed": true,
        "description": "The date and time when the monitor was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_name": {
        "computed": true,
        "description": "The name of the monitor.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_status": {
        "computed": true,
        "description": "The status of the monitor.",
        "description_kind": "plain",
        "type": "string"
      },
      "remote_resources": {
        "computed": true,
        "description": "The remote resources to monitor.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identifier": {
              "computed": true,
              "description": "The identifier of the remote resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of the remote resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "scope_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the scope for the monitor.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the monitor.",
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
      }
    },
    "description": "Data Source schema for AWS::NetworkFlowMonitor::Monitor",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkflowmonitorMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkflowmonitorMonitor), &result)
	return &result
}
