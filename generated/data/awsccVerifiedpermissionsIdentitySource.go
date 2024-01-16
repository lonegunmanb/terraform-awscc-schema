package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVerifiedpermissionsIdentitySource = `{
  "block": {
    "attributes": {
      "configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cognito_user_pool_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "user_pool_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "client_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "discovery_url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "open_id_issuer": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "user_pool_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_source_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "principal_entity_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::VerifiedPermissions::IdentitySource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccVerifiedpermissionsIdentitySourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVerifiedpermissionsIdentitySource), &result)
	return &result
}
