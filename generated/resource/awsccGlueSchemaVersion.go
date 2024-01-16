package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueSchemaVersion = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "schema": {
        "description": "Identifier for the schema where the schema version will be created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "registry_name": {
              "computed": true,
              "description": "Name of the registry to identify where the Schema is located.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schema_arn": {
              "computed": true,
              "description": "Amazon Resource Name for the Schema. This attribute can be used to uniquely represent the Schema.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schema_name": {
              "computed": true,
              "description": "Name of the schema. This parameter requires RegistryName to be provided.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "schema_definition": {
        "description": "Complete definition of the schema in plain-text.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version_id": {
        "computed": true,
        "description": "Represents the version ID associated with the schema version.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "This resource represents an individual schema version of a schema defined in Glue Schema Registry.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueSchemaVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueSchemaVersion), &result)
	return &result
}
