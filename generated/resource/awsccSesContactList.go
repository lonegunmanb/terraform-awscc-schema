package resource

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
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the contact list.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The tags (keys and values) associated with the contact list.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "topics": {
        "computed": true,
        "description": "The topics associated with the contact list.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_subscription_status": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description of the topic.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "display_name": {
              "description": "The display name of the topic.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "topic_name": {
              "description": "The name of the topic.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::SES::ContactList.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesContactListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesContactList), &result)
	return &result
}
