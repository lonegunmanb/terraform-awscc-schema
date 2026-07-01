package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCertificatemanagerAcmeDomainValidation = `{
  "block": {
    "attributes": {
      "acme_endpoint_arn": {
        "description": "The ARN of the ACME endpoint this domain validation is associated with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the domain validation.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description": "The domain name to validate.",
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
      "prevalidation_options": {
        "description": "Prevalidation method configuration. Currently only DNS-based prevalidation is supported.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dns_prevalidation": {
              "description": "DNS-based prevalidation options for the domain validation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain_scope": {
                    "computed": true,
                    "description": "Controls which certificate types are authorized to be issued for the domain via the ACME endpoint.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "exact_domain": {
                          "computed": true,
                          "description": "Whether certificates may be issued for the exact domain.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "subdomains": {
                          "computed": true,
                          "description": "Whether certificates may be issued for subdomains of the domain.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "wildcards": {
                          "computed": true,
                          "description": "Whether wildcard certificates may be issued for the domain.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "hosted_zone_id": {
                    "computed": true,
                    "description": "The Route 53 hosted zone ID for automatic DNS record management. When provided, the service creates the validation DNS record on the customer's behalf.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the domain validation.",
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
    "description": "Resource Type definition for AWS::CertificateManager::AcmeDomainValidation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCertificatemanagerAcmeDomainValidationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCertificatemanagerAcmeDomainValidation), &result)
	return &result
}
