package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsCertificate = `{
  "block": {
    "attributes": {
      "certificate_arn": {
        "computed": true,
        "description": "The certificate Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_identifier": {
        "computed": true,
        "description": "The certificate Identifier",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_pem": {
        "computed": true,
        "description": "The certificate Pem",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_wallet": {
        "computed": true,
        "description": "The certificate Wallet",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::DMS::Certificate",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDmsCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsCertificate), &result)
	return &result
}
