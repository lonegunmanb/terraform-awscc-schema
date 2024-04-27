package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPaymentcryptographyAlias = `{
  "block": {
    "attributes": {
      "alias_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::PaymentCryptography::Alias Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPaymentcryptographyAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPaymentcryptographyAlias), &result)
	return &result
}
