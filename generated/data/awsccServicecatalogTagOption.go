package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogTagOption = `{
  "block": {
    "attributes": {
      "active": {
        "computed": true,
        "description": "The TagOption active state.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key": {
        "computed": true,
        "description": "The TagOption key.",
        "description_kind": "plain",
        "type": "string"
      },
      "tag_option_id": {
        "computed": true,
        "description": "The TagOption identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "value": {
        "computed": true,
        "description": "The TagOption value.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceCatalog::TagOption",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogTagOptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogTagOption), &result)
	return &result
}
