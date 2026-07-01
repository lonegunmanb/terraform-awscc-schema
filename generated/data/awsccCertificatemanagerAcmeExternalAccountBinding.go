package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCertificatemanagerAcmeExternalAccountBinding = `{
  "block": {
    "attributes": {
      "acme_endpoint_arn": {
        "computed": true,
        "description": "The ARN of the ACME endpoint this binding is associated with.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The expiration value.",
              "description_kind": "plain",
              "type": "number"
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
      "role_arn": {
        "computed": true,
        "description": "The IAM role ARN for cross-account access.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::CertificateManager::AcmeExternalAccountBinding",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCertificatemanagerAcmeExternalAccountBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCertificatemanagerAcmeExternalAccountBinding), &result)
	return &result
}
