package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VerifiedAccessTrustProvider = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the Amazon Web Services Verified Access trust provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "device_options": {
        "computed": true,
        "description": "The options for device identity based trust providers.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "public_signing_key_url": {
              "computed": true,
              "description": "URL Verified Access will use to verify authenticity of the device tokens.",
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description": "The ID of the tenant application with the device-identity provider.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "device_trust_provider_type": {
        "computed": true,
        "description": "The type of device-based trust provider. Possible values: jamf|crowdstrike",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The last updated time.",
        "description_kind": "plain",
        "type": "string"
      },
      "native_application_oidc_options": {
        "computed": true,
        "description": "The OpenID Connect details for an oidc -type, user-identity based trust provider for L4.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authorization_endpoint": {
              "computed": true,
              "description": "The OIDC authorization endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "client_id": {
              "computed": true,
              "description": "The client identifier.",
              "description_kind": "plain",
              "type": "string"
            },
            "client_secret": {
              "computed": true,
              "description": "The client secret.",
              "description_kind": "plain",
              "type": "string"
            },
            "issuer": {
              "computed": true,
              "description": "The OIDC issuer.",
              "description_kind": "plain",
              "type": "string"
            },
            "public_signing_key_endpoint": {
              "computed": true,
              "description": "The public signing key for endpoint",
              "description_kind": "plain",
              "type": "string"
            },
            "scope": {
              "computed": true,
              "description": "OpenID Connect (OIDC) scopes are used by an application during authentication to authorize access to details of a user. Each scope returns a specific set of user attributes.",
              "description_kind": "plain",
              "type": "string"
            },
            "token_endpoint": {
              "computed": true,
              "description": "The OIDC token endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "user_info_endpoint": {
              "computed": true,
              "description": "The OIDC user info endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "oidc_options": {
        "computed": true,
        "description": "The OpenID Connect details for an oidc -type, user-identity based trust provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authorization_endpoint": {
              "computed": true,
              "description": "The OIDC authorization endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "client_id": {
              "computed": true,
              "description": "The client identifier.",
              "description_kind": "plain",
              "type": "string"
            },
            "client_secret": {
              "computed": true,
              "description": "The client secret.",
              "description_kind": "plain",
              "type": "string"
            },
            "issuer": {
              "computed": true,
              "description": "The OIDC issuer.",
              "description_kind": "plain",
              "type": "string"
            },
            "scope": {
              "computed": true,
              "description": "OpenID Connect (OIDC) scopes are used by an application during authentication to authorize access to details of a user. Each scope returns a specific set of user attributes.",
              "description_kind": "plain",
              "type": "string"
            },
            "token_endpoint": {
              "computed": true,
              "description": "The OIDC token endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "user_info_endpoint": {
              "computed": true,
              "description": "The OIDC user info endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "policy_reference_name": {
        "computed": true,
        "description": "The identifier to be used when working with policy rules.",
        "description_kind": "plain",
        "type": "string"
      },
      "sse_specification": {
        "computed": true,
        "description": "The configuration options for customer provided KMS encryption.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "customer_managed_key_enabled": {
              "computed": true,
              "description": "Whether to encrypt the policy with the provided key or disable encryption",
              "description_kind": "plain",
              "type": "bool"
            },
            "kms_key_arn": {
              "computed": true,
              "description": "KMS Key Arn used to encrypt the group policy",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
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
      "trust_provider_type": {
        "computed": true,
        "description": "Type of trust provider. Possible values: user|device",
        "description_kind": "plain",
        "type": "string"
      },
      "user_trust_provider_type": {
        "computed": true,
        "description": "The type of device-based trust provider. Possible values: oidc|iam-identity-center",
        "description_kind": "plain",
        "type": "string"
      },
      "verified_access_trust_provider_id": {
        "computed": true,
        "description": "The ID of the Amazon Web Services Verified Access trust provider.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::VerifiedAccessTrustProvider",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VerifiedAccessTrustProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VerifiedAccessTrustProvider), &result)
	return &result
}
