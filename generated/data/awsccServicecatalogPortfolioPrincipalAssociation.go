package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogPortfolioPrincipalAssociation = `{
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
      },
      "principal_arn": {
        "computed": true,
        "description": "The ARN of the principal (user, role, or group).",
        "description_kind": "plain",
        "type": "string"
      },
      "principal_type": {
        "computed": true,
        "description": "The principal type. The supported value is IAM if you use a fully defined Amazon Resource Name (ARN), or IAM_PATTERN if you use an ARN with no accountID, with or without wildcard characters.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceCatalog::PortfolioPrincipalAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogPortfolioPrincipalAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogPortfolioPrincipalAssociation), &result)
	return &result
}
