package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesEmailIdentityCertificate = `{
  "block": {
    "attributes": {
      "certificate_arn": {
        "description": "The ARN of the AWS Certificate Manager certificate to associate with the sender.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "email_identity": {
        "description": "The email identity that owns the sender the certificate is associated with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "from_address": {
        "description": "The sender the certificate signs for. For an email address identity this is the identity itself.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Associates an AWS Certificate Manager certificate with a sender under an SES email identity, so that SES signs the sender's mail with S/MIME. Removing this resource removes only the association. The ACM certificate itself stays intact. SES stores the sender lower-cased, and this provider matches senders case-insensitively and preserves the casing from the template.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesEmailIdentityCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesEmailIdentityCertificate), &result)
	return &result
}
