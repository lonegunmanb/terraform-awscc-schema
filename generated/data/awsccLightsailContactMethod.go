package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailContactMethod = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the contact method.",
        "description_kind": "plain",
        "type": "string"
      },
      "contact_endpoint": {
        "computed": true,
        "description": "The destination of the contact method, such as an email address or a mobile phone number.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the contact method was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the contact method.",
        "description_kind": "plain",
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description": "The protocol of the contact method, such as Email or SMS (text messaging).",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The Lightsail resource type of the contact method.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the contact method.",
        "description_kind": "plain",
        "type": "string"
      },
      "support_code": {
        "computed": true,
        "description": "The support code for the contact method.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Lightsail::ContactMethod",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLightsailContactMethodSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailContactMethod), &result)
	return &result
}
