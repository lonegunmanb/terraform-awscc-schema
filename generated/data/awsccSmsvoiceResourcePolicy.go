package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSmsvoiceResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "The JSON formatted resource-based policy to attach.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AWS End User Messaging SMS and Voice resource to attach the resource-based policy to.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SMSVOICE::ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSmsvoiceResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSmsvoiceResourcePolicy), &result)
	return &result
}
