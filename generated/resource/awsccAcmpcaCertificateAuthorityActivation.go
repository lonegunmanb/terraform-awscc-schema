package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAcmpcaCertificateAuthorityActivation = `{
  "block": {
    "attributes": {
      "certificate": {
        "description": "Certificate Authority certificate that will be installed in the Certificate Authority.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_authority_arn": {
        "description": "Arn of the Certificate Authority.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_chain": {
        "computed": true,
        "description": "Certificate chain for the Certificate Authority certificate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "complete_certificate_chain": {
        "computed": true,
        "description": "The complete certificate chain, including the Certificate Authority certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Certificate Authority.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Used to install the certificate authority certificate and update the certificate authority status.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAcmpcaCertificateAuthorityActivationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAcmpcaCertificateAuthorityActivation), &result)
	return &result
}
