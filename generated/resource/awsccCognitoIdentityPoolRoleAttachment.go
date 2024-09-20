package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoIdentityPoolRoleAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_pool_id": {
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "identity_provider": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "match_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "map"
        },
        "optional": true
      },
      "roles": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource Type definition for AWS::Cognito::IdentityPoolRoleAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoIdentityPoolRoleAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoIdentityPoolRoleAttachment), &result)
	return &result
}
