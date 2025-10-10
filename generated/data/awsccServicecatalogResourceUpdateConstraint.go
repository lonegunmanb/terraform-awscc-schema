package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogResourceUpdateConstraint = `{
  "block": {
    "attributes": {
      "accept_language": {
        "computed": true,
        "description": "The language code",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the constraint",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "portfolio_id": {
        "computed": true,
        "description": "The portfolio identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description": "The product identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_update_constraint_id": {
        "computed": true,
        "description": "Unique identifier for the constraint",
        "description_kind": "plain",
        "type": "string"
      },
      "tag_update_on_provisioned_product": {
        "computed": true,
        "description": "ALLOWED or NOT_ALLOWED, to permit or prevent changes to the tags on provisioned instances of the specified portfolio / product combination",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceCatalog::ResourceUpdateConstraint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogResourceUpdateConstraintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogResourceUpdateConstraint), &result)
	return &result
}
