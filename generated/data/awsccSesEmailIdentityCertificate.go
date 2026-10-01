package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesEmailIdentityCertificate = `{
  "block": {
    "attributes": {
      "certificate_arn": {
        "computed": true,
        "description": "The ARN of the AWS Certificate Manager certificate to associate with the sender.",
        "description_kind": "plain",
        "type": "string"
      },
      "email_identity": {
        "computed": true,
        "description": "The email identity that owns the sender the certificate is associated with.",
        "description_kind": "plain",
        "type": "string"
      },
      "from_address": {
        "computed": true,
        "description": "The sender the certificate signs for. For an email address identity this is the identity itself.",
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
    "description": "Data Source schema for AWS::SES::EmailIdentityCertificate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesEmailIdentityCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesEmailIdentityCertificate), &result)
	return &result
}
