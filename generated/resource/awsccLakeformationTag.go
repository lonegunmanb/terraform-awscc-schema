package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLakeformationTag = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "computed": true,
        "description": "The identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.",
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
      "tag_key": {
        "description": "The key-name for the LF-tag.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tag_values": {
        "description": "A list of possible values an attribute can take.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "A resource schema representing a Lake Formation Tag.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLakeformationTagSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLakeformationTag), &result)
	return &result
}
