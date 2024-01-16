package resource

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
        "optional": true,
        "type": "string"
      },
      "certificate_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_pem": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_signing_request": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Use the AWS::IoT::Certificate resource to declare an AWS IoT X.509 certificate.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotCertificate), &result)
	return &result
}
