package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCertificatemanagerAcmeEndpoint = `{
  "block": {
    "attributes": {
      "acme_endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the ACME endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorization_behavior": {
        "computed": true,
        "description": "The authorization behavior for the ACME endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_authority": {
        "computed": true,
        "description": "The certificate authority configuration for the ACME endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "public_certificate_authority": {
              "computed": true,
              "description": "Configuration for the public certificate authority.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_key_algorithms": {
                    "computed": true,
                    "description": "The allowed key algorithms for certificates issued via this endpoint.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "certificate_tags": {
        "computed": true,
        "description": "Tags applied to certificates issued via this endpoint.",
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
      },
      "contact": {
        "computed": true,
        "description": "Whether contact information is required for the ACME endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_url": {
        "computed": true,
        "description": "The ACME directory URL for the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the ACME endpoint.",
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
    "description": "Data Source schema for AWS::CertificateManager::AcmeEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCertificatemanagerAcmeEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCertificatemanagerAcmeEndpoint), &result)
	return &result
}
