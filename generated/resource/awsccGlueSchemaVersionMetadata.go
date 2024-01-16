package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueSchemaVersionMetadata = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key": {
        "description": "Metadata key",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema_version_id": {
        "description": "Represents the version ID associated with the schema version.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "value": {
        "description": "Metadata value",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "This resource adds Key-Value metadata to a Schema version of Glue Schema Registry.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueSchemaVersionMetadataSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueSchemaVersionMetadata), &result)
	return &result
}
