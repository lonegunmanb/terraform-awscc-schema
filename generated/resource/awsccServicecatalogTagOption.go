package resource

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
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key": {
        "description": "The TagOption key.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tag_option_id": {
        "computed": true,
        "description": "The TagOption identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "value": {
        "description": "The TagOption value.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::ServiceCatalog::TagOption",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogTagOptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogTagOption), &result)
	return &result
}
