package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneSubscriptionTarget = `{
  "block": {
    "attributes": {
      "applicable_asset_types": {
        "computed": true,
        "description": "The asset types that can be included in the subscription target.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "authorized_principals": {
        "computed": true,
        "description": "The authorized principals of the subscription target.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the subscription target was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The Amazon DataZone user who created the subscription target.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone domain in which subscription target is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description": "The ID of the Amazon DataZone domain in which subscription target would be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The ID of the environment in which subscription target is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_identifier": {
        "computed": true,
        "description": "The ID of the environment in which subscription target would be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "manage_access_role": {
        "computed": true,
        "description": "The manage access role that is used to create the subscription target.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the subscription target.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description": "The identifier of the project specified in the subscription target.",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_name": {
        "computed": true,
        "description": "The provider of the subscription target.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_target_config": {
        "computed": true,
        "description": "The configuration of the subscription target.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content": {
              "computed": true,
              "description": "The content of the subscription target configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "form_name": {
              "computed": true,
              "description": "The form name included in the subscription target configuration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "subscription_target_id": {
        "computed": true,
        "description": "The ID of the subscription target.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of the subscription target.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp of when the subscription target was updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_by": {
        "computed": true,
        "description": "The Amazon DataZone user who updated the subscription target.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataZone::SubscriptionTarget",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneSubscriptionTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneSubscriptionTarget), &result)
	return &result
}
