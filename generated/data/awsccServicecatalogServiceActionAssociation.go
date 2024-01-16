package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogServiceActionAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_artifact_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_action_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceCatalog::ServiceActionAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogServiceActionAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogServiceActionAssociation), &result)
	return &result
}
