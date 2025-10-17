package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogTagOptionAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "The CloudformationProduct or Portfolio identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "tag_option_id": {
        "computed": true,
        "description": "The TagOption identifier.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceCatalog::TagOptionAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogTagOptionAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogTagOptionAssociation), &result)
	return &result
}
