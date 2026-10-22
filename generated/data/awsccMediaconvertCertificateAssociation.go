package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconvertCertificateAssociation = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the ACM certificate to associate with MediaConvert.",
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
    "description": "Data Source schema for AWS::MediaConvert::CertificateAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediaconvertCertificateAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconvertCertificateAssociation), &result)
	return &result
}
