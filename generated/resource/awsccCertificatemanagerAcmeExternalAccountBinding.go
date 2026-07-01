package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCertificatemanagerAcmeExternalAccountBinding = `{
  "block": {
    "attributes": {
      "acme_endpoint_arn": {
        "description": "The ARN of the ACME endpoint this binding is associated with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "acme_external_account_binding_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the external account binding.",
        "description_kind": "plain",
        "type": "string"
      },
      "expiration": {
        "computed": true,
        "description": "The expiration configuration for the external account binding.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description": "The time unit for the expiration value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The expiration value.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "description": "The IAM role ARN for cross-account access.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the external account binding.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::CertificateManager::AcmeExternalAccountBinding",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCertificatemanagerAcmeExternalAccountBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCertificatemanagerAcmeExternalAccountBinding), &result)
	return &result
}
