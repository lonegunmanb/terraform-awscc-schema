package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccInspectorv2CodeSecurityScanConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Code Security Scan Configuration ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description": "Code Security Scan Configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "continuous_integration_scan_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "supported_events": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "periodic_scan_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "frequency": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "frequency_expression": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "rule_set_categories": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "level": {
        "computed": true,
        "description": "Configuration Level",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Code Security Scan Configuration name",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope_settings": {
        "computed": true,
        "description": "Scope Settings",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "project_selection_scope": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Inspector CodeSecurityScanConfiguration resource schema",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccInspectorv2CodeSecurityScanConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInspectorv2CodeSecurityScanConfiguration), &result)
	return &result
}
