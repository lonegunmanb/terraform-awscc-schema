package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventsv2EventBus = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time the event bus was created, as an ISO 8601 timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the event bus. Control characters and Unicode line separators are not allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "Encryption configuration for the event bus. The service stores and returns the customer managed key as its key ARN. The key ARN is recommended so that drift detection stays accurate. A key ID is also accepted and is matched to the returned key ARN when CloudFormation checks for drift; a key alias is accepted but can be reported as a value difference, because an alias cannot be matched to the key ARN it points to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_identifier": {
              "computed": true,
              "description": "The identifier of the AWS KMS customer managed key that the event bus uses to encrypt events. You can specify the key ARN, key ID, alias name, or alias ARN. If you do not specify a key, EventBridge uses an AWS owned key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "event_bus_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the event bus.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time the event bus was last modified, as an ISO 8601 timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the event bus. The first character must be alphanumeric; the remaining characters may also include '.', '-', and '_'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The lifecycle state of the event bus. The settled operational state is ACTIVE.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_configuration": {
        "computed": true,
        "description": "The event storage configuration for the event bus, which controls the number of days events are retained on the bus.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "retention_period_in_days": {
              "computed": true,
              "description": "The number of days events are retained on the event bus for replay, 1-365.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the event bus.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key. For each resource, each tag key must be unique and each key can have only one value; keys are case sensitive. A key cannot begin or end with a whitespace character; whitespace inside the key is allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value. May be empty. A value cannot begin or end with a whitespace character; whitespace inside the value is allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource type definition for AWS::EventsV2::EventBus, an Amazon EventBridge custom event bus that receives events and delivers them to matching subscribers.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventsv2EventBusSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsv2EventBus), &result)
	return &result
}
