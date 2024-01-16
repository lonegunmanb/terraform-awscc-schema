package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolClient = `{
  "block": {
    "attributes": {
      "access_token_validity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "allowed_o_auth_flows": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "allowed_o_auth_flows_user_pool_client": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "allowed_o_auth_scopes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "analytics_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "application_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "external_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "user_data_shared": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "auth_session_validity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "callback_ur_ls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "client_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_secret": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_redirect_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_propagate_additional_user_context_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_token_revocation": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "explicit_auth_flows": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "generate_secret": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id_token_validity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "logout_ur_ls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "prevent_user_existence_errors": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "read_attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "refresh_token_validity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "supported_identity_providers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "token_validity_units": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_token": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "id_token": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "refresh_token": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "write_attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Cognito::UserPoolClient",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoUserPoolClientSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolClient), &result)
	return &result
}
