package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2KeyPair = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key_fingerprint": {
        "computed": true,
        "description": "A short sequence of bytes used for public key verification",
        "description_kind": "plain",
        "type": "string"
      },
      "key_format": {
        "computed": true,
        "description": "The format of the private key",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_name": {
        "description": "The name of the SSH key pair",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_pair_id": {
        "computed": true,
        "description": "An AWS generated ID for the key pair",
        "description_kind": "plain",
        "type": "string"
      },
      "key_type": {
        "computed": true,
        "description": "The crypto-system used to generate a key pair.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_key_material": {
        "computed": true,
        "description": "Plain text public key to import",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      }
    },
    "description": "The AWS::EC2::KeyPair creates an SSH key pair",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2KeyPairSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2KeyPair), &result)
	return &result
}
