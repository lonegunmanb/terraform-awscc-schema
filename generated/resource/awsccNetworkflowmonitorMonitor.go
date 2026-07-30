package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_resources": {
        "description": "The local resources to monitor.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identifier": {
              "description": "The identifier of the local resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description": "The type of the local resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "modified_at": {
        "computed": true,
        "description": "The date and time when the monitor was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_name": {
        "description": "The name of the monitor.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of the remote resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "scope_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the scope for the monitor.",
        "description_kind": "plain",
        "optional": true,
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
      }
    },
    "description": "Creates a monitor for specific network flows between local and remote resources to monitor network performance for workloads.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkflowmonitorMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkflowmonitorMonitor), &result)
	return &result
}
