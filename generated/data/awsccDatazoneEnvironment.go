package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneEnvironment = `{
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
        "description": "The AWS region in which the Amazon DataZone environment is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the environment was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The Amazon DataZone user who created the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the Amazon DataZone environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the environment is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the environment would be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_account_identifier": {
        "computed": true,
        "description": "The AWS account in which the Amazon DataZone environment is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_account_region": {
        "computed": true,
        "description": "The AWS region in which the Amazon DataZone environment is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_blueprint_id": {
        "computed": true,
        "description": "The ID of the blueprint with which the Amazon DataZone environment was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_profile_id": {
        "computed": true,
        "description": "The ID of the environment profile with which the Amazon DataZone environment was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_profile_identifier": {
        "computed": true,
        "description": "The ID of the environment profile with which the Amazon DataZone environment would be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_role_arn": {
        "computed": true,
        "description": "Environment role arn for custom aws environment permissions",
        "description_kind": "plain",
        "type": "string"
      },
      "glossary_terms": {
        "computed": true,
        "description": "The glossary terms that can be used in the Amazon DataZone environment.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone project in which the environment is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_identifier": {
        "computed": true,
        "description": "The ID of the Amazon DataZone project in which the environment would be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_name": {
        "computed": true,
        "description": "The provider of the Amazon DataZone environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Amazon DataZone environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp of when the environment was updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_parameters": {
        "computed": true,
        "description": "The user parameters of the Amazon DataZone environment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of an environment parameter.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of an environment parameter.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::DataZone::Environment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneEnvironment), &result)
	return &result
}
