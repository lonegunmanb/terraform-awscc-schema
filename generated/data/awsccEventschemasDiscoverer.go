package data

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
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A description for the discoverer.",
        "description_kind": "plain",
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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_arn": {
        "computed": true,
        "description": "The ARN of the event bus.",
        "description_kind": "plain",
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
    "description": "Data Source schema for AWS::EventSchemas::Discoverer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEventschemasDiscovererSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventschemasDiscoverer), &result)
	return &result
}
