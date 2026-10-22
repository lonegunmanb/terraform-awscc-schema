package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconvertCertificateAssociation = `{
  "block": {
    "attributes": {
      "arn": {
        "description": "The ARN of the ACM certificate to associate with MediaConvert.",
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
    "description": "Associates an AWS Certificate Manager (ACM) certificate with AWS Elemental MediaConvert to enable encryption for content delivery.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconvertCertificateAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconvertCertificateAssociation), &result)
	return &result
}
