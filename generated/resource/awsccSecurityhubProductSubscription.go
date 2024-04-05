package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubProductSubscription = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_arn": {
        "description": "The generic ARN of the product being subscribed to",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_subscription_arn": {
        "computed": true,
        "description": "The ARN of the product subscription for the account",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The AWS::SecurityHub::ProductSubscription resource represents a subscription to a service that is allowed to generate findings for your Security Hub account. One product subscription resource is created for each product enabled.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubProductSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubProductSubscription), &result)
	return &result
}
