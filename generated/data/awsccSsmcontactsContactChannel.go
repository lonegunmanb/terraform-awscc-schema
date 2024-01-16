package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmcontactsContactChannel = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the engagement to a contact channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_address": {
        "computed": true,
        "description": "The details that SSM Incident Manager uses when trying to engage the contact channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_name": {
        "computed": true,
        "description": "The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_type": {
        "computed": true,
        "description": "Device type, which specify notification channel. Currently supported values: ?SMS?, ?VOICE?, ?EMAIL?, ?CHATBOT.",
        "description_kind": "plain",
        "type": "string"
      },
      "contact_id": {
        "computed": true,
        "description": "ARN of the contact resource",
        "description_kind": "plain",
        "type": "string"
      },
      "defer_activation": {
        "computed": true,
        "description": "If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSMContacts::ContactChannel",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmcontactsContactChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmcontactsContactChannel), &result)
	return &result
}
