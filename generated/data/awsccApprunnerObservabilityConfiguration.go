package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApprunnerObservabilityConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "latest": {
        "computed": true,
        "description": "It's set to true for the configuration with the highest Revision among all configurations that share the same Name. It's set to false otherwise.",
        "description_kind": "plain",
        "type": "bool"
      },
      "observability_configuration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of this ObservabilityConfiguration",
        "description_kind": "plain",
        "type": "string"
      },
      "observability_configuration_name": {
        "computed": true,
        "description": "A name for the observability configuration. When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "observability_configuration_revision": {
        "computed": true,
        "description": "The revision of this observability configuration. It's unique among all the active configurations ('Status': 'ACTIVE') that share the same ObservabilityConfigurationName.",
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "A list of metadata items that you can associate with your observability configuration resource. A tag is a key-value pair.",
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
      "trace_configuration": {
        "computed": true,
        "description": "The configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vendor": {
              "computed": true,
              "description": "The implementation provider chosen for tracing App Runner services.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::AppRunner::ObservabilityConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApprunnerObservabilityConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApprunnerObservabilityConfiguration), &result)
	return &result
}
