package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCeAnomalyMonitor = `{
  "block": {
    "attributes": {
      "creation_date": {
        "computed": true,
        "description": "The date when the monitor was created. ",
        "description_kind": "plain",
        "type": "string"
      },
      "dimensional_value_count": {
        "computed": true,
        "description": "The value for evaluated dimensions.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_evaluated_date": {
        "computed": true,
        "description": "The date when the monitor last evaluated for anomalies.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_date": {
        "computed": true,
        "description": "The date when the monitor was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_arn": {
        "computed": true,
        "description": "Monitor ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_dimension": {
        "computed": true,
        "description": "The dimensions to evaluate",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_name": {
        "computed": true,
        "description": "The name of the monitor.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_specification": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_tags": {
        "computed": true,
        "description": "Tags to assign to monitor.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name for the tag.",
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
      }
    },
    "description": "Data Source schema for AWS::CE::AnomalyMonitor",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCeAnomalyMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCeAnomalyMonitor), &result)
	return &result
}
