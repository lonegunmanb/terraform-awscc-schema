package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogTagOptionAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "The CloudformationProduct or Portfolio identifier.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tag_option_id": {
        "computed": true,
        "description": "The TagOption identifier.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ServiceCatalog::TagOptionAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogTagOptionAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogTagOptionAssociation), &result)
	return &result
}
