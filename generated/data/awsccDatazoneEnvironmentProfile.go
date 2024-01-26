package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneEnvironmentProfile = `{
  "block": {
    "attributes": {
      "aws_account_id": {
        "computed": true,
        "description": "The AWS account in which the Amazon DataZone environment is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_region": {
        "computed": true,
        "description": "The AWS region in which this environment profile is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp of when this environment profile was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The Amazon DataZone user who created this environment profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of this Amazon DataZone environment profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone domain in which this environment profile is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description": "The ID of the Amazon DataZone domain in which this environment profile is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_blueprint_id": {
        "computed": true,
        "description": "The ID of the blueprint with which this environment profile is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_blueprint_identifier": {
        "computed": true,
        "description": "The ID of the blueprint with which this environment profile is created.",
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
        "description": "The name of this Amazon DataZone environment profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description": "The identifier of the project in which to create the environment profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_identifier": {
        "computed": true,
        "description": "The identifier of the project in which to create the environment profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp of when this environment profile was updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_parameters": {
        "computed": true,
        "description": "The user parameters of this Amazon DataZone environment profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of an environment profile parameter.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of an environment profile parameter.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::DataZone::EnvironmentProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneEnvironmentProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneEnvironmentProfile), &result)
	return &result
}
