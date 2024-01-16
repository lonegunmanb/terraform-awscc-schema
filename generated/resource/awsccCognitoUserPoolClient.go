package resource

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
        "optional": true,
        "type": "number"
      },
      "allowed_o_auth_flows": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "allowed_o_auth_flows_user_pool_client": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "allowed_o_auth_scopes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "application_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "external_id": {
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
            "user_data_shared": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "auth_session_validity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "callback_ur_ls": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "enable_propagate_additional_user_context_data": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_token_revocation": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "explicit_auth_flows": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "generate_secret": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id_token_validity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "logout_ur_ls": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "read_attributes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "refresh_token_validity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "supported_identity_providers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "id_token": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "refresh_token": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "write_attributes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Resource Type definition for AWS::Cognito::UserPoolClient",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoUserPoolClientSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolClient), &result)
	return &result
}
