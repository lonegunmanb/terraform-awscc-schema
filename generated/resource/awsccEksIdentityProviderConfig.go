package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEksIdentityProviderConfig = `{
  "block": {
    "attributes": {
      "cluster_name": {
        "description": "The name of the identity provider configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_provider_config_arn": {
        "computed": true,
        "description": "The ARN of the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_provider_config_name": {
        "computed": true,
        "description": "The name of the OIDC provider configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oidc": {
        "computed": true,
        "description": "An object representing an OpenID Connect (OIDC) configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "client_id": {
              "description": "This is also known as audience. The ID for the client application that makes authentication requests to the OpenID identity provider.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "groups_claim": {
              "computed": true,
              "description": "The JWT claim that the provider uses to return your groups.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "groups_prefix": {
              "computed": true,
              "description": "The prefix that is prepended to group claims to prevent clashes with existing names (such as system: groups).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "issuer_url": {
              "description": "The URL of the OpenID identity provider that allows the API server to discover public signing keys for verifying tokens.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "required_claims": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "description": "The key of the requiredClaims.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The value for the requiredClaims.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "username_claim": {
              "computed": true,
              "description": "The JSON Web Token (JWT) claim to use as the username. The default is sub, which is expected to be a unique identifier of the end user. You can choose other claims, such as email or name, depending on the OpenID identity provider. Claims other than email are prefixed with the issuer URL to prevent naming clashes with other plug-ins.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "username_prefix": {
              "computed": true,
              "description": "The prefix that is prepended to username claims to prevent clashes with existing names. If you do not provide this field, and username is a value other than email, the prefix defaults to issuerurl#. You can use the value - to disable all prefixing.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "type": {
        "description": "The type of the identity provider configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "An object representing an Amazon EKS IdentityProviderConfig.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEksIdentityProviderConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEksIdentityProviderConfig), &result)
	return &result
}
