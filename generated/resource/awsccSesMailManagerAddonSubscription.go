package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesMailManagerAddonSubscription = `{
  "block": {
    "attributes": {
      "addon_name": {
        "description_kind": "plain",
        "required": true,
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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::SES::MailManagerAddonSubscription Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesMailManagerAddonSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMailManagerAddonSubscription), &result)
	return &result
}
