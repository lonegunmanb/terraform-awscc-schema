package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueSchema = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name for the Schema.",
        "description_kind": "plain",
        "type": "string"
      },
      "checkpoint_version": {
        "computed": true,
        "description": "Specify checkpoint version for update. This is only required to update the Compatibility.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "is_latest": {
              "computed": true,
              "description": "Indicates if the latest version needs to be updated.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "version_number": {
              "computed": true,
              "description": "Indicates the version number in the schema to update.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "compatibility": {
        "description": "Compatibility setting for the schema.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_format": {
        "description": "Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the schema. If description is not provided, there will not be any default value for this.",
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
      "initial_schema_version_id": {
        "computed": true,
        "description": "Represents the version ID associated with the initial schema version.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the schema.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "registry": {
        "computed": true,
        "description": "Identifier for the registry which the schema is part of.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "Amazon Resource Name for the Registry.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "Name of the registry in which the schema will be created.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "schema_definition": {
        "description": "Definition for the initial schema version in plain-text.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "List of tags to tag the schema",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A key to identify the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "Corresponding tag value for the key.",
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
    "description": "This resource represents a schema of Glue Schema Registry.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueSchemaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueSchema), &result)
	return &result
}
