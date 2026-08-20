package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesReceiptRuleSet = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_set_name": {
        "computed": true,
        "description": "The name of the rule set.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SES::ReceiptRuleSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesReceiptRuleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesReceiptRuleSet), &result)
	return &result
}
