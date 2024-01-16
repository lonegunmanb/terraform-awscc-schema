package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesContactList = `{
  "block": {
    "attributes": {
      "contact_list_name": {
        "computed": true,
        "description": "The name of the contact list.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the contact list.",
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
        "description": "The tags (keys and values) associated with the contact list.",
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
      },
      "topics": {
        "computed": true,
        "description": "The topics associated with the contact list.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_subscription_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description of the topic.",
              "description_kind": "plain",
              "type": "string"
            },
            "display_name": {
              "computed": true,
              "description": "The display name of the topic.",
              "description_kind": "plain",
              "type": "string"
            },
            "topic_name": {
              "computed": true,
              "description": "The name of the topic.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SES::ContactList",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesContactListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesContactList), &result)
	return &result
}
