package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventschemasRegistry = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description of the registry to be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "registry_arn": {
        "computed": true,
        "description": "The ARN of the registry.",
        "description_kind": "plain",
        "type": "string"
      },
      "registry_name": {
        "computed": true,
        "description": "The name of the schema registry.",
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
    "description": "Data Source schema for AWS::EventSchemas::Registry",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEventschemasRegistrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventschemasRegistry), &result)
	return &result
}
