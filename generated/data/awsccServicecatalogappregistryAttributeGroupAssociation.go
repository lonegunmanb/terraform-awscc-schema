package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogappregistryAttributeGroupAssociation = `{
  "block": {
    "attributes": {
      "application": {
        "computed": true,
        "description": "The name or the Id of the Application.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "attribute_group": {
        "computed": true,
        "description": "The name or the Id of the AttributeGroup.",
        "description_kind": "plain",
        "type": "string"
      },
      "attribute_group_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogappregistryAttributeGroupAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogappregistryAttributeGroupAssociation), &result)
	return &result
}
