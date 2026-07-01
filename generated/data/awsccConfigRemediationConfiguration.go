package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigRemediationConfiguration = `{
  "block": {
    "attributes": {
      "automatic": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "config_rule_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "execution_controls": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ssm_controls": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "concurrent_execution_rate_percentage": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "error_percentage": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
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
      "maximum_automatic_attempts": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "retry_attempt_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "target_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Config::RemediationConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConfigRemediationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigRemediationConfiguration), &result)
	return &result
}
