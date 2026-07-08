package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogAcceptedPortfolioShare = `{
  "block": {
    "attributes": {
      "accept_language": {
        "computed": true,
        "description": "The language code.",
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
        "description": "The portfolio identifier.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceCatalog::AcceptedPortfolioShare",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogAcceptedPortfolioShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogAcceptedPortfolioShare), &result)
	return &result
}
