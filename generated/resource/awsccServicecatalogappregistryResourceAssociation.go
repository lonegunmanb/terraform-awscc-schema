package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogappregistryResourceAssociation = `{
  "block": {
    "attributes": {
      "application": {
        "description": "The name or the Id of the Application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "application_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource": {
        "description": "The name or the Id of the Resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "description": "The type of the CFN Resource for now it's enum CFN_STACK.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Schema for AWS::ServiceCatalogAppRegistry::ResourceAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogappregistryResourceAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogappregistryResourceAssociation), &result)
	return &result
}
