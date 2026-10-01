package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmonitorMonitor = `{
  "block": {
    "attributes": {
      "aggregation_period": {
        "computed": true,
        "description": "The time, in seconds, that metrics are aggregated and sent to Amazon CloudWatch. Valid values are either 30 or 60.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "created_at": {
        "computed": true,
        "description": "The time and date when the monitor was created.",
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
        "description": "The time and date when the monitor was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_arn": {
        "computed": true,
        "description": "The ARN of the monitor.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_name": {
        "description": "The name identifying the monitor. It can contain only letters, underscores (_), or dashes (-), and can be up to 200 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the monitor.",
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
    "description": "An Amazon CloudWatch Network Monitor resource that monitors network traffic between a source VPC subnet and a destination IP address.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmonitorMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmonitorMonitor), &result)
	return &result
}
