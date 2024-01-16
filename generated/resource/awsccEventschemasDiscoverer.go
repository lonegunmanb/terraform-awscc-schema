package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventschemasDiscoverer = `{
  "block": {
    "attributes": {
      "cross_account": {
        "computed": true,
        "description": "Defines whether event schemas from other accounts are discovered. Default is True.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A description for the discoverer.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "discoverer_arn": {
        "computed": true,
        "description": "The ARN of the discoverer.",
        "description_kind": "plain",
        "type": "string"
      },
      "discoverer_id": {
        "computed": true,
        "description": "The Id of the discoverer.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_arn": {
        "description": "The ARN of the event bus.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Defines the current state of the discoverer.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the resource.",
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
    "description": "Resource Type definition for AWS::EventSchemas::Discoverer",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventschemasDiscovererSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventschemasDiscoverer), &result)
	return &result
}
