package resource

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
        "optional": true,
        "type": "bool"
      },
      "config_rule_name": {
        "description_kind": "plain",
        "required": true,
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
                    "optional": true,
                    "type": "number"
                  },
                  "error_percentage": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
      "maximum_automatic_attempts": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retry_attempt_seconds": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "target_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Config::RemediationConfiguration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConfigRemediationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigRemediationConfiguration), &result)
	return &result
}
