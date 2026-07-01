package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightOAuthClientApplication = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_secret": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_type": {
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
      "identity_provider_vpc_connection_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_connection_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "o_auth_authorization_endpoint_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "o_auth_client_application_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "o_auth_client_authentication_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "o_auth_scopes": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "o_auth_token_endpoint_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
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
    "description": "Data Source schema for AWS::QuickSight::OAuthClientApplication",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightOAuthClientApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightOAuthClientApplication), &result)
	return &result
}
