package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogServiceActionAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provisioning_artifact_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_action_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Schema for AWS::ServiceCatalog::ServiceActionAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogServiceActionAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogServiceActionAssociation), &result)
	return &result
}
