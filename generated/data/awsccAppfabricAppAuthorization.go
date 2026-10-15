package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppfabricAppAuthorization = `{
  "block": {
    "attributes": {
      "app": {
        "computed": true,
        "description": "The name of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "app_authorization_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the app authorization.",
        "description_kind": "plain",
        "type": "string"
      },
      "app_authorization_identifier": {
        "computed": true,
        "description": "The unique identifier of the app authorization.",
        "description_kind": "plain",
        "type": "string"
      },
      "app_bundle_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the app bundle for the app authorization.",
        "description_kind": "plain",
        "type": "string"
      },
      "app_bundle_identifier": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.",
        "description_kind": "plain",
        "type": "string"
      },
      "auth_type": {
        "computed": true,
        "description": "The authorization type.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the app authorization was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "credential": {
        "computed": true,
        "description": "Contains credentials for the application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "api_key_credential": {
              "computed": true,
              "description": "Contains API key credential information.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "api_key": {
                    "computed": true,
                    "description": "An API key for an application.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "oauth_2_credential": {
              "computed": true,
              "description": "Contains OAuth2 client credential information.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description": "The client ID of the client application.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_secret": {
                    "computed": true,
                    "description": "The client secret of the client application.",
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "persona": {
        "computed": true,
        "description": "The user persona of the app authorization.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The state of the app authorization.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to apply to the app authorization.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "tenant": {
        "computed": true,
        "description": "Contains information about an application tenant.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "tenant_display_name": {
              "computed": true,
              "description": "The display name of the tenant.",
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_identifier": {
              "computed": true,
              "description": "The ID of the application tenant.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp of when the app authorization was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AppFabric::AppAuthorization",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppfabricAppAuthorizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppfabricAppAuthorization), &result)
	return &result
}
