package data

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
        "type": "string"
      },
      "certificate_pem": {
        "computed": true,
        "description": "The certificate Pem",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_wallet": {
        "computed": true,
        "description": "The certificate Wallet",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DMS::Certificate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDmsCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsCertificate), &result)
	return &result
}
