package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalProductManagedApiAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "managed_api_ids": {
        "description": "The identifiers of the Managed APIs to associate with the product.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "product_id": {
        "description": "The identifier of the product. The product must use CUSTOM API selection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::APIPortal::ProductManagedApiAssociation. Associates a set of AWS::APIPortal::ManagedApi resources with an AWS::APIPortal::Product. The resource owns all of the product's Managed API associations.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApiportalProductManagedApiAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalProductManagedApiAssociation), &result)
	return &result
}
