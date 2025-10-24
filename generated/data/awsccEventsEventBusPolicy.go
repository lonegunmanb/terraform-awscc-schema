package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventsEventBusPolicy = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description": "The action that you are enabling the other account to perform.",
        "description_kind": "plain",
        "type": "string"
      },
      "condition": {
        "computed": true,
        "description": "This parameter enables you to limit the permission to accounts that fulfill a certain condition, such as being a member of a certain AWS organization.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Specifies the value for the key. Currently, this must be the ID of the organization.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "Specifies the type of condition. Currently the only supported value is StringEquals.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Specifies the key for the condition. Currently the only supported key is aws:PrincipalOrgID.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "event_bus_name": {
        "computed": true,
        "description": "The name of the event bus associated with the rule. If you omit this, the default event bus is used.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description": "The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify \"*\" to permit any account to put events to your default event bus.",
        "description_kind": "plain",
        "type": "string"
      },
      "statement": {
        "computed": true,
        "description": "A JSON string that describes the permission policy statement. You can include a Policy parameter in the request instead of using the StatementId, Action, Principal, or Condition parameters.",
        "description_kind": "plain",
        "type": "string"
      },
      "statement_id": {
        "computed": true,
        "description": "An identifier string for the external account that you are granting permissions to",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Events::EventBusPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEventsEventBusPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsEventBusPolicy), &result)
	return &result
}
