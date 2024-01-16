package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotCertificate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ca_certificate_pem": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_pem": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_signing_request": {
        "computed": true,
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
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoT::Certificate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotCertificate), &result)
	return &result
}
