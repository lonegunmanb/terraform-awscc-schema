package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesReceiptRuleSet = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_set_name": {
        "computed": true,
        "description": "The name of the rule set.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SES::ReceiptRuleSet",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesReceiptRuleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesReceiptRuleSet), &result)
	return &result
}
