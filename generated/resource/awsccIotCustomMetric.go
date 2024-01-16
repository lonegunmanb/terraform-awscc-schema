package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotCustomMetric = `{
  "block": {
    "attributes": {
      "display_name": {
        "computed": true,
        "description": "Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.",
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
      "metric_arn": {
        "computed": true,
        "description": "The Amazon Resource Number (ARN) of the custom metric.",
        "description_kind": "plain",
        "type": "string"
      },
      "metric_name": {
        "computed": true,
        "description": "The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metric_type": {
        "description": "The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The tag's key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The tag's value.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "A custom metric published by your devices to Device Defender.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotCustomMetricSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotCustomMetric), &result)
	return &result
}
