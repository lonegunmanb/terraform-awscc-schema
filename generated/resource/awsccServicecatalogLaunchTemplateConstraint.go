package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogLaunchTemplateConstraint = `{
  "block": {
    "attributes": {
      "accept_language": {
        "computed": true,
        "description": "The language code.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the constraint.",
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
      "launch_template_constraint_id": {
        "computed": true,
        "description": "Unique identifier for the constraint",
        "description_kind": "plain",
        "type": "string"
      },
      "portfolio_id": {
        "description": "The portfolio identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_id": {
        "description": "The product identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rules": {
        "description": "A json encoded string of the template constraint rules",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ServiceCatalog::LaunchTemplateConstraint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogLaunchTemplateConstraintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogLaunchTemplateConstraint), &result)
	return &result
}
