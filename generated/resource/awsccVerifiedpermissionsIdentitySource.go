package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVerifiedpermissionsIdentitySource = `{
  "block": {
    "attributes": {
      "configuration": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cognito_user_pool_configuration": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "group_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "group_entity_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "user_pool_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_source_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_entity_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::VerifiedPermissions::IdentitySource Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccVerifiedpermissionsIdentitySourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVerifiedpermissionsIdentitySource), &result)
	return &result
}
