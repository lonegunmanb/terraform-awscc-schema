package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneSubscriptionTarget = `{
  "block": {
    "attributes": {
      "applicable_asset_types": {
        "description": "The asset types that can be included in the subscription target.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "authorized_principals": {
        "description": "The authorized principals of the subscription target.",
        "description_kind": "plain",
        "required": true,
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
        "description": "The ID of the Amazon DataZone domain in which subscription target would be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The ID of the environment in which subscription target is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_identifier": {
        "description": "The ID of the environment in which subscription target would be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "manage_access_role": {
        "computed": true,
        "description": "The manage access role that is used to create the subscription target.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the subscription target.",
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      },
      "subscription_target_config": {
        "description": "The configuration of the subscription target.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content": {
              "description": "The content of the subscription target configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "form_name": {
              "description": "The form name included in the subscription target configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "subscription_target_id": {
        "computed": true,
        "description": "The ID of the subscription target.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "The type of the subscription target.",
        "description_kind": "plain",
        "required": true,
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
    "description": "Subscription targets enables one to access the data to which you have subscribed in your projects.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneSubscriptionTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneSubscriptionTarget), &result)
	return &result
}
