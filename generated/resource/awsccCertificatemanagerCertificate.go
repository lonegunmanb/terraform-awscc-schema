package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCertificatemanagerCertificate = `{
  "block": {
    "attributes": {
      "certificate_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the private certificate authority (CA) that will be used to issue the certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_authority_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the private certificate authority (CA) that will be used to issue the certificate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_export": {
        "computed": true,
        "description": "Specifies whether the certificate can be exported. ENABLED allows the certificate to be exported, DISABLED prevents export.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_transparency_logging_preference": {
        "computed": true,
        "description": "You can opt out of certificate transparency logging by specifying the DISABLED option. Opt in by specifying ENABLED.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "description": "The fully qualified domain name (FQDN), such as www.example.com, with which you want to secure an ACM certificate",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_validation_options": {
        "computed": true,
        "description": "Domain information that domain name registrars use to verify your identity.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "domain_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hosted_zone_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "validation_domain": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key_algorithm": {
        "computed": true,
        "description": "Specifies the algorithm of the public and private key pair that your certificate uses to encrypt data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subject_alternative_names": {
        "computed": true,
        "description": "Additional FQDNs to be included in the Subject Alternative Name extension of the ACM certificate.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs that can identify the certificate.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "validation_method": {
        "computed": true,
        "description": "The method you want to use to validate that you own or control the domain associated with a public certificate. Valid values are DNS, EMAIL or HTTP",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CertificateManager::Certificate",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCertificatemanagerCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCertificatemanagerCertificate), &result)
	return &result
}
