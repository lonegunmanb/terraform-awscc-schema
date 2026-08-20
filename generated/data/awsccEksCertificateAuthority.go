package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEksCertificateAuthority = `{
  "block": {
    "attributes": {
      "activated_at": {
        "computed": true,
        "description": "The timestamp when the certificate authority was activated.",
        "description_kind": "plain",
        "type": "string"
      },
      "activated_by": {
        "computed": true,
        "description": "The entity that activated the certificate authority.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_authority_id": {
        "computed": true,
        "description": "The unique identifier of the certificate authority.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description": "The name of the EKS cluster that the certificate authority belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the certificate authority was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The entity that created the certificate authority.",
        "description_kind": "plain",
        "type": "string"
      },
      "data": {
        "computed": true,
        "description": "The Base64 encoded certificate authority data.",
        "description_kind": "plain",
        "type": "string"
      },
      "distribution_status": {
        "computed": true,
        "description": "The distribution status of the certificate authority.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rollback_available": {
        "computed": true,
        "description": "Whether activation of this certificate authority can still be rolled back.",
        "description_kind": "plain",
        "type": "bool"
      },
      "scheduled_events": {
        "computed": true,
        "description": "The scheduled auto-activation events for the certificate authority.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "final_auto_activation": {
              "computed": true,
              "description": "The deadline by which EKS will auto-activate this certificate authority (notAfter minus 45 days).",
              "description_kind": "plain",
              "type": "string"
            },
            "first_auto_activation": {
              "computed": true,
              "description": "The earliest date EKS may auto-activate this certificate authority (notAfter minus 6 months).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "signing_status": {
        "computed": true,
        "description": "The signing status of the certificate authority.",
        "description_kind": "plain",
        "type": "string"
      },
      "validity": {
        "computed": true,
        "description": "The validity period of the certificate authority.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "not_after": {
              "computed": true,
              "description": "The end of the validity period for the certificate authority.",
              "description_kind": "plain",
              "type": "string"
            },
            "not_before": {
              "computed": true,
              "description": "The start of the validity period for the certificate authority.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::EKS::CertificateAuthority",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEksCertificateAuthoritySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEksCertificateAuthority), &result)
	return &result
}
