package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueSchemaVersionMetadata = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key": {
        "computed": true,
        "description": "Metadata key",
        "description_kind": "plain",
        "type": "string"
      },
      "schema_version_id": {
        "computed": true,
        "description": "Represents the version ID associated with the schema version.",
        "description_kind": "plain",
        "type": "string"
      },
      "value": {
        "computed": true,
        "description": "Metadata value",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Glue::SchemaVersionMetadata",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueSchemaVersionMetadataSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueSchemaVersionMetadata), &result)
	return &result
}
