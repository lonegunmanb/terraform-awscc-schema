package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppintegrationsEventIntegration = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The event integration description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_bridge_bus": {
        "description": "The Amazon Eventbridge bus for the event integration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_filter": {
        "description": "The EventFilter (source) associated with the event integration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source": {
              "description": "The source of the events.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "event_integration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the event integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the event integration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the event integration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A key to identify the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Corresponding tag value for the key.",
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
    "description": "Resource Type definition for AWS::AppIntegrations::EventIntegration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppintegrationsEventIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppintegrationsEventIntegration), &result)
	return &result
}
