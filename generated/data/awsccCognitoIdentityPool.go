package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoIdentityPool = `{
  "block": {
    "attributes": {
      "allow_classic_flow": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "allow_unauthenticated_identities": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "cognito_events": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cognito_identity_providers": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "client_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "provider_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "server_side_token_check": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "list"
        }
      },
      "cognito_streams": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "stream_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "streaming_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "developer_provider_name": {
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
      "identity_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_pool_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_pool_tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "open_id_connect_provider_ar_ns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "push_sync": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_arns": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "saml_provider_ar_ns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "supported_login_providers": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::IdentityPool",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoIdentityPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoIdentityPool), &result)
	return &result
}
