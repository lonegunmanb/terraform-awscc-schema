package resource

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
        "optional": true,
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the environment is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "description": "The identifier of the Amazon DataZone domain in which the environment would be created.",
        "description_kind": "plain",
        "required": true,
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
        "description": "The ID of the environment profile with which the Amazon DataZone environment would be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "glossary_terms": {
        "computed": true,
        "description": "The glossary terms that can be used in the Amazon DataZone environment.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the environment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone project in which the environment is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_identifier": {
        "description": "The ID of the Amazon DataZone project in which the environment would be created.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of an environment parameter.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::DataZone::Environment Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneEnvironment), &result)
	return &result
}
