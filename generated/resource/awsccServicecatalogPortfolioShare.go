package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogPortfolioShare = `{
  "block": {
    "attributes": {
      "accept_language": {
        "computed": true,
        "description": "The language code.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "account_id": {
        "description": "The AWS account ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "portfolio_id": {
        "description": "The portfolio identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "share_tag_options": {
        "computed": true,
        "description": "Enables or disables TagOptions sharing when creating the portfolio share.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Resource Type definition for AWS::ServiceCatalog::PortfolioShare",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogPortfolioShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogPortfolioShare), &result)
	return &result
}
