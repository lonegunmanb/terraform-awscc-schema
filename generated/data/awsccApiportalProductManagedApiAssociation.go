package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalProductManagedApiAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_api_ids": {
        "computed": true,
        "description": "The identifiers of the Managed APIs to associate with the product.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "product_id": {
        "computed": true,
        "description": "The identifier of the product. The product must use CUSTOM API selection.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::APIPortal::ProductManagedApiAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApiportalProductManagedApiAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalProductManagedApiAssociation), &result)
	return &result
}
