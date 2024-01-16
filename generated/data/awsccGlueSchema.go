package data

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
              "type": "bool"
            },
            "version_number": {
              "computed": true,
              "description": "Indicates the version number in the schema to update.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "compatibility": {
        "computed": true,
        "description": "Compatibility setting for the schema.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_format": {
        "computed": true,
        "description": "Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the schema. If description is not provided, there will not be any default value for this.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "initial_schema_version_id": {
        "computed": true,
        "description": "Represents the version ID associated with the initial schema version.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the schema.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "Name of the registry in which the schema will be created.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "schema_definition": {
        "computed": true,
        "description": "Definition for the initial schema version in plain-text.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "List of tags to tag the schema",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A key to identify the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Corresponding tag value for the key.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Glue::Schema",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueSchemaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueSchema), &result)
	return &result
}
