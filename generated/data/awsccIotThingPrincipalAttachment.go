package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotThingPrincipalAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description": "The principal, which can be a certificate ARN (as returned from the CreateCertificate operation) or an Amazon Cognito ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "thing_name": {
        "computed": true,
        "description": "The name of the AWS IoT thing.",
        "description_kind": "plain",
        "type": "string"
      },
      "thing_principal_type": {
        "computed": true,
        "description": "The type of the relation you want to specify when you attach a principal to a thing.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoT::ThingPrincipalAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotThingPrincipalAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotThingPrincipalAttachment), &result)
	return &result
}
