package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoManagedLoginBranding = `{
  "block": {
    "attributes": {
      "assets": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bytes": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "category": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "color_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "extension": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "resource_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "client_id": {
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
      "managed_login_branding_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "return_merged_resources": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "settings": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "use_cognito_provided_values": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::ManagedLoginBranding",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoManagedLoginBrandingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoManagedLoginBranding), &result)
	return &result
}
