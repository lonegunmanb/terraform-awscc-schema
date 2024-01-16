package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
        "optional": true,
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
    "description": "Resource Type definition for AWS::EventSchemas::Registry",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventschemasRegistrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventschemasRegistry), &result)
	return &result
}
