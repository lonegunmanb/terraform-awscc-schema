package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoIdentityPoolRoleAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_pool_role_attachment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_mappings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ambiguous_role_resolution": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "identity_provider": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rules_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rules": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "claim": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "match_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "map"
        }
      },
      "roles": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Cognito::IdentityPoolRoleAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoIdentityPoolRoleAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoIdentityPoolRoleAttachment), &result)
	return &result
}
