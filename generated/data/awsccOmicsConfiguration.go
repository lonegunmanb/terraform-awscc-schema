package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOmicsConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Unique resource identifier for the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "Configuration creation timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Optional description for the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "User-friendly name for the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "run_configurations": {
        "computed": true,
        "description": "Required run-specific configurations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_ids": {
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "Current configuration status.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for the configuration.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "uuid": {
        "computed": true,
        "description": "Unique identifier for the configuration.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Omics::Configuration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOmicsConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOmicsConfiguration), &result)
	return &result
}
