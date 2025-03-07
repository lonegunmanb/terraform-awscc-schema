package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamSamlProvider = `{
  "block": {
    "attributes": {
      "add_private_key": {
        "computed": true,
        "description": "The private key from your external identity provider",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the SAML provider",
        "description_kind": "plain",
        "type": "string"
      },
      "assertion_encryption_mode": {
        "computed": true,
        "description": "The encryption setting for the SAML provider",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_key_list": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key_id": {
              "computed": true,
              "description": "The unique identifier for the SAML private key.",
              "description_kind": "plain",
              "type": "string"
            },
            "timestamp": {
              "computed": true,
              "description": "The date and time, in \u003ca href=\\\"http://www.iso.org/iso/iso8601\\\"\u003eISO 8601 date-time \u003c/a\u003e format, when the private key was uploaded.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "remove_private_key": {
        "computed": true,
        "description": "The Key ID of the private key to remove",
        "description_kind": "plain",
        "type": "string"
      },
      "saml_metadata_document": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "saml_provider_uuid": {
        "computed": true,
        "description": "The unique identifier assigned to the SAML provider",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::IAM::SAMLProvider",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIamSamlProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamSamlProvider), &result)
	return &result
}
