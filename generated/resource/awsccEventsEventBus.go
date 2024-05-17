package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventsEventBus = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the event bus.",
        "description_kind": "plain",
        "type": "string"
      },
      "dead_letter_config": {
        "computed": true,
        "description": "Dead Letter Queue for the event bus.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description": "The description of the event bus.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_source_name": {
        "computed": true,
        "description": "If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.",
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
      "kms_key_identifier": {
        "computed": true,
        "description": "Kms Key Identifier used to encrypt events at rest in the event bus.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the event bus.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description": "A JSON string that describes the permission policy statement for the event bus.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the event bus.",
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
      }
    },
    "description": "Resource type definition for AWS::Events::EventBus",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventsEventBusSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsEventBus), &result)
	return &result
}
