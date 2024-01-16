package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tag_key": {
        "computed": true,
        "description": "The key-name for the LF-tag.",
        "description_kind": "plain",
        "type": "string"
      },
      "tag_values": {
        "computed": true,
        "description": "A list of possible values an attribute can take.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::LakeFormation::Tag",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLakeformationTagSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLakeformationTag), &result)
	return &result
}
