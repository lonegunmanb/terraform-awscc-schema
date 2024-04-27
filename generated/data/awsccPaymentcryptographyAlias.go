package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPaymentcryptographyAlias = `{
  "block": {
    "attributes": {
      "alias_name": {
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
      "key_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::PaymentCryptography::Alias",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPaymentcryptographyAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPaymentcryptographyAlias), &result)
	return &result
}
