package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTransferCertificate = `{
  "block": {
    "attributes": {
      "active_date": {
        "computed": true,
        "description": "Specifies the active date for the certificate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Specifies the unique Amazon Resource Name (ARN) for the agreement.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate": {
        "description": "Specifies the certificate body to be imported.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_chain": {
        "computed": true,
        "description": "Specifies the certificate chain to be imported.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_id": {
        "computed": true,
        "description": "A unique identifier for the certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A textual description for the certificate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "inactive_date": {
        "computed": true,
        "description": "Specifies the inactive date for the certificate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "not_after_date": {
        "computed": true,
        "description": "Specifies the not after date for the certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "not_before_date": {
        "computed": true,
        "description": "Specifies the not before date for the certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_key": {
        "computed": true,
        "description": "Specifies the private key for the certificate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "serial": {
        "computed": true,
        "description": "Specifies Certificate's serial.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "A status description for the certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs that can be used to group and search for certificates. Tags are metadata attached to certificates for any purpose.",
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
        "computed": true,
        "description": "Describing the type of certificate. With or without a private key.",
        "description_kind": "plain",
        "type": "string"
      },
      "usage": {
        "description": "Specifies the usage type for the certificate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Transfer::Certificate",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTransferCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferCertificate), &result)
	return &result
}
