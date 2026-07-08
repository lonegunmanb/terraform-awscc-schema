package resource

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
        "optional": true,
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
      }
    },
    "description": "Resource Type definition for AWS::ServiceCatalog::AcceptedPortfolioShare",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogAcceptedPortfolioShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogAcceptedPortfolioShare), &result)
	return &result
}
