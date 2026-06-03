package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserverlessCollectionIndex = `{
  "block": {
    "attributes": {
      "collection_index_id": {
        "computed": true,
        "description": "The identifier of the collection",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "index_name": {
        "computed": true,
        "description": "The name of the collection index",
        "description_kind": "plain",
        "type": "string"
      },
      "index_schema": {
        "computed": true,
        "description": "The Mappings for the collection index",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::OpenSearchServerless::CollectionIndex",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchserverlessCollectionIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserverlessCollectionIndex), &result)
	return &result
}
