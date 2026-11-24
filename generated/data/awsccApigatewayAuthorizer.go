package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayAuthorizer = `{
  "block": {
    "attributes": {
      "auth_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_credentials": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_result_ttl_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "authorizer_uri": {
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
      "identity_source": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_validation_expression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "jwt_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_algorithms": {
              "computed": true,
              "description": "The signing algorithms the authorizer accepts: RS256, RS384, RS512, ES256, ES384 or ES512.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "issuers": {
              "computed": true,
              "description": "The token issuers the authorizer accepts.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "additional_claims": {
                    "computed": true,
                    "description": "Other claims the token must carry.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "claim_name": {
                          "computed": true,
                          "description": "The claim name. Can't be iss, aud, exp, nbf, iat, scope or scp.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "expected_values": {
                          "computed": true,
                          "description": "The value the claim must have. Leave it out to check only that the claim is present.",
                          "description_kind": "plain",
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "format": {
                          "computed": true,
                          "description": "The claim's format. SINGLE_STRING is the only format.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "required": {
                          "computed": true,
                          "description": "Whether the token must carry the claim. Defaults to false.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "audiences": {
                    "computed": true,
                    "description": "The values that the token's aud claim must contain at least one of.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "issuer": {
                    "computed": true,
                    "description": "The HTTPS URL that the token's iss claim must match.",
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provider_ar_ns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "rest_api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::Authorizer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayAuthorizerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayAuthorizer), &result)
	return &result
}
