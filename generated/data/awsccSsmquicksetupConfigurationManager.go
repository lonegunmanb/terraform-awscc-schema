package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmquicksetupConfigurationManager = `{
  "block": {
    "attributes": {
      "configuration_definitions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "local_deployment_administration_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "local_deployment_execution_role_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "parameters": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "manager_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_summaries": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "last_updated_at": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status_details": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "status_message": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
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
    "description": "Data Source schema for AWS::SSMQuickSetup::ConfigurationManager",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmquicksetupConfigurationManagerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmquicksetupConfigurationManager), &result)
	return &result
}
