package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubProductSubscription = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_arn": {
        "computed": true,
        "description": "The generic ARN of the product being subscribed to",
        "description_kind": "plain",
        "type": "string"
      },
      "product_subscription_arn": {
        "computed": true,
        "description": "The ARN of the product subscription for the account",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityHub::ProductSubscription",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityhubProductSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubProductSubscription), &result)
	return &result
}
