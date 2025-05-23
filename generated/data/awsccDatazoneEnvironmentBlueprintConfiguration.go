package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneEnvironmentBlueprintConfiguration = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled_regions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "environment_blueprint_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_blueprint_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_role_permission_boundary": {
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
      "manage_access_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_configurations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "lake_formation_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "location_registration_exclude_s3_locations": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "location_registration_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "provisioning_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "regional_parameters": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "parameters": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataZone::EnvironmentBlueprintConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneEnvironmentBlueprintConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneEnvironmentBlueprintConfiguration), &result)
	return &result
}
