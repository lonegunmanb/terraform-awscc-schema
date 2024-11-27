package resource

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
              "optional": true,
              "type": "string"
            },
            "category": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "color_mode": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "extension": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "client_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
        "optional": true,
        "type": "bool"
      },
      "settings": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "use_cognito_provided_values": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Cognito::ManagedLoginBranding",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoManagedLoginBrandingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoManagedLoginBranding), &result)
	return &result
}
