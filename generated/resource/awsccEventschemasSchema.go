package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventschemasSchema = `{
  "block": {
    "attributes": {
      "content": {
        "description": "The source of the schema definition.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the schema.",
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
      "last_modified": {
        "computed": true,
        "description": "The last modified time of the schema.",
        "description_kind": "plain",
        "type": "string"
      },
      "registry_name": {
        "description": "The name of the schema registry.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema_arn": {
        "computed": true,
        "description": "The ARN of the schema.",
        "description_kind": "plain",
        "type": "string"
      },
      "schema_name": {
        "computed": true,
        "description": "The name of the schema.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schema_version": {
        "computed": true,
        "description": "The version number of the schema.",
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
      },
      "type": {
        "description": "The type of schema. Valid types include OpenApi3 and JSONSchemaDraft4.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version_created_date": {
        "computed": true,
        "description": "The date the schema version was created.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EventSchemas::Schema",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventschemasSchemaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventschemasSchema), &result)
	return &result
}
