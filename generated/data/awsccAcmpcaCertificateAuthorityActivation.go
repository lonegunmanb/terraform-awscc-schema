package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAcmpcaCertificateAuthorityActivation = `{
  "block": {
    "attributes": {
      "certificate": {
        "computed": true,
        "description": "Certificate Authority certificate that will be installed in the Certificate Authority.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_authority_arn": {
        "computed": true,
        "description": "Arn of the Certificate Authority.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_chain": {
        "computed": true,
        "description": "Certificate chain for the Certificate Authority certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "complete_certificate_chain": {
        "computed": true,
        "description": "The complete certificate chain, including the Certificate Authority certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Certificate Authority.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ACMPCA::CertificateAuthorityActivation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAcmpcaCertificateAuthorityActivationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAcmpcaCertificateAuthorityActivation), &result)
	return &result
}
