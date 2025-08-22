package data

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
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "periodic_scan_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "frequency": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "frequency_expression": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "rule_set_categories": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "level": {
        "computed": true,
        "description": "Configuration Level",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Code Security Scan Configuration name",
        "description_kind": "plain",
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
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::InspectorV2::CodeSecurityScanConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccInspectorv2CodeSecurityScanConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInspectorv2CodeSecurityScanConfiguration), &result)
	return &result
}
