package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesMailManagerAddonSubscription = `{
  "block": {
    "attributes": {
      "addon_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "addon_subscription_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "addon_subscription_id": {
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SES::MailManagerAddonSubscription",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesMailManagerAddonSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMailManagerAddonSubscription), &result)
	return &result
}
