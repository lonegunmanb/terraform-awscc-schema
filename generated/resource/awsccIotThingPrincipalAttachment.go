package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotThingPrincipalAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "description": "The principal, which can be a certificate ARN (as returned from the CreateCertificate operation) or an Amazon Cognito ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "thing_name": {
        "description": "The name of the AWS IoT thing.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "thing_principal_type": {
        "computed": true,
        "description": "The type of the relation you want to specify when you attach a principal to a thing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::IoT::ThingPrincipalAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotThingPrincipalAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotThingPrincipalAttachment), &result)
	return &result
}
